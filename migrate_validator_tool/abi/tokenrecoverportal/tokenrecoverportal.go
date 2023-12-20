// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenrecoverportal

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

// TokenrecoverportalMetaData contains all meta data concerning the Tokenrecoverportal contract.
var TokenrecoverportalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyRecovered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InBlackList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidApprovalSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwnerPubKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwnerSignatureLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleRootAlreadyInitiated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleRootNotInitialize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAssetProtector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCoinbase\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"systemContract\",\"type\":\"address\"}],\"name\":\"OnlySystemContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyZeroGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenRecoverPortalPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"UnknownParam\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"ParamChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenSymbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRecoverRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetProtector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenSymbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"attacker\",\"type\":\"address\"}],\"name\":\"cancelTokenRecover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"isRecovered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRootAlreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenSymbol\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"ownerPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ownerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approvalSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenrecoverportalABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenrecoverportalMetaData.ABI instead.
var TokenrecoverportalABI = TokenrecoverportalMetaData.ABI

// Tokenrecoverportal is an auto generated Go binding around an Ethereum contract.
type Tokenrecoverportal struct {
	TokenrecoverportalCaller     // Read-only binding to the contract
	TokenrecoverportalTransactor // Write-only binding to the contract
	TokenrecoverportalFilterer   // Log filterer for contract events
}

// TokenrecoverportalCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenrecoverportalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenrecoverportalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenrecoverportalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenrecoverportalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenrecoverportalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenrecoverportalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenrecoverportalSession struct {
	Contract     *Tokenrecoverportal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TokenrecoverportalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenrecoverportalCallerSession struct {
	Contract *TokenrecoverportalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TokenrecoverportalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenrecoverportalTransactorSession struct {
	Contract     *TokenrecoverportalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TokenrecoverportalRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenrecoverportalRaw struct {
	Contract *Tokenrecoverportal // Generic contract binding to access the raw methods on
}

// TokenrecoverportalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenrecoverportalCallerRaw struct {
	Contract *TokenrecoverportalCaller // Generic read-only contract binding to access the raw methods on
}

// TokenrecoverportalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenrecoverportalTransactorRaw struct {
	Contract *TokenrecoverportalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenrecoverportal creates a new instance of Tokenrecoverportal, bound to a specific deployed contract.
func NewTokenrecoverportal(address common.Address, backend bind.ContractBackend) (*Tokenrecoverportal, error) {
	contract, err := bindTokenrecoverportal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenrecoverportal{TokenrecoverportalCaller: TokenrecoverportalCaller{contract: contract}, TokenrecoverportalTransactor: TokenrecoverportalTransactor{contract: contract}, TokenrecoverportalFilterer: TokenrecoverportalFilterer{contract: contract}}, nil
}

// NewTokenrecoverportalCaller creates a new read-only instance of Tokenrecoverportal, bound to a specific deployed contract.
func NewTokenrecoverportalCaller(address common.Address, caller bind.ContractCaller) (*TokenrecoverportalCaller, error) {
	contract, err := bindTokenrecoverportal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalCaller{contract: contract}, nil
}

// NewTokenrecoverportalTransactor creates a new write-only instance of Tokenrecoverportal, bound to a specific deployed contract.
func NewTokenrecoverportalTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenrecoverportalTransactor, error) {
	contract, err := bindTokenrecoverportal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalTransactor{contract: contract}, nil
}

// NewTokenrecoverportalFilterer creates a new log filterer instance of Tokenrecoverportal, bound to a specific deployed contract.
func NewTokenrecoverportalFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenrecoverportalFilterer, error) {
	contract, err := bindTokenrecoverportal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalFilterer{contract: contract}, nil
}

// bindTokenrecoverportal binds a generic wrapper to an already deployed contract.
func bindTokenrecoverportal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenrecoverportalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenrecoverportal *TokenrecoverportalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenrecoverportal.Contract.TokenrecoverportalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenrecoverportal *TokenrecoverportalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.TokenrecoverportalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenrecoverportal *TokenrecoverportalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.TokenrecoverportalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenrecoverportal *TokenrecoverportalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenrecoverportal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenrecoverportal *TokenrecoverportalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenrecoverportal *TokenrecoverportalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.contract.Transact(opts, method, params...)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(string)
func (_Tokenrecoverportal *TokenrecoverportalCaller) SOURCECHAINID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(string)
func (_Tokenrecoverportal *TokenrecoverportalSession) SOURCECHAINID() (string, error) {
	return _Tokenrecoverportal.Contract.SOURCECHAINID(&_Tokenrecoverportal.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(string)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) SOURCECHAINID() (string, error) {
	return _Tokenrecoverportal.Contract.SOURCECHAINID(&_Tokenrecoverportal.CallOpts)
}

// ApprovalAddress is a free data retrieval call binding the contract method 0xe842426a.
//
// Solidity: function approvalAddress() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalCaller) ApprovalAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "approvalAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApprovalAddress is a free data retrieval call binding the contract method 0xe842426a.
//
// Solidity: function approvalAddress() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalSession) ApprovalAddress() (common.Address, error) {
	return _Tokenrecoverportal.Contract.ApprovalAddress(&_Tokenrecoverportal.CallOpts)
}

// ApprovalAddress is a free data retrieval call binding the contract method 0xe842426a.
//
// Solidity: function approvalAddress() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) ApprovalAddress() (common.Address, error) {
	return _Tokenrecoverportal.Contract.ApprovalAddress(&_Tokenrecoverportal.CallOpts)
}

// AssetProtector is a free data retrieval call binding the contract method 0xde88700b.
//
// Solidity: function assetProtector() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalCaller) AssetProtector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "assetProtector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetProtector is a free data retrieval call binding the contract method 0xde88700b.
//
// Solidity: function assetProtector() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalSession) AssetProtector() (common.Address, error) {
	return _Tokenrecoverportal.Contract.AssetProtector(&_Tokenrecoverportal.CallOpts)
}

// AssetProtector is a free data retrieval call binding the contract method 0xde88700b.
//
// Solidity: function assetProtector() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) AssetProtector() (common.Address, error) {
	return _Tokenrecoverportal.Contract.AssetProtector(&_Tokenrecoverportal.CallOpts)
}

// IsRecovered is a free data retrieval call binding the contract method 0xe33f8d32.
//
// Solidity: function isRecovered(bytes32 node) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCaller) IsRecovered(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "isRecovered", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRecovered is a free data retrieval call binding the contract method 0xe33f8d32.
//
// Solidity: function isRecovered(bytes32 node) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalSession) IsRecovered(node [32]byte) (bool, error) {
	return _Tokenrecoverportal.Contract.IsRecovered(&_Tokenrecoverportal.CallOpts, node)
}

// IsRecovered is a free data retrieval call binding the contract method 0xe33f8d32.
//
// Solidity: function isRecovered(bytes32 node) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) IsRecovered(node [32]byte) (bool, error) {
	return _Tokenrecoverportal.Contract.IsRecovered(&_Tokenrecoverportal.CallOpts, node)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Tokenrecoverportal *TokenrecoverportalCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Tokenrecoverportal *TokenrecoverportalSession) MerkleRoot() ([32]byte, error) {
	return _Tokenrecoverportal.Contract.MerkleRoot(&_Tokenrecoverportal.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) MerkleRoot() ([32]byte, error) {
	return _Tokenrecoverportal.Contract.MerkleRoot(&_Tokenrecoverportal.CallOpts)
}

// MerkleRootAlreadyInit is a free data retrieval call binding the contract method 0x9fcb5012.
//
// Solidity: function merkleRootAlreadyInit() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCaller) MerkleRootAlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "merkleRootAlreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MerkleRootAlreadyInit is a free data retrieval call binding the contract method 0x9fcb5012.
//
// Solidity: function merkleRootAlreadyInit() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalSession) MerkleRootAlreadyInit() (bool, error) {
	return _Tokenrecoverportal.Contract.MerkleRootAlreadyInit(&_Tokenrecoverportal.CallOpts)
}

// MerkleRootAlreadyInit is a free data retrieval call binding the contract method 0x9fcb5012.
//
// Solidity: function merkleRootAlreadyInit() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) MerkleRootAlreadyInit() (bool, error) {
	return _Tokenrecoverportal.Contract.MerkleRootAlreadyInit(&_Tokenrecoverportal.CallOpts)
}

// CancelTokenRecover is a paid mutator transaction binding the contract method 0x572c9980.
//
// Solidity: function cancelTokenRecover(bytes32 tokenSymbol, address attacker) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) CancelTokenRecover(opts *bind.TransactOpts, tokenSymbol [32]byte, attacker common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "cancelTokenRecover", tokenSymbol, attacker)
}

// CancelTokenRecover is a paid mutator transaction binding the contract method 0x572c9980.
//
// Solidity: function cancelTokenRecover(bytes32 tokenSymbol, address attacker) returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) CancelTokenRecover(tokenSymbol [32]byte, attacker common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.CancelTokenRecover(&_Tokenrecoverportal.TransactOpts, tokenSymbol, attacker)
}

// CancelTokenRecover is a paid mutator transaction binding the contract method 0x572c9980.
//
// Solidity: function cancelTokenRecover(bytes32 tokenSymbol, address attacker) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) CancelTokenRecover(tokenSymbol [32]byte, attacker common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.CancelTokenRecover(&_Tokenrecoverportal.TransactOpts, tokenSymbol, attacker)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) Pause() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Pause(&_Tokenrecoverportal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) Pause() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Pause(&_Tokenrecoverportal.TransactOpts)
}

// Recover is a paid mutator transaction binding the contract method 0xbfb5a6a1.
//
// Solidity: function recover(bytes32 tokenSymbol, uint256 amount, bytes ownerPubKey, bytes ownerSignature, bytes approvalSignature, bytes32[] merkleProof) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) Recover(opts *bind.TransactOpts, tokenSymbol [32]byte, amount *big.Int, ownerPubKey []byte, ownerSignature []byte, approvalSignature []byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "recover", tokenSymbol, amount, ownerPubKey, ownerSignature, approvalSignature, merkleProof)
}

// Recover is a paid mutator transaction binding the contract method 0xbfb5a6a1.
//
// Solidity: function recover(bytes32 tokenSymbol, uint256 amount, bytes ownerPubKey, bytes ownerSignature, bytes approvalSignature, bytes32[] merkleProof) returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) Recover(tokenSymbol [32]byte, amount *big.Int, ownerPubKey []byte, ownerSignature []byte, approvalSignature []byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Recover(&_Tokenrecoverportal.TransactOpts, tokenSymbol, amount, ownerPubKey, ownerSignature, approvalSignature, merkleProof)
}

// Recover is a paid mutator transaction binding the contract method 0xbfb5a6a1.
//
// Solidity: function recover(bytes32 tokenSymbol, uint256 amount, bytes ownerPubKey, bytes ownerSignature, bytes approvalSignature, bytes32[] merkleProof) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) Recover(tokenSymbol [32]byte, amount *big.Int, ownerPubKey []byte, ownerSignature []byte, approvalSignature []byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Recover(&_Tokenrecoverportal.TransactOpts, tokenSymbol, amount, ownerPubKey, ownerSignature, approvalSignature, merkleProof)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) Resume() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Resume(&_Tokenrecoverportal.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) Resume() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Resume(&_Tokenrecoverportal.TransactOpts)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.UpdateParam(&_Tokenrecoverportal.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.UpdateParam(&_Tokenrecoverportal.TransactOpts, key, value)
}

// TokenrecoverportalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tokenrecoverportal contract.
type TokenrecoverportalInitializedIterator struct {
	Event *TokenrecoverportalInitialized // Event containing the contract specifics and raw log

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
func (it *TokenrecoverportalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalInitialized)
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
		it.Event = new(TokenrecoverportalInitialized)
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
func (it *TokenrecoverportalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalInitialized represents a Initialized event raised by the Tokenrecoverportal contract.
type TokenrecoverportalInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenrecoverportalInitializedIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalInitializedIterator{contract: _Tokenrecoverportal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalInitialized) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalInitialized)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseInitialized(log types.Log) (*TokenrecoverportalInitialized, error) {
	event := new(TokenrecoverportalInitialized)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Tokenrecoverportal contract.
type TokenrecoverportalParamChangeIterator struct {
	Event *TokenrecoverportalParamChange // Event containing the contract specifics and raw log

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
func (it *TokenrecoverportalParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalParamChange)
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
		it.Event = new(TokenrecoverportalParamChange)
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
func (it *TokenrecoverportalParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalParamChange represents a ParamChange event raised by the Tokenrecoverportal contract.
type TokenrecoverportalParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterParamChange(opts *bind.FilterOpts) (*TokenrecoverportalParamChangeIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalParamChangeIterator{contract: _Tokenrecoverportal.contract, event: "ParamChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalParamChange) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalParamChange)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "ParamChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseParamChange(log types.Log) (*TokenrecoverportalParamChange, error) {
	event := new(TokenrecoverportalParamChange)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "ParamChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Tokenrecoverportal contract.
type TokenrecoverportalPausedIterator struct {
	Event *TokenrecoverportalPaused // Event containing the contract specifics and raw log

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
func (it *TokenrecoverportalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalPaused)
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
		it.Event = new(TokenrecoverportalPaused)
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
func (it *TokenrecoverportalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalPaused represents a Paused event raised by the Tokenrecoverportal contract.
type TokenrecoverportalPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterPaused(opts *bind.FilterOpts) (*TokenrecoverportalPausedIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalPausedIterator{contract: _Tokenrecoverportal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalPaused) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalPaused)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParsePaused(log types.Log) (*TokenrecoverportalPaused, error) {
	event := new(TokenrecoverportalPaused)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Tokenrecoverportal contract.
type TokenrecoverportalResumedIterator struct {
	Event *TokenrecoverportalResumed // Event containing the contract specifics and raw log

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
func (it *TokenrecoverportalResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalResumed)
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
		it.Event = new(TokenrecoverportalResumed)
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
func (it *TokenrecoverportalResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalResumed represents a Resumed event raised by the Tokenrecoverportal contract.
type TokenrecoverportalResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterResumed(opts *bind.FilterOpts) (*TokenrecoverportalResumedIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalResumedIterator{contract: _Tokenrecoverportal.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalResumed) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalResumed)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseResumed(log types.Log) (*TokenrecoverportalResumed, error) {
	event := new(TokenrecoverportalResumed)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalTokenRecoverRequestedIterator is returned from FilterTokenRecoverRequested and is used to iterate over the raw logs and unpacked data for TokenRecoverRequested events raised by the Tokenrecoverportal contract.
type TokenrecoverportalTokenRecoverRequestedIterator struct {
	Event *TokenrecoverportalTokenRecoverRequested // Event containing the contract specifics and raw log

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
func (it *TokenrecoverportalTokenRecoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalTokenRecoverRequested)
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
		it.Event = new(TokenrecoverportalTokenRecoverRequested)
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
func (it *TokenrecoverportalTokenRecoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalTokenRecoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalTokenRecoverRequested represents a TokenRecoverRequested event raised by the Tokenrecoverportal contract.
type TokenrecoverportalTokenRecoverRequested struct {
	TokenSymbol [32]byte
	Account     common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenRecoverRequested is a free log retrieval operation binding the contract event 0x717c40d7a8dbe67bb6371e964c24423645711d1624980e1c74e6956fc2108805.
//
// Solidity: event TokenRecoverRequested(bytes32 tokenSymbol, address account, uint256 amount)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterTokenRecoverRequested(opts *bind.FilterOpts) (*TokenrecoverportalTokenRecoverRequestedIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "TokenRecoverRequested")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalTokenRecoverRequestedIterator{contract: _Tokenrecoverportal.contract, event: "TokenRecoverRequested", logs: logs, sub: sub}, nil
}

// WatchTokenRecoverRequested is a free log subscription operation binding the contract event 0x717c40d7a8dbe67bb6371e964c24423645711d1624980e1c74e6956fc2108805.
//
// Solidity: event TokenRecoverRequested(bytes32 tokenSymbol, address account, uint256 amount)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchTokenRecoverRequested(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalTokenRecoverRequested) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "TokenRecoverRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalTokenRecoverRequested)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "TokenRecoverRequested", log); err != nil {
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

// ParseTokenRecoverRequested is a log parse operation binding the contract event 0x717c40d7a8dbe67bb6371e964c24423645711d1624980e1c74e6956fc2108805.
//
// Solidity: event TokenRecoverRequested(bytes32 tokenSymbol, address account, uint256 amount)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseTokenRecoverRequested(log types.Log) (*TokenrecoverportalTokenRecoverRequested, error) {
	event := new(TokenrecoverportalTokenRecoverRequested)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "TokenRecoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
