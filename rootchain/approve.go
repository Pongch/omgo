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
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pongch/omgo/abi"
)

// ERC20Transaction is a collection of functions and data needed to approve ERC20 tokens to be spendable
type ApproveTransaction struct {
	*bind.TransactOpts
	bind.ContractBackend
	Erc20Binder Erc20Binder
	Erc20Address common.Address
	PlasmaVault common.Address
	Amount *big.Int
}

type Erc20Binder func(common.Address, bind.ContractBackend) (*abi.Erc20, error)

func (c *Client) NewApprove(vaultAddress, erc20Address common.Address, amount *big.Int) *ApproveTransaction {
	return &ApproveTransaction{
		ContractBackend: c.ContractBackend,
		Erc20Address: erc20Address,
		Erc20Binder: abi.NewErc20,
		PlasmaVault: vaultAddress,
		Amount: amount,
	}
}

// ERC20 token approve transaction
func (a *ApproveTransaction) Options(t *bind.TransactOpts) error {
	a.TransactOpts = t
	return nil
}

func (a *ApproveTransaction) Build() error {
	//TODO validation
	return nil
}

func (a *ApproveTransaction) Submit() (*types.Transaction, error) {
	instance, err := a.Erc20Binder(a.Erc20Address, a.ContractBackend)
	if err != nil {
		return nil, err
	}
	tx, err := instance.Approve(
		a.TransactOpts,
		a.PlasmaVault,
		a.Amount,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

