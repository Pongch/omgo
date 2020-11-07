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
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pongch/omgo/abi"
)

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
