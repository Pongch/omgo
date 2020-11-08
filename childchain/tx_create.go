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
	"encoding/json"
	"math/big"
)

// CreateTransactionResponse is the response to the transaction.create
type CreateTransactionResponse struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    data   `json:"data"`
}

type data struct {
	Result       string         `json:"result"`
	Transactions []Transactions `json:"transactions"`
	Object       string         `json:"object"`
	Code         string         `json:"code"`
	Description  string         `json:"description"`
}

//Transactions to be signed by the client
type Transactions struct {
	Inputs    []Inputs  `json:"inputs"`
	Outputs   []Outputs `json:"outputs"`
	Fee       Fee       `json:"fee"`
	Metadata  string    `json:"metadata"`
	Txbytes   string    `json:"txbytes"`
	SignHash  string    `json:"sign_hash"`
	TypedData TypedData `json:"typed_data"`
}

// Inputs to the transaction
type Inputs struct {
	Blknum   json.Number `json:"blknum"`
	Txindex  json.Number `json:"txindex"`
	Oindex   json.Number `json:"oindex"`
	UtxoPos  *big.Int    `json:"utxo_pos"`
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}

// Outputs to the transaction
type Outputs struct {
	Amount   json.Number `json:"amount"`
	Currency string      `json:"currency"`
	Owner    string      `json:"owner"`
}

// TypedData of the transaction to be signed
type TypedData struct {
	Types       types   `json:"types"`
	PrimaryType string  `json:"primaryType"`
	Domain      domain  `json:"domain"`
	Message     message `json:"message"`
}

type types struct {
	EIP712Domain []eip712Domain `json:"EIP712Domain"`
	Transaction  []transaction  `json:"Transaction"`
	Input        []input        `json:"Input"`
	Output       []output       `json:"Output"`
}

type domain struct {
	Name              string `json:"name"`
	Salt              string `json:"salt"`
	VerifyingContract string `json:"verifyingContract"`
	Version           string `json:"version"`
}

type message struct {
	Input0   input0  `json:"input0"`
	Input1   input1  `json:"input1"`
	Input2   input2  `json:"input2"`
	Input3   input3  `json:"input3"`
	Output0  output0 `json:"output0"`
	Output1  output1 `json:"output1"`
	Output2  output2 `json:"output2"`
	Output3  output3 `json:"output3"`
	Metadata string  `json:"metadata"`
}

type eip712Domain struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type transaction struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type input struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type output struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type input0 struct {
	Blknum  json.Number `json:"blknum"`
	Txindex json.Number `json:"txindex"`
	Oindex  json.Number `json:"oindex"`
}
type input1 struct {
	Blknum  json.Number `json:"blknum"`
	Txindex json.Number `json:"txindex"`
	Oindex  json.Number `json:"oindex"`
}
type input2 struct {
	Blknum  json.Number `json:"blknum"`
	Txindex json.Number `json:"txindex"`
	Oindex  json.Number `json:"oindex"`
}
type input3 struct {
	Blknum  json.Number `json:"blknum"`
	Txindex json.Number `json:"txindex"`
	Oindex  json.Number `json:"oindex"`
}
type output0 struct {
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}
type output1 struct {
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}
type output2 struct {
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}
type output3 struct {
	Owner    string      `json:"owner"`
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}
