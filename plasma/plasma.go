// Copyright 2019 OmiseGO Pte Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plasma

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	// "github.com/omisego/plasma-cli/rootchain"
	"github.com/omisego/plasma-cli/abi"
	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
)

type RootchainTransaction interface {
	Submit() (*types.Transaction, error)
	Build() (error)
	Options(*bind.TransactOpts) error
}
type Client struct {
	bind.ContractBackend
}


//TODO fix transaction builder signature, make the API less brittle 
// form a deposit transaction 
type DepositBuilder  func(common.Address, common.Address, uint64, uint64) ([]byte, error)

// TODO implement vault binder for ERC20
type VaultBinder func(common.Address, bind.ContractBackend) (*abi.Ethvault, error) 

type DepositTransaction struct {
	*bind.TransactOpts
	bind.ContractBackend
	VaultAddress common.Address
	Amount string
	Owner common.Address
	Currency common.Address
	VaultBinder VaultBinder
	Builder DepositBuilder
	RlpTransaction []byte
}

type ExitBinder func(common.Address, bind.ContractBackend) (*abi.PaymentExitgame, error)

type StandardExitTransaction struct {
	*bind.TransactOpts
	bind.ContractBackend
	ExitGameAddress common.Address
	ExitBinder ExitBinder
	Proof []byte
	TxBytes []byte
	UtxoPos *big.Int
	UtxoData StandardExitUTXOData
}

// TODO: make binder return struct a more generic reusable contract 
type PlasmaFrameworkBinder func(common.Address, bind.ContractBackend) (*abi.PlasmaFramework, error)

type ProcessExitTransaction struct {
	*bind.TransactOpts
	bind.ContractBackend
	*ProcessExitOptions
	PlasmaFrameworkBinder PlasmaFrameworkBinder
}

type ProcessExitOptions struct {
	PlasmaAddress common.Address
	TokenAddress common.Address
	VaultId string
	TopExit string
	NumberOfExits string
}

type ProcessExitOption func(*ProcessExitOptions)


type StandardExitUTXOData struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		UtxoPos *big.Int `json:"utxo_pos"`
		Txbytes string   `json:"txbytes"`
		Proof   string   `json:"proof"`
	} `json:"data"`
}


//Retrieve the UTXO exit data from the UTXO position
func GetUTXOExitData(watcher string, utxoPosition int) (StandardExitUTXOData, error) {
	// Build request
	var url strings.Builder
	url.WriteString(watcher)
	url.WriteString("/utxo.get_exit_data")
	postData := map[string]interface{}{"utxo_pos": utxoPosition}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := StandardExitUTXOData{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return response, errors.New("No exit data found for UTXO provided")
	}

	return response, nil
}


// initialize new root chain client 
func NewClient(ethclient bind.ContractBackend) *Client{
	return &Client{ethclient}
}

// create a new deposit 
func (c *Client) NewDeposit(vaultAddress, owner, currency common.Address, amount string) *DepositTransaction{
	return &DepositTransaction{
		ContractBackend: c.ContractBackend,
		VaultAddress: vaultAddress,
		VaultBinder: abi.NewEthvault,
		Amount: amount,
		Owner: owner,
		Builder: util.BuildRLPDeposit,
		Currency: currency,
	}
}


// set Ethereum transaction options, ensure fields are valid 
func (d *DepositTransaction) Options(t *bind.TransactOpts) error {
	//TODO: add validations
	d.TransactOpts = t
	return nil
}

// build deposit transaction 
func (d *DepositTransaction) Build() error {
	amount, err := strconv.ParseUint(d.Amount, 10, 64)
	if err != nil {
		return err
	}
	rlpencoded, err := d.Builder(
		d.Owner,
		d.Currency,
		amount,
		1,
	)
	if err != nil {
		return err
	}
	d.RlpTransaction = rlpencoded
	return nil
}

// TODO: return something else 
func (d *DepositTransaction) Submit() (*types.Transaction, error) {
	instance, err := d.VaultBinder(d.VaultAddress, d.ContractBackend)
	if err != nil {
		return nil, err
	}
	res, err := instance.Deposit(d.TransactOpts, d.RlpTransaction)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Client) NewStandardExit(vaultAddress common.Address) *StandardExitTransaction {
	return &StandardExitTransaction{
		ContractBackend: c.ContractBackend,
		ExitGameAddress: vaultAddress,
		ExitBinder: abi.NewPaymentExitgame,
	}
}

// get standard exit bond
func (c *StandardExitTransaction) GetStandardExitBond(opts *bind.CallOpts) (*big.Int, error) {
	instance, err := c.ExitBinder(c.ExitGameAddress, c.ContractBackend)
	if err != nil {
		return nil, err
	}
	amount, err := instance.StartStandardExitBondSize(opts)
	if err != nil {
		return nil, err
	}
	return amount, nil
}

//TODO check validations
func (s *StandardExitTransaction) Options(t *bind.TransactOpts) error {
	s.TransactOpts = t
	return nil
}

// set UTXO data to standard exit
func (s *StandardExitTransaction) Utxo(sed StandardExitUTXOData) error {
	s.UtxoData = sed
	return nil
}

// deserialize transaction txbytes and proofs 
func (s *StandardExitTransaction) Build() error {
	proof, err := hex.DecodeString(util.RemoveLeadingZeroX(s.UtxoData.Data.Proof))
	if err != nil {
		return err
	}
	txbytes, err := hex.DecodeString(util.RemoveLeadingZeroX(s.UtxoData.Data.Txbytes))
	if err != nil {
		return err
	}
	s.Proof = proof
	s.TxBytes = txbytes
	s.UtxoPos = s.UtxoData.Data.UtxoPos
	return nil
}

//submit standard exit transaction 
func (s *StandardExitTransaction) Submit() (*types.Transaction, error) {
	instance, err := s.ExitBinder(s.ExitGameAddress, s.ContractBackend)
	if err != nil {
		return nil, err
	}

	seargs := abi.PaymentStandardExitRouterArgsStartStandardExitArgs{
		UtxoPos: s.UtxoPos,
		RlpOutputTx: s.TxBytes,
		OutputTxInclusionProof: s.Proof,
	}

	res, err := instance.StartStandardExit(s.TransactOpts, seargs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func PlasmaAddress(address common.Address) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.PlasmaAddress = address
	}
}
func TokenAddress(address common.Address) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.TokenAddress = address
	}
}
func VaultId(vid string) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.VaultId = vid
	}
}
func TopExit(te string) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.TopExit = te
	}
}
func NumberOfExits(no string) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.NumberOfExits = no
	}
}

func (c *Client) NewProcessExit(setters ...ProcessExitOption) *ProcessExitTransaction  {
	peargs := &ProcessExitOptions{}
	for _, setter := range setters {
		setter(peargs)
	}
	return &ProcessExitTransaction{
		ProcessExitOptions: peargs,
		ContractBackend: c.ContractBackend,
		PlasmaFrameworkBinder: abi.NewPlasmaFramework,
	}
}

// set options for process exit transaction
func(p *ProcessExitTransaction) Options(t *bind.TransactOpts) error {
	p.TransactOpts = t
	return nil
}


// build the transaction
func (p *ProcessExitTransaction) Build() error {
	// TODO validation
	return nil
}

// submit process exit transaction
func (p *ProcessExitTransaction) Submit() (*types.Transaction, error) {
	instance, err := p.PlasmaFrameworkBinder(p.PlasmaAddress, p.ContractBackend)
	if err != nil {
		return nil, err
	}
	exitnum, ok := new(big.Int).SetString(p.NumberOfExits, 10)
	if !ok {
		return nil, errors.New("cannot convert to big int in process exit")
	}
	topexit, ok :=  new(big.Int).SetString(p.TopExit, 10)
	if !ok {
		return nil, errors.New("cannot convert to big int in process exit")
	}
	vaultId, ok :=  new(big.Int).SetString(p.VaultId, 10)
	if !ok {
		return nil, errors.New("cannot convert to big int in process exit")
	}
	tx, err := instance.ProcessExits(
		p.TransactOpts,
		vaultId,
		p.TokenAddress,
		topexit,
		exitnum,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func Options(rtx  RootchainTransaction, t *bind.TransactOpts) error {
	return rtx.Options(t)
}

func Submit(rtx RootchainTransaction) ( *types.Transaction, error ) {
	return rtx.Submit()
}

func Build(rtx RootchainTransaction) error {
	return rtx.Build()
}







