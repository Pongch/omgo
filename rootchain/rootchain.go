//
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

package rootchain

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
	"github.com/pongch/omgo/abi"
	"github.com/pongch/omgo/util"
	log "github.com/sirupsen/logrus"
)

// RootChainTransaction is an interface for a write operations on the root chain contract
type RootchainTransaction interface {
	Submit() (*types.Transaction, error)
	Build() (error)
	Options(*bind.TransactOpts) error
}

// Client defines a struct that can read and write to a contract
type Client struct {
	bind.ContractBackend
}


// DepositBuilder is a function that builds an RLP input from a set of inputs
// TODO clean up transaction builder signature to return a curried function instead
// form a deposit transaction
type DepositBuilder  func(common.Address, common.Address, uint64, uint64) ([]byte, error)

// VaultBinder binds the go code to an Ethvault contract on the root chain
// TODO implement vault binder for ERC20
type VaultBinder func(common.Address, bind.ContractBackend) (*abi.Ethvault, error) 

// DepositTransaction is a collection of functions and data needed to make a deposit transaction on the root chain
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

// ExitBinder binds the go code to a payment exit game contract on the root chain
type ExitBinder func(common.Address, bind.ContractBackend) (*abi.PaymentExitgame, error)

// StandardExitTransaction is a collection of functions and data needed to make a payment exit transaction on the root chain
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

// PlasmaFrameworkBinder binds the go code to a plasma framework contract on the root chain
// TODO: make binder return struct a more generic reusable contract
type PlasmaFrameworkBinder func(common.Address, bind.ContractBackend) (*abi.PlasmaFramework, error)

// ProcessExitTransaction is a collection of functions and data needed to make a process exit transaction on the root chain
type ProcessExitTransaction struct {
	*bind.TransactOpts
	bind.ContractBackend
	*ProcessExitOptions
	PlasmaFrameworkBinder PlasmaFrameworkBinder
}

// ProcessExitOptions holds the input data required to make a process exit
type ProcessExitOptions struct {
	PlasmaAddress common.Address
	TokenAddress common.Address
	VaultId string
	TopExit string
	NumberOfExits string
}

//  ProcessExitOption is a functional option that performs an operation on ProcessExitOptions
type ProcessExitOption func(*ProcessExitOptions)

// StandardExitUTXOData is a returned data from calling /utxo.get_exit_data from the watcher, the data is required to start a standard exit on the utxo
type StandardExitUTXOData struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		UtxoPos *big.Int `json:"utxo_pos"`
		Txbytes string   `json:"txbytes"`
		Proof   string   `json:"proof"`
	} `json:"data"`
}


// GetUTXOExitData is a helper function to retrieve the UTXO exit data from a given UTXO position
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


// NewClient initializes new root chain client
func NewClient(ethclient bind.ContractBackend) *Client{
	return &Client{ethclient}
}

// NewDeposit is a method on root chain client that create a new deposit
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


// Options is a method on DepositTransaction that set root chain transaction options
func (d *DepositTransaction) Options(t *bind.TransactOpts) error {
	//TODO: add validations
	d.TransactOpts = t
	return nil
}

// Bulid is a method on DepositTransaction that encodes the transaction to RLP with the RLP Builder
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

// Submit is a method on DepositTransaction that calls a deposit on the root chain vault
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

// NewStandardExit is a method on the root chain client that creates new standard payment exit
func (c *Client) NewStandardExit(vaultAddress common.Address) *StandardExitTransaction {
	return &StandardExitTransaction{
		ContractBackend: c.ContractBackend,
		ExitGameAddress: vaultAddress,
		ExitBinder: abi.NewPaymentExitgame,
	}
}

// GetStandardExitBond is a method on StandardExitTransaction that gets standard exit bond
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

// Options is a method on StandardExitTransaction that sets root chain transaction options
func (s *StandardExitTransaction) Options(t *bind.TransactOpts) error {
	s.TransactOpts = t
	return nil
}

// Utxo is a method on StandardExitTransaction that sets UTXO data returned from /utxo.get_exit_data to standard exit
func (s *StandardExitTransaction) Utxo(sed StandardExitUTXOData) error {
	s.UtxoData = sed
	return nil
}

// Build is a method on StandardExitTransaction that deserialize transaction txbytes and proofs to golang bytes
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

// Submit is a method on StandardExitTransaction that calls standard exit on the root chain contract
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

// PlasmaAddress sets the plasma contract address to process exit option
func PlasmaAddress(address common.Address) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.PlasmaAddress = address
	}
}

// TokenAddress sets the token address to process exit option
func TokenAddress(address common.Address) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.TokenAddress = address
	}
}

// VaultId sets the vault ID to process exit option
func VaultId(vid string) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.VaultId = vid
	}
}

// TopExit sets the top exit to process to process exit option
func TopExit(te string) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.TopExit = te
	}
}

// NumberOfExits sets the number of exits to process to process exit option
func NumberOfExits(no string) ProcessExitOption {
	return func(p *ProcessExitOptions) {
		p.NumberOfExits = no
	}
}

// NewProcessExit is a method on root chain client that returns a new instance of process exit transaction
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

// Option is a method on ProcessExitTransaction set root chain transaction options for process exit transaction
func(p *ProcessExitTransaction) Options(t *bind.TransactOpts) error {
	p.TransactOpts = t
	return nil
}


// Option is a method on ProcessExitTransaction that builds process exit transaction
func (p *ProcessExitTransaction) Build() error {
	// TODO validation
	return nil
}

// Submit makes a root chain process exit call on the contract with given arguement
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

// Options is a wrapper function that takes any transaction that satisfies RootchainTransaction interface and dispatch Opt on it
func Options(rtx  RootchainTransaction, t *bind.TransactOpts) error {
	return rtx.Options(t)
}

// Submit is a wrapper function that takes any transaction that satisfies RootchainTransaction interface and dispatch Submit on it
func Submit(rtx RootchainTransaction) ( *types.Transaction, error ) {
	return rtx.Submit()
}

// Build is a wrapper function that takes any transaction that satisfies RootchainTransaction interface and dispatch Build on it
func Build(rtx RootchainTransaction) error {
	return rtx.Build()
}







