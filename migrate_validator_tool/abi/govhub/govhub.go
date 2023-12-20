// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package govhub

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

// GovhubMetaData contains all meta data concerning the Govhub contract.
var GovhubMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BIND_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CODE_OK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_STAKE_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_FAIL_DECODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_TARGET_CONTRACT_FAIL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_TARGET_NOT_CONTRACT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_TOKEN_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INCENTIVIZE_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_UPDATE_MESSAGE_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_CREDIT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYSTEM_REWARD_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TIMELOCK_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_MANAGER_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_RECOVER_PORTAL_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_IN_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_OUT_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"alreadyInit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bscChainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleAckPackage\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleFailAckPackage\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleSynPackage\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"responsePayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"failReasonWithBytes\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"failReasonWithStr\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"paramChange\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// GovhubABI is the input ABI used to generate the binding from.
// Deprecated: Use GovhubMetaData.ABI instead.
var GovhubABI = GovhubMetaData.ABI

// Govhub is an auto generated Go binding around an Ethereum contract.
type Govhub struct {
	GovhubCaller     // Read-only binding to the contract
	GovhubTransactor // Write-only binding to the contract
	GovhubFilterer   // Log filterer for contract events
}

// GovhubCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovhubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovhubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovhubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovhubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovhubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovhubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovhubSession struct {
	Contract     *Govhub           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovhubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovhubCallerSession struct {
	Contract *GovhubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovhubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovhubTransactorSession struct {
	Contract     *GovhubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovhubRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovhubRaw struct {
	Contract *Govhub // Generic contract binding to access the raw methods on
}

// GovhubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovhubCallerRaw struct {
	Contract *GovhubCaller // Generic read-only contract binding to access the raw methods on
}

// GovhubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovhubTransactorRaw struct {
	Contract *GovhubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovhub creates a new instance of Govhub, bound to a specific deployed contract.
func NewGovhub(address common.Address, backend bind.ContractBackend) (*Govhub, error) {
	contract, err := bindGovhub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Govhub{GovhubCaller: GovhubCaller{contract: contract}, GovhubTransactor: GovhubTransactor{contract: contract}, GovhubFilterer: GovhubFilterer{contract: contract}}, nil
}

// NewGovhubCaller creates a new read-only instance of Govhub, bound to a specific deployed contract.
func NewGovhubCaller(address common.Address, caller bind.ContractCaller) (*GovhubCaller, error) {
	contract, err := bindGovhub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovhubCaller{contract: contract}, nil
}

// NewGovhubTransactor creates a new write-only instance of Govhub, bound to a specific deployed contract.
func NewGovhubTransactor(address common.Address, transactor bind.ContractTransactor) (*GovhubTransactor, error) {
	contract, err := bindGovhub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovhubTransactor{contract: contract}, nil
}

// NewGovhubFilterer creates a new log filterer instance of Govhub, bound to a specific deployed contract.
func NewGovhubFilterer(address common.Address, filterer bind.ContractFilterer) (*GovhubFilterer, error) {
	contract, err := bindGovhub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovhubFilterer{contract: contract}, nil
}

// bindGovhub binds a generic wrapper to an already deployed contract.
func bindGovhub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GovhubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Govhub *GovhubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Govhub.Contract.GovhubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Govhub *GovhubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Govhub.Contract.GovhubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Govhub *GovhubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Govhub.Contract.GovhubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Govhub *GovhubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Govhub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Govhub *GovhubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Govhub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Govhub *GovhubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Govhub.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Govhub *GovhubSession) BINDCHANNELID() (uint8, error) {
	return _Govhub.Contract.BINDCHANNELID(&_Govhub.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCallerSession) BINDCHANNELID() (uint8, error) {
	return _Govhub.Contract.BINDCHANNELID(&_Govhub.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Govhub *GovhubCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Govhub *GovhubSession) CODEOK() (uint32, error) {
	return _Govhub.Contract.CODEOK(&_Govhub.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Govhub *GovhubCallerSession) CODEOK() (uint32, error) {
	return _Govhub.Contract.CODEOK(&_Govhub.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.CROSSCHAINCONTRACTADDR(&_Govhub.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.CROSSCHAINCONTRACTADDR(&_Govhub.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCaller) CROSSSTAKECHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "CROSS_STAKE_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Govhub *GovhubSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Govhub.Contract.CROSSSTAKECHANNELID(&_Govhub.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCallerSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Govhub.Contract.CROSSSTAKECHANNELID(&_Govhub.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Govhub *GovhubCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Govhub *GovhubSession) ERRORFAILDECODE() (uint32, error) {
	return _Govhub.Contract.ERRORFAILDECODE(&_Govhub.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Govhub *GovhubCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Govhub.Contract.ERRORFAILDECODE(&_Govhub.CallOpts)
}

// ERRORTARGETCONTRACTFAIL is a free data retrieval call binding the contract method 0x3a21baae.
//
// Solidity: function ERROR_TARGET_CONTRACT_FAIL() view returns(uint32)
func (_Govhub *GovhubCaller) ERRORTARGETCONTRACTFAIL(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "ERROR_TARGET_CONTRACT_FAIL")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORTARGETCONTRACTFAIL is a free data retrieval call binding the contract method 0x3a21baae.
//
// Solidity: function ERROR_TARGET_CONTRACT_FAIL() view returns(uint32)
func (_Govhub *GovhubSession) ERRORTARGETCONTRACTFAIL() (uint32, error) {
	return _Govhub.Contract.ERRORTARGETCONTRACTFAIL(&_Govhub.CallOpts)
}

// ERRORTARGETCONTRACTFAIL is a free data retrieval call binding the contract method 0x3a21baae.
//
// Solidity: function ERROR_TARGET_CONTRACT_FAIL() view returns(uint32)
func (_Govhub *GovhubCallerSession) ERRORTARGETCONTRACTFAIL() (uint32, error) {
	return _Govhub.Contract.ERRORTARGETCONTRACTFAIL(&_Govhub.CallOpts)
}

// ERRORTARGETNOTCONTRACT is a free data retrieval call binding the contract method 0x9ab1a373.
//
// Solidity: function ERROR_TARGET_NOT_CONTRACT() view returns(uint32)
func (_Govhub *GovhubCaller) ERRORTARGETNOTCONTRACT(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "ERROR_TARGET_NOT_CONTRACT")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORTARGETNOTCONTRACT is a free data retrieval call binding the contract method 0x9ab1a373.
//
// Solidity: function ERROR_TARGET_NOT_CONTRACT() view returns(uint32)
func (_Govhub *GovhubSession) ERRORTARGETNOTCONTRACT() (uint32, error) {
	return _Govhub.Contract.ERRORTARGETNOTCONTRACT(&_Govhub.CallOpts)
}

// ERRORTARGETNOTCONTRACT is a free data retrieval call binding the contract method 0x9ab1a373.
//
// Solidity: function ERROR_TARGET_NOT_CONTRACT() view returns(uint32)
func (_Govhub *GovhubCallerSession) ERRORTARGETNOTCONTRACT() (uint32, error) {
	return _Govhub.Contract.ERRORTARGETNOTCONTRACT(&_Govhub.CallOpts)
}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Govhub *GovhubCaller) GOVERNORADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "GOVERNOR_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Govhub *GovhubSession) GOVERNORADDR() (common.Address, error) {
	return _Govhub.Contract.GOVERNORADDR(&_Govhub.CallOpts)
}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) GOVERNORADDR() (common.Address, error) {
	return _Govhub.Contract.GOVERNORADDR(&_Govhub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Govhub *GovhubSession) GOVCHANNELID() (uint8, error) {
	return _Govhub.Contract.GOVCHANNELID(&_Govhub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCallerSession) GOVCHANNELID() (uint8, error) {
	return _Govhub.Contract.GOVCHANNELID(&_Govhub.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Govhub *GovhubCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Govhub *GovhubSession) GOVHUBADDR() (common.Address, error) {
	return _Govhub.Contract.GOVHUBADDR(&_Govhub.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Govhub.Contract.GOVHUBADDR(&_Govhub.CallOpts)
}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Govhub *GovhubCaller) GOVTOKENADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "GOV_TOKEN_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Govhub *GovhubSession) GOVTOKENADDR() (common.Address, error) {
	return _Govhub.Contract.GOVTOKENADDR(&_Govhub.CallOpts)
}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) GOVTOKENADDR() (common.Address, error) {
	return _Govhub.Contract.GOVTOKENADDR(&_Govhub.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Govhub *GovhubCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Govhub *GovhubSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Govhub.Contract.INCENTIVIZEADDR(&_Govhub.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Govhub.Contract.INCENTIVIZEADDR(&_Govhub.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Govhub *GovhubCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Govhub *GovhubSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Govhub.Contract.LIGHTCLIENTADDR(&_Govhub.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Govhub.Contract.LIGHTCLIENTADDR(&_Govhub.CallOpts)
}

// PARAMUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x4900c4ea.
//
// Solidity: function PARAM_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_Govhub *GovhubCaller) PARAMUPDATEMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "PARAM_UPDATE_MESSAGE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PARAMUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x4900c4ea.
//
// Solidity: function PARAM_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_Govhub *GovhubSession) PARAMUPDATEMESSAGETYPE() (uint8, error) {
	return _Govhub.Contract.PARAMUPDATEMESSAGETYPE(&_Govhub.CallOpts)
}

// PARAMUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x4900c4ea.
//
// Solidity: function PARAM_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_Govhub *GovhubCallerSession) PARAMUPDATEMESSAGETYPE() (uint8, error) {
	return _Govhub.Contract.PARAMUPDATEMESSAGETYPE(&_Govhub.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.RELAYERHUBCONTRACTADDR(&_Govhub.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.RELAYERHUBCONTRACTADDR(&_Govhub.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Govhub *GovhubSession) SLASHCHANNELID() (uint8, error) {
	return _Govhub.Contract.SLASHCHANNELID(&_Govhub.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Govhub.Contract.SLASHCHANNELID(&_Govhub.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.SLASHCONTRACTADDR(&_Govhub.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.SLASHCONTRACTADDR(&_Govhub.CallOpts)
}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Govhub *GovhubCaller) STAKECREDITADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "STAKE_CREDIT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Govhub *GovhubSession) STAKECREDITADDR() (common.Address, error) {
	return _Govhub.Contract.STAKECREDITADDR(&_Govhub.CallOpts)
}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) STAKECREDITADDR() (common.Address, error) {
	return _Govhub.Contract.STAKECREDITADDR(&_Govhub.CallOpts)
}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Govhub *GovhubCaller) STAKEHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "STAKE_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Govhub *GovhubSession) STAKEHUBADDR() (common.Address, error) {
	return _Govhub.Contract.STAKEHUBADDR(&_Govhub.CallOpts)
}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) STAKEHUBADDR() (common.Address, error) {
	return _Govhub.Contract.STAKEHUBADDR(&_Govhub.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Govhub *GovhubSession) STAKINGCHANNELID() (uint8, error) {
	return _Govhub.Contract.STAKINGCHANNELID(&_Govhub.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Govhub.Contract.STAKINGCHANNELID(&_Govhub.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCaller) STAKINGCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "STAKING_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.STAKINGCONTRACTADDR(&_Govhub.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.STAKINGCONTRACTADDR(&_Govhub.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Govhub *GovhubCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Govhub *GovhubSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Govhub.Contract.SYSTEMREWARDADDR(&_Govhub.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Govhub.Contract.SYSTEMREWARDADDR(&_Govhub.CallOpts)
}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Govhub *GovhubCaller) TIMELOCKADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "TIMELOCK_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Govhub *GovhubSession) TIMELOCKADDR() (common.Address, error) {
	return _Govhub.Contract.TIMELOCKADDR(&_Govhub.CallOpts)
}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) TIMELOCKADDR() (common.Address, error) {
	return _Govhub.Contract.TIMELOCKADDR(&_Govhub.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Govhub *GovhubCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Govhub *GovhubSession) TOKENHUBADDR() (common.Address, error) {
	return _Govhub.Contract.TOKENHUBADDR(&_Govhub.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Govhub.Contract.TOKENHUBADDR(&_Govhub.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Govhub *GovhubCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Govhub *GovhubSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Govhub.Contract.TOKENMANAGERADDR(&_Govhub.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Govhub.Contract.TOKENMANAGERADDR(&_Govhub.CallOpts)
}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Govhub *GovhubCaller) TOKENRECOVERPORTALADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "TOKEN_RECOVER_PORTAL_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Govhub *GovhubSession) TOKENRECOVERPORTALADDR() (common.Address, error) {
	return _Govhub.Contract.TOKENRECOVERPORTALADDR(&_Govhub.CallOpts)
}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) TOKENRECOVERPORTALADDR() (common.Address, error) {
	return _Govhub.Contract.TOKENRECOVERPORTALADDR(&_Govhub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Govhub *GovhubSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Govhub.Contract.TRANSFERINCHANNELID(&_Govhub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Govhub.Contract.TRANSFERINCHANNELID(&_Govhub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Govhub *GovhubSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Govhub.Contract.TRANSFEROUTCHANNELID(&_Govhub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Govhub *GovhubCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Govhub.Contract.TRANSFEROUTCHANNELID(&_Govhub.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.VALIDATORCONTRACTADDR(&_Govhub.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Govhub *GovhubCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Govhub.Contract.VALIDATORCONTRACTADDR(&_Govhub.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Govhub *GovhubCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Govhub *GovhubSession) AlreadyInit() (bool, error) {
	return _Govhub.Contract.AlreadyInit(&_Govhub.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Govhub *GovhubCallerSession) AlreadyInit() (bool, error) {
	return _Govhub.Contract.AlreadyInit(&_Govhub.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Govhub *GovhubCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Govhub.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Govhub *GovhubSession) BscChainID() (uint16, error) {
	return _Govhub.Contract.BscChainID(&_Govhub.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Govhub *GovhubCallerSession) BscChainID() (uint16, error) {
	return _Govhub.Contract.BscChainID(&_Govhub.CallOpts)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes ) returns()
func (_Govhub *GovhubTransactor) HandleAckPackage(opts *bind.TransactOpts, arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Govhub.contract.Transact(opts, "handleAckPackage", arg0, arg1)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes ) returns()
func (_Govhub *GovhubSession) HandleAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Govhub.Contract.HandleAckPackage(&_Govhub.TransactOpts, arg0, arg1)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes ) returns()
func (_Govhub *GovhubTransactorSession) HandleAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Govhub.Contract.HandleAckPackage(&_Govhub.TransactOpts, arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Govhub *GovhubTransactor) HandleFailAckPackage(opts *bind.TransactOpts, arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Govhub.contract.Transact(opts, "handleFailAckPackage", arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Govhub *GovhubSession) HandleFailAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Govhub.Contract.HandleFailAckPackage(&_Govhub.TransactOpts, arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Govhub *GovhubTransactorSession) HandleFailAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Govhub.Contract.HandleFailAckPackage(&_Govhub.TransactOpts, arg0, arg1)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Govhub *GovhubTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Govhub.contract.Transact(opts, "handleSynPackage", arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Govhub *GovhubSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Govhub.Contract.HandleSynPackage(&_Govhub.TransactOpts, arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Govhub *GovhubTransactorSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Govhub.Contract.HandleSynPackage(&_Govhub.TransactOpts, arg0, msgBytes)
}

// UpdateParam is a paid mutator transaction binding the contract method 0x88e4194e.
//
// Solidity: function updateParam(string key, bytes value, address target) returns()
func (_Govhub *GovhubTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte, target common.Address) (*types.Transaction, error) {
	return _Govhub.contract.Transact(opts, "updateParam", key, value, target)
}

// UpdateParam is a paid mutator transaction binding the contract method 0x88e4194e.
//
// Solidity: function updateParam(string key, bytes value, address target) returns()
func (_Govhub *GovhubSession) UpdateParam(key string, value []byte, target common.Address) (*types.Transaction, error) {
	return _Govhub.Contract.UpdateParam(&_Govhub.TransactOpts, key, value, target)
}

// UpdateParam is a paid mutator transaction binding the contract method 0x88e4194e.
//
// Solidity: function updateParam(string key, bytes value, address target) returns()
func (_Govhub *GovhubTransactorSession) UpdateParam(key string, value []byte, target common.Address) (*types.Transaction, error) {
	return _Govhub.Contract.UpdateParam(&_Govhub.TransactOpts, key, value, target)
}

// GovhubFailReasonWithBytesIterator is returned from FilterFailReasonWithBytes and is used to iterate over the raw logs and unpacked data for FailReasonWithBytes events raised by the Govhub contract.
type GovhubFailReasonWithBytesIterator struct {
	Event *GovhubFailReasonWithBytes // Event containing the contract specifics and raw log

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
func (it *GovhubFailReasonWithBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovhubFailReasonWithBytes)
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
		it.Event = new(GovhubFailReasonWithBytes)
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
func (it *GovhubFailReasonWithBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovhubFailReasonWithBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovhubFailReasonWithBytes represents a FailReasonWithBytes event raised by the Govhub contract.
type GovhubFailReasonWithBytes struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailReasonWithBytes is a free log retrieval operation binding the contract event 0x1279f84165b4fd69c35e1f338ff107231b036c655cd1688851e011ce617c4e8d.
//
// Solidity: event failReasonWithBytes(bytes message)
func (_Govhub *GovhubFilterer) FilterFailReasonWithBytes(opts *bind.FilterOpts) (*GovhubFailReasonWithBytesIterator, error) {

	logs, sub, err := _Govhub.contract.FilterLogs(opts, "failReasonWithBytes")
	if err != nil {
		return nil, err
	}
	return &GovhubFailReasonWithBytesIterator{contract: _Govhub.contract, event: "failReasonWithBytes", logs: logs, sub: sub}, nil
}

// WatchFailReasonWithBytes is a free log subscription operation binding the contract event 0x1279f84165b4fd69c35e1f338ff107231b036c655cd1688851e011ce617c4e8d.
//
// Solidity: event failReasonWithBytes(bytes message)
func (_Govhub *GovhubFilterer) WatchFailReasonWithBytes(opts *bind.WatchOpts, sink chan<- *GovhubFailReasonWithBytes) (event.Subscription, error) {

	logs, sub, err := _Govhub.contract.WatchLogs(opts, "failReasonWithBytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovhubFailReasonWithBytes)
				if err := _Govhub.contract.UnpackLog(event, "failReasonWithBytes", log); err != nil {
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

// ParseFailReasonWithBytes is a log parse operation binding the contract event 0x1279f84165b4fd69c35e1f338ff107231b036c655cd1688851e011ce617c4e8d.
//
// Solidity: event failReasonWithBytes(bytes message)
func (_Govhub *GovhubFilterer) ParseFailReasonWithBytes(log types.Log) (*GovhubFailReasonWithBytes, error) {
	event := new(GovhubFailReasonWithBytes)
	if err := _Govhub.contract.UnpackLog(event, "failReasonWithBytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovhubFailReasonWithStrIterator is returned from FilterFailReasonWithStr and is used to iterate over the raw logs and unpacked data for FailReasonWithStr events raised by the Govhub contract.
type GovhubFailReasonWithStrIterator struct {
	Event *GovhubFailReasonWithStr // Event containing the contract specifics and raw log

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
func (it *GovhubFailReasonWithStrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovhubFailReasonWithStr)
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
		it.Event = new(GovhubFailReasonWithStr)
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
func (it *GovhubFailReasonWithStrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovhubFailReasonWithStrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovhubFailReasonWithStr represents a FailReasonWithStr event raised by the Govhub contract.
type GovhubFailReasonWithStr struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailReasonWithStr is a free log retrieval operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Govhub *GovhubFilterer) FilterFailReasonWithStr(opts *bind.FilterOpts) (*GovhubFailReasonWithStrIterator, error) {

	logs, sub, err := _Govhub.contract.FilterLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return &GovhubFailReasonWithStrIterator{contract: _Govhub.contract, event: "failReasonWithStr", logs: logs, sub: sub}, nil
}

// WatchFailReasonWithStr is a free log subscription operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Govhub *GovhubFilterer) WatchFailReasonWithStr(opts *bind.WatchOpts, sink chan<- *GovhubFailReasonWithStr) (event.Subscription, error) {

	logs, sub, err := _Govhub.contract.WatchLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovhubFailReasonWithStr)
				if err := _Govhub.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
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

// ParseFailReasonWithStr is a log parse operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Govhub *GovhubFilterer) ParseFailReasonWithStr(log types.Log) (*GovhubFailReasonWithStr, error) {
	event := new(GovhubFailReasonWithStr)
	if err := _Govhub.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovhubParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Govhub contract.
type GovhubParamChangeIterator struct {
	Event *GovhubParamChange // Event containing the contract specifics and raw log

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
func (it *GovhubParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovhubParamChange)
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
		it.Event = new(GovhubParamChange)
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
func (it *GovhubParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovhubParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovhubParamChange represents a ParamChange event raised by the Govhub contract.
type GovhubParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Govhub *GovhubFilterer) FilterParamChange(opts *bind.FilterOpts) (*GovhubParamChangeIterator, error) {

	logs, sub, err := _Govhub.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &GovhubParamChangeIterator{contract: _Govhub.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Govhub *GovhubFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *GovhubParamChange) (event.Subscription, error) {

	logs, sub, err := _Govhub.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovhubParamChange)
				if err := _Govhub.contract.UnpackLog(event, "paramChange", log); err != nil {
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
func (_Govhub *GovhubFilterer) ParseParamChange(log types.Log) (*GovhubParamChange, error) {
	event := new(GovhubParamChange)
	if err := _Govhub.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
