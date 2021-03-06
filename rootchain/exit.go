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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pongch/omgo/abi"
	"github.com/pongch/omgo/util"
	log "github.com/sirupsen/logrus"
)

// ExitBinder binds the go code to a payment exit game contract on the root chain
type ExitBinder func(common.Address, bind.ContractBackend) (*abi.PaymentExitgame, error)

// StandardExitTransaction is a collection of functions and data needed to make a payment exit transaction on the root chain
type StandardExitTransaction struct {
	*bind.TransactOpts
	bind.ContractBackend
	ExitGameAddress common.Address
	ExitBinder      ExitBinder
	Proof           []byte
	TxBytes         []byte
	UtxoPos         *big.Int
	UtxoData        StandardExitUTXOData
}

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

// NewStandardExit is a method on the root chain client that creates new standard payment exit
func (c *Client) NewStandardExit(vaultAddress common.Address) *StandardExitTransaction {
	return &StandardExitTransaction{
		ContractBackend: c.ContractBackend,
		ExitGameAddress: vaultAddress,
		ExitBinder:      abi.NewPaymentExitgame,
	}
}

// GetStandardExitBond is a method on StandardExitTransaction that gets standard exit bond
func (s *StandardExitTransaction) GetStandardExitBond(opts *bind.CallOpts) (*big.Int, error) {
	instance, err := s.ExitBinder(s.ExitGameAddress, s.ContractBackend)
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
		UtxoPos:                s.UtxoPos,
		RlpOutputTx:            s.TxBytes,
		OutputTxInclusionProof: s.Proof,
	}

	res, err := instance.StartStandardExit(s.TransactOpts, seargs)
	if err != nil {
		return nil, err
	}
	return res, nil
}
