// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bep20Test

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

// Bep20TestMetaData contains all meta data concerning the Bep20Test contract.
var Bep20TestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintMore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Bep20TestABI is the input ABI used to generate the binding from.
// Deprecated: Use Bep20TestMetaData.ABI instead.
var Bep20TestABI = Bep20TestMetaData.ABI

// Bep20Test is an auto generated Go binding around an Ethereum contract.
type Bep20Test struct {
	Bep20TestCaller     // Read-only binding to the contract
	Bep20TestTransactor // Write-only binding to the contract
	Bep20TestFilterer   // Log filterer for contract events
}

// Bep20TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type Bep20TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bep20TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Bep20TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bep20TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bep20TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bep20TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bep20TestSession struct {
	Contract     *Bep20Test        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bep20TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bep20TestCallerSession struct {
	Contract *Bep20TestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Bep20TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bep20TestTransactorSession struct {
	Contract     *Bep20TestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Bep20TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type Bep20TestRaw struct {
	Contract *Bep20Test // Generic contract binding to access the raw methods on
}

// Bep20TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bep20TestCallerRaw struct {
	Contract *Bep20TestCaller // Generic read-only contract binding to access the raw methods on
}

// Bep20TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bep20TestTransactorRaw struct {
	Contract *Bep20TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBep20Test creates a new instance of Bep20Test, bound to a specific deployed contract.
func NewBep20Test(address common.Address, backend bind.ContractBackend) (*Bep20Test, error) {
	contract, err := bindBep20Test(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bep20Test{Bep20TestCaller: Bep20TestCaller{contract: contract}, Bep20TestTransactor: Bep20TestTransactor{contract: contract}, Bep20TestFilterer: Bep20TestFilterer{contract: contract}}, nil
}

// NewBep20TestCaller creates a new read-only instance of Bep20Test, bound to a specific deployed contract.
func NewBep20TestCaller(address common.Address, caller bind.ContractCaller) (*Bep20TestCaller, error) {
	contract, err := bindBep20Test(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bep20TestCaller{contract: contract}, nil
}

// NewBep20TestTransactor creates a new write-only instance of Bep20Test, bound to a specific deployed contract.
func NewBep20TestTransactor(address common.Address, transactor bind.ContractTransactor) (*Bep20TestTransactor, error) {
	contract, err := bindBep20Test(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bep20TestTransactor{contract: contract}, nil
}

// NewBep20TestFilterer creates a new log filterer instance of Bep20Test, bound to a specific deployed contract.
func NewBep20TestFilterer(address common.Address, filterer bind.ContractFilterer) (*Bep20TestFilterer, error) {
	contract, err := bindBep20Test(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bep20TestFilterer{contract: contract}, nil
}

// bindBep20Test binds a generic wrapper to an already deployed contract.
func bindBep20Test(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Bep20TestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bep20Test *Bep20TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bep20Test.Contract.Bep20TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bep20Test *Bep20TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bep20Test.Contract.Bep20TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bep20Test *Bep20TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bep20Test.Contract.Bep20TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bep20Test *Bep20TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bep20Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bep20Test *Bep20TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bep20Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bep20Test *Bep20TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bep20Test.Contract.contract.Transact(opts, method, params...)
}

// NewDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_Bep20Test *Bep20TestCaller) NewDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "_decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NewDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_Bep20Test *Bep20TestSession) NewDecimals() (uint8, error) {
	return _Bep20Test.Contract.NewDecimals(&_Bep20Test.CallOpts)
}

// NewDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_Bep20Test *Bep20TestCallerSession) NewDecimals() (uint8, error) {
	return _Bep20Test.Contract.NewDecimals(&_Bep20Test.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_Bep20Test *Bep20TestCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "_name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_Bep20Test *Bep20TestSession) Name() (string, error) {
	return _Bep20Test.Contract.Name(&_Bep20Test.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_Bep20Test *Bep20TestCallerSession) Name() (string, error) {
	return _Bep20Test.Contract.Name(&_Bep20Test.CallOpts)
}

// NewSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_Bep20Test *Bep20TestCaller) NewSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "_symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NewSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_Bep20Test *Bep20TestSession) NewSymbol() (string, error) {
	return _Bep20Test.Contract.NewSymbol(&_Bep20Test.CallOpts)
}

// NewSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_Bep20Test *Bep20TestCallerSession) NewSymbol() (string, error) {
	return _Bep20Test.Contract.NewSymbol(&_Bep20Test.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bep20Test *Bep20TestCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bep20Test *Bep20TestSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bep20Test.Contract.Allowance(&_Bep20Test.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bep20Test *Bep20TestCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bep20Test.Contract.Allowance(&_Bep20Test.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bep20Test *Bep20TestCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bep20Test *Bep20TestSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bep20Test.Contract.BalanceOf(&_Bep20Test.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bep20Test *Bep20TestCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bep20Test.Contract.BalanceOf(&_Bep20Test.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bep20Test *Bep20TestCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bep20Test *Bep20TestSession) Decimals() (uint8, error) {
	return _Bep20Test.Contract.Decimals(&_Bep20Test.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bep20Test *Bep20TestCallerSession) Decimals() (uint8, error) {
	return _Bep20Test.Contract.Decimals(&_Bep20Test.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Bep20Test *Bep20TestCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Bep20Test *Bep20TestSession) GetOwner() (common.Address, error) {
	return _Bep20Test.Contract.GetOwner(&_Bep20Test.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Bep20Test *Bep20TestCallerSession) GetOwner() (common.Address, error) {
	return _Bep20Test.Contract.GetOwner(&_Bep20Test.CallOpts)
}

// NewName is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bep20Test *Bep20TestCaller) NewName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NewName is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bep20Test *Bep20TestSession) NewName() (string, error) {
	return _Bep20Test.Contract.NewName(&_Bep20Test.CallOpts)
}

// NewName is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bep20Test *Bep20TestCallerSession) NewName() (string, error) {
	return _Bep20Test.Contract.NewName(&_Bep20Test.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bep20Test *Bep20TestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bep20Test *Bep20TestSession) Owner() (common.Address, error) {
	return _Bep20Test.Contract.Owner(&_Bep20Test.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bep20Test *Bep20TestCallerSession) Owner() (common.Address, error) {
	return _Bep20Test.Contract.Owner(&_Bep20Test.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bep20Test *Bep20TestCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bep20Test *Bep20TestSession) Symbol() (string, error) {
	return _Bep20Test.Contract.Symbol(&_Bep20Test.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bep20Test *Bep20TestCallerSession) Symbol() (string, error) {
	return _Bep20Test.Contract.Symbol(&_Bep20Test.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bep20Test *Bep20TestCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bep20Test.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bep20Test *Bep20TestSession) TotalSupply() (*big.Int, error) {
	return _Bep20Test.Contract.TotalSupply(&_Bep20Test.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bep20Test *Bep20TestCallerSession) TotalSupply() (*big.Int, error) {
	return _Bep20Test.Contract.TotalSupply(&_Bep20Test.CallOpts)
}

// MintMore is a paid mutator transaction binding the contract method 0x0255ac3f.
//
// Solidity: function MintMore(uint256 amount) returns()
func (_Bep20Test *Bep20TestTransactor) MintMore(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.contract.Transact(opts, "MintMore", amount)
}

// MintMore is a paid mutator transaction binding the contract method 0x0255ac3f.
//
// Solidity: function MintMore(uint256 amount) returns()
func (_Bep20Test *Bep20TestSession) MintMore(amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.MintMore(&_Bep20Test.TransactOpts, amount)
}

// MintMore is a paid mutator transaction binding the contract method 0x0255ac3f.
//
// Solidity: function MintMore(uint256 amount) returns()
func (_Bep20Test *Bep20TestTransactorSession) MintMore(amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.MintMore(&_Bep20Test.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.Approve(&_Bep20Test.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.Approve(&_Bep20Test.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bep20Test *Bep20TestTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bep20Test.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bep20Test *Bep20TestSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.DecreaseAllowance(&_Bep20Test.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bep20Test *Bep20TestTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.DecreaseAllowance(&_Bep20Test.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bep20Test *Bep20TestTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bep20Test.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bep20Test *Bep20TestSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.IncreaseAllowance(&_Bep20Test.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bep20Test *Bep20TestTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.IncreaseAllowance(&_Bep20Test.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bep20Test *Bep20TestTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bep20Test.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bep20Test *Bep20TestSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bep20Test.Contract.RenounceOwnership(&_Bep20Test.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bep20Test *Bep20TestTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bep20Test.Contract.RenounceOwnership(&_Bep20Test.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.Transfer(&_Bep20Test.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.Transfer(&_Bep20Test.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.TransferFrom(&_Bep20Test.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bep20Test *Bep20TestTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20Test.Contract.TransferFrom(&_Bep20Test.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bep20Test *Bep20TestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bep20Test.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bep20Test *Bep20TestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bep20Test.Contract.TransferOwnership(&_Bep20Test.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bep20Test *Bep20TestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bep20Test.Contract.TransferOwnership(&_Bep20Test.TransactOpts, newOwner)
}

// Bep20TestApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bep20Test contract.
type Bep20TestApprovalIterator struct {
	Event *Bep20TestApproval // Event containing the contract specifics and raw log

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
func (it *Bep20TestApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bep20TestApproval)
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
		it.Event = new(Bep20TestApproval)
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
func (it *Bep20TestApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bep20TestApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bep20TestApproval represents a Approval event raised by the Bep20Test contract.
type Bep20TestApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bep20Test *Bep20TestFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Bep20TestApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bep20Test.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Bep20TestApprovalIterator{contract: _Bep20Test.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bep20Test *Bep20TestFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Bep20TestApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bep20Test.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bep20TestApproval)
				if err := _Bep20Test.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bep20Test *Bep20TestFilterer) ParseApproval(log types.Log) (*Bep20TestApproval, error) {
	event := new(Bep20TestApproval)
	if err := _Bep20Test.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bep20TestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bep20Test contract.
type Bep20TestOwnershipTransferredIterator struct {
	Event *Bep20TestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Bep20TestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bep20TestOwnershipTransferred)
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
		it.Event = new(Bep20TestOwnershipTransferred)
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
func (it *Bep20TestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bep20TestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bep20TestOwnershipTransferred represents a OwnershipTransferred event raised by the Bep20Test contract.
type Bep20TestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bep20Test *Bep20TestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Bep20TestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bep20Test.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Bep20TestOwnershipTransferredIterator{contract: _Bep20Test.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bep20Test *Bep20TestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Bep20TestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bep20Test.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bep20TestOwnershipTransferred)
				if err := _Bep20Test.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bep20Test *Bep20TestFilterer) ParseOwnershipTransferred(log types.Log) (*Bep20TestOwnershipTransferred, error) {
	event := new(Bep20TestOwnershipTransferred)
	if err := _Bep20Test.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bep20TestTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bep20Test contract.
type Bep20TestTransferIterator struct {
	Event *Bep20TestTransfer // Event containing the contract specifics and raw log

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
func (it *Bep20TestTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bep20TestTransfer)
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
		it.Event = new(Bep20TestTransfer)
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
func (it *Bep20TestTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bep20TestTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bep20TestTransfer represents a Transfer event raised by the Bep20Test contract.
type Bep20TestTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bep20Test *Bep20TestFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Bep20TestTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bep20Test.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Bep20TestTransferIterator{contract: _Bep20Test.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bep20Test *Bep20TestFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Bep20TestTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bep20Test.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bep20TestTransfer)
				if err := _Bep20Test.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bep20Test *Bep20TestFilterer) ParseTransfer(log types.Log) (*Bep20TestTransfer, error) {
	event := new(Bep20TestTransfer)
	if err := _Bep20Test.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
