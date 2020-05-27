// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Erc20vaultABI is the input ABI used to generate the binding from.
const Erc20vaultABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"setDepositVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEffectiveDepositVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newDepositVerifierMaturityTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPlasmaFramework\",\"name\":\"_framework\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Erc20Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blknum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nextDepositVerifier\",\"type\":\"address\"}],\"name\":\"SetDepositVerifierCalled\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"depositTx\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Erc20vault is an auto generated Go binding around an Ethereum contract.
type Erc20vault struct {
	Erc20vaultCaller     // Read-only binding to the contract
	Erc20vaultTransactor // Write-only binding to the contract
}

// Erc20vaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20vaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20vaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20vaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20vaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20vaultSession struct {
	Contract     *Erc20vault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20vaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20vaultCallerSession struct {
	Contract *Erc20vaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Erc20vaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20vaultTransactorSession struct {
	Contract     *Erc20vaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Erc20vaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20vaultRaw struct {
	Contract *Erc20vault // Generic contract binding to access the raw methods on
}

// Erc20vaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20vaultCallerRaw struct {
	Contract *Erc20vaultCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20vaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20vaultTransactorRaw struct {
	Contract *Erc20vaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20vault creates a new instance of Erc20vault, bound to a specific deployed contract.
func NewErc20vault(address common.Address, backend bind.ContractBackend) (*Erc20vault, error) {
	contract, err := bindErc20vault(address, backend, backend, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20vault{Erc20vaultCaller: Erc20vaultCaller{contract: contract}, Erc20vaultTransactor: Erc20vaultTransactor{contract: contract}}, nil
}

// NewErc20vaultCaller creates a new read-only instance of Erc20vault, bound to a specific deployed contract.
func NewErc20vaultCaller(address common.Address, caller bind.ContractCaller) (*Erc20vaultCaller, error) {
	contract, err := bindErc20vault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20vaultCaller{contract: contract}, nil
}

// NewErc20vaultTransactor creates a new write-only instance of Erc20vault, bound to a specific deployed contract.
func NewErc20vaultTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20vaultTransactor, error) {
	contract, err := bindErc20vault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20vaultTransactor{contract: contract}, nil
}

// bindErc20vault binds a generic wrapper to an already deployed contract.
func bindErc20vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20vaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20vault *Erc20vaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc20vault.Contract.Erc20vaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20vault *Erc20vaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20vault.Contract.Erc20vaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20vault *Erc20vaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20vault.Contract.Erc20vaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20vault *Erc20vaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc20vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20vault *Erc20vaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20vault *Erc20vaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20vault.Contract.contract.Transact(opts, method, params...)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_Erc20vault *Erc20vaultCaller) DepositVerifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Erc20vault.contract.Call(opts, out, "depositVerifiers", arg0)
	return *ret0, err
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_Erc20vault *Erc20vaultSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _Erc20vault.Contract.DepositVerifiers(&_Erc20vault.CallOpts, arg0)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_Erc20vault *Erc20vaultCallerSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _Erc20vault.Contract.DepositVerifiers(&_Erc20vault.CallOpts, arg0)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Erc20vault *Erc20vaultCaller) GetEffectiveDepositVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Erc20vault.contract.Call(opts, out, "getEffectiveDepositVerifier")
	return *ret0, err
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Erc20vault *Erc20vaultSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _Erc20vault.Contract.GetEffectiveDepositVerifier(&_Erc20vault.CallOpts)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Erc20vault *Erc20vaultCallerSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _Erc20vault.Contract.GetEffectiveDepositVerifier(&_Erc20vault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Erc20vault *Erc20vaultCaller) NewDepositVerifierMaturityTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc20vault.contract.Call(opts, out, "newDepositVerifierMaturityTimestamp")
	return *ret0, err
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Erc20vault *Erc20vaultSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _Erc20vault.Contract.NewDepositVerifierMaturityTimestamp(&_Erc20vault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Erc20vault *Erc20vaultCallerSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _Erc20vault.Contract.NewDepositVerifierMaturityTimestamp(&_Erc20vault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(depositTx bytes) returns()
func (_Erc20vault *Erc20vaultTransactor) Deposit(opts *bind.TransactOpts, depositTx []byte) (*types.Transaction, error) {
	return _Erc20vault.contract.Transact(opts, "deposit", depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(depositTx bytes) returns()
func (_Erc20vault *Erc20vaultSession) Deposit(depositTx []byte) (*types.Transaction, error) {
	return _Erc20vault.Contract.Deposit(&_Erc20vault.TransactOpts, depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(depositTx bytes) returns()
func (_Erc20vault *Erc20vaultTransactorSession) Deposit(depositTx []byte) (*types.Transaction, error) {
	return _Erc20vault.Contract.Deposit(&_Erc20vault.TransactOpts, depositTx)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_Erc20vault *Erc20vaultTransactor) SetDepositVerifier(opts *bind.TransactOpts, _verifier common.Address) (*types.Transaction, error) {
	return _Erc20vault.contract.Transact(opts, "setDepositVerifier", _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_Erc20vault *Erc20vaultSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Erc20vault.Contract.SetDepositVerifier(&_Erc20vault.TransactOpts, _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_Erc20vault *Erc20vaultTransactorSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Erc20vault.Contract.SetDepositVerifier(&_Erc20vault.TransactOpts, _verifier)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(receiver address, token address, amount uint256) returns()
func (_Erc20vault *Erc20vaultTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20vault.contract.Transact(opts, "withdraw", receiver, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(receiver address, token address, amount uint256) returns()
func (_Erc20vault *Erc20vaultSession) Withdraw(receiver common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20vault.Contract.Withdraw(&_Erc20vault.TransactOpts, receiver, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(receiver address, token address, amount uint256) returns()
func (_Erc20vault *Erc20vaultTransactorSession) Withdraw(receiver common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20vault.Contract.Withdraw(&_Erc20vault.TransactOpts, receiver, token, amount)
}
