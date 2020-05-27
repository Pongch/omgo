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

// PaymentExitgameABI is the input ABI used to generate the binding from.
const PaymentExitgameABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint160[]\",\"name\":\"exitIds\",\"type\":\"uint160[]\"}],\"name\":\"standardExits\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exitable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"outputId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondSize\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitDataModel.StandardExit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBondSize\",\"type\":\"uint128\"}],\"name\":\"updateStartStandardExitBondSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"inFlightTxInclusionProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"outputUtxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"challengingTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"challengingTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"challengingTxWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"senderData\",\"type\":\"bytes32\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.ChallengeOutputSpent\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"challengeInFlightExitOutputSpent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBondSize\",\"type\":\"uint128\"}],\"name\":\"updatePiggybackBondSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inFlightTxPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"inFlightTxInclusionProof\",\"type\":\"bytes\"}],\"name\":\"respondToNonCanonicalChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"piggybackBondSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputTxs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"inputUtxosPos\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"inputTxsInclusionProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"inFlightTxWitnesses\",\"type\":\"bytes[]\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.StartExitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"startInFlightExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpOutputTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"outputTxInclusionProof\",\"type\":\"bytes\"}],\"internalType\":\"structPaymentStandardExitRouterArgs.StartStandardExitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"startStandardExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startIFEBondSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_PB_BOND_SIZE\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOND_LOWER_BOUND_DIVISOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOND_UPPER_BOUND_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"outputIndex\",\"type\":\"uint16\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.PiggybackInFlightExitOnOutputArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"piggybackInFlightExitOnOutput\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBondSize\",\"type\":\"uint128\"}],\"name\":\"updateStartIFEBondSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"deleteNonPiggybackedInFlightExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"inFlightTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"challengingTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"challengingTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"challengingTxWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"inputTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inputUtxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderData\",\"type\":\"bytes32\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.ChallengeInputSpentArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"challengeInFlightExitInputSpent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_BOND_SIZE\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_IFE_BOND_SIZE\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint160[]\",\"name\":\"exitIds\",\"type\":\"uint160[]\"}],\"name\":\"inFlightExits\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isCanonical\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"exitStartTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"exitMap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"piggybackBondSize\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitDataModel.WithdrawData[4]\",\"name\":\"inputs\",\"type\":\"tuple[4]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"piggybackBondSize\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitDataModel.WithdrawData[4]\",\"name\":\"outputs\",\"type\":\"tuple[4]\"},{\"internalType\":\"addresspayable\",\"name\":\"bondOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bondSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldestCompetitorPosition\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitDataModel.InFlightExit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"exitingTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"witness\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"senderData\",\"type\":\"bytes32\"}],\"internalType\":\"structPaymentStandardExitRouterArgs.ChallengeStandardExitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"challengeStandardExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.PiggybackInFlightExitOnInputArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"piggybackInFlightExitOnInput\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"inputTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inputUtxoPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"inFlightTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"inFlightTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"competingTx\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"competingTxInputIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"competingTxPos\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"competingTxInclusionProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"competingTxWitness\",\"type\":\"bytes\"}],\"internalType\":\"structPaymentInFlightExitRouterArgs.ChallengeCanonicityArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"challengeInFlightExitNotCanonical\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startStandardExitBondSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractPlasmaFramework\",\"name\":\"framework\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ethVaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"erc20VaultId\",\"type\":\"uint256\"},{\"internalType\":\"contractSpendingConditionRegistry\",\"name\":\"spendingConditionRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIStateTransitionVerifier\",\"name\":\"stateTransitionVerifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supportTxType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"safeGasStipend\",\"type\":\"uint256\"}],\"internalType\":\"structPaymentExitGameArgs.Args\",\"name\":\"args\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"bondSize\",\"type\":\"uint128\"}],\"name\":\"IFEBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"bondSize\",\"type\":\"uint128\"}],\"name\":\"PiggybackBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"InFlightExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitInputPiggybacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InFlightExitOmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InFlightBondReturnFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"outputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitOutputWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitInputWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitTarget\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"outputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitOutputPiggybacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeTxPosition\",\"type\":\"uint256\"}],\"name\":\"InFlightExitChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeTxPosition\",\"type\":\"uint256\"}],\"name\":\"InFlightExitChallengeResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"inputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitInputBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"outputIndex\",\"type\":\"uint16\"}],\"name\":\"InFlightExitOutputBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"InFlightExitDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"bondSize\",\"type\":\"uint128\"}],\"name\":\"StandardExitBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"ExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"name\":\"ExitChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"ExitOmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"}],\"name\":\"ExitFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondReturnFailed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"exitId\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"processExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isDeposit\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_txBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_utxoPos\",\"type\":\"uint256\"}],\"name\":\"getStandardExitId\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_txBytes\",\"type\":\"bytes\"}],\"name\":\"getInFlightExitId\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PaymentExitgame is an auto generated Go binding around an Ethereum contract.
type PaymentExitgame struct {
	PaymentExitgameCaller     // Read-only binding to the contract
	PaymentExitgameTransactor // Write-only binding to the contract
	PaymentExitgameFilterer   // Log filterer for contract events
}

// PaymentExitgameCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentExitgameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentExitgameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentExitgameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentExitgameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentExitgameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentExitgameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentExitgameSession struct {
	Contract     *PaymentExitgame  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentExitgameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentExitgameCallerSession struct {
	Contract *PaymentExitgameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PaymentExitgameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentExitgameTransactorSession struct {
	Contract     *PaymentExitgameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PaymentExitgameRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentExitgameRaw struct {
	Contract *PaymentExitgame // Generic contract binding to access the raw methods on
}

// PaymentExitgameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentExitgameCallerRaw struct {
	Contract *PaymentExitgameCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentExitgameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentExitgameTransactorRaw struct {
	Contract *PaymentExitgameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentExitgame creates a new instance of PaymentExitgame, bound to a specific deployed contract.
func NewPaymentExitgame(address common.Address, backend bind.ContractBackend) (*PaymentExitgame, error) {
	contract, err := bindPaymentExitgame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgame{PaymentExitgameCaller: PaymentExitgameCaller{contract: contract}, PaymentExitgameTransactor: PaymentExitgameTransactor{contract: contract}, PaymentExitgameFilterer: PaymentExitgameFilterer{contract: contract}}, nil
}

// NewPaymentExitgameCaller creates a new read-only instance of PaymentExitgame, bound to a specific deployed contract.
func NewPaymentExitgameCaller(address common.Address, caller bind.ContractCaller) (*PaymentExitgameCaller, error) {
	contract, err := bindPaymentExitgame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameCaller{contract: contract}, nil
}

// NewPaymentExitgameTransactor creates a new write-only instance of PaymentExitgame, bound to a specific deployed contract.
func NewPaymentExitgameTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentExitgameTransactor, error) {
	contract, err := bindPaymentExitgame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameTransactor{contract: contract}, nil
}

// NewPaymentExitgameFilterer creates a new log filterer instance of PaymentExitgame, bound to a specific deployed contract.
func NewPaymentExitgameFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentExitgameFilterer, error) {
	contract, err := bindPaymentExitgame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameFilterer{contract: contract}, nil
}

// bindPaymentExitgame binds a generic wrapper to an already deployed contract.
func bindPaymentExitgame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentExitgameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentExitgame *PaymentExitgameRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentExitgame.Contract.PaymentExitgameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentExitgame *PaymentExitgameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.PaymentExitgameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentExitgame *PaymentExitgameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.PaymentExitgameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentExitgame *PaymentExitgameCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentExitgame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentExitgame *PaymentExitgameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentExitgame *PaymentExitgameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.contract.Transact(opts, method, params...)
}

// PaymentExitDataModelInFlightExit is an auto generated low-level Go binding around an user-defined struct.
type PaymentExitDataModelInFlightExit struct {
	IsCanonical              bool
	ExitStartTimestamp       uint64
	ExitMap                  *big.Int
	Position                 *big.Int
	Inputs                   [4]PaymentExitDataModelWithdrawData
	Outputs                  [4]PaymentExitDataModelWithdrawData
	BondOwner                common.Address
	BondSize                 *big.Int
	OldestCompetitorPosition *big.Int
}

// PaymentExitDataModelStandardExit is an auto generated low-level Go binding around an user-defined struct.
type PaymentExitDataModelStandardExit struct {
	Exitable   bool
	UtxoPos    *big.Int
	OutputId   [32]byte
	ExitTarget common.Address
	Amount     *big.Int
	BondSize   *big.Int
}

// PaymentExitDataModelWithdrawData is an auto generated low-level Go binding around an user-defined struct.
type PaymentExitDataModelWithdrawData struct {
	OutputId          [32]byte
	ExitTarget        common.Address
	Token             common.Address
	Amount            *big.Int
	PiggybackBondSize *big.Int
}

// PaymentInFlightExitRouterArgsChallengeCanonicityArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsChallengeCanonicityArgs struct {
	InputTx                   []byte
	InputUtxoPos              *big.Int
	InFlightTx                []byte
	InFlightTxInputIndex      uint16
	CompetingTx               []byte
	CompetingTxInputIndex     uint16
	CompetingTxPos            *big.Int
	CompetingTxInclusionProof []byte
	CompetingTxWitness        []byte
}

// PaymentInFlightExitRouterArgsChallengeInputSpentArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsChallengeInputSpentArgs struct {
	InFlightTx              []byte
	InFlightTxInputIndex    uint16
	ChallengingTx           []byte
	ChallengingTxInputIndex uint16
	ChallengingTxWitness    []byte
	InputTx                 []byte
	InputUtxoPos            *big.Int
	SenderData              [32]byte
}

// PaymentInFlightExitRouterArgsChallengeOutputSpent is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsChallengeOutputSpent struct {
	InFlightTx               []byte
	InFlightTxInclusionProof []byte
	OutputUtxoPos            *big.Int
	ChallengingTx            []byte
	ChallengingTxInputIndex  uint16
	ChallengingTxWitness     []byte
	SenderData               [32]byte
}

// PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs struct {
	InFlightTx []byte
	InputIndex uint16
}

// PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs struct {
	InFlightTx  []byte
	OutputIndex uint16
}

// PaymentInFlightExitRouterArgsStartExitArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentInFlightExitRouterArgsStartExitArgs struct {
	InFlightTx              []byte
	InputTxs                [][]byte
	InputUtxosPos           []*big.Int
	InputTxsInclusionProofs [][]byte
	InFlightTxWitnesses     [][]byte
}

// PaymentStandardExitRouterArgsChallengeStandardExitArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentStandardExitRouterArgsChallengeStandardExitArgs struct {
	ExitId      *big.Int
	ExitingTx   []byte
	ChallengeTx []byte
	InputIndex  uint16
	Witness     []byte
	SenderData  [32]byte
}

// PaymentStandardExitRouterArgsStartStandardExitArgs is an auto generated low-level Go binding around an user-defined struct.
type PaymentStandardExitRouterArgsStartStandardExitArgs struct {
	UtxoPos                *big.Int
	RlpOutputTx            []byte
	OutputTxInclusionProof []byte
}

// BONDLOWERBOUNDDIVISOR is a free data retrieval call binding the contract method 0x905d6a99.
//
// Solidity: function BOND_LOWER_BOUND_DIVISOR() constant returns(uint16)
func (_PaymentExitgame *PaymentExitgameCaller) BONDLOWERBOUNDDIVISOR(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "BOND_LOWER_BOUND_DIVISOR")
	return *ret0, err
}

// BONDLOWERBOUNDDIVISOR is a free data retrieval call binding the contract method 0x905d6a99.
//
// Solidity: function BOND_LOWER_BOUND_DIVISOR() constant returns(uint16)
func (_PaymentExitgame *PaymentExitgameSession) BONDLOWERBOUNDDIVISOR() (uint16, error) {
	return _PaymentExitgame.Contract.BONDLOWERBOUNDDIVISOR(&_PaymentExitgame.CallOpts)
}

// BONDLOWERBOUNDDIVISOR is a free data retrieval call binding the contract method 0x905d6a99.
//
// Solidity: function BOND_LOWER_BOUND_DIVISOR() constant returns(uint16)
func (_PaymentExitgame *PaymentExitgameCallerSession) BONDLOWERBOUNDDIVISOR() (uint16, error) {
	return _PaymentExitgame.Contract.BONDLOWERBOUNDDIVISOR(&_PaymentExitgame.CallOpts)
}

// BONDUPPERBOUNDMULTIPLIER is a free data retrieval call binding the contract method 0xa0e403b1.
//
// Solidity: function BOND_UPPER_BOUND_MULTIPLIER() constant returns(uint16)
func (_PaymentExitgame *PaymentExitgameCaller) BONDUPPERBOUNDMULTIPLIER(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "BOND_UPPER_BOUND_MULTIPLIER")
	return *ret0, err
}

// BONDUPPERBOUNDMULTIPLIER is a free data retrieval call binding the contract method 0xa0e403b1.
//
// Solidity: function BOND_UPPER_BOUND_MULTIPLIER() constant returns(uint16)
func (_PaymentExitgame *PaymentExitgameSession) BONDUPPERBOUNDMULTIPLIER() (uint16, error) {
	return _PaymentExitgame.Contract.BONDUPPERBOUNDMULTIPLIER(&_PaymentExitgame.CallOpts)
}

// BONDUPPERBOUNDMULTIPLIER is a free data retrieval call binding the contract method 0xa0e403b1.
//
// Solidity: function BOND_UPPER_BOUND_MULTIPLIER() constant returns(uint16)
func (_PaymentExitgame *PaymentExitgameCallerSession) BONDUPPERBOUNDMULTIPLIER() (uint16, error) {
	return _PaymentExitgame.Contract.BONDUPPERBOUNDMULTIPLIER(&_PaymentExitgame.CallOpts)
}

// INITIALBONDSIZE is a free data retrieval call binding the contract method 0xc170ecf5.
//
// Solidity: function INITIAL_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCaller) INITIALBONDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "INITIAL_BOND_SIZE")
	return *ret0, err
}

// INITIALBONDSIZE is a free data retrieval call binding the contract method 0xc170ecf5.
//
// Solidity: function INITIAL_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameSession) INITIALBONDSIZE() (*big.Int, error) {
	return _PaymentExitgame.Contract.INITIALBONDSIZE(&_PaymentExitgame.CallOpts)
}

// INITIALBONDSIZE is a free data retrieval call binding the contract method 0xc170ecf5.
//
// Solidity: function INITIAL_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCallerSession) INITIALBONDSIZE() (*big.Int, error) {
	return _PaymentExitgame.Contract.INITIALBONDSIZE(&_PaymentExitgame.CallOpts)
}

// INITIALIFEBONDSIZE is a free data retrieval call binding the contract method 0xc96e7c05.
//
// Solidity: function INITIAL_IFE_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCaller) INITIALIFEBONDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "INITIAL_IFE_BOND_SIZE")
	return *ret0, err
}

// INITIALIFEBONDSIZE is a free data retrieval call binding the contract method 0xc96e7c05.
//
// Solidity: function INITIAL_IFE_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameSession) INITIALIFEBONDSIZE() (*big.Int, error) {
	return _PaymentExitgame.Contract.INITIALIFEBONDSIZE(&_PaymentExitgame.CallOpts)
}

// INITIALIFEBONDSIZE is a free data retrieval call binding the contract method 0xc96e7c05.
//
// Solidity: function INITIAL_IFE_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCallerSession) INITIALIFEBONDSIZE() (*big.Int, error) {
	return _PaymentExitgame.Contract.INITIALIFEBONDSIZE(&_PaymentExitgame.CallOpts)
}

// INITIALPBBONDSIZE is a free data retrieval call binding the contract method 0x7e2b2efe.
//
// Solidity: function INITIAL_PB_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCaller) INITIALPBBONDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "INITIAL_PB_BOND_SIZE")
	return *ret0, err
}

// INITIALPBBONDSIZE is a free data retrieval call binding the contract method 0x7e2b2efe.
//
// Solidity: function INITIAL_PB_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameSession) INITIALPBBONDSIZE() (*big.Int, error) {
	return _PaymentExitgame.Contract.INITIALPBBONDSIZE(&_PaymentExitgame.CallOpts)
}

// INITIALPBBONDSIZE is a free data retrieval call binding the contract method 0x7e2b2efe.
//
// Solidity: function INITIAL_PB_BOND_SIZE() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCallerSession) INITIALPBBONDSIZE() (*big.Int, error) {
	return _PaymentExitgame.Contract.INITIALPBBONDSIZE(&_PaymentExitgame.CallOpts)
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(bytes _txBytes) constant returns(uint160)
func (_PaymentExitgame *PaymentExitgameCaller) GetInFlightExitId(opts *bind.CallOpts, _txBytes []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "getInFlightExitId", _txBytes)
	return *ret0, err
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(bytes _txBytes) constant returns(uint160)
func (_PaymentExitgame *PaymentExitgameSession) GetInFlightExitId(_txBytes []byte) (*big.Int, error) {
	return _PaymentExitgame.Contract.GetInFlightExitId(&_PaymentExitgame.CallOpts, _txBytes)
}

// GetInFlightExitId is a free data retrieval call binding the contract method 0x839c78f9.
//
// Solidity: function getInFlightExitId(bytes _txBytes) constant returns(uint160)
func (_PaymentExitgame *PaymentExitgameCallerSession) GetInFlightExitId(_txBytes []byte) (*big.Int, error) {
	return _PaymentExitgame.Contract.GetInFlightExitId(&_PaymentExitgame.CallOpts, _txBytes)
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xb22d3e14.
//
// Solidity: function getStandardExitId(bool _isDeposit, bytes _txBytes, uint256 _utxoPos) constant returns(uint160)
func (_PaymentExitgame *PaymentExitgameCaller) GetStandardExitId(opts *bind.CallOpts, _isDeposit bool, _txBytes []byte, _utxoPos *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "getStandardExitId", _isDeposit, _txBytes, _utxoPos)
	return *ret0, err
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xb22d3e14.
//
// Solidity: function getStandardExitId(bool _isDeposit, bytes _txBytes, uint256 _utxoPos) constant returns(uint160)
func (_PaymentExitgame *PaymentExitgameSession) GetStandardExitId(_isDeposit bool, _txBytes []byte, _utxoPos *big.Int) (*big.Int, error) {
	return _PaymentExitgame.Contract.GetStandardExitId(&_PaymentExitgame.CallOpts, _isDeposit, _txBytes, _utxoPos)
}

// GetStandardExitId is a free data retrieval call binding the contract method 0xb22d3e14.
//
// Solidity: function getStandardExitId(bool _isDeposit, bytes _txBytes, uint256 _utxoPos) constant returns(uint160)
func (_PaymentExitgame *PaymentExitgameCallerSession) GetStandardExitId(_isDeposit bool, _txBytes []byte, _utxoPos *big.Int) (*big.Int, error) {
	return _PaymentExitgame.Contract.GetStandardExitId(&_PaymentExitgame.CallOpts, _isDeposit, _txBytes, _utxoPos)
}

// InFlightExits is a free data retrieval call binding the contract method 0xcec9e1a7.
//
// Solidity: function inFlightExits(uint160[] exitIds) constant returns([]PaymentExitDataModelInFlightExit)
func (_PaymentExitgame *PaymentExitgameCaller) InFlightExits(opts *bind.CallOpts, exitIds []*big.Int) ([]PaymentExitDataModelInFlightExit, error) {
	var (
		ret0 = new([]PaymentExitDataModelInFlightExit)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "inFlightExits", exitIds)
	return *ret0, err
}

// InFlightExits is a free data retrieval call binding the contract method 0xcec9e1a7.
//
// Solidity: function inFlightExits(uint160[] exitIds) constant returns([]PaymentExitDataModelInFlightExit)
func (_PaymentExitgame *PaymentExitgameSession) InFlightExits(exitIds []*big.Int) ([]PaymentExitDataModelInFlightExit, error) {
	return _PaymentExitgame.Contract.InFlightExits(&_PaymentExitgame.CallOpts, exitIds)
}

// InFlightExits is a free data retrieval call binding the contract method 0xcec9e1a7.
//
// Solidity: function inFlightExits(uint160[] exitIds) constant returns([]PaymentExitDataModelInFlightExit)
func (_PaymentExitgame *PaymentExitgameCallerSession) InFlightExits(exitIds []*big.Int) ([]PaymentExitDataModelInFlightExit, error) {
	return _PaymentExitgame.Contract.InFlightExits(&_PaymentExitgame.CallOpts, exitIds)
}

// PiggybackBondSize is a free data retrieval call binding the contract method 0x3f1214cf.
//
// Solidity: function piggybackBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCaller) PiggybackBondSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "piggybackBondSize")
	return *ret0, err
}

// PiggybackBondSize is a free data retrieval call binding the contract method 0x3f1214cf.
//
// Solidity: function piggybackBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameSession) PiggybackBondSize() (*big.Int, error) {
	return _PaymentExitgame.Contract.PiggybackBondSize(&_PaymentExitgame.CallOpts)
}

// PiggybackBondSize is a free data retrieval call binding the contract method 0x3f1214cf.
//
// Solidity: function piggybackBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCallerSession) PiggybackBondSize() (*big.Int, error) {
	return _PaymentExitgame.Contract.PiggybackBondSize(&_PaymentExitgame.CallOpts)
}

// StandardExits is a free data retrieval call binding the contract method 0x0ca5b676.
//
// Solidity: function standardExits(uint160[] exitIds) constant returns([]PaymentExitDataModelStandardExit)
func (_PaymentExitgame *PaymentExitgameCaller) StandardExits(opts *bind.CallOpts, exitIds []*big.Int) ([]PaymentExitDataModelStandardExit, error) {
	var (
		ret0 = new([]PaymentExitDataModelStandardExit)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "standardExits", exitIds)
	return *ret0, err
}

// StandardExits is a free data retrieval call binding the contract method 0x0ca5b676.
//
// Solidity: function standardExits(uint160[] exitIds) constant returns([]PaymentExitDataModelStandardExit)
func (_PaymentExitgame *PaymentExitgameSession) StandardExits(exitIds []*big.Int) ([]PaymentExitDataModelStandardExit, error) {
	return _PaymentExitgame.Contract.StandardExits(&_PaymentExitgame.CallOpts, exitIds)
}

// StandardExits is a free data retrieval call binding the contract method 0x0ca5b676.
//
// Solidity: function standardExits(uint160[] exitIds) constant returns([]PaymentExitDataModelStandardExit)
func (_PaymentExitgame *PaymentExitgameCallerSession) StandardExits(exitIds []*big.Int) ([]PaymentExitDataModelStandardExit, error) {
	return _PaymentExitgame.Contract.StandardExits(&_PaymentExitgame.CallOpts, exitIds)
}

// StartIFEBondSize is a free data retrieval call binding the contract method 0x7702fa17.
//
// Solidity: function startIFEBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCaller) StartIFEBondSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "startIFEBondSize")
	return *ret0, err
}

// StartIFEBondSize is a free data retrieval call binding the contract method 0x7702fa17.
//
// Solidity: function startIFEBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameSession) StartIFEBondSize() (*big.Int, error) {
	return _PaymentExitgame.Contract.StartIFEBondSize(&_PaymentExitgame.CallOpts)
}

// StartIFEBondSize is a free data retrieval call binding the contract method 0x7702fa17.
//
// Solidity: function startIFEBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCallerSession) StartIFEBondSize() (*big.Int, error) {
	return _PaymentExitgame.Contract.StartIFEBondSize(&_PaymentExitgame.CallOpts)
}

// StartStandardExitBondSize is a free data retrieval call binding the contract method 0xfe32a124.
//
// Solidity: function startStandardExitBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCaller) StartStandardExitBondSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentExitgame.contract.Call(opts, out, "startStandardExitBondSize")
	return *ret0, err
}

// StartStandardExitBondSize is a free data retrieval call binding the contract method 0xfe32a124.
//
// Solidity: function startStandardExitBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameSession) StartStandardExitBondSize() (*big.Int, error) {
	return _PaymentExitgame.Contract.StartStandardExitBondSize(&_PaymentExitgame.CallOpts)
}

// StartStandardExitBondSize is a free data retrieval call binding the contract method 0xfe32a124.
//
// Solidity: function startStandardExitBondSize() constant returns(uint128)
func (_PaymentExitgame *PaymentExitgameCallerSession) StartStandardExitBondSize() (*big.Int, error) {
	return _PaymentExitgame.Contract.StartStandardExitBondSize(&_PaymentExitgame.CallOpts)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0xb9a3a28e.
//
// Solidity: function challengeInFlightExitInputSpent(PaymentInFlightExitRouterArgsChallengeInputSpentArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) ChallengeInFlightExitInputSpent(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsChallengeInputSpentArgs) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "challengeInFlightExitInputSpent", args)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0xb9a3a28e.
//
// Solidity: function challengeInFlightExitInputSpent(PaymentInFlightExitRouterArgsChallengeInputSpentArgs args) returns()
func (_PaymentExitgame *PaymentExitgameSession) ChallengeInFlightExitInputSpent(args PaymentInFlightExitRouterArgsChallengeInputSpentArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ChallengeInFlightExitInputSpent(&_PaymentExitgame.TransactOpts, args)
}

// ChallengeInFlightExitInputSpent is a paid mutator transaction binding the contract method 0xb9a3a28e.
//
// Solidity: function challengeInFlightExitInputSpent(PaymentInFlightExitRouterArgsChallengeInputSpentArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) ChallengeInFlightExitInputSpent(args PaymentInFlightExitRouterArgsChallengeInputSpentArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ChallengeInFlightExitInputSpent(&_PaymentExitgame.TransactOpts, args)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0xe8362298.
//
// Solidity: function challengeInFlightExitNotCanonical(PaymentInFlightExitRouterArgsChallengeCanonicityArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) ChallengeInFlightExitNotCanonical(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsChallengeCanonicityArgs) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "challengeInFlightExitNotCanonical", args)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0xe8362298.
//
// Solidity: function challengeInFlightExitNotCanonical(PaymentInFlightExitRouterArgsChallengeCanonicityArgs args) returns()
func (_PaymentExitgame *PaymentExitgameSession) ChallengeInFlightExitNotCanonical(args PaymentInFlightExitRouterArgsChallengeCanonicityArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ChallengeInFlightExitNotCanonical(&_PaymentExitgame.TransactOpts, args)
}

// ChallengeInFlightExitNotCanonical is a paid mutator transaction binding the contract method 0xe8362298.
//
// Solidity: function challengeInFlightExitNotCanonical(PaymentInFlightExitRouterArgsChallengeCanonicityArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) ChallengeInFlightExitNotCanonical(args PaymentInFlightExitRouterArgsChallengeCanonicityArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ChallengeInFlightExitNotCanonical(&_PaymentExitgame.TransactOpts, args)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0x2c3bad95.
//
// Solidity: function challengeInFlightExitOutputSpent(PaymentInFlightExitRouterArgsChallengeOutputSpent args) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) ChallengeInFlightExitOutputSpent(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsChallengeOutputSpent) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "challengeInFlightExitOutputSpent", args)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0x2c3bad95.
//
// Solidity: function challengeInFlightExitOutputSpent(PaymentInFlightExitRouterArgsChallengeOutputSpent args) returns()
func (_PaymentExitgame *PaymentExitgameSession) ChallengeInFlightExitOutputSpent(args PaymentInFlightExitRouterArgsChallengeOutputSpent) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ChallengeInFlightExitOutputSpent(&_PaymentExitgame.TransactOpts, args)
}

// ChallengeInFlightExitOutputSpent is a paid mutator transaction binding the contract method 0x2c3bad95.
//
// Solidity: function challengeInFlightExitOutputSpent(PaymentInFlightExitRouterArgsChallengeOutputSpent args) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) ChallengeInFlightExitOutputSpent(args PaymentInFlightExitRouterArgsChallengeOutputSpent) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ChallengeInFlightExitOutputSpent(&_PaymentExitgame.TransactOpts, args)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0xdf2933e2.
//
// Solidity: function challengeStandardExit(PaymentStandardExitRouterArgsChallengeStandardExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) ChallengeStandardExit(opts *bind.TransactOpts, args PaymentStandardExitRouterArgsChallengeStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "challengeStandardExit", args)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0xdf2933e2.
//
// Solidity: function challengeStandardExit(PaymentStandardExitRouterArgsChallengeStandardExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameSession) ChallengeStandardExit(args PaymentStandardExitRouterArgsChallengeStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ChallengeStandardExit(&_PaymentExitgame.TransactOpts, args)
}

// ChallengeStandardExit is a paid mutator transaction binding the contract method 0xdf2933e2.
//
// Solidity: function challengeStandardExit(PaymentStandardExitRouterArgsChallengeStandardExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) ChallengeStandardExit(args PaymentStandardExitRouterArgsChallengeStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ChallengeStandardExit(&_PaymentExitgame.TransactOpts, args)
}

// DeleteNonPiggybackedInFlightExit is a paid mutator transaction binding the contract method 0xb3c2ec76.
//
// Solidity: function deleteNonPiggybackedInFlightExit(uint160 exitId) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) DeleteNonPiggybackedInFlightExit(opts *bind.TransactOpts, exitId *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "deleteNonPiggybackedInFlightExit", exitId)
}

// DeleteNonPiggybackedInFlightExit is a paid mutator transaction binding the contract method 0xb3c2ec76.
//
// Solidity: function deleteNonPiggybackedInFlightExit(uint160 exitId) returns()
func (_PaymentExitgame *PaymentExitgameSession) DeleteNonPiggybackedInFlightExit(exitId *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.DeleteNonPiggybackedInFlightExit(&_PaymentExitgame.TransactOpts, exitId)
}

// DeleteNonPiggybackedInFlightExit is a paid mutator transaction binding the contract method 0xb3c2ec76.
//
// Solidity: function deleteNonPiggybackedInFlightExit(uint160 exitId) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) DeleteNonPiggybackedInFlightExit(exitId *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.DeleteNonPiggybackedInFlightExit(&_PaymentExitgame.TransactOpts, exitId)
}

// PiggybackInFlightExitOnInput is a paid mutator transaction binding the contract method 0xe5bc60c7.
//
// Solidity: function piggybackInFlightExitOnInput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) PiggybackInFlightExitOnInput(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "piggybackInFlightExitOnInput", args)
}

// PiggybackInFlightExitOnInput is a paid mutator transaction binding the contract method 0xe5bc60c7.
//
// Solidity: function piggybackInFlightExitOnInput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs args) returns()
func (_PaymentExitgame *PaymentExitgameSession) PiggybackInFlightExitOnInput(args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.PiggybackInFlightExitOnInput(&_PaymentExitgame.TransactOpts, args)
}

// PiggybackInFlightExitOnInput is a paid mutator transaction binding the contract method 0xe5bc60c7.
//
// Solidity: function piggybackInFlightExitOnInput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) PiggybackInFlightExitOnInput(args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnInputArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.PiggybackInFlightExitOnInput(&_PaymentExitgame.TransactOpts, args)
}

// PiggybackInFlightExitOnOutput is a paid mutator transaction binding the contract method 0xa4a25441.
//
// Solidity: function piggybackInFlightExitOnOutput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) PiggybackInFlightExitOnOutput(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "piggybackInFlightExitOnOutput", args)
}

// PiggybackInFlightExitOnOutput is a paid mutator transaction binding the contract method 0xa4a25441.
//
// Solidity: function piggybackInFlightExitOnOutput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs args) returns()
func (_PaymentExitgame *PaymentExitgameSession) PiggybackInFlightExitOnOutput(args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.PiggybackInFlightExitOnOutput(&_PaymentExitgame.TransactOpts, args)
}

// PiggybackInFlightExitOnOutput is a paid mutator transaction binding the contract method 0xa4a25441.
//
// Solidity: function piggybackInFlightExitOnOutput(PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) PiggybackInFlightExitOnOutput(args PaymentInFlightExitRouterArgsPiggybackInFlightExitOnOutputArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.PiggybackInFlightExitOnOutput(&_PaymentExitgame.TransactOpts, args)
}

// ProcessExit is a paid mutator transaction binding the contract method 0x68496351.
//
// Solidity: function processExit(uint160 exitId, uint256 , address token) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) ProcessExit(opts *bind.TransactOpts, exitId *big.Int, arg1 *big.Int, token common.Address) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "processExit", exitId, arg1, token)
}

// ProcessExit is a paid mutator transaction binding the contract method 0x68496351.
//
// Solidity: function processExit(uint160 exitId, uint256 , address token) returns()
func (_PaymentExitgame *PaymentExitgameSession) ProcessExit(exitId *big.Int, arg1 *big.Int, token common.Address) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ProcessExit(&_PaymentExitgame.TransactOpts, exitId, arg1, token)
}

// ProcessExit is a paid mutator transaction binding the contract method 0x68496351.
//
// Solidity: function processExit(uint160 exitId, uint256 , address token) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) ProcessExit(exitId *big.Int, arg1 *big.Int, token common.Address) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.ProcessExit(&_PaymentExitgame.TransactOpts, exitId, arg1, token)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(bytes inFlightTx, uint256 inFlightTxPos, bytes inFlightTxInclusionProof) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) RespondToNonCanonicalChallenge(opts *bind.TransactOpts, inFlightTx []byte, inFlightTxPos *big.Int, inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "respondToNonCanonicalChallenge", inFlightTx, inFlightTxPos, inFlightTxInclusionProof)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(bytes inFlightTx, uint256 inFlightTxPos, bytes inFlightTxInclusionProof) returns()
func (_PaymentExitgame *PaymentExitgameSession) RespondToNonCanonicalChallenge(inFlightTx []byte, inFlightTxPos *big.Int, inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.RespondToNonCanonicalChallenge(&_PaymentExitgame.TransactOpts, inFlightTx, inFlightTxPos, inFlightTxInclusionProof)
}

// RespondToNonCanonicalChallenge is a paid mutator transaction binding the contract method 0x36660de4.
//
// Solidity: function respondToNonCanonicalChallenge(bytes inFlightTx, uint256 inFlightTxPos, bytes inFlightTxInclusionProof) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) RespondToNonCanonicalChallenge(inFlightTx []byte, inFlightTxPos *big.Int, inFlightTxInclusionProof []byte) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.RespondToNonCanonicalChallenge(&_PaymentExitgame.TransactOpts, inFlightTx, inFlightTxPos, inFlightTxInclusionProof)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0x5a528514.
//
// Solidity: function startInFlightExit(PaymentInFlightExitRouterArgsStartExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) StartInFlightExit(opts *bind.TransactOpts, args PaymentInFlightExitRouterArgsStartExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "startInFlightExit", args)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0x5a528514.
//
// Solidity: function startInFlightExit(PaymentInFlightExitRouterArgsStartExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameSession) StartInFlightExit(args PaymentInFlightExitRouterArgsStartExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.StartInFlightExit(&_PaymentExitgame.TransactOpts, args)
}

// StartInFlightExit is a paid mutator transaction binding the contract method 0x5a528514.
//
// Solidity: function startInFlightExit(PaymentInFlightExitRouterArgsStartExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) StartInFlightExit(args PaymentInFlightExitRouterArgsStartExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.StartInFlightExit(&_PaymentExitgame.TransactOpts, args)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0x70e01462.
//
// Solidity: function startStandardExit(PaymentStandardExitRouterArgsStartStandardExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) StartStandardExit(opts *bind.TransactOpts, args PaymentStandardExitRouterArgsStartStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "startStandardExit", args)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0x70e01462.
//
// Solidity: function startStandardExit(PaymentStandardExitRouterArgsStartStandardExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameSession) StartStandardExit(args PaymentStandardExitRouterArgsStartStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.StartStandardExit(&_PaymentExitgame.TransactOpts, args)
}

// StartStandardExit is a paid mutator transaction binding the contract method 0x70e01462.
//
// Solidity: function startStandardExit(PaymentStandardExitRouterArgsStartStandardExitArgs args) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) StartStandardExit(args PaymentStandardExitRouterArgsStartStandardExitArgs) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.StartStandardExit(&_PaymentExitgame.TransactOpts, args)
}

// UpdatePiggybackBondSize is a paid mutator transaction binding the contract method 0x315fb7da.
//
// Solidity: function updatePiggybackBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) UpdatePiggybackBondSize(opts *bind.TransactOpts, newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "updatePiggybackBondSize", newBondSize)
}

// UpdatePiggybackBondSize is a paid mutator transaction binding the contract method 0x315fb7da.
//
// Solidity: function updatePiggybackBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameSession) UpdatePiggybackBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.UpdatePiggybackBondSize(&_PaymentExitgame.TransactOpts, newBondSize)
}

// UpdatePiggybackBondSize is a paid mutator transaction binding the contract method 0x315fb7da.
//
// Solidity: function updatePiggybackBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) UpdatePiggybackBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.UpdatePiggybackBondSize(&_PaymentExitgame.TransactOpts, newBondSize)
}

// UpdateStartIFEBondSize is a paid mutator transaction binding the contract method 0xb177ba23.
//
// Solidity: function updateStartIFEBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) UpdateStartIFEBondSize(opts *bind.TransactOpts, newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "updateStartIFEBondSize", newBondSize)
}

// UpdateStartIFEBondSize is a paid mutator transaction binding the contract method 0xb177ba23.
//
// Solidity: function updateStartIFEBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameSession) UpdateStartIFEBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.UpdateStartIFEBondSize(&_PaymentExitgame.TransactOpts, newBondSize)
}

// UpdateStartIFEBondSize is a paid mutator transaction binding the contract method 0xb177ba23.
//
// Solidity: function updateStartIFEBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) UpdateStartIFEBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.UpdateStartIFEBondSize(&_PaymentExitgame.TransactOpts, newBondSize)
}

// UpdateStartStandardExitBondSize is a paid mutator transaction binding the contract method 0x163d7bc7.
//
// Solidity: function updateStartStandardExitBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameTransactor) UpdateStartStandardExitBondSize(opts *bind.TransactOpts, newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.contract.Transact(opts, "updateStartStandardExitBondSize", newBondSize)
}

// UpdateStartStandardExitBondSize is a paid mutator transaction binding the contract method 0x163d7bc7.
//
// Solidity: function updateStartStandardExitBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameSession) UpdateStartStandardExitBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.UpdateStartStandardExitBondSize(&_PaymentExitgame.TransactOpts, newBondSize)
}

// UpdateStartStandardExitBondSize is a paid mutator transaction binding the contract method 0x163d7bc7.
//
// Solidity: function updateStartStandardExitBondSize(uint128 newBondSize) returns()
func (_PaymentExitgame *PaymentExitgameTransactorSession) UpdateStartStandardExitBondSize(newBondSize *big.Int) (*types.Transaction, error) {
	return _PaymentExitgame.Contract.UpdateStartStandardExitBondSize(&_PaymentExitgame.TransactOpts, newBondSize)
}

// PaymentExitgameBondReturnFailedIterator is returned from FilterBondReturnFailed and is used to iterate over the raw logs and unpacked data for BondReturnFailed events raised by the PaymentExitgame contract.
type PaymentExitgameBondReturnFailedIterator struct {
	Event *PaymentExitgameBondReturnFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameBondReturnFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameBondReturnFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameBondReturnFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameBondReturnFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameBondReturnFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameBondReturnFailed represents a BondReturnFailed event raised by the PaymentExitgame contract.
type PaymentExitgameBondReturnFailed struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBondReturnFailed is a free log retrieval operation binding the contract event 0xc128a1281b55cd407680883dcfc3d0fffb69ec5176f175cefa465d34968547e8.
//
// Solidity: event BondReturnFailed(address indexed receiver, uint256 amount)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterBondReturnFailed(opts *bind.FilterOpts, receiver []common.Address) (*PaymentExitgameBondReturnFailedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "BondReturnFailed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameBondReturnFailedIterator{contract: _PaymentExitgame.contract, event: "BondReturnFailed", logs: logs, sub: sub}, nil
}

// WatchBondReturnFailed is a free log subscription operation binding the contract event 0xc128a1281b55cd407680883dcfc3d0fffb69ec5176f175cefa465d34968547e8.
//
// Solidity: event BondReturnFailed(address indexed receiver, uint256 amount)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchBondReturnFailed(opts *bind.WatchOpts, sink chan<- *PaymentExitgameBondReturnFailed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "BondReturnFailed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameBondReturnFailed)
				if err := _PaymentExitgame.contract.UnpackLog(event, "BondReturnFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBondReturnFailed is a log parse operation binding the contract event 0xc128a1281b55cd407680883dcfc3d0fffb69ec5176f175cefa465d34968547e8.
//
// Solidity: event BondReturnFailed(address indexed receiver, uint256 amount)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseBondReturnFailed(log types.Log) (*PaymentExitgameBondReturnFailed, error) {
	event := new(PaymentExitgameBondReturnFailed)
	if err := _PaymentExitgame.contract.UnpackLog(event, "BondReturnFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameExitChallengedIterator is returned from FilterExitChallenged and is used to iterate over the raw logs and unpacked data for ExitChallenged events raised by the PaymentExitgame contract.
type PaymentExitgameExitChallengedIterator struct {
	Event *PaymentExitgameExitChallenged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameExitChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameExitChallenged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameExitChallenged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameExitChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameExitChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameExitChallenged represents a ExitChallenged event raised by the PaymentExitgame contract.
type PaymentExitgameExitChallenged struct {
	UtxoPos *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitChallenged is a free log retrieval operation binding the contract event 0x5dfba526c59b25f899f935c5b0d5b8739e97e4d89c38c158eca3192ea34b87d8.
//
// Solidity: event ExitChallenged(uint256 indexed utxoPos)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterExitChallenged(opts *bind.FilterOpts, utxoPos []*big.Int) (*PaymentExitgameExitChallengedIterator, error) {

	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "ExitChallenged", utxoPosRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameExitChallengedIterator{contract: _PaymentExitgame.contract, event: "ExitChallenged", logs: logs, sub: sub}, nil
}

// WatchExitChallenged is a free log subscription operation binding the contract event 0x5dfba526c59b25f899f935c5b0d5b8739e97e4d89c38c158eca3192ea34b87d8.
//
// Solidity: event ExitChallenged(uint256 indexed utxoPos)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchExitChallenged(opts *bind.WatchOpts, sink chan<- *PaymentExitgameExitChallenged, utxoPos []*big.Int) (event.Subscription, error) {

	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "ExitChallenged", utxoPosRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameExitChallenged)
				if err := _PaymentExitgame.contract.UnpackLog(event, "ExitChallenged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExitChallenged is a log parse operation binding the contract event 0x5dfba526c59b25f899f935c5b0d5b8739e97e4d89c38c158eca3192ea34b87d8.
//
// Solidity: event ExitChallenged(uint256 indexed utxoPos)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseExitChallenged(log types.Log) (*PaymentExitgameExitChallenged, error) {
	event := new(PaymentExitgameExitChallenged)
	if err := _PaymentExitgame.contract.UnpackLog(event, "ExitChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameExitFinalizedIterator is returned from FilterExitFinalized and is used to iterate over the raw logs and unpacked data for ExitFinalized events raised by the PaymentExitgame contract.
type PaymentExitgameExitFinalizedIterator struct {
	Event *PaymentExitgameExitFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameExitFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameExitFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameExitFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameExitFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameExitFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameExitFinalized represents a ExitFinalized event raised by the PaymentExitgame contract.
type PaymentExitgameExitFinalized struct {
	ExitId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitFinalized is a free log retrieval operation binding the contract event 0x0adb29b0831e081044cefe31155c1f2b2b85ad3613a480a5f901ee287addef55.
//
// Solidity: event ExitFinalized(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterExitFinalized(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitgameExitFinalizedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "ExitFinalized", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameExitFinalizedIterator{contract: _PaymentExitgame.contract, event: "ExitFinalized", logs: logs, sub: sub}, nil
}

// WatchExitFinalized is a free log subscription operation binding the contract event 0x0adb29b0831e081044cefe31155c1f2b2b85ad3613a480a5f901ee287addef55.
//
// Solidity: event ExitFinalized(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchExitFinalized(opts *bind.WatchOpts, sink chan<- *PaymentExitgameExitFinalized, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "ExitFinalized", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameExitFinalized)
				if err := _PaymentExitgame.contract.UnpackLog(event, "ExitFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExitFinalized is a log parse operation binding the contract event 0x0adb29b0831e081044cefe31155c1f2b2b85ad3613a480a5f901ee287addef55.
//
// Solidity: event ExitFinalized(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseExitFinalized(log types.Log) (*PaymentExitgameExitFinalized, error) {
	event := new(PaymentExitgameExitFinalized)
	if err := _PaymentExitgame.contract.UnpackLog(event, "ExitFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameExitOmittedIterator is returned from FilterExitOmitted and is used to iterate over the raw logs and unpacked data for ExitOmitted events raised by the PaymentExitgame contract.
type PaymentExitgameExitOmittedIterator struct {
	Event *PaymentExitgameExitOmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameExitOmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameExitOmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameExitOmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameExitOmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameExitOmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameExitOmitted represents a ExitOmitted event raised by the PaymentExitgame contract.
type PaymentExitgameExitOmitted struct {
	ExitId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitOmitted is a free log retrieval operation binding the contract event 0x91b69d9a0a91ef4f81d0232886cb1880ada4ccdadc8981a3fede280c284bc46a.
//
// Solidity: event ExitOmitted(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterExitOmitted(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitgameExitOmittedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "ExitOmitted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameExitOmittedIterator{contract: _PaymentExitgame.contract, event: "ExitOmitted", logs: logs, sub: sub}, nil
}

// WatchExitOmitted is a free log subscription operation binding the contract event 0x91b69d9a0a91ef4f81d0232886cb1880ada4ccdadc8981a3fede280c284bc46a.
//
// Solidity: event ExitOmitted(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchExitOmitted(opts *bind.WatchOpts, sink chan<- *PaymentExitgameExitOmitted, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "ExitOmitted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameExitOmitted)
				if err := _PaymentExitgame.contract.UnpackLog(event, "ExitOmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExitOmitted is a log parse operation binding the contract event 0x91b69d9a0a91ef4f81d0232886cb1880ada4ccdadc8981a3fede280c284bc46a.
//
// Solidity: event ExitOmitted(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseExitOmitted(log types.Log) (*PaymentExitgameExitOmitted, error) {
	event := new(PaymentExitgameExitOmitted)
	if err := _PaymentExitgame.contract.UnpackLog(event, "ExitOmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameExitStartedIterator is returned from FilterExitStarted and is used to iterate over the raw logs and unpacked data for ExitStarted events raised by the PaymentExitgame contract.
type PaymentExitgameExitStartedIterator struct {
	Event *PaymentExitgameExitStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameExitStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameExitStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameExitStarted represents a ExitStarted event raised by the PaymentExitgame contract.
type PaymentExitgameExitStarted struct {
	Owner  common.Address
	ExitId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitStarted is a free log retrieval operation binding the contract event 0xdd6f755cba05d0a420007aef6afc05e4889ab424505e2e440ecd1c434ba7082e.
//
// Solidity: event ExitStarted(address indexed owner, uint160 exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterExitStarted(opts *bind.FilterOpts, owner []common.Address) (*PaymentExitgameExitStartedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "ExitStarted", ownerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameExitStartedIterator{contract: _PaymentExitgame.contract, event: "ExitStarted", logs: logs, sub: sub}, nil
}

// WatchExitStarted is a free log subscription operation binding the contract event 0xdd6f755cba05d0a420007aef6afc05e4889ab424505e2e440ecd1c434ba7082e.
//
// Solidity: event ExitStarted(address indexed owner, uint160 exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchExitStarted(opts *bind.WatchOpts, sink chan<- *PaymentExitgameExitStarted, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "ExitStarted", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameExitStarted)
				if err := _PaymentExitgame.contract.UnpackLog(event, "ExitStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExitStarted is a log parse operation binding the contract event 0xdd6f755cba05d0a420007aef6afc05e4889ab424505e2e440ecd1c434ba7082e.
//
// Solidity: event ExitStarted(address indexed owner, uint160 exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseExitStarted(log types.Log) (*PaymentExitgameExitStarted, error) {
	event := new(PaymentExitgameExitStarted)
	if err := _PaymentExitgame.contract.UnpackLog(event, "ExitStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameIFEBondUpdatedIterator is returned from FilterIFEBondUpdated and is used to iterate over the raw logs and unpacked data for IFEBondUpdated events raised by the PaymentExitgame contract.
type PaymentExitgameIFEBondUpdatedIterator struct {
	Event *PaymentExitgameIFEBondUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameIFEBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameIFEBondUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameIFEBondUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameIFEBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameIFEBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameIFEBondUpdated represents a IFEBondUpdated event raised by the PaymentExitgame contract.
type PaymentExitgameIFEBondUpdated struct {
	BondSize *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIFEBondUpdated is a free log retrieval operation binding the contract event 0x63ccac61f03626a0e717e96da239ff8996745b6c4ddcaffdedc88d7ef938f5ec.
//
// Solidity: event IFEBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterIFEBondUpdated(opts *bind.FilterOpts) (*PaymentExitgameIFEBondUpdatedIterator, error) {

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "IFEBondUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameIFEBondUpdatedIterator{contract: _PaymentExitgame.contract, event: "IFEBondUpdated", logs: logs, sub: sub}, nil
}

// WatchIFEBondUpdated is a free log subscription operation binding the contract event 0x63ccac61f03626a0e717e96da239ff8996745b6c4ddcaffdedc88d7ef938f5ec.
//
// Solidity: event IFEBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchIFEBondUpdated(opts *bind.WatchOpts, sink chan<- *PaymentExitgameIFEBondUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "IFEBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameIFEBondUpdated)
				if err := _PaymentExitgame.contract.UnpackLog(event, "IFEBondUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIFEBondUpdated is a log parse operation binding the contract event 0x63ccac61f03626a0e717e96da239ff8996745b6c4ddcaffdedc88d7ef938f5ec.
//
// Solidity: event IFEBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseIFEBondUpdated(log types.Log) (*PaymentExitgameIFEBondUpdated, error) {
	event := new(PaymentExitgameIFEBondUpdated)
	if err := _PaymentExitgame.contract.UnpackLog(event, "IFEBondUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightBondReturnFailedIterator is returned from FilterInFlightBondReturnFailed and is used to iterate over the raw logs and unpacked data for InFlightBondReturnFailed events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightBondReturnFailedIterator struct {
	Event *PaymentExitgameInFlightBondReturnFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightBondReturnFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightBondReturnFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightBondReturnFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightBondReturnFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightBondReturnFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightBondReturnFailed represents a InFlightBondReturnFailed event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightBondReturnFailed struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInFlightBondReturnFailed is a free log retrieval operation binding the contract event 0x6e9bc14d15fc350181ead2878efb011f527c6984ebf42e3c0dbf72a2d5edc9f2.
//
// Solidity: event InFlightBondReturnFailed(address indexed receiver, uint256 amount)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightBondReturnFailed(opts *bind.FilterOpts, receiver []common.Address) (*PaymentExitgameInFlightBondReturnFailedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightBondReturnFailed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightBondReturnFailedIterator{contract: _PaymentExitgame.contract, event: "InFlightBondReturnFailed", logs: logs, sub: sub}, nil
}

// WatchInFlightBondReturnFailed is a free log subscription operation binding the contract event 0x6e9bc14d15fc350181ead2878efb011f527c6984ebf42e3c0dbf72a2d5edc9f2.
//
// Solidity: event InFlightBondReturnFailed(address indexed receiver, uint256 amount)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightBondReturnFailed(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightBondReturnFailed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightBondReturnFailed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightBondReturnFailed)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightBondReturnFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightBondReturnFailed is a log parse operation binding the contract event 0x6e9bc14d15fc350181ead2878efb011f527c6984ebf42e3c0dbf72a2d5edc9f2.
//
// Solidity: event InFlightBondReturnFailed(address indexed receiver, uint256 amount)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightBondReturnFailed(log types.Log) (*PaymentExitgameInFlightBondReturnFailed, error) {
	event := new(PaymentExitgameInFlightBondReturnFailed)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightBondReturnFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitChallengeRespondedIterator is returned from FilterInFlightExitChallengeResponded and is used to iterate over the raw logs and unpacked data for InFlightExitChallengeResponded events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitChallengeRespondedIterator struct {
	Event *PaymentExitgameInFlightExitChallengeResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitChallengeRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitChallengeResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitChallengeResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitChallengeRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitChallengeRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitChallengeResponded represents a InFlightExitChallengeResponded event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitChallengeResponded struct {
	Challenger          common.Address
	TxHash              [32]byte
	ChallengeTxPosition *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitChallengeResponded is a free log retrieval operation binding the contract event 0x637cc4a7148767df19331a5c7dfb6d31f0a7e159a3dbb28a716be18c8c74f768.
//
// Solidity: event InFlightExitChallengeResponded(address indexed challenger, bytes32 indexed txHash, uint256 challengeTxPosition)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitChallengeResponded(opts *bind.FilterOpts, challenger []common.Address, txHash [][32]byte) (*PaymentExitgameInFlightExitChallengeRespondedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitChallengeResponded", challengerRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitChallengeRespondedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitChallengeResponded", logs: logs, sub: sub}, nil
}

// WatchInFlightExitChallengeResponded is a free log subscription operation binding the contract event 0x637cc4a7148767df19331a5c7dfb6d31f0a7e159a3dbb28a716be18c8c74f768.
//
// Solidity: event InFlightExitChallengeResponded(address indexed challenger, bytes32 indexed txHash, uint256 challengeTxPosition)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitChallengeResponded(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitChallengeResponded, challenger []common.Address, txHash [][32]byte) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitChallengeResponded", challengerRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitChallengeResponded)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitChallengeResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitChallengeResponded is a log parse operation binding the contract event 0x637cc4a7148767df19331a5c7dfb6d31f0a7e159a3dbb28a716be18c8c74f768.
//
// Solidity: event InFlightExitChallengeResponded(address indexed challenger, bytes32 indexed txHash, uint256 challengeTxPosition)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitChallengeResponded(log types.Log) (*PaymentExitgameInFlightExitChallengeResponded, error) {
	event := new(PaymentExitgameInFlightExitChallengeResponded)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitChallengeResponded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitChallengedIterator is returned from FilterInFlightExitChallenged and is used to iterate over the raw logs and unpacked data for InFlightExitChallenged events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitChallengedIterator struct {
	Event *PaymentExitgameInFlightExitChallenged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitChallenged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitChallenged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitChallenged represents a InFlightExitChallenged event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitChallenged struct {
	Challenger          common.Address
	TxHash              [32]byte
	ChallengeTxPosition *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitChallenged is a free log retrieval operation binding the contract event 0x687401968e501bda2d2d6f880dd1a0a56ff50b1787185ee0b6f4c3fb9fc417ab.
//
// Solidity: event InFlightExitChallenged(address indexed challenger, bytes32 indexed txHash, uint256 challengeTxPosition)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitChallenged(opts *bind.FilterOpts, challenger []common.Address, txHash [][32]byte) (*PaymentExitgameInFlightExitChallengedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitChallenged", challengerRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitChallengedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitChallenged", logs: logs, sub: sub}, nil
}

// WatchInFlightExitChallenged is a free log subscription operation binding the contract event 0x687401968e501bda2d2d6f880dd1a0a56ff50b1787185ee0b6f4c3fb9fc417ab.
//
// Solidity: event InFlightExitChallenged(address indexed challenger, bytes32 indexed txHash, uint256 challengeTxPosition)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitChallenged(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitChallenged, challenger []common.Address, txHash [][32]byte) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitChallenged", challengerRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitChallenged)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitChallenged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitChallenged is a log parse operation binding the contract event 0x687401968e501bda2d2d6f880dd1a0a56ff50b1787185ee0b6f4c3fb9fc417ab.
//
// Solidity: event InFlightExitChallenged(address indexed challenger, bytes32 indexed txHash, uint256 challengeTxPosition)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitChallenged(log types.Log) (*PaymentExitgameInFlightExitChallenged, error) {
	event := new(PaymentExitgameInFlightExitChallenged)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitDeletedIterator is returned from FilterInFlightExitDeleted and is used to iterate over the raw logs and unpacked data for InFlightExitDeleted events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitDeletedIterator struct {
	Event *PaymentExitgameInFlightExitDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitDeleted represents a InFlightExitDeleted event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitDeleted struct {
	ExitId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitDeleted is a free log retrieval operation binding the contract event 0x1991c4c350498b0cc937c6a08bc5bdecf2e4fdd9d918052a880f102e43dbe45c.
//
// Solidity: event InFlightExitDeleted(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitDeleted(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitgameInFlightExitDeletedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitDeleted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitDeletedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitDeleted", logs: logs, sub: sub}, nil
}

// WatchInFlightExitDeleted is a free log subscription operation binding the contract event 0x1991c4c350498b0cc937c6a08bc5bdecf2e4fdd9d918052a880f102e43dbe45c.
//
// Solidity: event InFlightExitDeleted(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitDeleted(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitDeleted, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitDeleted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitDeleted)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitDeleted is a log parse operation binding the contract event 0x1991c4c350498b0cc937c6a08bc5bdecf2e4fdd9d918052a880f102e43dbe45c.
//
// Solidity: event InFlightExitDeleted(uint160 indexed exitId)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitDeleted(log types.Log) (*PaymentExitgameInFlightExitDeleted, error) {
	event := new(PaymentExitgameInFlightExitDeleted)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitInputBlockedIterator is returned from FilterInFlightExitInputBlocked and is used to iterate over the raw logs and unpacked data for InFlightExitInputBlocked events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitInputBlockedIterator struct {
	Event *PaymentExitgameInFlightExitInputBlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitInputBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitInputBlocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitInputBlocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitInputBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitInputBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitInputBlocked represents a InFlightExitInputBlocked event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitInputBlocked struct {
	Challenger common.Address
	TxHash     [32]byte
	InputIndex uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitInputBlocked is a free log retrieval operation binding the contract event 0x4794045885b2e249895bc730deb7caedc2e6b4b49d56ccd696b69d1f4944b8ec.
//
// Solidity: event InFlightExitInputBlocked(address indexed challenger, bytes32 indexed txHash, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitInputBlocked(opts *bind.FilterOpts, challenger []common.Address, txHash [][32]byte) (*PaymentExitgameInFlightExitInputBlockedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitInputBlocked", challengerRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitInputBlockedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitInputBlocked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitInputBlocked is a free log subscription operation binding the contract event 0x4794045885b2e249895bc730deb7caedc2e6b4b49d56ccd696b69d1f4944b8ec.
//
// Solidity: event InFlightExitInputBlocked(address indexed challenger, bytes32 indexed txHash, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitInputBlocked(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitInputBlocked, challenger []common.Address, txHash [][32]byte) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitInputBlocked", challengerRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitInputBlocked)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitInputBlocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitInputBlocked is a log parse operation binding the contract event 0x4794045885b2e249895bc730deb7caedc2e6b4b49d56ccd696b69d1f4944b8ec.
//
// Solidity: event InFlightExitInputBlocked(address indexed challenger, bytes32 indexed txHash, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitInputBlocked(log types.Log) (*PaymentExitgameInFlightExitInputBlocked, error) {
	event := new(PaymentExitgameInFlightExitInputBlocked)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitInputBlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitInputPiggybackedIterator is returned from FilterInFlightExitInputPiggybacked and is used to iterate over the raw logs and unpacked data for InFlightExitInputPiggybacked events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitInputPiggybackedIterator struct {
	Event *PaymentExitgameInFlightExitInputPiggybacked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitInputPiggybackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitInputPiggybacked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitInputPiggybacked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitInputPiggybackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitInputPiggybackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitInputPiggybacked represents a InFlightExitInputPiggybacked event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitInputPiggybacked struct {
	ExitTarget common.Address
	TxHash     [32]byte
	InputIndex uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitInputPiggybacked is a free log retrieval operation binding the contract event 0xa93c0e9b202feaf554acf6ef1185b898c9f214da16e51740b06b5f7487b018e5.
//
// Solidity: event InFlightExitInputPiggybacked(address indexed exitTarget, bytes32 indexed txHash, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitInputPiggybacked(opts *bind.FilterOpts, exitTarget []common.Address, txHash [][32]byte) (*PaymentExitgameInFlightExitInputPiggybackedIterator, error) {

	var exitTargetRule []interface{}
	for _, exitTargetItem := range exitTarget {
		exitTargetRule = append(exitTargetRule, exitTargetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitInputPiggybacked", exitTargetRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitInputPiggybackedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitInputPiggybacked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitInputPiggybacked is a free log subscription operation binding the contract event 0xa93c0e9b202feaf554acf6ef1185b898c9f214da16e51740b06b5f7487b018e5.
//
// Solidity: event InFlightExitInputPiggybacked(address indexed exitTarget, bytes32 indexed txHash, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitInputPiggybacked(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitInputPiggybacked, exitTarget []common.Address, txHash [][32]byte) (event.Subscription, error) {

	var exitTargetRule []interface{}
	for _, exitTargetItem := range exitTarget {
		exitTargetRule = append(exitTargetRule, exitTargetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitInputPiggybacked", exitTargetRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitInputPiggybacked)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitInputPiggybacked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitInputPiggybacked is a log parse operation binding the contract event 0xa93c0e9b202feaf554acf6ef1185b898c9f214da16e51740b06b5f7487b018e5.
//
// Solidity: event InFlightExitInputPiggybacked(address indexed exitTarget, bytes32 indexed txHash, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitInputPiggybacked(log types.Log) (*PaymentExitgameInFlightExitInputPiggybacked, error) {
	event := new(PaymentExitgameInFlightExitInputPiggybacked)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitInputPiggybacked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitInputWithdrawnIterator is returned from FilterInFlightExitInputWithdrawn and is used to iterate over the raw logs and unpacked data for InFlightExitInputWithdrawn events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitInputWithdrawnIterator struct {
	Event *PaymentExitgameInFlightExitInputWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitInputWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitInputWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitInputWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitInputWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitInputWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitInputWithdrawn represents a InFlightExitInputWithdrawn event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitInputWithdrawn struct {
	ExitId     *big.Int
	InputIndex uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitInputWithdrawn is a free log retrieval operation binding the contract event 0x4446ec118df0062851669a051944a56f276ef134203ac88cbb6032bd10b6aeec.
//
// Solidity: event InFlightExitInputWithdrawn(uint160 indexed exitId, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitInputWithdrawn(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitgameInFlightExitInputWithdrawnIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitInputWithdrawn", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitInputWithdrawnIterator{contract: _PaymentExitgame.contract, event: "InFlightExitInputWithdrawn", logs: logs, sub: sub}, nil
}

// WatchInFlightExitInputWithdrawn is a free log subscription operation binding the contract event 0x4446ec118df0062851669a051944a56f276ef134203ac88cbb6032bd10b6aeec.
//
// Solidity: event InFlightExitInputWithdrawn(uint160 indexed exitId, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitInputWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitInputWithdrawn, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitInputWithdrawn", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitInputWithdrawn)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitInputWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitInputWithdrawn is a log parse operation binding the contract event 0x4446ec118df0062851669a051944a56f276ef134203ac88cbb6032bd10b6aeec.
//
// Solidity: event InFlightExitInputWithdrawn(uint160 indexed exitId, uint16 inputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitInputWithdrawn(log types.Log) (*PaymentExitgameInFlightExitInputWithdrawn, error) {
	event := new(PaymentExitgameInFlightExitInputWithdrawn)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitInputWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitOmittedIterator is returned from FilterInFlightExitOmitted and is used to iterate over the raw logs and unpacked data for InFlightExitOmitted events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitOmittedIterator struct {
	Event *PaymentExitgameInFlightExitOmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitOmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitOmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitOmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitOmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitOmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitOmitted represents a InFlightExitOmitted event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitOmitted struct {
	ExitId *big.Int
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOmitted is a free log retrieval operation binding the contract event 0x438fb52bbaecea88ef59bcef337f0f8d5ba28dc74f727c0c176f81ff2c7ab81f.
//
// Solidity: event InFlightExitOmitted(uint160 indexed exitId, address token)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitOmitted(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitgameInFlightExitOmittedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitOmitted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitOmittedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitOmitted", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOmitted is a free log subscription operation binding the contract event 0x438fb52bbaecea88ef59bcef337f0f8d5ba28dc74f727c0c176f81ff2c7ab81f.
//
// Solidity: event InFlightExitOmitted(uint160 indexed exitId, address token)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitOmitted(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitOmitted, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitOmitted", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitOmitted)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitOmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitOmitted is a log parse operation binding the contract event 0x438fb52bbaecea88ef59bcef337f0f8d5ba28dc74f727c0c176f81ff2c7ab81f.
//
// Solidity: event InFlightExitOmitted(uint160 indexed exitId, address token)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitOmitted(log types.Log) (*PaymentExitgameInFlightExitOmitted, error) {
	event := new(PaymentExitgameInFlightExitOmitted)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitOmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitOutputBlockedIterator is returned from FilterInFlightExitOutputBlocked and is used to iterate over the raw logs and unpacked data for InFlightExitOutputBlocked events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitOutputBlockedIterator struct {
	Event *PaymentExitgameInFlightExitOutputBlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitOutputBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitOutputBlocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitOutputBlocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitOutputBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitOutputBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitOutputBlocked represents a InFlightExitOutputBlocked event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitOutputBlocked struct {
	Challenger  common.Address
	TxHash      [32]byte
	OutputIndex uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOutputBlocked is a free log retrieval operation binding the contract event 0xcbe8dad2e7fcbfe0dcba2f9b2e44f122c66cd26dc0808a0f7e9ec41e4fe285bf.
//
// Solidity: event InFlightExitOutputBlocked(address indexed challenger, bytes32 indexed txHash, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitOutputBlocked(opts *bind.FilterOpts, challenger []common.Address, txHash [][32]byte) (*PaymentExitgameInFlightExitOutputBlockedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitOutputBlocked", challengerRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitOutputBlockedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitOutputBlocked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOutputBlocked is a free log subscription operation binding the contract event 0xcbe8dad2e7fcbfe0dcba2f9b2e44f122c66cd26dc0808a0f7e9ec41e4fe285bf.
//
// Solidity: event InFlightExitOutputBlocked(address indexed challenger, bytes32 indexed txHash, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitOutputBlocked(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitOutputBlocked, challenger []common.Address, txHash [][32]byte) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitOutputBlocked", challengerRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitOutputBlocked)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitOutputBlocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitOutputBlocked is a log parse operation binding the contract event 0xcbe8dad2e7fcbfe0dcba2f9b2e44f122c66cd26dc0808a0f7e9ec41e4fe285bf.
//
// Solidity: event InFlightExitOutputBlocked(address indexed challenger, bytes32 indexed txHash, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitOutputBlocked(log types.Log) (*PaymentExitgameInFlightExitOutputBlocked, error) {
	event := new(PaymentExitgameInFlightExitOutputBlocked)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitOutputBlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitOutputPiggybackedIterator is returned from FilterInFlightExitOutputPiggybacked and is used to iterate over the raw logs and unpacked data for InFlightExitOutputPiggybacked events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitOutputPiggybackedIterator struct {
	Event *PaymentExitgameInFlightExitOutputPiggybacked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitOutputPiggybackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitOutputPiggybacked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitOutputPiggybacked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitOutputPiggybackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitOutputPiggybackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitOutputPiggybacked represents a InFlightExitOutputPiggybacked event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitOutputPiggybacked struct {
	ExitTarget  common.Address
	TxHash      [32]byte
	OutputIndex uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOutputPiggybacked is a free log retrieval operation binding the contract event 0x6ecd8e79a5f67f6c12b54371ada2ffb41bc128c61d9ac1e969f0aa2aca46cd78.
//
// Solidity: event InFlightExitOutputPiggybacked(address indexed exitTarget, bytes32 indexed txHash, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitOutputPiggybacked(opts *bind.FilterOpts, exitTarget []common.Address, txHash [][32]byte) (*PaymentExitgameInFlightExitOutputPiggybackedIterator, error) {

	var exitTargetRule []interface{}
	for _, exitTargetItem := range exitTarget {
		exitTargetRule = append(exitTargetRule, exitTargetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitOutputPiggybacked", exitTargetRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitOutputPiggybackedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitOutputPiggybacked", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOutputPiggybacked is a free log subscription operation binding the contract event 0x6ecd8e79a5f67f6c12b54371ada2ffb41bc128c61d9ac1e969f0aa2aca46cd78.
//
// Solidity: event InFlightExitOutputPiggybacked(address indexed exitTarget, bytes32 indexed txHash, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitOutputPiggybacked(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitOutputPiggybacked, exitTarget []common.Address, txHash [][32]byte) (event.Subscription, error) {

	var exitTargetRule []interface{}
	for _, exitTargetItem := range exitTarget {
		exitTargetRule = append(exitTargetRule, exitTargetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitOutputPiggybacked", exitTargetRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitOutputPiggybacked)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitOutputPiggybacked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitOutputPiggybacked is a log parse operation binding the contract event 0x6ecd8e79a5f67f6c12b54371ada2ffb41bc128c61d9ac1e969f0aa2aca46cd78.
//
// Solidity: event InFlightExitOutputPiggybacked(address indexed exitTarget, bytes32 indexed txHash, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitOutputPiggybacked(log types.Log) (*PaymentExitgameInFlightExitOutputPiggybacked, error) {
	event := new(PaymentExitgameInFlightExitOutputPiggybacked)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitOutputPiggybacked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitOutputWithdrawnIterator is returned from FilterInFlightExitOutputWithdrawn and is used to iterate over the raw logs and unpacked data for InFlightExitOutputWithdrawn events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitOutputWithdrawnIterator struct {
	Event *PaymentExitgameInFlightExitOutputWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitOutputWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitOutputWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitOutputWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitOutputWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitOutputWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitOutputWithdrawn represents a InFlightExitOutputWithdrawn event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitOutputWithdrawn struct {
	ExitId      *big.Int
	OutputIndex uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitOutputWithdrawn is a free log retrieval operation binding the contract event 0xa241c6deaf193e53a1b002d779e4f247bf5d57ba0be5a753e628dfcee645a4f7.
//
// Solidity: event InFlightExitOutputWithdrawn(uint160 indexed exitId, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitOutputWithdrawn(opts *bind.FilterOpts, exitId []*big.Int) (*PaymentExitgameInFlightExitOutputWithdrawnIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitOutputWithdrawn", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitOutputWithdrawnIterator{contract: _PaymentExitgame.contract, event: "InFlightExitOutputWithdrawn", logs: logs, sub: sub}, nil
}

// WatchInFlightExitOutputWithdrawn is a free log subscription operation binding the contract event 0xa241c6deaf193e53a1b002d779e4f247bf5d57ba0be5a753e628dfcee645a4f7.
//
// Solidity: event InFlightExitOutputWithdrawn(uint160 indexed exitId, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitOutputWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitOutputWithdrawn, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitOutputWithdrawn", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitOutputWithdrawn)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitOutputWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitOutputWithdrawn is a log parse operation binding the contract event 0xa241c6deaf193e53a1b002d779e4f247bf5d57ba0be5a753e628dfcee645a4f7.
//
// Solidity: event InFlightExitOutputWithdrawn(uint160 indexed exitId, uint16 outputIndex)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitOutputWithdrawn(log types.Log) (*PaymentExitgameInFlightExitOutputWithdrawn, error) {
	event := new(PaymentExitgameInFlightExitOutputWithdrawn)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitOutputWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameInFlightExitStartedIterator is returned from FilterInFlightExitStarted and is used to iterate over the raw logs and unpacked data for InFlightExitStarted events raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitStartedIterator struct {
	Event *PaymentExitgameInFlightExitStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameInFlightExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameInFlightExitStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameInFlightExitStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameInFlightExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameInFlightExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameInFlightExitStarted represents a InFlightExitStarted event raised by the PaymentExitgame contract.
type PaymentExitgameInFlightExitStarted struct {
	Initiator common.Address
	TxHash    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInFlightExitStarted is a free log retrieval operation binding the contract event 0xd5f1fe9d48880b57daa227004b16d320c0eb885d6c49d472d54c16a05fa3179e.
//
// Solidity: event InFlightExitStarted(address indexed initiator, bytes32 indexed txHash)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterInFlightExitStarted(opts *bind.FilterOpts, initiator []common.Address, txHash [][32]byte) (*PaymentExitgameInFlightExitStartedIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "InFlightExitStarted", initiatorRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameInFlightExitStartedIterator{contract: _PaymentExitgame.contract, event: "InFlightExitStarted", logs: logs, sub: sub}, nil
}

// WatchInFlightExitStarted is a free log subscription operation binding the contract event 0xd5f1fe9d48880b57daa227004b16d320c0eb885d6c49d472d54c16a05fa3179e.
//
// Solidity: event InFlightExitStarted(address indexed initiator, bytes32 indexed txHash)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchInFlightExitStarted(opts *bind.WatchOpts, sink chan<- *PaymentExitgameInFlightExitStarted, initiator []common.Address, txHash [][32]byte) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "InFlightExitStarted", initiatorRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameInFlightExitStarted)
				if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInFlightExitStarted is a log parse operation binding the contract event 0xd5f1fe9d48880b57daa227004b16d320c0eb885d6c49d472d54c16a05fa3179e.
//
// Solidity: event InFlightExitStarted(address indexed initiator, bytes32 indexed txHash)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseInFlightExitStarted(log types.Log) (*PaymentExitgameInFlightExitStarted, error) {
	event := new(PaymentExitgameInFlightExitStarted)
	if err := _PaymentExitgame.contract.UnpackLog(event, "InFlightExitStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgamePiggybackBondUpdatedIterator is returned from FilterPiggybackBondUpdated and is used to iterate over the raw logs and unpacked data for PiggybackBondUpdated events raised by the PaymentExitgame contract.
type PaymentExitgamePiggybackBondUpdatedIterator struct {
	Event *PaymentExitgamePiggybackBondUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgamePiggybackBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgamePiggybackBondUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgamePiggybackBondUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgamePiggybackBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgamePiggybackBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgamePiggybackBondUpdated represents a PiggybackBondUpdated event raised by the PaymentExitgame contract.
type PaymentExitgamePiggybackBondUpdated struct {
	BondSize *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPiggybackBondUpdated is a free log retrieval operation binding the contract event 0x30ce291291dca42b8168ab211fafc36d3d4868292ed3259874ed00864432e20e.
//
// Solidity: event PiggybackBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterPiggybackBondUpdated(opts *bind.FilterOpts) (*PaymentExitgamePiggybackBondUpdatedIterator, error) {

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "PiggybackBondUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentExitgamePiggybackBondUpdatedIterator{contract: _PaymentExitgame.contract, event: "PiggybackBondUpdated", logs: logs, sub: sub}, nil
}

// WatchPiggybackBondUpdated is a free log subscription operation binding the contract event 0x30ce291291dca42b8168ab211fafc36d3d4868292ed3259874ed00864432e20e.
//
// Solidity: event PiggybackBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchPiggybackBondUpdated(opts *bind.WatchOpts, sink chan<- *PaymentExitgamePiggybackBondUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "PiggybackBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgamePiggybackBondUpdated)
				if err := _PaymentExitgame.contract.UnpackLog(event, "PiggybackBondUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePiggybackBondUpdated is a log parse operation binding the contract event 0x30ce291291dca42b8168ab211fafc36d3d4868292ed3259874ed00864432e20e.
//
// Solidity: event PiggybackBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) ParsePiggybackBondUpdated(log types.Log) (*PaymentExitgamePiggybackBondUpdated, error) {
	event := new(PaymentExitgamePiggybackBondUpdated)
	if err := _PaymentExitgame.contract.UnpackLog(event, "PiggybackBondUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentExitgameStandardExitBondUpdatedIterator is returned from FilterStandardExitBondUpdated and is used to iterate over the raw logs and unpacked data for StandardExitBondUpdated events raised by the PaymentExitgame contract.
type PaymentExitgameStandardExitBondUpdatedIterator struct {
	Event *PaymentExitgameStandardExitBondUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentExitgameStandardExitBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentExitgameStandardExitBondUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentExitgameStandardExitBondUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentExitgameStandardExitBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentExitgameStandardExitBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentExitgameStandardExitBondUpdated represents a StandardExitBondUpdated event raised by the PaymentExitgame contract.
type PaymentExitgameStandardExitBondUpdated struct {
	BondSize *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStandardExitBondUpdated is a free log retrieval operation binding the contract event 0xd21ccbbae3d29826d53310efdc993c64e27496dd31694b292f1a120344ec37ab.
//
// Solidity: event StandardExitBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) FilterStandardExitBondUpdated(opts *bind.FilterOpts) (*PaymentExitgameStandardExitBondUpdatedIterator, error) {

	logs, sub, err := _PaymentExitgame.contract.FilterLogs(opts, "StandardExitBondUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentExitgameStandardExitBondUpdatedIterator{contract: _PaymentExitgame.contract, event: "StandardExitBondUpdated", logs: logs, sub: sub}, nil
}

// WatchStandardExitBondUpdated is a free log subscription operation binding the contract event 0xd21ccbbae3d29826d53310efdc993c64e27496dd31694b292f1a120344ec37ab.
//
// Solidity: event StandardExitBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) WatchStandardExitBondUpdated(opts *bind.WatchOpts, sink chan<- *PaymentExitgameStandardExitBondUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentExitgame.contract.WatchLogs(opts, "StandardExitBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentExitgameStandardExitBondUpdated)
				if err := _PaymentExitgame.contract.UnpackLog(event, "StandardExitBondUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStandardExitBondUpdated is a log parse operation binding the contract event 0xd21ccbbae3d29826d53310efdc993c64e27496dd31694b292f1a120344ec37ab.
//
// Solidity: event StandardExitBondUpdated(uint128 bondSize)
func (_PaymentExitgame *PaymentExitgameFilterer) ParseStandardExitBondUpdated(log types.Log) (*PaymentExitgameStandardExitBondUpdated, error) {
	event := new(PaymentExitgameStandardExitBondUpdated)
	if err := _PaymentExitgame.contract.UnpackLog(event, "StandardExitBondUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
