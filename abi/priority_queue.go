// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PriorityQueueABI is the input ABI used to generate the binding from.
const PriorityQueueABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"heapList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriorityQueue is an auto generated Go binding around an Ethereum contract.
type PriorityQueue struct {
	PriorityQueueCaller     // Read-only binding to the contract
	PriorityQueueTransactor // Write-only binding to the contract
	PriorityQueueFilterer   // Log filterer for contract events
}

// PriorityQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriorityQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriorityQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriorityQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriorityQueueSession struct {
	Contract     *PriorityQueue    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriorityQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriorityQueueCallerSession struct {
	Contract *PriorityQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PriorityQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriorityQueueTransactorSession struct {
	Contract     *PriorityQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PriorityQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriorityQueueRaw struct {
	Contract *PriorityQueue // Generic contract binding to access the raw methods on
}

// PriorityQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriorityQueueCallerRaw struct {
	Contract *PriorityQueueCaller // Generic read-only contract binding to access the raw methods on
}

// PriorityQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriorityQueueTransactorRaw struct {
	Contract *PriorityQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriorityQueue creates a new instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueue(address common.Address, backend bind.ContractBackend) (*PriorityQueue, error) {
	contract, err := bindPriorityQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}, PriorityQueueFilterer: PriorityQueueFilterer{contract: contract}}, nil
}

// NewPriorityQueueCaller creates a new read-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueCaller(address common.Address, caller bind.ContractCaller) (*PriorityQueueCaller, error) {
	contract, err := bindPriorityQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueCaller{contract: contract}, nil
}

// NewPriorityQueueTransactor creates a new write-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*PriorityQueueTransactor, error) {
	contract, err := bindPriorityQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueTransactor{contract: contract}, nil
}

// NewPriorityQueueFilterer creates a new log filterer instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*PriorityQueueFilterer, error) {
	contract, err := bindPriorityQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueFilterer{contract: contract}, nil
}

// bindPriorityQueue binds a generic wrapper to an already deployed contract.
func bindPriorityQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriorityQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriorityQueue *PriorityQueueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriorityQueue.Contract.PriorityQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriorityQueue *PriorityQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.Contract.PriorityQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriorityQueue *PriorityQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriorityQueue.Contract.PriorityQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriorityQueue *PriorityQueueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriorityQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriorityQueue *PriorityQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriorityQueue *PriorityQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriorityQueue.Contract.contract.Transact(opts, method, params...)
}

// HeapList is a free data retrieval call binding the contract method 0x446b2f88.
//
// Solidity: function heapList() constant returns(uint256[])
func (_PriorityQueue *PriorityQueueCaller) HeapList(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "heapList")
	return *ret0, err
}

// HeapList is a free data retrieval call binding the contract method 0x446b2f88.
//
// Solidity: function heapList() constant returns(uint256[])
func (_PriorityQueue *PriorityQueueSession) HeapList() ([]*big.Int, error) {
	return _PriorityQueue.Contract.HeapList(&_PriorityQueue.CallOpts)
}

// HeapList is a free data retrieval call binding the contract method 0x446b2f88.
//
// Solidity: function heapList() constant returns(uint256[])
func (_PriorityQueue *PriorityQueueCallerSession) HeapList() ([]*big.Int, error) {
	return _PriorityQueue.Contract.HeapList(&_PriorityQueue.CallOpts)
}
