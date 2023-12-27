// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenmanager

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TokenmanagerMetaData contains all meta data concerning the Tokenmanager contract.
var TokenmanagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BEP2_TOKEN_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_PACKAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_STATUS_ALREADY_BOUND_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_STATUS_DECIMALS_MISMATCH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_STATUS_REJECTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_STATUS_SYMBOL_MISMATCH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_STATUS_TIMEOUT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_STATUS_TOO_MUCH_TOKENHUB_BALANCE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_STATUS_TOTAL_SUPPLY_MISMATCH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CODE_OK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_STAKE_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_FAIL_DECODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_TOKEN_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INCENTIVIZE_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LOG_MAX_UINT256\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_BEP20_SYMBOL_LEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BEP2_TOTAL_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_GAS_FOR_TRANSFER_BNB\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINIMUM_BEP20_SYMBOL_LEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIRROR_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIRROR_STATUS_ALREADY_BOUND\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIRROR_STATUS_DUPLICATED_BEP2_SYMBOL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIRROR_STATUS_TIMEOUT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_CREDIT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYNC_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYNC_STATUS_NOT_BOUND_MIRROR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYNC_STATUS_TIMEOUT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYSTEM_REWARD_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TEN_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TIMELOCK_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_MANAGER_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_RECOVER_PORTAL_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_IN_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_OUT_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNBIND_PACKAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"alreadyInit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveBind\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bep2Symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bindPackageRecord\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"packageType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"peggyAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bep20Decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"expireTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"boundByMirror\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bscChainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"expireBind\",\"inputs\":[{\"name\":\"bep2Symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"handleAckPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleFailAckPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleSynPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mirror\",\"inputs\":[{\"name\":\"bep20Addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expireTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"mirrorFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mirrorPendingRecord\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryRequiredLockAmountForBind\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rejectBind\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bep2Symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sync\",\"inputs\":[{\"name\":\"bep20Addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expireTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"syncFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"bindFailure\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bep2Symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"failedReason\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"bindSuccess\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bep2Symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"peggyAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"mirrorFailure\",\"inputs\":[{\"name\":\"bep20Addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"errCode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"mirrorSuccess\",\"inputs\":[{\"name\":\"bep20Addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bep2Symbol\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"paramChange\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"syncFailure\",\"inputs\":[{\"name\":\"bep20Addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"errCode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"syncSuccess\",\"inputs\":[{\"name\":\"bep20Addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"unexpectedPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// TokenmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenmanagerMetaData.ABI instead.
var TokenmanagerABI = TokenmanagerMetaData.ABI

// Tokenmanager is an auto generated Go binding around an Ethereum contract.
type Tokenmanager struct {
	TokenmanagerCaller     // Read-only binding to the contract
	TokenmanagerTransactor // Write-only binding to the contract
	TokenmanagerFilterer   // Log filterer for contract events
}

// TokenmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenmanagerSession struct {
	Contract     *Tokenmanager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenmanagerCallerSession struct {
	Contract *TokenmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenmanagerTransactorSession struct {
	Contract     *TokenmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenmanagerRaw struct {
	Contract *Tokenmanager // Generic contract binding to access the raw methods on
}

// TokenmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenmanagerCallerRaw struct {
	Contract *TokenmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenmanagerTransactorRaw struct {
	Contract *TokenmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenmanager creates a new instance of Tokenmanager, bound to a specific deployed contract.
func NewTokenmanager(address common.Address, backend bind.ContractBackend) (*Tokenmanager, error) {
	contract, err := bindTokenmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenmanager{TokenmanagerCaller: TokenmanagerCaller{contract: contract}, TokenmanagerTransactor: TokenmanagerTransactor{contract: contract}, TokenmanagerFilterer: TokenmanagerFilterer{contract: contract}}, nil
}

// NewTokenmanagerCaller creates a new read-only instance of Tokenmanager, bound to a specific deployed contract.
func NewTokenmanagerCaller(address common.Address, caller bind.ContractCaller) (*TokenmanagerCaller, error) {
	contract, err := bindTokenmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerCaller{contract: contract}, nil
}

// NewTokenmanagerTransactor creates a new write-only instance of Tokenmanager, bound to a specific deployed contract.
func NewTokenmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenmanagerTransactor, error) {
	contract, err := bindTokenmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerTransactor{contract: contract}, nil
}

// NewTokenmanagerFilterer creates a new log filterer instance of Tokenmanager, bound to a specific deployed contract.
func NewTokenmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenmanagerFilterer, error) {
	contract, err := bindTokenmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerFilterer{contract: contract}, nil
}

// bindTokenmanager binds a generic wrapper to an already deployed contract.
func bindTokenmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenmanager *TokenmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenmanager.Contract.TokenmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenmanager *TokenmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenmanager.Contract.TokenmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenmanager *TokenmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenmanager.Contract.TokenmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenmanager *TokenmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenmanager *TokenmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenmanager *TokenmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenmanager.Contract.contract.Transact(opts, method, params...)
}

// BEP2TOKENDECIMALS is a free data retrieval call binding the contract method 0x61368475.
//
// Solidity: function BEP2_TOKEN_DECIMALS() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BEP2TOKENDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BEP2_TOKEN_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BEP2TOKENDECIMALS is a free data retrieval call binding the contract method 0x61368475.
//
// Solidity: function BEP2_TOKEN_DECIMALS() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BEP2TOKENDECIMALS() (uint8, error) {
	return _Tokenmanager.Contract.BEP2TOKENDECIMALS(&_Tokenmanager.CallOpts)
}

// BEP2TOKENDECIMALS is a free data retrieval call binding the contract method 0x61368475.
//
// Solidity: function BEP2_TOKEN_DECIMALS() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BEP2TOKENDECIMALS() (uint8, error) {
	return _Tokenmanager.Contract.BEP2TOKENDECIMALS(&_Tokenmanager.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.BINDCHANNELID(&_Tokenmanager.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.BINDCHANNELID(&_Tokenmanager.CallOpts)
}

// BINDPACKAGE is a free data retrieval call binding the contract method 0xfe3a2af5.
//
// Solidity: function BIND_PACKAGE() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDPACKAGE is a free data retrieval call binding the contract method 0xfe3a2af5.
//
// Solidity: function BIND_PACKAGE() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDPACKAGE() (uint8, error) {
	return _Tokenmanager.Contract.BINDPACKAGE(&_Tokenmanager.CallOpts)
}

// BINDPACKAGE is a free data retrieval call binding the contract method 0xfe3a2af5.
//
// Solidity: function BIND_PACKAGE() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDPACKAGE() (uint8, error) {
	return _Tokenmanager.Contract.BINDPACKAGE(&_Tokenmanager.CallOpts)
}

// BINDSTATUSALREADYBOUNDTOKEN is a free data retrieval call binding the contract method 0x0f212b1b.
//
// Solidity: function BIND_STATUS_ALREADY_BOUND_TOKEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDSTATUSALREADYBOUNDTOKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_STATUS_ALREADY_BOUND_TOKEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDSTATUSALREADYBOUNDTOKEN is a free data retrieval call binding the contract method 0x0f212b1b.
//
// Solidity: function BIND_STATUS_ALREADY_BOUND_TOKEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDSTATUSALREADYBOUNDTOKEN() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSALREADYBOUNDTOKEN(&_Tokenmanager.CallOpts)
}

// BINDSTATUSALREADYBOUNDTOKEN is a free data retrieval call binding the contract method 0x0f212b1b.
//
// Solidity: function BIND_STATUS_ALREADY_BOUND_TOKEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDSTATUSALREADYBOUNDTOKEN() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSALREADYBOUNDTOKEN(&_Tokenmanager.CallOpts)
}

// BINDSTATUSDECIMALSMISMATCH is a free data retrieval call binding the contract method 0x4bc81c00.
//
// Solidity: function BIND_STATUS_DECIMALS_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDSTATUSDECIMALSMISMATCH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_STATUS_DECIMALS_MISMATCH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDSTATUSDECIMALSMISMATCH is a free data retrieval call binding the contract method 0x4bc81c00.
//
// Solidity: function BIND_STATUS_DECIMALS_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDSTATUSDECIMALSMISMATCH() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSDECIMALSMISMATCH(&_Tokenmanager.CallOpts)
}

// BINDSTATUSDECIMALSMISMATCH is a free data retrieval call binding the contract method 0x4bc81c00.
//
// Solidity: function BIND_STATUS_DECIMALS_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDSTATUSDECIMALSMISMATCH() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSDECIMALSMISMATCH(&_Tokenmanager.CallOpts)
}

// BINDSTATUSREJECTED is a free data retrieval call binding the contract method 0x95b9ad26.
//
// Solidity: function BIND_STATUS_REJECTED() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDSTATUSREJECTED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_STATUS_REJECTED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDSTATUSREJECTED is a free data retrieval call binding the contract method 0x95b9ad26.
//
// Solidity: function BIND_STATUS_REJECTED() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDSTATUSREJECTED() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSREJECTED(&_Tokenmanager.CallOpts)
}

// BINDSTATUSREJECTED is a free data retrieval call binding the contract method 0x95b9ad26.
//
// Solidity: function BIND_STATUS_REJECTED() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDSTATUSREJECTED() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSREJECTED(&_Tokenmanager.CallOpts)
}

// BINDSTATUSSYMBOLMISMATCH is a free data retrieval call binding the contract method 0x5f558f86.
//
// Solidity: function BIND_STATUS_SYMBOL_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDSTATUSSYMBOLMISMATCH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_STATUS_SYMBOL_MISMATCH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDSTATUSSYMBOLMISMATCH is a free data retrieval call binding the contract method 0x5f558f86.
//
// Solidity: function BIND_STATUS_SYMBOL_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDSTATUSSYMBOLMISMATCH() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSSYMBOLMISMATCH(&_Tokenmanager.CallOpts)
}

// BINDSTATUSSYMBOLMISMATCH is a free data retrieval call binding the contract method 0x5f558f86.
//
// Solidity: function BIND_STATUS_SYMBOL_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDSTATUSSYMBOLMISMATCH() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSSYMBOLMISMATCH(&_Tokenmanager.CallOpts)
}

// BINDSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x7d078e13.
//
// Solidity: function BIND_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDSTATUSTIMEOUT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_STATUS_TIMEOUT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x7d078e13.
//
// Solidity: function BIND_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDSTATUSTIMEOUT() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSTIMEOUT(&_Tokenmanager.CallOpts)
}

// BINDSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x7d078e13.
//
// Solidity: function BIND_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDSTATUSTIMEOUT() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSTIMEOUT(&_Tokenmanager.CallOpts)
}

// BINDSTATUSTOOMUCHTOKENHUBBALANCE is a free data retrieval call binding the contract method 0xc8e704a4.
//
// Solidity: function BIND_STATUS_TOO_MUCH_TOKENHUB_BALANCE() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDSTATUSTOOMUCHTOKENHUBBALANCE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_STATUS_TOO_MUCH_TOKENHUB_BALANCE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDSTATUSTOOMUCHTOKENHUBBALANCE is a free data retrieval call binding the contract method 0xc8e704a4.
//
// Solidity: function BIND_STATUS_TOO_MUCH_TOKENHUB_BALANCE() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDSTATUSTOOMUCHTOKENHUBBALANCE() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSTOOMUCHTOKENHUBBALANCE(&_Tokenmanager.CallOpts)
}

// BINDSTATUSTOOMUCHTOKENHUBBALANCE is a free data retrieval call binding the contract method 0xc8e704a4.
//
// Solidity: function BIND_STATUS_TOO_MUCH_TOKENHUB_BALANCE() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDSTATUSTOOMUCHTOKENHUBBALANCE() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSTOOMUCHTOKENHUBBALANCE(&_Tokenmanager.CallOpts)
}

// BINDSTATUSTOTALSUPPLYMISMATCH is a free data retrieval call binding the contract method 0x1f91600b.
//
// Solidity: function BIND_STATUS_TOTAL_SUPPLY_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) BINDSTATUSTOTALSUPPLYMISMATCH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "BIND_STATUS_TOTAL_SUPPLY_MISMATCH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDSTATUSTOTALSUPPLYMISMATCH is a free data retrieval call binding the contract method 0x1f91600b.
//
// Solidity: function BIND_STATUS_TOTAL_SUPPLY_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) BINDSTATUSTOTALSUPPLYMISMATCH() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSTOTALSUPPLYMISMATCH(&_Tokenmanager.CallOpts)
}

// BINDSTATUSTOTALSUPPLYMISMATCH is a free data retrieval call binding the contract method 0x1f91600b.
//
// Solidity: function BIND_STATUS_TOTAL_SUPPLY_MISMATCH() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) BINDSTATUSTOTALSUPPLYMISMATCH() (uint8, error) {
	return _Tokenmanager.Contract.BINDSTATUSTOTALSUPPLYMISMATCH(&_Tokenmanager.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tokenmanager *TokenmanagerCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tokenmanager *TokenmanagerSession) CODEOK() (uint32, error) {
	return _Tokenmanager.Contract.CODEOK(&_Tokenmanager.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tokenmanager *TokenmanagerCallerSession) CODEOK() (uint32, error) {
	return _Tokenmanager.Contract.CODEOK(&_Tokenmanager.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.CROSSCHAINCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.CROSSCHAINCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) CROSSSTAKECHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "CROSS_STAKE_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.CROSSSTAKECHANNELID(&_Tokenmanager.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.CROSSSTAKECHANNELID(&_Tokenmanager.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tokenmanager *TokenmanagerCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tokenmanager *TokenmanagerSession) ERRORFAILDECODE() (uint32, error) {
	return _Tokenmanager.Contract.ERRORFAILDECODE(&_Tokenmanager.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tokenmanager *TokenmanagerCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Tokenmanager.Contract.ERRORFAILDECODE(&_Tokenmanager.CallOpts)
}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) GOVERNORADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "GOVERNOR_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) GOVERNORADDR() (common.Address, error) {
	return _Tokenmanager.Contract.GOVERNORADDR(&_Tokenmanager.CallOpts)
}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) GOVERNORADDR() (common.Address, error) {
	return _Tokenmanager.Contract.GOVERNORADDR(&_Tokenmanager.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) GOVCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.GOVCHANNELID(&_Tokenmanager.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) GOVCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.GOVCHANNELID(&_Tokenmanager.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) GOVHUBADDR() (common.Address, error) {
	return _Tokenmanager.Contract.GOVHUBADDR(&_Tokenmanager.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Tokenmanager.Contract.GOVHUBADDR(&_Tokenmanager.CallOpts)
}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) GOVTOKENADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "GOV_TOKEN_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) GOVTOKENADDR() (common.Address, error) {
	return _Tokenmanager.Contract.GOVTOKENADDR(&_Tokenmanager.CallOpts)
}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) GOVTOKENADDR() (common.Address, error) {
	return _Tokenmanager.Contract.GOVTOKENADDR(&_Tokenmanager.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Tokenmanager.Contract.INCENTIVIZEADDR(&_Tokenmanager.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Tokenmanager.Contract.INCENTIVIZEADDR(&_Tokenmanager.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.LIGHTCLIENTADDR(&_Tokenmanager.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.LIGHTCLIENTADDR(&_Tokenmanager.CallOpts)
}

// LOGMAXUINT256 is a free data retrieval call binding the contract method 0xd7109ce6.
//
// Solidity: function LOG_MAX_UINT256() view returns(uint256)
func (_Tokenmanager *TokenmanagerCaller) LOGMAXUINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "LOG_MAX_UINT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOGMAXUINT256 is a free data retrieval call binding the contract method 0xd7109ce6.
//
// Solidity: function LOG_MAX_UINT256() view returns(uint256)
func (_Tokenmanager *TokenmanagerSession) LOGMAXUINT256() (*big.Int, error) {
	return _Tokenmanager.Contract.LOGMAXUINT256(&_Tokenmanager.CallOpts)
}

// LOGMAXUINT256 is a free data retrieval call binding the contract method 0xd7109ce6.
//
// Solidity: function LOG_MAX_UINT256() view returns(uint256)
func (_Tokenmanager *TokenmanagerCallerSession) LOGMAXUINT256() (*big.Int, error) {
	return _Tokenmanager.Contract.LOGMAXUINT256(&_Tokenmanager.CallOpts)
}

// MAXIMUMBEP20SYMBOLLEN is a free data retrieval call binding the contract method 0xd9e6dae9.
//
// Solidity: function MAXIMUM_BEP20_SYMBOL_LEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) MAXIMUMBEP20SYMBOLLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "MAXIMUM_BEP20_SYMBOL_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMBEP20SYMBOLLEN is a free data retrieval call binding the contract method 0xd9e6dae9.
//
// Solidity: function MAXIMUM_BEP20_SYMBOL_LEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) MAXIMUMBEP20SYMBOLLEN() (uint8, error) {
	return _Tokenmanager.Contract.MAXIMUMBEP20SYMBOLLEN(&_Tokenmanager.CallOpts)
}

// MAXIMUMBEP20SYMBOLLEN is a free data retrieval call binding the contract method 0xd9e6dae9.
//
// Solidity: function MAXIMUM_BEP20_SYMBOL_LEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) MAXIMUMBEP20SYMBOLLEN() (uint8, error) {
	return _Tokenmanager.Contract.MAXIMUMBEP20SYMBOLLEN(&_Tokenmanager.CallOpts)
}

// MAXBEP2TOTALSUPPLY is a free data retrieval call binding the contract method 0x9a854bbd.
//
// Solidity: function MAX_BEP2_TOTAL_SUPPLY() view returns(uint256)
func (_Tokenmanager *TokenmanagerCaller) MAXBEP2TOTALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "MAX_BEP2_TOTAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBEP2TOTALSUPPLY is a free data retrieval call binding the contract method 0x9a854bbd.
//
// Solidity: function MAX_BEP2_TOTAL_SUPPLY() view returns(uint256)
func (_Tokenmanager *TokenmanagerSession) MAXBEP2TOTALSUPPLY() (*big.Int, error) {
	return _Tokenmanager.Contract.MAXBEP2TOTALSUPPLY(&_Tokenmanager.CallOpts)
}

// MAXBEP2TOTALSUPPLY is a free data retrieval call binding the contract method 0x9a854bbd.
//
// Solidity: function MAX_BEP2_TOTAL_SUPPLY() view returns(uint256)
func (_Tokenmanager *TokenmanagerCallerSession) MAXBEP2TOTALSUPPLY() (*big.Int, error) {
	return _Tokenmanager.Contract.MAXBEP2TOTALSUPPLY(&_Tokenmanager.CallOpts)
}

// MAXGASFORTRANSFERBNB is a free data retrieval call binding the contract method 0xfa9e9159.
//
// Solidity: function MAX_GAS_FOR_TRANSFER_BNB() view returns(uint256)
func (_Tokenmanager *TokenmanagerCaller) MAXGASFORTRANSFERBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "MAX_GAS_FOR_TRANSFER_BNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXGASFORTRANSFERBNB is a free data retrieval call binding the contract method 0xfa9e9159.
//
// Solidity: function MAX_GAS_FOR_TRANSFER_BNB() view returns(uint256)
func (_Tokenmanager *TokenmanagerSession) MAXGASFORTRANSFERBNB() (*big.Int, error) {
	return _Tokenmanager.Contract.MAXGASFORTRANSFERBNB(&_Tokenmanager.CallOpts)
}

// MAXGASFORTRANSFERBNB is a free data retrieval call binding the contract method 0xfa9e9159.
//
// Solidity: function MAX_GAS_FOR_TRANSFER_BNB() view returns(uint256)
func (_Tokenmanager *TokenmanagerCallerSession) MAXGASFORTRANSFERBNB() (*big.Int, error) {
	return _Tokenmanager.Contract.MAXGASFORTRANSFERBNB(&_Tokenmanager.CallOpts)
}

// MINIMUMBEP20SYMBOLLEN is a free data retrieval call binding the contract method 0x66dea52a.
//
// Solidity: function MINIMUM_BEP20_SYMBOL_LEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) MINIMUMBEP20SYMBOLLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "MINIMUM_BEP20_SYMBOL_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINIMUMBEP20SYMBOLLEN is a free data retrieval call binding the contract method 0x66dea52a.
//
// Solidity: function MINIMUM_BEP20_SYMBOL_LEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) MINIMUMBEP20SYMBOLLEN() (uint8, error) {
	return _Tokenmanager.Contract.MINIMUMBEP20SYMBOLLEN(&_Tokenmanager.CallOpts)
}

// MINIMUMBEP20SYMBOLLEN is a free data retrieval call binding the contract method 0x66dea52a.
//
// Solidity: function MINIMUM_BEP20_SYMBOL_LEN() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) MINIMUMBEP20SYMBOLLEN() (uint8, error) {
	return _Tokenmanager.Contract.MINIMUMBEP20SYMBOLLEN(&_Tokenmanager.CallOpts)
}

// MIRRORCHANNELID is a free data retrieval call binding the contract method 0xbd32d3f9.
//
// Solidity: function MIRROR_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) MIRRORCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "MIRROR_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MIRRORCHANNELID is a free data retrieval call binding the contract method 0xbd32d3f9.
//
// Solidity: function MIRROR_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) MIRRORCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.MIRRORCHANNELID(&_Tokenmanager.CallOpts)
}

// MIRRORCHANNELID is a free data retrieval call binding the contract method 0xbd32d3f9.
//
// Solidity: function MIRROR_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) MIRRORCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.MIRRORCHANNELID(&_Tokenmanager.CallOpts)
}

// MIRRORSTATUSALREADYBOUND is a free data retrieval call binding the contract method 0x401809f9.
//
// Solidity: function MIRROR_STATUS_ALREADY_BOUND() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) MIRRORSTATUSALREADYBOUND(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "MIRROR_STATUS_ALREADY_BOUND")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MIRRORSTATUSALREADYBOUND is a free data retrieval call binding the contract method 0x401809f9.
//
// Solidity: function MIRROR_STATUS_ALREADY_BOUND() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) MIRRORSTATUSALREADYBOUND() (uint8, error) {
	return _Tokenmanager.Contract.MIRRORSTATUSALREADYBOUND(&_Tokenmanager.CallOpts)
}

// MIRRORSTATUSALREADYBOUND is a free data retrieval call binding the contract method 0x401809f9.
//
// Solidity: function MIRROR_STATUS_ALREADY_BOUND() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) MIRRORSTATUSALREADYBOUND() (uint8, error) {
	return _Tokenmanager.Contract.MIRRORSTATUSALREADYBOUND(&_Tokenmanager.CallOpts)
}

// MIRRORSTATUSDUPLICATEDBEP2SYMBOL is a free data retrieval call binding the contract method 0xab67a485.
//
// Solidity: function MIRROR_STATUS_DUPLICATED_BEP2_SYMBOL() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) MIRRORSTATUSDUPLICATEDBEP2SYMBOL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "MIRROR_STATUS_DUPLICATED_BEP2_SYMBOL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MIRRORSTATUSDUPLICATEDBEP2SYMBOL is a free data retrieval call binding the contract method 0xab67a485.
//
// Solidity: function MIRROR_STATUS_DUPLICATED_BEP2_SYMBOL() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) MIRRORSTATUSDUPLICATEDBEP2SYMBOL() (uint8, error) {
	return _Tokenmanager.Contract.MIRRORSTATUSDUPLICATEDBEP2SYMBOL(&_Tokenmanager.CallOpts)
}

// MIRRORSTATUSDUPLICATEDBEP2SYMBOL is a free data retrieval call binding the contract method 0xab67a485.
//
// Solidity: function MIRROR_STATUS_DUPLICATED_BEP2_SYMBOL() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) MIRRORSTATUSDUPLICATEDBEP2SYMBOL() (uint8, error) {
	return _Tokenmanager.Contract.MIRRORSTATUSDUPLICATEDBEP2SYMBOL(&_Tokenmanager.CallOpts)
}

// MIRRORSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x2d89ac32.
//
// Solidity: function MIRROR_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) MIRRORSTATUSTIMEOUT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "MIRROR_STATUS_TIMEOUT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MIRRORSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x2d89ac32.
//
// Solidity: function MIRROR_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) MIRRORSTATUSTIMEOUT() (uint8, error) {
	return _Tokenmanager.Contract.MIRRORSTATUSTIMEOUT(&_Tokenmanager.CallOpts)
}

// MIRRORSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x2d89ac32.
//
// Solidity: function MIRROR_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) MIRRORSTATUSTIMEOUT() (uint8, error) {
	return _Tokenmanager.Contract.MIRRORSTATUSTIMEOUT(&_Tokenmanager.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.RELAYERHUBCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.RELAYERHUBCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) SLASHCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.SLASHCHANNELID(&_Tokenmanager.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.SLASHCHANNELID(&_Tokenmanager.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.SLASHCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.SLASHCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) STAKECREDITADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "STAKE_CREDIT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) STAKECREDITADDR() (common.Address, error) {
	return _Tokenmanager.Contract.STAKECREDITADDR(&_Tokenmanager.CallOpts)
}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) STAKECREDITADDR() (common.Address, error) {
	return _Tokenmanager.Contract.STAKECREDITADDR(&_Tokenmanager.CallOpts)
}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) STAKEHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "STAKE_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) STAKEHUBADDR() (common.Address, error) {
	return _Tokenmanager.Contract.STAKEHUBADDR(&_Tokenmanager.CallOpts)
}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) STAKEHUBADDR() (common.Address, error) {
	return _Tokenmanager.Contract.STAKEHUBADDR(&_Tokenmanager.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) STAKINGCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.STAKINGCHANNELID(&_Tokenmanager.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.STAKINGCHANNELID(&_Tokenmanager.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) STAKINGCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "STAKING_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.STAKINGCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.STAKINGCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// SYNCCHANNELID is a free data retrieval call binding the contract method 0x859180fb.
//
// Solidity: function SYNC_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) SYNCCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "SYNC_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SYNCCHANNELID is a free data retrieval call binding the contract method 0x859180fb.
//
// Solidity: function SYNC_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) SYNCCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.SYNCCHANNELID(&_Tokenmanager.CallOpts)
}

// SYNCCHANNELID is a free data retrieval call binding the contract method 0x859180fb.
//
// Solidity: function SYNC_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) SYNCCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.SYNCCHANNELID(&_Tokenmanager.CallOpts)
}

// SYNCSTATUSNOTBOUNDMIRROR is a free data retrieval call binding the contract method 0xb7950317.
//
// Solidity: function SYNC_STATUS_NOT_BOUND_MIRROR() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) SYNCSTATUSNOTBOUNDMIRROR(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "SYNC_STATUS_NOT_BOUND_MIRROR")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SYNCSTATUSNOTBOUNDMIRROR is a free data retrieval call binding the contract method 0xb7950317.
//
// Solidity: function SYNC_STATUS_NOT_BOUND_MIRROR() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) SYNCSTATUSNOTBOUNDMIRROR() (uint8, error) {
	return _Tokenmanager.Contract.SYNCSTATUSNOTBOUNDMIRROR(&_Tokenmanager.CallOpts)
}

// SYNCSTATUSNOTBOUNDMIRROR is a free data retrieval call binding the contract method 0xb7950317.
//
// Solidity: function SYNC_STATUS_NOT_BOUND_MIRROR() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) SYNCSTATUSNOTBOUNDMIRROR() (uint8, error) {
	return _Tokenmanager.Contract.SYNCSTATUSNOTBOUNDMIRROR(&_Tokenmanager.CallOpts)
}

// SYNCSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x487c88ac.
//
// Solidity: function SYNC_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) SYNCSTATUSTIMEOUT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "SYNC_STATUS_TIMEOUT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SYNCSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x487c88ac.
//
// Solidity: function SYNC_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) SYNCSTATUSTIMEOUT() (uint8, error) {
	return _Tokenmanager.Contract.SYNCSTATUSTIMEOUT(&_Tokenmanager.CallOpts)
}

// SYNCSTATUSTIMEOUT is a free data retrieval call binding the contract method 0x487c88ac.
//
// Solidity: function SYNC_STATUS_TIMEOUT() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) SYNCSTATUSTIMEOUT() (uint8, error) {
	return _Tokenmanager.Contract.SYNCSTATUSTIMEOUT(&_Tokenmanager.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Tokenmanager.Contract.SYSTEMREWARDADDR(&_Tokenmanager.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Tokenmanager.Contract.SYSTEMREWARDADDR(&_Tokenmanager.CallOpts)
}

// TENDECIMALS is a free data retrieval call binding the contract method 0x5d499b1b.
//
// Solidity: function TEN_DECIMALS() view returns(uint256)
func (_Tokenmanager *TokenmanagerCaller) TENDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "TEN_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TENDECIMALS is a free data retrieval call binding the contract method 0x5d499b1b.
//
// Solidity: function TEN_DECIMALS() view returns(uint256)
func (_Tokenmanager *TokenmanagerSession) TENDECIMALS() (*big.Int, error) {
	return _Tokenmanager.Contract.TENDECIMALS(&_Tokenmanager.CallOpts)
}

// TENDECIMALS is a free data retrieval call binding the contract method 0x5d499b1b.
//
// Solidity: function TEN_DECIMALS() view returns(uint256)
func (_Tokenmanager *TokenmanagerCallerSession) TENDECIMALS() (*big.Int, error) {
	return _Tokenmanager.Contract.TENDECIMALS(&_Tokenmanager.CallOpts)
}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) TIMELOCKADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "TIMELOCK_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) TIMELOCKADDR() (common.Address, error) {
	return _Tokenmanager.Contract.TIMELOCKADDR(&_Tokenmanager.CallOpts)
}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) TIMELOCKADDR() (common.Address, error) {
	return _Tokenmanager.Contract.TIMELOCKADDR(&_Tokenmanager.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) TOKENHUBADDR() (common.Address, error) {
	return _Tokenmanager.Contract.TOKENHUBADDR(&_Tokenmanager.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Tokenmanager.Contract.TOKENHUBADDR(&_Tokenmanager.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Tokenmanager.Contract.TOKENMANAGERADDR(&_Tokenmanager.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Tokenmanager.Contract.TOKENMANAGERADDR(&_Tokenmanager.CallOpts)
}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) TOKENRECOVERPORTALADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "TOKEN_RECOVER_PORTAL_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) TOKENRECOVERPORTALADDR() (common.Address, error) {
	return _Tokenmanager.Contract.TOKENRECOVERPORTALADDR(&_Tokenmanager.CallOpts)
}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) TOKENRECOVERPORTALADDR() (common.Address, error) {
	return _Tokenmanager.Contract.TOKENRECOVERPORTALADDR(&_Tokenmanager.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.TRANSFERINCHANNELID(&_Tokenmanager.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.TRANSFERINCHANNELID(&_Tokenmanager.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.TRANSFEROUTCHANNELID(&_Tokenmanager.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Tokenmanager.Contract.TRANSFEROUTCHANNELID(&_Tokenmanager.CallOpts)
}

// UNBINDPACKAGE is a free data retrieval call binding the contract method 0x23996b53.
//
// Solidity: function UNBIND_PACKAGE() view returns(uint8)
func (_Tokenmanager *TokenmanagerCaller) UNBINDPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "UNBIND_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UNBINDPACKAGE is a free data retrieval call binding the contract method 0x23996b53.
//
// Solidity: function UNBIND_PACKAGE() view returns(uint8)
func (_Tokenmanager *TokenmanagerSession) UNBINDPACKAGE() (uint8, error) {
	return _Tokenmanager.Contract.UNBINDPACKAGE(&_Tokenmanager.CallOpts)
}

// UNBINDPACKAGE is a free data retrieval call binding the contract method 0x23996b53.
//
// Solidity: function UNBIND_PACKAGE() view returns(uint8)
func (_Tokenmanager *TokenmanagerCallerSession) UNBINDPACKAGE() (uint8, error) {
	return _Tokenmanager.Contract.UNBINDPACKAGE(&_Tokenmanager.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.VALIDATORCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Tokenmanager *TokenmanagerCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Tokenmanager.Contract.VALIDATORCONTRACTADDR(&_Tokenmanager.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Tokenmanager *TokenmanagerCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Tokenmanager *TokenmanagerSession) AlreadyInit() (bool, error) {
	return _Tokenmanager.Contract.AlreadyInit(&_Tokenmanager.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Tokenmanager *TokenmanagerCallerSession) AlreadyInit() (bool, error) {
	return _Tokenmanager.Contract.AlreadyInit(&_Tokenmanager.CallOpts)
}

// BindPackageRecord is a free data retrieval call binding the contract method 0xd117a110.
//
// Solidity: function bindPackageRecord(bytes32 ) view returns(uint8 packageType, bytes32 bep2TokenSymbol, address contractAddr, uint256 totalSupply, uint256 peggyAmount, uint8 bep20Decimals, uint64 expireTime)
func (_Tokenmanager *TokenmanagerCaller) BindPackageRecord(opts *bind.CallOpts, arg0 [32]byte) (struct {
	PackageType     uint8
	Bep2TokenSymbol [32]byte
	ContractAddr    common.Address
	TotalSupply     *big.Int
	PeggyAmount     *big.Int
	Bep20Decimals   uint8
	ExpireTime      uint64
}, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "bindPackageRecord", arg0)

	outstruct := new(struct {
		PackageType     uint8
		Bep2TokenSymbol [32]byte
		ContractAddr    common.Address
		TotalSupply     *big.Int
		PeggyAmount     *big.Int
		Bep20Decimals   uint8
		ExpireTime      uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PackageType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Bep2TokenSymbol = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ContractAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TotalSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PeggyAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Bep20Decimals = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.ExpireTime = *abi.ConvertType(out[6], new(uint64)).(*uint64)

	return *outstruct, err

}

// BindPackageRecord is a free data retrieval call binding the contract method 0xd117a110.
//
// Solidity: function bindPackageRecord(bytes32 ) view returns(uint8 packageType, bytes32 bep2TokenSymbol, address contractAddr, uint256 totalSupply, uint256 peggyAmount, uint8 bep20Decimals, uint64 expireTime)
func (_Tokenmanager *TokenmanagerSession) BindPackageRecord(arg0 [32]byte) (struct {
	PackageType     uint8
	Bep2TokenSymbol [32]byte
	ContractAddr    common.Address
	TotalSupply     *big.Int
	PeggyAmount     *big.Int
	Bep20Decimals   uint8
	ExpireTime      uint64
}, error) {
	return _Tokenmanager.Contract.BindPackageRecord(&_Tokenmanager.CallOpts, arg0)
}

// BindPackageRecord is a free data retrieval call binding the contract method 0xd117a110.
//
// Solidity: function bindPackageRecord(bytes32 ) view returns(uint8 packageType, bytes32 bep2TokenSymbol, address contractAddr, uint256 totalSupply, uint256 peggyAmount, uint8 bep20Decimals, uint64 expireTime)
func (_Tokenmanager *TokenmanagerCallerSession) BindPackageRecord(arg0 [32]byte) (struct {
	PackageType     uint8
	Bep2TokenSymbol [32]byte
	ContractAddr    common.Address
	TotalSupply     *big.Int
	PeggyAmount     *big.Int
	Bep20Decimals   uint8
	ExpireTime      uint64
}, error) {
	return _Tokenmanager.Contract.BindPackageRecord(&_Tokenmanager.CallOpts, arg0)
}

// BoundByMirror is a free data retrieval call binding the contract method 0x2e02d776.
//
// Solidity: function boundByMirror(address ) view returns(bool)
func (_Tokenmanager *TokenmanagerCaller) BoundByMirror(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "boundByMirror", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BoundByMirror is a free data retrieval call binding the contract method 0x2e02d776.
//
// Solidity: function boundByMirror(address ) view returns(bool)
func (_Tokenmanager *TokenmanagerSession) BoundByMirror(arg0 common.Address) (bool, error) {
	return _Tokenmanager.Contract.BoundByMirror(&_Tokenmanager.CallOpts, arg0)
}

// BoundByMirror is a free data retrieval call binding the contract method 0x2e02d776.
//
// Solidity: function boundByMirror(address ) view returns(bool)
func (_Tokenmanager *TokenmanagerCallerSession) BoundByMirror(arg0 common.Address) (bool, error) {
	return _Tokenmanager.Contract.BoundByMirror(&_Tokenmanager.CallOpts, arg0)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Tokenmanager *TokenmanagerCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Tokenmanager *TokenmanagerSession) BscChainID() (uint16, error) {
	return _Tokenmanager.Contract.BscChainID(&_Tokenmanager.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Tokenmanager *TokenmanagerCallerSession) BscChainID() (uint16, error) {
	return _Tokenmanager.Contract.BscChainID(&_Tokenmanager.CallOpts)
}

// MirrorFee is a free data retrieval call binding the contract method 0x7ec816dd.
//
// Solidity: function mirrorFee() view returns(uint256)
func (_Tokenmanager *TokenmanagerCaller) MirrorFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "mirrorFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MirrorFee is a free data retrieval call binding the contract method 0x7ec816dd.
//
// Solidity: function mirrorFee() view returns(uint256)
func (_Tokenmanager *TokenmanagerSession) MirrorFee() (*big.Int, error) {
	return _Tokenmanager.Contract.MirrorFee(&_Tokenmanager.CallOpts)
}

// MirrorFee is a free data retrieval call binding the contract method 0x7ec816dd.
//
// Solidity: function mirrorFee() view returns(uint256)
func (_Tokenmanager *TokenmanagerCallerSession) MirrorFee() (*big.Int, error) {
	return _Tokenmanager.Contract.MirrorFee(&_Tokenmanager.CallOpts)
}

// MirrorPendingRecord is a free data retrieval call binding the contract method 0x37e6ecda.
//
// Solidity: function mirrorPendingRecord(address ) view returns(bool)
func (_Tokenmanager *TokenmanagerCaller) MirrorPendingRecord(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "mirrorPendingRecord", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MirrorPendingRecord is a free data retrieval call binding the contract method 0x37e6ecda.
//
// Solidity: function mirrorPendingRecord(address ) view returns(bool)
func (_Tokenmanager *TokenmanagerSession) MirrorPendingRecord(arg0 common.Address) (bool, error) {
	return _Tokenmanager.Contract.MirrorPendingRecord(&_Tokenmanager.CallOpts, arg0)
}

// MirrorPendingRecord is a free data retrieval call binding the contract method 0x37e6ecda.
//
// Solidity: function mirrorPendingRecord(address ) view returns(bool)
func (_Tokenmanager *TokenmanagerCallerSession) MirrorPendingRecord(arg0 common.Address) (bool, error) {
	return _Tokenmanager.Contract.MirrorPendingRecord(&_Tokenmanager.CallOpts, arg0)
}

// QueryRequiredLockAmountForBind is a free data retrieval call binding the contract method 0x445fcefe.
//
// Solidity: function queryRequiredLockAmountForBind(string symbol) view returns(uint256)
func (_Tokenmanager *TokenmanagerCaller) QueryRequiredLockAmountForBind(opts *bind.CallOpts, symbol string) (*big.Int, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "queryRequiredLockAmountForBind", symbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryRequiredLockAmountForBind is a free data retrieval call binding the contract method 0x445fcefe.
//
// Solidity: function queryRequiredLockAmountForBind(string symbol) view returns(uint256)
func (_Tokenmanager *TokenmanagerSession) QueryRequiredLockAmountForBind(symbol string) (*big.Int, error) {
	return _Tokenmanager.Contract.QueryRequiredLockAmountForBind(&_Tokenmanager.CallOpts, symbol)
}

// QueryRequiredLockAmountForBind is a free data retrieval call binding the contract method 0x445fcefe.
//
// Solidity: function queryRequiredLockAmountForBind(string symbol) view returns(uint256)
func (_Tokenmanager *TokenmanagerCallerSession) QueryRequiredLockAmountForBind(symbol string) (*big.Int, error) {
	return _Tokenmanager.Contract.QueryRequiredLockAmountForBind(&_Tokenmanager.CallOpts, symbol)
}

// SyncFee is a free data retrieval call binding the contract method 0xe605bca0.
//
// Solidity: function syncFee() view returns(uint256)
func (_Tokenmanager *TokenmanagerCaller) SyncFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenmanager.contract.Call(opts, &out, "syncFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SyncFee is a free data retrieval call binding the contract method 0xe605bca0.
//
// Solidity: function syncFee() view returns(uint256)
func (_Tokenmanager *TokenmanagerSession) SyncFee() (*big.Int, error) {
	return _Tokenmanager.Contract.SyncFee(&_Tokenmanager.CallOpts)
}

// SyncFee is a free data retrieval call binding the contract method 0xe605bca0.
//
// Solidity: function syncFee() view returns(uint256)
func (_Tokenmanager *TokenmanagerCallerSession) SyncFee() (*big.Int, error) {
	return _Tokenmanager.Contract.SyncFee(&_Tokenmanager.CallOpts)
}

// ApproveBind is a paid mutator transaction binding the contract method 0x6b3f1307.
//
// Solidity: function approveBind(address contractAddr, string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactor) ApproveBind(opts *bind.TransactOpts, contractAddr common.Address, bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "approveBind", contractAddr, bep2Symbol)
}

// ApproveBind is a paid mutator transaction binding the contract method 0x6b3f1307.
//
// Solidity: function approveBind(address contractAddr, string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerSession) ApproveBind(contractAddr common.Address, bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.Contract.ApproveBind(&_Tokenmanager.TransactOpts, contractAddr, bep2Symbol)
}

// ApproveBind is a paid mutator transaction binding the contract method 0x6b3f1307.
//
// Solidity: function approveBind(address contractAddr, string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactorSession) ApproveBind(contractAddr common.Address, bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.Contract.ApproveBind(&_Tokenmanager.TransactOpts, contractAddr, bep2Symbol)
}

// ExpireBind is a paid mutator transaction binding the contract method 0x72c4e086.
//
// Solidity: function expireBind(string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactor) ExpireBind(opts *bind.TransactOpts, bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "expireBind", bep2Symbol)
}

// ExpireBind is a paid mutator transaction binding the contract method 0x72c4e086.
//
// Solidity: function expireBind(string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerSession) ExpireBind(bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.Contract.ExpireBind(&_Tokenmanager.TransactOpts, bep2Symbol)
}

// ExpireBind is a paid mutator transaction binding the contract method 0x72c4e086.
//
// Solidity: function expireBind(string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactorSession) ExpireBind(bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.Contract.ExpireBind(&_Tokenmanager.TransactOpts, bep2Symbol)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenmanager *TokenmanagerTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenmanager *TokenmanagerSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.Contract.HandleAckPackage(&_Tokenmanager.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenmanager *TokenmanagerTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.Contract.HandleAckPackage(&_Tokenmanager.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenmanager *TokenmanagerTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenmanager *TokenmanagerSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.Contract.HandleFailAckPackage(&_Tokenmanager.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenmanager *TokenmanagerTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.Contract.HandleFailAckPackage(&_Tokenmanager.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes)
func (_Tokenmanager *TokenmanagerTransactor) HandleSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "handleSynPackage", channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes)
func (_Tokenmanager *TokenmanagerSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.Contract.HandleSynPackage(&_Tokenmanager.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes)
func (_Tokenmanager *TokenmanagerTransactorSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenmanager.Contract.HandleSynPackage(&_Tokenmanager.TransactOpts, channelId, msgBytes)
}

// Mirror is a paid mutator transaction binding the contract method 0x94553a4e.
//
// Solidity: function mirror(address bep20Addr, uint64 expireTime) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactor) Mirror(opts *bind.TransactOpts, bep20Addr common.Address, expireTime uint64) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "mirror", bep20Addr, expireTime)
}

// Mirror is a paid mutator transaction binding the contract method 0x94553a4e.
//
// Solidity: function mirror(address bep20Addr, uint64 expireTime) payable returns(bool)
func (_Tokenmanager *TokenmanagerSession) Mirror(bep20Addr common.Address, expireTime uint64) (*types.Transaction, error) {
	return _Tokenmanager.Contract.Mirror(&_Tokenmanager.TransactOpts, bep20Addr, expireTime)
}

// Mirror is a paid mutator transaction binding the contract method 0x94553a4e.
//
// Solidity: function mirror(address bep20Addr, uint64 expireTime) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactorSession) Mirror(bep20Addr common.Address, expireTime uint64) (*types.Transaction, error) {
	return _Tokenmanager.Contract.Mirror(&_Tokenmanager.TransactOpts, bep20Addr, expireTime)
}

// RejectBind is a paid mutator transaction binding the contract method 0x77d9dae8.
//
// Solidity: function rejectBind(address contractAddr, string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactor) RejectBind(opts *bind.TransactOpts, contractAddr common.Address, bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "rejectBind", contractAddr, bep2Symbol)
}

// RejectBind is a paid mutator transaction binding the contract method 0x77d9dae8.
//
// Solidity: function rejectBind(address contractAddr, string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerSession) RejectBind(contractAddr common.Address, bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.Contract.RejectBind(&_Tokenmanager.TransactOpts, contractAddr, bep2Symbol)
}

// RejectBind is a paid mutator transaction binding the contract method 0x77d9dae8.
//
// Solidity: function rejectBind(address contractAddr, string bep2Symbol) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactorSession) RejectBind(contractAddr common.Address, bep2Symbol string) (*types.Transaction, error) {
	return _Tokenmanager.Contract.RejectBind(&_Tokenmanager.TransactOpts, contractAddr, bep2Symbol)
}

// Sync is a paid mutator transaction binding the contract method 0x25c751b7.
//
// Solidity: function sync(address bep20Addr, uint64 expireTime) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactor) Sync(opts *bind.TransactOpts, bep20Addr common.Address, expireTime uint64) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "sync", bep20Addr, expireTime)
}

// Sync is a paid mutator transaction binding the contract method 0x25c751b7.
//
// Solidity: function sync(address bep20Addr, uint64 expireTime) payable returns(bool)
func (_Tokenmanager *TokenmanagerSession) Sync(bep20Addr common.Address, expireTime uint64) (*types.Transaction, error) {
	return _Tokenmanager.Contract.Sync(&_Tokenmanager.TransactOpts, bep20Addr, expireTime)
}

// Sync is a paid mutator transaction binding the contract method 0x25c751b7.
//
// Solidity: function sync(address bep20Addr, uint64 expireTime) payable returns(bool)
func (_Tokenmanager *TokenmanagerTransactorSession) Sync(bep20Addr common.Address, expireTime uint64) (*types.Transaction, error) {
	return _Tokenmanager.Contract.Sync(&_Tokenmanager.TransactOpts, bep20Addr, expireTime)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenmanager *TokenmanagerTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Tokenmanager.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenmanager *TokenmanagerSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Tokenmanager.Contract.UpdateParam(&_Tokenmanager.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenmanager *TokenmanagerTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Tokenmanager.Contract.UpdateParam(&_Tokenmanager.TransactOpts, key, value)
}

// TokenmanagerBindFailureIterator is returned from FilterBindFailure and is used to iterate over the raw logs and unpacked data for BindFailure events raised by the Tokenmanager contract.
type TokenmanagerBindFailureIterator struct {
	Event *TokenmanagerBindFailure // Event containing the contract specifics and raw log

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
func (it *TokenmanagerBindFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmanagerBindFailure)
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
		it.Event = new(TokenmanagerBindFailure)
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
func (it *TokenmanagerBindFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmanagerBindFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmanagerBindFailure represents a BindFailure event raised by the Tokenmanager contract.
type TokenmanagerBindFailure struct {
	ContractAddr common.Address
	Bep2Symbol   string
	FailedReason uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBindFailure is a free log retrieval operation binding the contract event 0x831c0ef4d93bda3bce08b69ae3f29ef1a6e052b833200988554158494405a107.
//
// Solidity: event bindFailure(address indexed contractAddr, string bep2Symbol, uint32 failedReason)
func (_Tokenmanager *TokenmanagerFilterer) FilterBindFailure(opts *bind.FilterOpts, contractAddr []common.Address) (*TokenmanagerBindFailureIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.FilterLogs(opts, "bindFailure", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerBindFailureIterator{contract: _Tokenmanager.contract, event: "bindFailure", logs: logs, sub: sub}, nil
}

// WatchBindFailure is a free log subscription operation binding the contract event 0x831c0ef4d93bda3bce08b69ae3f29ef1a6e052b833200988554158494405a107.
//
// Solidity: event bindFailure(address indexed contractAddr, string bep2Symbol, uint32 failedReason)
func (_Tokenmanager *TokenmanagerFilterer) WatchBindFailure(opts *bind.WatchOpts, sink chan<- *TokenmanagerBindFailure, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.WatchLogs(opts, "bindFailure", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmanagerBindFailure)
				if err := _Tokenmanager.contract.UnpackLog(event, "bindFailure", log); err != nil {
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

// ParseBindFailure is a log parse operation binding the contract event 0x831c0ef4d93bda3bce08b69ae3f29ef1a6e052b833200988554158494405a107.
//
// Solidity: event bindFailure(address indexed contractAddr, string bep2Symbol, uint32 failedReason)
func (_Tokenmanager *TokenmanagerFilterer) ParseBindFailure(log types.Log) (*TokenmanagerBindFailure, error) {
	event := new(TokenmanagerBindFailure)
	if err := _Tokenmanager.contract.UnpackLog(event, "bindFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenmanagerBindSuccessIterator is returned from FilterBindSuccess and is used to iterate over the raw logs and unpacked data for BindSuccess events raised by the Tokenmanager contract.
type TokenmanagerBindSuccessIterator struct {
	Event *TokenmanagerBindSuccess // Event containing the contract specifics and raw log

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
func (it *TokenmanagerBindSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmanagerBindSuccess)
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
		it.Event = new(TokenmanagerBindSuccess)
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
func (it *TokenmanagerBindSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmanagerBindSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmanagerBindSuccess represents a BindSuccess event raised by the Tokenmanager contract.
type TokenmanagerBindSuccess struct {
	ContractAddr common.Address
	Bep2Symbol   string
	TotalSupply  *big.Int
	PeggyAmount  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBindSuccess is a free log retrieval operation binding the contract event 0x78e7dd9aefcdbf795c4936a66f7dc6d41bb56637b54f561a6bf7829dca3348a8.
//
// Solidity: event bindSuccess(address indexed contractAddr, string bep2Symbol, uint256 totalSupply, uint256 peggyAmount)
func (_Tokenmanager *TokenmanagerFilterer) FilterBindSuccess(opts *bind.FilterOpts, contractAddr []common.Address) (*TokenmanagerBindSuccessIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.FilterLogs(opts, "bindSuccess", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerBindSuccessIterator{contract: _Tokenmanager.contract, event: "bindSuccess", logs: logs, sub: sub}, nil
}

// WatchBindSuccess is a free log subscription operation binding the contract event 0x78e7dd9aefcdbf795c4936a66f7dc6d41bb56637b54f561a6bf7829dca3348a8.
//
// Solidity: event bindSuccess(address indexed contractAddr, string bep2Symbol, uint256 totalSupply, uint256 peggyAmount)
func (_Tokenmanager *TokenmanagerFilterer) WatchBindSuccess(opts *bind.WatchOpts, sink chan<- *TokenmanagerBindSuccess, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.WatchLogs(opts, "bindSuccess", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmanagerBindSuccess)
				if err := _Tokenmanager.contract.UnpackLog(event, "bindSuccess", log); err != nil {
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

// ParseBindSuccess is a log parse operation binding the contract event 0x78e7dd9aefcdbf795c4936a66f7dc6d41bb56637b54f561a6bf7829dca3348a8.
//
// Solidity: event bindSuccess(address indexed contractAddr, string bep2Symbol, uint256 totalSupply, uint256 peggyAmount)
func (_Tokenmanager *TokenmanagerFilterer) ParseBindSuccess(log types.Log) (*TokenmanagerBindSuccess, error) {
	event := new(TokenmanagerBindSuccess)
	if err := _Tokenmanager.contract.UnpackLog(event, "bindSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenmanagerMirrorFailureIterator is returned from FilterMirrorFailure and is used to iterate over the raw logs and unpacked data for MirrorFailure events raised by the Tokenmanager contract.
type TokenmanagerMirrorFailureIterator struct {
	Event *TokenmanagerMirrorFailure // Event containing the contract specifics and raw log

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
func (it *TokenmanagerMirrorFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmanagerMirrorFailure)
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
		it.Event = new(TokenmanagerMirrorFailure)
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
func (it *TokenmanagerMirrorFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmanagerMirrorFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmanagerMirrorFailure represents a MirrorFailure event raised by the Tokenmanager contract.
type TokenmanagerMirrorFailure struct {
	Bep20Addr common.Address
	ErrCode   uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMirrorFailure is a free log retrieval operation binding the contract event 0xefe400ad0042ebf81a245de9ae669616105e0ca9fc946352c085da0c2bc524e5.
//
// Solidity: event mirrorFailure(address indexed bep20Addr, uint8 errCode)
func (_Tokenmanager *TokenmanagerFilterer) FilterMirrorFailure(opts *bind.FilterOpts, bep20Addr []common.Address) (*TokenmanagerMirrorFailureIterator, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.FilterLogs(opts, "mirrorFailure", bep20AddrRule)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerMirrorFailureIterator{contract: _Tokenmanager.contract, event: "mirrorFailure", logs: logs, sub: sub}, nil
}

// WatchMirrorFailure is a free log subscription operation binding the contract event 0xefe400ad0042ebf81a245de9ae669616105e0ca9fc946352c085da0c2bc524e5.
//
// Solidity: event mirrorFailure(address indexed bep20Addr, uint8 errCode)
func (_Tokenmanager *TokenmanagerFilterer) WatchMirrorFailure(opts *bind.WatchOpts, sink chan<- *TokenmanagerMirrorFailure, bep20Addr []common.Address) (event.Subscription, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.WatchLogs(opts, "mirrorFailure", bep20AddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmanagerMirrorFailure)
				if err := _Tokenmanager.contract.UnpackLog(event, "mirrorFailure", log); err != nil {
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

// ParseMirrorFailure is a log parse operation binding the contract event 0xefe400ad0042ebf81a245de9ae669616105e0ca9fc946352c085da0c2bc524e5.
//
// Solidity: event mirrorFailure(address indexed bep20Addr, uint8 errCode)
func (_Tokenmanager *TokenmanagerFilterer) ParseMirrorFailure(log types.Log) (*TokenmanagerMirrorFailure, error) {
	event := new(TokenmanagerMirrorFailure)
	if err := _Tokenmanager.contract.UnpackLog(event, "mirrorFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenmanagerMirrorSuccessIterator is returned from FilterMirrorSuccess and is used to iterate over the raw logs and unpacked data for MirrorSuccess events raised by the Tokenmanager contract.
type TokenmanagerMirrorSuccessIterator struct {
	Event *TokenmanagerMirrorSuccess // Event containing the contract specifics and raw log

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
func (it *TokenmanagerMirrorSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmanagerMirrorSuccess)
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
		it.Event = new(TokenmanagerMirrorSuccess)
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
func (it *TokenmanagerMirrorSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmanagerMirrorSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmanagerMirrorSuccess represents a MirrorSuccess event raised by the Tokenmanager contract.
type TokenmanagerMirrorSuccess struct {
	Bep20Addr  common.Address
	Bep2Symbol [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMirrorSuccess is a free log retrieval operation binding the contract event 0x41787d7db08fc5907641ee8343379f28215727eb123d4b462099afab4300b036.
//
// Solidity: event mirrorSuccess(address indexed bep20Addr, bytes32 bep2Symbol)
func (_Tokenmanager *TokenmanagerFilterer) FilterMirrorSuccess(opts *bind.FilterOpts, bep20Addr []common.Address) (*TokenmanagerMirrorSuccessIterator, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.FilterLogs(opts, "mirrorSuccess", bep20AddrRule)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerMirrorSuccessIterator{contract: _Tokenmanager.contract, event: "mirrorSuccess", logs: logs, sub: sub}, nil
}

// WatchMirrorSuccess is a free log subscription operation binding the contract event 0x41787d7db08fc5907641ee8343379f28215727eb123d4b462099afab4300b036.
//
// Solidity: event mirrorSuccess(address indexed bep20Addr, bytes32 bep2Symbol)
func (_Tokenmanager *TokenmanagerFilterer) WatchMirrorSuccess(opts *bind.WatchOpts, sink chan<- *TokenmanagerMirrorSuccess, bep20Addr []common.Address) (event.Subscription, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.WatchLogs(opts, "mirrorSuccess", bep20AddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmanagerMirrorSuccess)
				if err := _Tokenmanager.contract.UnpackLog(event, "mirrorSuccess", log); err != nil {
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

// ParseMirrorSuccess is a log parse operation binding the contract event 0x41787d7db08fc5907641ee8343379f28215727eb123d4b462099afab4300b036.
//
// Solidity: event mirrorSuccess(address indexed bep20Addr, bytes32 bep2Symbol)
func (_Tokenmanager *TokenmanagerFilterer) ParseMirrorSuccess(log types.Log) (*TokenmanagerMirrorSuccess, error) {
	event := new(TokenmanagerMirrorSuccess)
	if err := _Tokenmanager.contract.UnpackLog(event, "mirrorSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenmanagerParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Tokenmanager contract.
type TokenmanagerParamChangeIterator struct {
	Event *TokenmanagerParamChange // Event containing the contract specifics and raw log

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
func (it *TokenmanagerParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmanagerParamChange)
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
		it.Event = new(TokenmanagerParamChange)
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
func (it *TokenmanagerParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmanagerParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmanagerParamChange represents a ParamChange event raised by the Tokenmanager contract.
type TokenmanagerParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Tokenmanager *TokenmanagerFilterer) FilterParamChange(opts *bind.FilterOpts) (*TokenmanagerParamChangeIterator, error) {

	logs, sub, err := _Tokenmanager.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &TokenmanagerParamChangeIterator{contract: _Tokenmanager.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Tokenmanager *TokenmanagerFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *TokenmanagerParamChange) (event.Subscription, error) {

	logs, sub, err := _Tokenmanager.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmanagerParamChange)
				if err := _Tokenmanager.contract.UnpackLog(event, "paramChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Tokenmanager *TokenmanagerFilterer) ParseParamChange(log types.Log) (*TokenmanagerParamChange, error) {
	event := new(TokenmanagerParamChange)
	if err := _Tokenmanager.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenmanagerSyncFailureIterator is returned from FilterSyncFailure and is used to iterate over the raw logs and unpacked data for SyncFailure events raised by the Tokenmanager contract.
type TokenmanagerSyncFailureIterator struct {
	Event *TokenmanagerSyncFailure // Event containing the contract specifics and raw log

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
func (it *TokenmanagerSyncFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmanagerSyncFailure)
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
		it.Event = new(TokenmanagerSyncFailure)
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
func (it *TokenmanagerSyncFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmanagerSyncFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmanagerSyncFailure represents a SyncFailure event raised by the Tokenmanager contract.
type TokenmanagerSyncFailure struct {
	Bep20Addr common.Address
	ErrCode   uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSyncFailure is a free log retrieval operation binding the contract event 0xf1e25fa536da84053644fd788b1d6a27ea6edd1f3af80c7d36ca253c4c82c9c3.
//
// Solidity: event syncFailure(address indexed bep20Addr, uint8 errCode)
func (_Tokenmanager *TokenmanagerFilterer) FilterSyncFailure(opts *bind.FilterOpts, bep20Addr []common.Address) (*TokenmanagerSyncFailureIterator, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.FilterLogs(opts, "syncFailure", bep20AddrRule)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerSyncFailureIterator{contract: _Tokenmanager.contract, event: "syncFailure", logs: logs, sub: sub}, nil
}

// WatchSyncFailure is a free log subscription operation binding the contract event 0xf1e25fa536da84053644fd788b1d6a27ea6edd1f3af80c7d36ca253c4c82c9c3.
//
// Solidity: event syncFailure(address indexed bep20Addr, uint8 errCode)
func (_Tokenmanager *TokenmanagerFilterer) WatchSyncFailure(opts *bind.WatchOpts, sink chan<- *TokenmanagerSyncFailure, bep20Addr []common.Address) (event.Subscription, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.WatchLogs(opts, "syncFailure", bep20AddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmanagerSyncFailure)
				if err := _Tokenmanager.contract.UnpackLog(event, "syncFailure", log); err != nil {
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

// ParseSyncFailure is a log parse operation binding the contract event 0xf1e25fa536da84053644fd788b1d6a27ea6edd1f3af80c7d36ca253c4c82c9c3.
//
// Solidity: event syncFailure(address indexed bep20Addr, uint8 errCode)
func (_Tokenmanager *TokenmanagerFilterer) ParseSyncFailure(log types.Log) (*TokenmanagerSyncFailure, error) {
	event := new(TokenmanagerSyncFailure)
	if err := _Tokenmanager.contract.UnpackLog(event, "syncFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenmanagerSyncSuccessIterator is returned from FilterSyncSuccess and is used to iterate over the raw logs and unpacked data for SyncSuccess events raised by the Tokenmanager contract.
type TokenmanagerSyncSuccessIterator struct {
	Event *TokenmanagerSyncSuccess // Event containing the contract specifics and raw log

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
func (it *TokenmanagerSyncSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmanagerSyncSuccess)
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
		it.Event = new(TokenmanagerSyncSuccess)
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
func (it *TokenmanagerSyncSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmanagerSyncSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmanagerSyncSuccess represents a SyncSuccess event raised by the Tokenmanager contract.
type TokenmanagerSyncSuccess struct {
	Bep20Addr common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSyncSuccess is a free log retrieval operation binding the contract event 0xbb7d3a9a559080d8281b0e4fb39dacbe2fdcafd5ef765e9a6ed871c9167dc60f.
//
// Solidity: event syncSuccess(address indexed bep20Addr)
func (_Tokenmanager *TokenmanagerFilterer) FilterSyncSuccess(opts *bind.FilterOpts, bep20Addr []common.Address) (*TokenmanagerSyncSuccessIterator, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.FilterLogs(opts, "syncSuccess", bep20AddrRule)
	if err != nil {
		return nil, err
	}
	return &TokenmanagerSyncSuccessIterator{contract: _Tokenmanager.contract, event: "syncSuccess", logs: logs, sub: sub}, nil
}

// WatchSyncSuccess is a free log subscription operation binding the contract event 0xbb7d3a9a559080d8281b0e4fb39dacbe2fdcafd5ef765e9a6ed871c9167dc60f.
//
// Solidity: event syncSuccess(address indexed bep20Addr)
func (_Tokenmanager *TokenmanagerFilterer) WatchSyncSuccess(opts *bind.WatchOpts, sink chan<- *TokenmanagerSyncSuccess, bep20Addr []common.Address) (event.Subscription, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}

	logs, sub, err := _Tokenmanager.contract.WatchLogs(opts, "syncSuccess", bep20AddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmanagerSyncSuccess)
				if err := _Tokenmanager.contract.UnpackLog(event, "syncSuccess", log); err != nil {
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

// ParseSyncSuccess is a log parse operation binding the contract event 0xbb7d3a9a559080d8281b0e4fb39dacbe2fdcafd5ef765e9a6ed871c9167dc60f.
//
// Solidity: event syncSuccess(address indexed bep20Addr)
func (_Tokenmanager *TokenmanagerFilterer) ParseSyncSuccess(log types.Log) (*TokenmanagerSyncSuccess, error) {
	event := new(TokenmanagerSyncSuccess)
	if err := _Tokenmanager.contract.UnpackLog(event, "syncSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenmanagerUnexpectedPackageIterator is returned from FilterUnexpectedPackage and is used to iterate over the raw logs and unpacked data for UnexpectedPackage events raised by the Tokenmanager contract.
type TokenmanagerUnexpectedPackageIterator struct {
	Event *TokenmanagerUnexpectedPackage // Event containing the contract specifics and raw log

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
func (it *TokenmanagerUnexpectedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmanagerUnexpectedPackage)
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
		it.Event = new(TokenmanagerUnexpectedPackage)
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
func (it *TokenmanagerUnexpectedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmanagerUnexpectedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmanagerUnexpectedPackage represents a UnexpectedPackage event raised by the Tokenmanager contract.
type TokenmanagerUnexpectedPackage struct {
	ChannelId uint8
	MsgBytes  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedPackage is a free log retrieval operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Tokenmanager *TokenmanagerFilterer) FilterUnexpectedPackage(opts *bind.FilterOpts) (*TokenmanagerUnexpectedPackageIterator, error) {

	logs, sub, err := _Tokenmanager.contract.FilterLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return &TokenmanagerUnexpectedPackageIterator{contract: _Tokenmanager.contract, event: "unexpectedPackage", logs: logs, sub: sub}, nil
}

// WatchUnexpectedPackage is a free log subscription operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Tokenmanager *TokenmanagerFilterer) WatchUnexpectedPackage(opts *bind.WatchOpts, sink chan<- *TokenmanagerUnexpectedPackage) (event.Subscription, error) {

	logs, sub, err := _Tokenmanager.contract.WatchLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmanagerUnexpectedPackage)
				if err := _Tokenmanager.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
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

// ParseUnexpectedPackage is a log parse operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Tokenmanager *TokenmanagerFilterer) ParseUnexpectedPackage(log types.Log) (*TokenmanagerUnexpectedPackage, error) {
	event := new(TokenmanagerUnexpectedPackage)
	if err := _Tokenmanager.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
