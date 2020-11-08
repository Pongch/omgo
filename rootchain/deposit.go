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
	"github.com/pongch/omgo/util"
	"strconv"
)

// DepositBuilder is a function that builds an RLP input from a set of inputs
// TODO clean up transaction builder signature to return a curried function instead
// form a deposit transaction
type DepositBuilder func(common.Address, common.Address, uint64, uint64) ([]byte, error)

// VaultBinder binds the go code to an Ethvault contract on the root chain
// TODO implement vault binder for ERC20
type VaultBinder func(common.Address, bind.ContractBackend) (*abi.Ethvault, error)

// Erc20VaultBinder binds the go code to an Erc20Vault contract on the root chain
type Erc20VaultBinder func(common.Address, bind.ContractBackend) (*abi.Erc20vault, error)

// DepositTransaction is a collection of functions and data needed to make a deposit transaction on the root chain
type DepositTransaction struct {
	*bind.TransactOpts
	bind.ContractBackend
	VaultAddress     common.Address
	Amount           string
	Owner            common.Address
	Currency         common.Address
	VaultBinder      VaultBinder
	Erc20VaultBinder Erc20VaultBinder
	Builder          DepositBuilder
	RlpTransaction   []byte
}

// NewDeposit is a method on root chain client that create a new deposit
func (c *Client) NewDeposit(vaultAddress, owner, currency common.Address, amount string) *DepositTransaction {
	return &DepositTransaction{
		ContractBackend:  c.ContractBackend,
		VaultAddress:     vaultAddress,
		VaultBinder:      abi.NewEthvault,
		Erc20VaultBinder: abi.NewErc20vault,
		Amount:           amount,
		Owner:            owner,
		Builder:          util.BuildRLPDeposit,
		Currency:         currency,
	}
}

// Options is a method on DepositTransaction that set root chain transaction options
func (d *DepositTransaction) Options(t *bind.TransactOpts) error {
	//TODO: add validations
	d.TransactOpts = t
	return nil
}

// Build is a method on DepositTransaction that encodes the transaction to RLP with the RLP Builder
func (d *DepositTransaction) Build() error {
	amount, err := strconv.ParseUint(d.Amount, 10, 64)
	if err != nil {
		return err
	}
	//must supports erc20 vault if currency is not ETH
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
	var res *types.Transaction
	if d.Currency == common.HexToAddress(util.EthCurrency) {
		instance, err := d.VaultBinder(d.VaultAddress, d.ContractBackend)
		if err != nil {
			return nil, err
		}
		res, err = instance.Deposit(d.TransactOpts, d.RlpTransaction)
		if err != nil {
			return nil, err
		}
	} else {
		instance, err := d.Erc20VaultBinder(d.VaultAddress, d.ContractBackend)
		if err != nil {
			return nil, err
		}
		res, err = instance.Deposit(d.TransactOpts, d.RlpTransaction)
		if err != nil {
			return nil, err
		}

	}
	return res, nil
}
