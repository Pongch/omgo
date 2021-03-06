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

package childchain

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pongch/omgo/util"
)

// PaymentTx is a general P2P payment transaction
// implemented using Watcher's set of convenience API
// it can be used to transfer funds from one to many
type PaymentTx struct {
	Client                    *Client
	CreateTransaction         *CreateTransaction
	CreateTransactionResponse *CreateTransactionResponse
	Signatures                [][]byte
	TypedTransaction          TypedTransaction
}

// CreateTransaction is a payload to be submitted
// via transaction.create endpoint
type CreateTransaction struct {
	Owner    common.Address `json:"owner"`
	Payments []Payment      `json:"payments"`
	Fee      Fee            `json:"fee"`
	Metadata string         `json:"metadata"`
}

// Payment specifies the amount,
// currency and recipient of the output
type Payment struct {
	Amount   *big.Int       `json:"amount"`
	Currency common.Address `json:"currency"`
	Owner    common.Address `json:"owner"`
}

// TypedTransaction is a transaction
// data to be submitted to
// submit_typed endpoint
type TypedTransaction struct {
	Domain     domain   `json:"domain"`
	Message    message  `json:"message"`
	Signatures []string `json:"signatures"`
}

type PaymentOption func(*CreateTransaction)

func (c *Client) NewPaymentTx(popts ...PaymentOption) (p *PaymentTx) {
	p = &PaymentTx{Client: c}
	ctx := &CreateTransaction{}
	for _, setOption := range popts {
		setOption(ctx)
	}
	p.CreateTransaction = ctx
	return p
}

func AddOwner(o common.Address) PaymentOption {
	return func(c *CreateTransaction) {
		c.Owner = o
	}
}

func AddFee(curr common.Address) PaymentOption {
	return func(c *CreateTransaction) {
		c.Fee = Fee{Currency: curr}
	}
}

func AddPayment(amount string, addr, curr common.Address) PaymentOption {
	return func(c *CreateTransaction) {
		bamount, _ := new(big.Int).SetString(amount, 0)
		payment := Payment{
			Amount:   bamount,
			Currency: curr,
			Owner:    addr,
		}
		c.Payments = append(c.Payments, payment)
	}
}

func AddMetadata(m string) PaymentOption {
	return func(c *CreateTransaction) {
		c.Metadata = m
	}
}

// validate payment transaction structs
func validate(p *PaymentTx) error {
	if len(p.CreateTransaction.Metadata) != 66 {
		return fmt.Errorf("invalid length metadata, got %v, wanted %v", len(p.CreateTransaction.Metadata), 66)
	}
	if len(p.CreateTransaction.Payments) >= MaxOutputs {
		return fmt.Errorf("error too many payments, max outputs are %v, got %v", MaxOutputs, len(p.CreateTransaction.Payments))
	}
	if len(p.CreateTransaction.Payments) == 0 {
		return errors.New("no payment specified")
	}

	return nil
}

// BuildTransaction forms a transaction to be signed via transaction.create endpoint
// NOTE: if response.Data.Result == "intermediate"
// means payment cannot be completed in one transaction, this tx will automatically
// performs a merge instead
// TODO: clients should be able to respond to intermediate transactions
func (p *PaymentTx) BuildTransaction() error {
	if err := validate(p); err != nil {
		return err
	}
	if len(p.CreateTransaction.Payments) == 0 {
		return errors.New("not enough argument to build transaction")
	}
	rstring, err := p.Client.do(
		"/transaction.create",
		p.CreateTransaction,
	)
	response := CreateTransactionResponse{}
	if err != nil {
		return err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return jsonErr
	}

	if response.Success == false {
		return fmt.Errorf(
			"Error creating transaction. \n Code: %v \n Description: %v \n Object: %v \n",
			response.Data.Code,
			response.Data.Description,
			response.Data.Object,
		)
	}

	p.CreateTransactionResponse = &response
	return nil
}

// SignTransaction takes in any Signer function
// and uses it on the SignHash
// a custom signing function can be used as input
func (p *PaymentTx) SignTransaction(signer SignerFunc) ([][]byte, error) {
	signHashBytes, err := hex.DecodeString(util.FilterZeroX(p.CreateTransactionResponse.Data.Transactions[0].SignHash))
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %v", err)
	}
	sigs, err := signer(signHashBytes)
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %v", err)
	}
	p.Signatures = sigs
	return sigs, nil
}

// buildTypedTransaction is a private function to
// form the TypedTransaction to be submitted
//  on /transaction.submit_typed endpoint
// TODO currently accepts a single signer & tx, should work with arbitrary number of inputs and signatures
func (p *PaymentTx) buildTypedTransaction() error {
	for _, _ = range p.CreateTransactionResponse.Data.Transactions[0].Inputs {
		p.TypedTransaction.Signatures = append(p.TypedTransaction.Signatures, fmt.Sprintf("0x%v", hex.EncodeToString(p.Signatures[0])))
	}
	tx := p.CreateTransactionResponse.Data.Transactions[0].TypedData
	p.TypedTransaction.Domain = tx.Domain
	p.TypedTransaction.Message = tx.Message
	return nil
}

// SubmitTransaction submit payment transaction through
// /transaction.submit_typed endpoint
func (p *PaymentTx) SubmitTransaction() (*TransactionSubmitResponse, error) {
	if err := p.buildTypedTransaction(); err != nil {
		return nil, errors.New("error building typed transaction before submitting")
	}
	rstring, err := p.Client.do(
		"/transaction.submit_typed",
		p.TypedTransaction,
	)
	response := TransactionSubmitResponse{}
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if response.Success == false {
		return nil, fmt.Errorf(
			"Error submitting transaction. \n Code: %v \n Description: %v",
			response.Data.Code,
			response.Data.Description,
		)
	}

	return &response, nil
}
