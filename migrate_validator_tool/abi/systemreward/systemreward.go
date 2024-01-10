// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package systemreward

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

// SystemrewardMetaData contains all meta data concerning the Systemreward contract.
var SystemrewardMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BC_FUSION_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIND_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CODE_OK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_STAKE_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_FAIL_DECODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_TOKEN_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INCENTIVIZE_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REWARDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_CREDIT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYSTEM_REWARD_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TIMELOCK_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_MANAGER_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_RECOVER_PORTAL_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_IN_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_OUT_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"alreadyInit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bscChainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimRewards\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"addOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"deleteOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"paramChange\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"receiveDeposit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"rewardEmpty\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"rewardTo\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// SystemrewardABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemrewardMetaData.ABI instead.
var SystemrewardABI = SystemrewardMetaData.ABI

// Systemreward is an auto generated Go binding around an Ethereum contract.
type Systemreward struct {
	SystemrewardCaller     // Read-only binding to the contract
	SystemrewardTransactor // Write-only binding to the contract
	SystemrewardFilterer   // Log filterer for contract events
}

// SystemrewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemrewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemrewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemrewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemrewardSession struct {
	Contract     *Systemreward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemrewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemrewardCallerSession struct {
	Contract *SystemrewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SystemrewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemrewardTransactorSession struct {
	Contract     *SystemrewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SystemrewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemrewardRaw struct {
	Contract *Systemreward // Generic contract binding to access the raw methods on
}

// SystemrewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemrewardCallerRaw struct {
	Contract *SystemrewardCaller // Generic read-only contract binding to access the raw methods on
}

// SystemrewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemrewardTransactorRaw struct {
	Contract *SystemrewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemreward creates a new instance of Systemreward, bound to a specific deployed contract.
func NewSystemreward(address common.Address, backend bind.ContractBackend) (*Systemreward, error) {
	contract, err := bindSystemreward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Systemreward{SystemrewardCaller: SystemrewardCaller{contract: contract}, SystemrewardTransactor: SystemrewardTransactor{contract: contract}, SystemrewardFilterer: SystemrewardFilterer{contract: contract}}, nil
}

// NewSystemrewardCaller creates a new read-only instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardCaller(address common.Address, caller bind.ContractCaller) (*SystemrewardCaller, error) {
	contract, err := bindSystemreward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemrewardCaller{contract: contract}, nil
}

// NewSystemrewardTransactor creates a new write-only instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemrewardTransactor, error) {
	contract, err := bindSystemreward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemrewardTransactor{contract: contract}, nil
}

// NewSystemrewardFilterer creates a new log filterer instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemrewardFilterer, error) {
	contract, err := bindSystemreward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemrewardFilterer{contract: contract}, nil
}

// bindSystemreward binds a generic wrapper to an already deployed contract.
func bindSystemreward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SystemrewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Systemreward *SystemrewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Systemreward.Contract.SystemrewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Systemreward *SystemrewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.Contract.SystemrewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Systemreward *SystemrewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Systemreward.Contract.SystemrewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Systemreward *SystemrewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Systemreward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Systemreward *SystemrewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Systemreward *SystemrewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Systemreward.Contract.contract.Transact(opts, method, params...)
}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCaller) BCFUSIONCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "BC_FUSION_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardSession) BCFUSIONCHANNELID() (uint8, error) {
	return _Systemreward.Contract.BCFUSIONCHANNELID(&_Systemreward.CallOpts)
}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCallerSession) BCFUSIONCHANNELID() (uint8, error) {
	return _Systemreward.Contract.BCFUSIONCHANNELID(&_Systemreward.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardSession) BINDCHANNELID() (uint8, error) {
	return _Systemreward.Contract.BINDCHANNELID(&_Systemreward.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCallerSession) BINDCHANNELID() (uint8, error) {
	return _Systemreward.Contract.BINDCHANNELID(&_Systemreward.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Systemreward *SystemrewardCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Systemreward *SystemrewardSession) CODEOK() (uint32, error) {
	return _Systemreward.Contract.CODEOK(&_Systemreward.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Systemreward *SystemrewardCallerSession) CODEOK() (uint32, error) {
	return _Systemreward.Contract.CODEOK(&_Systemreward.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.CROSSCHAINCONTRACTADDR(&_Systemreward.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.CROSSCHAINCONTRACTADDR(&_Systemreward.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCaller) CROSSSTAKECHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "CROSS_STAKE_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Systemreward.Contract.CROSSSTAKECHANNELID(&_Systemreward.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCallerSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Systemreward.Contract.CROSSSTAKECHANNELID(&_Systemreward.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Systemreward *SystemrewardCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Systemreward *SystemrewardSession) ERRORFAILDECODE() (uint32, error) {
	return _Systemreward.Contract.ERRORFAILDECODE(&_Systemreward.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Systemreward *SystemrewardCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Systemreward.Contract.ERRORFAILDECODE(&_Systemreward.CallOpts)
}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) GOVERNORADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "GOVERNOR_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) GOVERNORADDR() (common.Address, error) {
	return _Systemreward.Contract.GOVERNORADDR(&_Systemreward.CallOpts)
}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) GOVERNORADDR() (common.Address, error) {
	return _Systemreward.Contract.GOVERNORADDR(&_Systemreward.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardSession) GOVCHANNELID() (uint8, error) {
	return _Systemreward.Contract.GOVCHANNELID(&_Systemreward.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCallerSession) GOVCHANNELID() (uint8, error) {
	return _Systemreward.Contract.GOVCHANNELID(&_Systemreward.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) GOVHUBADDR() (common.Address, error) {
	return _Systemreward.Contract.GOVHUBADDR(&_Systemreward.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Systemreward.Contract.GOVHUBADDR(&_Systemreward.CallOpts)
}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) GOVTOKENADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "GOV_TOKEN_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) GOVTOKENADDR() (common.Address, error) {
	return _Systemreward.Contract.GOVTOKENADDR(&_Systemreward.CallOpts)
}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) GOVTOKENADDR() (common.Address, error) {
	return _Systemreward.Contract.GOVTOKENADDR(&_Systemreward.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Systemreward.Contract.INCENTIVIZEADDR(&_Systemreward.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Systemreward.Contract.INCENTIVIZEADDR(&_Systemreward.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Systemreward.Contract.LIGHTCLIENTADDR(&_Systemreward.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Systemreward.Contract.LIGHTCLIENTADDR(&_Systemreward.CallOpts)
}

// MAXREWARDS is a free data retrieval call binding the contract method 0xfb5478b3.
//
// Solidity: function MAX_REWARDS() view returns(uint256)
func (_Systemreward *SystemrewardCaller) MAXREWARDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "MAX_REWARDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREWARDS is a free data retrieval call binding the contract method 0xfb5478b3.
//
// Solidity: function MAX_REWARDS() view returns(uint256)
func (_Systemreward *SystemrewardSession) MAXREWARDS() (*big.Int, error) {
	return _Systemreward.Contract.MAXREWARDS(&_Systemreward.CallOpts)
}

// MAXREWARDS is a free data retrieval call binding the contract method 0xfb5478b3.
//
// Solidity: function MAX_REWARDS() view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) MAXREWARDS() (*big.Int, error) {
	return _Systemreward.Contract.MAXREWARDS(&_Systemreward.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.RELAYERHUBCONTRACTADDR(&_Systemreward.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.RELAYERHUBCONTRACTADDR(&_Systemreward.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardSession) SLASHCHANNELID() (uint8, error) {
	return _Systemreward.Contract.SLASHCHANNELID(&_Systemreward.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Systemreward.Contract.SLASHCHANNELID(&_Systemreward.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.SLASHCONTRACTADDR(&_Systemreward.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.SLASHCONTRACTADDR(&_Systemreward.CallOpts)
}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) STAKECREDITADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "STAKE_CREDIT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) STAKECREDITADDR() (common.Address, error) {
	return _Systemreward.Contract.STAKECREDITADDR(&_Systemreward.CallOpts)
}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) STAKECREDITADDR() (common.Address, error) {
	return _Systemreward.Contract.STAKECREDITADDR(&_Systemreward.CallOpts)
}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) STAKEHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "STAKE_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) STAKEHUBADDR() (common.Address, error) {
	return _Systemreward.Contract.STAKEHUBADDR(&_Systemreward.CallOpts)
}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) STAKEHUBADDR() (common.Address, error) {
	return _Systemreward.Contract.STAKEHUBADDR(&_Systemreward.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardSession) STAKINGCHANNELID() (uint8, error) {
	return _Systemreward.Contract.STAKINGCHANNELID(&_Systemreward.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Systemreward.Contract.STAKINGCHANNELID(&_Systemreward.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) STAKINGCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "STAKING_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.STAKINGCONTRACTADDR(&_Systemreward.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.STAKINGCONTRACTADDR(&_Systemreward.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Systemreward.Contract.SYSTEMREWARDADDR(&_Systemreward.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Systemreward.Contract.SYSTEMREWARDADDR(&_Systemreward.CallOpts)
}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) TIMELOCKADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "TIMELOCK_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) TIMELOCKADDR() (common.Address, error) {
	return _Systemreward.Contract.TIMELOCKADDR(&_Systemreward.CallOpts)
}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) TIMELOCKADDR() (common.Address, error) {
	return _Systemreward.Contract.TIMELOCKADDR(&_Systemreward.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) TOKENHUBADDR() (common.Address, error) {
	return _Systemreward.Contract.TOKENHUBADDR(&_Systemreward.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Systemreward.Contract.TOKENHUBADDR(&_Systemreward.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Systemreward.Contract.TOKENMANAGERADDR(&_Systemreward.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Systemreward.Contract.TOKENMANAGERADDR(&_Systemreward.CallOpts)
}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) TOKENRECOVERPORTALADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "TOKEN_RECOVER_PORTAL_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) TOKENRECOVERPORTALADDR() (common.Address, error) {
	return _Systemreward.Contract.TOKENRECOVERPORTALADDR(&_Systemreward.CallOpts)
}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) TOKENRECOVERPORTALADDR() (common.Address, error) {
	return _Systemreward.Contract.TOKENRECOVERPORTALADDR(&_Systemreward.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Systemreward.Contract.TRANSFERINCHANNELID(&_Systemreward.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Systemreward.Contract.TRANSFERINCHANNELID(&_Systemreward.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Systemreward.Contract.TRANSFEROUTCHANNELID(&_Systemreward.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Systemreward *SystemrewardCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Systemreward.Contract.TRANSFEROUTCHANNELID(&_Systemreward.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.VALIDATORCONTRACTADDR(&_Systemreward.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Systemreward *SystemrewardCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Systemreward.Contract.VALIDATORCONTRACTADDR(&_Systemreward.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardSession) AlreadyInit() (bool, error) {
	return _Systemreward.Contract.AlreadyInit(&_Systemreward.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardCallerSession) AlreadyInit() (bool, error) {
	return _Systemreward.Contract.AlreadyInit(&_Systemreward.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Systemreward *SystemrewardCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Systemreward *SystemrewardSession) BscChainID() (uint16, error) {
	return _Systemreward.Contract.BscChainID(&_Systemreward.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Systemreward *SystemrewardCallerSession) BscChainID() (uint16, error) {
	return _Systemreward.Contract.BscChainID(&_Systemreward.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address addr) view returns(bool)
func (_Systemreward *SystemrewardCaller) IsOperator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "isOperator", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address addr) view returns(bool)
func (_Systemreward *SystemrewardSession) IsOperator(addr common.Address) (bool, error) {
	return _Systemreward.Contract.IsOperator(&_Systemreward.CallOpts, addr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address addr) view returns(bool)
func (_Systemreward *SystemrewardCallerSession) IsOperator(addr common.Address) (bool, error) {
	return _Systemreward.Contract.IsOperator(&_Systemreward.CallOpts, addr)
}

// NumOperator is a free data retrieval call binding the contract method 0x3a0b0eff.
//
// Solidity: function numOperator() view returns(uint256)
func (_Systemreward *SystemrewardCaller) NumOperator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "numOperator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOperator is a free data retrieval call binding the contract method 0x3a0b0eff.
//
// Solidity: function numOperator() view returns(uint256)
func (_Systemreward *SystemrewardSession) NumOperator() (*big.Int, error) {
	return _Systemreward.Contract.NumOperator(&_Systemreward.CallOpts)
}

// NumOperator is a free data retrieval call binding the contract method 0x3a0b0eff.
//
// Solidity: function numOperator() view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) NumOperator() (*big.Int, error) {
	return _Systemreward.Contract.NumOperator(&_Systemreward.CallOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256)
func (_Systemreward *SystemrewardTransactor) ClaimRewards(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "claimRewards", to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256)
func (_Systemreward *SystemrewardSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Systemreward.Contract.ClaimRewards(&_Systemreward.TransactOpts, to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256)
func (_Systemreward *SystemrewardTransactorSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Systemreward.Contract.ClaimRewards(&_Systemreward.TransactOpts, to, amount)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Systemreward *SystemrewardTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Systemreward *SystemrewardSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Systemreward.Contract.UpdateParam(&_Systemreward.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Systemreward *SystemrewardTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Systemreward.Contract.UpdateParam(&_Systemreward.TransactOpts, key, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Systemreward *SystemrewardTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Systemreward *SystemrewardSession) Receive() (*types.Transaction, error) {
	return _Systemreward.Contract.Receive(&_Systemreward.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Systemreward *SystemrewardTransactorSession) Receive() (*types.Transaction, error) {
	return _Systemreward.Contract.Receive(&_Systemreward.TransactOpts)
}

// SystemrewardAddOperatorIterator is returned from FilterAddOperator and is used to iterate over the raw logs and unpacked data for AddOperator events raised by the Systemreward contract.
type SystemrewardAddOperatorIterator struct {
	Event *SystemrewardAddOperator // Event containing the contract specifics and raw log

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
func (it *SystemrewardAddOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemrewardAddOperator)
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
		it.Event = new(SystemrewardAddOperator)
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
func (it *SystemrewardAddOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemrewardAddOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemrewardAddOperator represents a AddOperator event raised by the Systemreward contract.
type SystemrewardAddOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddOperator is a free log retrieval operation binding the contract event 0x9870d7fe5d112134c55844951dedf365363006d9c588db07c4c85af6322a0619.
//
// Solidity: event addOperator(address indexed operator)
func (_Systemreward *SystemrewardFilterer) FilterAddOperator(opts *bind.FilterOpts, operator []common.Address) (*SystemrewardAddOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Systemreward.contract.FilterLogs(opts, "addOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SystemrewardAddOperatorIterator{contract: _Systemreward.contract, event: "addOperator", logs: logs, sub: sub}, nil
}

// WatchAddOperator is a free log subscription operation binding the contract event 0x9870d7fe5d112134c55844951dedf365363006d9c588db07c4c85af6322a0619.
//
// Solidity: event addOperator(address indexed operator)
func (_Systemreward *SystemrewardFilterer) WatchAddOperator(opts *bind.WatchOpts, sink chan<- *SystemrewardAddOperator, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Systemreward.contract.WatchLogs(opts, "addOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemrewardAddOperator)
				if err := _Systemreward.contract.UnpackLog(event, "addOperator", log); err != nil {
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

// ParseAddOperator is a log parse operation binding the contract event 0x9870d7fe5d112134c55844951dedf365363006d9c588db07c4c85af6322a0619.
//
// Solidity: event addOperator(address indexed operator)
func (_Systemreward *SystemrewardFilterer) ParseAddOperator(log types.Log) (*SystemrewardAddOperator, error) {
	event := new(SystemrewardAddOperator)
	if err := _Systemreward.contract.UnpackLog(event, "addOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemrewardDeleteOperatorIterator is returned from FilterDeleteOperator and is used to iterate over the raw logs and unpacked data for DeleteOperator events raised by the Systemreward contract.
type SystemrewardDeleteOperatorIterator struct {
	Event *SystemrewardDeleteOperator // Event containing the contract specifics and raw log

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
func (it *SystemrewardDeleteOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemrewardDeleteOperator)
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
		it.Event = new(SystemrewardDeleteOperator)
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
func (it *SystemrewardDeleteOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemrewardDeleteOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemrewardDeleteOperator represents a DeleteOperator event raised by the Systemreward contract.
type SystemrewardDeleteOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteOperator is a free log retrieval operation binding the contract event 0xb40992a19dba61ea600e87fce607102bf5908dc89076217b6ca6ae195224f702.
//
// Solidity: event deleteOperator(address indexed operator)
func (_Systemreward *SystemrewardFilterer) FilterDeleteOperator(opts *bind.FilterOpts, operator []common.Address) (*SystemrewardDeleteOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Systemreward.contract.FilterLogs(opts, "deleteOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SystemrewardDeleteOperatorIterator{contract: _Systemreward.contract, event: "deleteOperator", logs: logs, sub: sub}, nil
}

// WatchDeleteOperator is a free log subscription operation binding the contract event 0xb40992a19dba61ea600e87fce607102bf5908dc89076217b6ca6ae195224f702.
//
// Solidity: event deleteOperator(address indexed operator)
func (_Systemreward *SystemrewardFilterer) WatchDeleteOperator(opts *bind.WatchOpts, sink chan<- *SystemrewardDeleteOperator, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Systemreward.contract.WatchLogs(opts, "deleteOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemrewardDeleteOperator)
				if err := _Systemreward.contract.UnpackLog(event, "deleteOperator", log); err != nil {
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

// ParseDeleteOperator is a log parse operation binding the contract event 0xb40992a19dba61ea600e87fce607102bf5908dc89076217b6ca6ae195224f702.
//
// Solidity: event deleteOperator(address indexed operator)
func (_Systemreward *SystemrewardFilterer) ParseDeleteOperator(log types.Log) (*SystemrewardDeleteOperator, error) {
	event := new(SystemrewardDeleteOperator)
	if err := _Systemreward.contract.UnpackLog(event, "deleteOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemrewardParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Systemreward contract.
type SystemrewardParamChangeIterator struct {
	Event *SystemrewardParamChange // Event containing the contract specifics and raw log

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
func (it *SystemrewardParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemrewardParamChange)
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
		it.Event = new(SystemrewardParamChange)
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
func (it *SystemrewardParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemrewardParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemrewardParamChange represents a ParamChange event raised by the Systemreward contract.
type SystemrewardParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Systemreward *SystemrewardFilterer) FilterParamChange(opts *bind.FilterOpts) (*SystemrewardParamChangeIterator, error) {

	logs, sub, err := _Systemreward.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &SystemrewardParamChangeIterator{contract: _Systemreward.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Systemreward *SystemrewardFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *SystemrewardParamChange) (event.Subscription, error) {

	logs, sub, err := _Systemreward.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemrewardParamChange)
				if err := _Systemreward.contract.UnpackLog(event, "paramChange", log); err != nil {
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
func (_Systemreward *SystemrewardFilterer) ParseParamChange(log types.Log) (*SystemrewardParamChange, error) {
	event := new(SystemrewardParamChange)
	if err := _Systemreward.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemrewardReceiveDepositIterator is returned from FilterReceiveDeposit and is used to iterate over the raw logs and unpacked data for ReceiveDeposit events raised by the Systemreward contract.
type SystemrewardReceiveDepositIterator struct {
	Event *SystemrewardReceiveDeposit // Event containing the contract specifics and raw log

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
func (it *SystemrewardReceiveDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemrewardReceiveDeposit)
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
		it.Event = new(SystemrewardReceiveDeposit)
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
func (it *SystemrewardReceiveDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemrewardReceiveDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemrewardReceiveDeposit represents a ReceiveDeposit event raised by the Systemreward contract.
type SystemrewardReceiveDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceiveDeposit is a free log retrieval operation binding the contract event 0x6c98249d85d88c3753a04a22230f595e4dc8d3dc86c34af35deeeedc861b89db.
//
// Solidity: event receiveDeposit(address indexed from, uint256 amount)
func (_Systemreward *SystemrewardFilterer) FilterReceiveDeposit(opts *bind.FilterOpts, from []common.Address) (*SystemrewardReceiveDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Systemreward.contract.FilterLogs(opts, "receiveDeposit", fromRule)
	if err != nil {
		return nil, err
	}
	return &SystemrewardReceiveDepositIterator{contract: _Systemreward.contract, event: "receiveDeposit", logs: logs, sub: sub}, nil
}

// WatchReceiveDeposit is a free log subscription operation binding the contract event 0x6c98249d85d88c3753a04a22230f595e4dc8d3dc86c34af35deeeedc861b89db.
//
// Solidity: event receiveDeposit(address indexed from, uint256 amount)
func (_Systemreward *SystemrewardFilterer) WatchReceiveDeposit(opts *bind.WatchOpts, sink chan<- *SystemrewardReceiveDeposit, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Systemreward.contract.WatchLogs(opts, "receiveDeposit", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemrewardReceiveDeposit)
				if err := _Systemreward.contract.UnpackLog(event, "receiveDeposit", log); err != nil {
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

// ParseReceiveDeposit is a log parse operation binding the contract event 0x6c98249d85d88c3753a04a22230f595e4dc8d3dc86c34af35deeeedc861b89db.
//
// Solidity: event receiveDeposit(address indexed from, uint256 amount)
func (_Systemreward *SystemrewardFilterer) ParseReceiveDeposit(log types.Log) (*SystemrewardReceiveDeposit, error) {
	event := new(SystemrewardReceiveDeposit)
	if err := _Systemreward.contract.UnpackLog(event, "receiveDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemrewardRewardEmptyIterator is returned from FilterRewardEmpty and is used to iterate over the raw logs and unpacked data for RewardEmpty events raised by the Systemreward contract.
type SystemrewardRewardEmptyIterator struct {
	Event *SystemrewardRewardEmpty // Event containing the contract specifics and raw log

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
func (it *SystemrewardRewardEmptyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemrewardRewardEmpty)
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
		it.Event = new(SystemrewardRewardEmpty)
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
func (it *SystemrewardRewardEmptyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemrewardRewardEmptyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemrewardRewardEmpty represents a RewardEmpty event raised by the Systemreward contract.
type SystemrewardRewardEmpty struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRewardEmpty is a free log retrieval operation binding the contract event 0xe589651933c2457488cc0d8e0941518abf748e799435e4e396d9c4d0b2db2d4d.
//
// Solidity: event rewardEmpty()
func (_Systemreward *SystemrewardFilterer) FilterRewardEmpty(opts *bind.FilterOpts) (*SystemrewardRewardEmptyIterator, error) {

	logs, sub, err := _Systemreward.contract.FilterLogs(opts, "rewardEmpty")
	if err != nil {
		return nil, err
	}
	return &SystemrewardRewardEmptyIterator{contract: _Systemreward.contract, event: "rewardEmpty", logs: logs, sub: sub}, nil
}

// WatchRewardEmpty is a free log subscription operation binding the contract event 0xe589651933c2457488cc0d8e0941518abf748e799435e4e396d9c4d0b2db2d4d.
//
// Solidity: event rewardEmpty()
func (_Systemreward *SystemrewardFilterer) WatchRewardEmpty(opts *bind.WatchOpts, sink chan<- *SystemrewardRewardEmpty) (event.Subscription, error) {

	logs, sub, err := _Systemreward.contract.WatchLogs(opts, "rewardEmpty")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemrewardRewardEmpty)
				if err := _Systemreward.contract.UnpackLog(event, "rewardEmpty", log); err != nil {
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

// ParseRewardEmpty is a log parse operation binding the contract event 0xe589651933c2457488cc0d8e0941518abf748e799435e4e396d9c4d0b2db2d4d.
//
// Solidity: event rewardEmpty()
func (_Systemreward *SystemrewardFilterer) ParseRewardEmpty(log types.Log) (*SystemrewardRewardEmpty, error) {
	event := new(SystemrewardRewardEmpty)
	if err := _Systemreward.contract.UnpackLog(event, "rewardEmpty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemrewardRewardToIterator is returned from FilterRewardTo and is used to iterate over the raw logs and unpacked data for RewardTo events raised by the Systemreward contract.
type SystemrewardRewardToIterator struct {
	Event *SystemrewardRewardTo // Event containing the contract specifics and raw log

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
func (it *SystemrewardRewardToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemrewardRewardTo)
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
		it.Event = new(SystemrewardRewardTo)
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
func (it *SystemrewardRewardToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemrewardRewardToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemrewardRewardTo represents a RewardTo event raised by the Systemreward contract.
type SystemrewardRewardTo struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardTo is a free log retrieval operation binding the contract event 0xf8b71c64315fc33b2ead2adfa487955065152a8ac33d9d5193aafd7f45dc15a0.
//
// Solidity: event rewardTo(address indexed to, uint256 amount)
func (_Systemreward *SystemrewardFilterer) FilterRewardTo(opts *bind.FilterOpts, to []common.Address) (*SystemrewardRewardToIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Systemreward.contract.FilterLogs(opts, "rewardTo", toRule)
	if err != nil {
		return nil, err
	}
	return &SystemrewardRewardToIterator{contract: _Systemreward.contract, event: "rewardTo", logs: logs, sub: sub}, nil
}

// WatchRewardTo is a free log subscription operation binding the contract event 0xf8b71c64315fc33b2ead2adfa487955065152a8ac33d9d5193aafd7f45dc15a0.
//
// Solidity: event rewardTo(address indexed to, uint256 amount)
func (_Systemreward *SystemrewardFilterer) WatchRewardTo(opts *bind.WatchOpts, sink chan<- *SystemrewardRewardTo, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Systemreward.contract.WatchLogs(opts, "rewardTo", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemrewardRewardTo)
				if err := _Systemreward.contract.UnpackLog(event, "rewardTo", log); err != nil {
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

// ParseRewardTo is a log parse operation binding the contract event 0xf8b71c64315fc33b2ead2adfa487955065152a8ac33d9d5193aafd7f45dc15a0.
//
// Solidity: event rewardTo(address indexed to, uint256 amount)
func (_Systemreward *SystemrewardFilterer) ParseRewardTo(log types.Log) (*SystemrewardRewardTo, error) {
	event := new(SystemrewardRewardTo)
	if err := _Systemreward.contract.UnpackLog(event, "rewardTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
