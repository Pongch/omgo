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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pongch/omgo/abi"
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

// PlasmaFrameworkBinder binds the go code to a plasma framework contract on the root chain
// TODO: make binder return struct a more generic reusable contract
type PlasmaFrameworkBinder func(common.Address, bind.ContractBackend) (*abi.PlasmaFramework, error)


// NewClient initializes new root chain client
func NewClient(ethclient bind.ContractBackend) *Client{
	return &Client{ethclient}
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







