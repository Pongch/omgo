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

// EthvaultABI is the input ABI used to generate the binding from.
const EthvaultABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"setDepositVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEffectiveDepositVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newDepositVerifierMaturityTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"safeGasStipend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPlasmaFramework\",\"name\":\"_framework\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_safeGasStipend\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blknum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nextDepositVerifier\",\"type\":\"address\"}],\"name\":\"SetDepositVerifierCalled\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_depositTx\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ethvault is an auto generated Go binding around an Ethereum contract.
type Ethvault struct {
	EthvaultCaller     // Read-only binding to the contract
	EthvaultTransactor // Write-only binding to the contract
}

// EthvaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthvaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthvaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthvaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthvaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthvaultSession struct {
	Contract     *Ethvault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthvaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthvaultCallerSession struct {
	Contract *EthvaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthvaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthvaultTransactorSession struct {
	Contract     *EthvaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthvaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthvaultRaw struct {
	Contract *Ethvault // Generic contract binding to access the raw methods on
}

// EthvaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthvaultCallerRaw struct {
	Contract *EthvaultCaller // Generic read-only contract binding to access the raw methods on
}

// EthvaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthvaultTransactorRaw struct {
	Contract *EthvaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthvault creates a new instance of Ethvault, bound to a specific deployed contract.
func NewEthvault(address common.Address, backend bind.ContractBackend) (*Ethvault, error) {
	contract, err := bindEthvault(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethvault{EthvaultCaller: EthvaultCaller{contract: contract}, EthvaultTransactor: EthvaultTransactor{contract: contract}}, nil
}

// NewEthvaultCaller creates a new read-only instance of Ethvault, bound to a specific deployed contract.
func NewEthvaultCaller(address common.Address, caller bind.ContractCaller) (*EthvaultCaller, error) {
	contract, err := bindEthvault(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EthvaultCaller{contract: contract}, nil
}

// NewEthvaultTransactor creates a new write-only instance of Ethvault, bound to a specific deployed contract.
func NewEthvaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EthvaultTransactor, error) {
	contract, err := bindEthvault(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EthvaultTransactor{contract: contract}, nil
}

// bindEthvault binds a generic wrapper to an already deployed contract.
func bindEthvault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, ) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthvaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethvault *EthvaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethvault.Contract.EthvaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethvault *EthvaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethvault.Contract.EthvaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethvault *EthvaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethvault.Contract.EthvaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethvault *EthvaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethvault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethvault *EthvaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethvault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethvault *EthvaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethvault.Contract.contract.Transact(opts, method, params...)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_Ethvault *EthvaultCaller) DepositVerifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethvault.contract.Call(opts, out, "depositVerifiers", arg0)
	return *ret0, err
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_Ethvault *EthvaultSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _Ethvault.Contract.DepositVerifiers(&_Ethvault.CallOpts, arg0)
}

// DepositVerifiers is a free data retrieval call binding the contract method 0x02e9b07b.
//
// Solidity: function depositVerifiers( uint256) constant returns(address)
func (_Ethvault *EthvaultCallerSession) DepositVerifiers(arg0 *big.Int) (common.Address, error) {
	return _Ethvault.Contract.DepositVerifiers(&_Ethvault.CallOpts, arg0)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Ethvault *EthvaultCaller) GetEffectiveDepositVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethvault.contract.Call(opts, out, "getEffectiveDepositVerifier")
	return *ret0, err
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Ethvault *EthvaultSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _Ethvault.Contract.GetEffectiveDepositVerifier(&_Ethvault.CallOpts)
}

// GetEffectiveDepositVerifier is a free data retrieval call binding the contract method 0x52518b7f.
//
// Solidity: function getEffectiveDepositVerifier() constant returns(address)
func (_Ethvault *EthvaultCallerSession) GetEffectiveDepositVerifier() (common.Address, error) {
	return _Ethvault.Contract.GetEffectiveDepositVerifier(&_Ethvault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Ethvault *EthvaultCaller) NewDepositVerifierMaturityTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethvault.contract.Call(opts, out, "newDepositVerifierMaturityTimestamp")
	return *ret0, err
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Ethvault *EthvaultSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _Ethvault.Contract.NewDepositVerifierMaturityTimestamp(&_Ethvault.CallOpts)
}

// NewDepositVerifierMaturityTimestamp is a free data retrieval call binding the contract method 0xa7e5994f.
//
// Solidity: function newDepositVerifierMaturityTimestamp() constant returns(uint256)
func (_Ethvault *EthvaultCallerSession) NewDepositVerifierMaturityTimestamp() (*big.Int, error) {
	return _Ethvault.Contract.NewDepositVerifierMaturityTimestamp(&_Ethvault.CallOpts)
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_Ethvault *EthvaultCaller) SafeGasStipend(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethvault.contract.Call(opts, out, "safeGasStipend")
	return *ret0, err
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_Ethvault *EthvaultSession) SafeGasStipend() (*big.Int, error) {
	return _Ethvault.Contract.SafeGasStipend(&_Ethvault.CallOpts)
}

// SafeGasStipend is a free data retrieval call binding the contract method 0xccbd2176.
//
// Solidity: function safeGasStipend() constant returns(uint256)
func (_Ethvault *EthvaultCallerSession) SafeGasStipend() (*big.Int, error) {
	return _Ethvault.Contract.SafeGasStipend(&_Ethvault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_Ethvault *EthvaultTransactor) Deposit(opts *bind.TransactOpts, _depositTx []byte) (*types.Transaction, error) {
	return _Ethvault.contract.Transact(opts, "deposit", _depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_Ethvault *EthvaultSession) Deposit(_depositTx []byte) (*types.Transaction, error) {
	return _Ethvault.Contract.Deposit(&_Ethvault.TransactOpts, _depositTx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(_depositTx bytes) returns()
func (_Ethvault *EthvaultTransactorSession) Deposit(_depositTx []byte) (*types.Transaction, error) {
	return _Ethvault.Contract.Deposit(&_Ethvault.TransactOpts, _depositTx)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_Ethvault *EthvaultTransactor) SetDepositVerifier(opts *bind.TransactOpts, _verifier common.Address) (*types.Transaction, error) {
	return _Ethvault.contract.Transact(opts, "setDepositVerifier", _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_Ethvault *EthvaultSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Ethvault.Contract.SetDepositVerifier(&_Ethvault.TransactOpts, _verifier)
}

// SetDepositVerifier is a paid mutator transaction binding the contract method 0x2c6ce78b.
//
// Solidity: function setDepositVerifier(_verifier address) returns()
func (_Ethvault *EthvaultTransactorSession) SetDepositVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Ethvault.Contract.SetDepositVerifier(&_Ethvault.TransactOpts, _verifier)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns()
func (_Ethvault *EthvaultTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethvault.contract.Transact(opts, "withdraw", receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns()
func (_Ethvault *EthvaultSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethvault.Contract.Withdraw(&_Ethvault.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns()
func (_Ethvault *EthvaultTransactorSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethvault.Contract.Withdraw(&_Ethvault.TransactOpts, receiver, amount)
}
