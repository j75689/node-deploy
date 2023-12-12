// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakehub

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

// StakeHubCommission is an auto generated low-level Go binding around an user-defined struct.
type StakeHubCommission struct {
	Rate          uint64
	MaxRate       uint64
	MaxChangeRate uint64
}

// StakeHubDescription is an auto generated low-level Go binding around an user-defined struct.
type StakeHubDescription struct {
	Moniker  string
	Identity string
	Website  string
	Details  string
}

// StakehubMetaData contains all meta data concerning the Stakehub contract.
var StakehubMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEAD_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_LOCK_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REDELEGATE_FEE_RATE_BASE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assetProtector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blackList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createValidator\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"voteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxChangeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"description\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegateVotePower\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distributeReward\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"doubleSignSlash\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"downtimeJailTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"downtimeSlash\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"downtimeSlashAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"editCommissionRate\",\"inputs\":[{\"name\":\"commissionRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"editConsensusAddress\",\"inputs\":[{\"name\":\"newConsensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"editDescription\",\"inputs\":[{\"name\":\"description\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"editVoteAddress\",\"inputs\":[{\"name\":\"newVoteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"felonyJailTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"felonySlashAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAddressByConsensusAddress\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAddressByVoteAddress\",\"inputs\":[{\"name\":\"voteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorBasicInfo\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"voteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"jailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"jailUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCommission\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxChangeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDescription\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structStakeHub.Description\",\"components\":[{\"name\":\"moniker\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"identity\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorElectionInfo\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"consensusAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"votingPowers\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"voteAddrs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorRewardRecord\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dayIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorTotalPooledBNBRecord\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dayIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maliciousVoteSlash\",\"inputs\":[{\"name\":\"_voteAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxElectedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minDelegationBNBChange\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minSelfDelegationBNB\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfJailed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redelegate\",\"inputs\":[{\"name\":\"srcValidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dstValidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateVotePower\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redelegateFeeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFromBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"syncGovToken\",\"inputs\":[{\"name\":\"operatorAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unbondPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unjail\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommissionRateEdited\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"commissionRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConsensusAddressEdited\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newConsensusAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delegated\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DescriptionEdited\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParamChange\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Redelegated\",\"inputs\":[{\"name\":\"srcValidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dstValidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Resumed\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardDistributeFailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"failReason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardDistributed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Undelegated\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorCreated\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creditContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"voteAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorEmptyJailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorJailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSlashed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jailUntil\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slashType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumStakeHub.SlashType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnjailed\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteAddressEdited\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newVoteAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadySlashed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DelegationAmountTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateConsensusAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateVoteAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InBlackList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommission\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMoniker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVoteAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"JailTimeNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMoreFelonyToday\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAssetProtector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelfDelegation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"systemContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SelfDelegationNotEnough\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakeHubPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UpdateTooFrequently\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorExisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotJailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroShares\",\"inputs\":[]}]",
}

// StakehubABI is the input ABI used to generate the binding from.
// Deprecated: Use StakehubMetaData.ABI instead.
var StakehubABI = StakehubMetaData.ABI

// Stakehub is an auto generated Go binding around an Ethereum contract.
type Stakehub struct {
	StakehubCaller     // Read-only binding to the contract
	StakehubTransactor // Write-only binding to the contract
	StakehubFilterer   // Log filterer for contract events
}

// StakehubCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakehubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakehubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakehubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakehubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakehubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakehubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakehubSession struct {
	Contract     *Stakehub         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakehubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakehubCallerSession struct {
	Contract *StakehubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StakehubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakehubTransactorSession struct {
	Contract     *StakehubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakehubRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakehubRaw struct {
	Contract *Stakehub // Generic contract binding to access the raw methods on
}

// StakehubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakehubCallerRaw struct {
	Contract *StakehubCaller // Generic read-only contract binding to access the raw methods on
}

// StakehubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakehubTransactorRaw struct {
	Contract *StakehubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakehub creates a new instance of Stakehub, bound to a specific deployed contract.
func NewStakehub(address common.Address, backend bind.ContractBackend) (*Stakehub, error) {
	contract, err := bindStakehub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakehub{StakehubCaller: StakehubCaller{contract: contract}, StakehubTransactor: StakehubTransactor{contract: contract}, StakehubFilterer: StakehubFilterer{contract: contract}}, nil
}

// NewStakehubCaller creates a new read-only instance of Stakehub, bound to a specific deployed contract.
func NewStakehubCaller(address common.Address, caller bind.ContractCaller) (*StakehubCaller, error) {
	contract, err := bindStakehub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakehubCaller{contract: contract}, nil
}

// NewStakehubTransactor creates a new write-only instance of Stakehub, bound to a specific deployed contract.
func NewStakehubTransactor(address common.Address, transactor bind.ContractTransactor) (*StakehubTransactor, error) {
	contract, err := bindStakehub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakehubTransactor{contract: contract}, nil
}

// NewStakehubFilterer creates a new log filterer instance of Stakehub, bound to a specific deployed contract.
func NewStakehubFilterer(address common.Address, filterer bind.ContractFilterer) (*StakehubFilterer, error) {
	contract, err := bindStakehub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakehubFilterer{contract: contract}, nil
}

// bindStakehub binds a generic wrapper to an already deployed contract.
func bindStakehub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakehubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakehub *StakehubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakehub.Contract.StakehubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakehub *StakehubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakehub.Contract.StakehubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakehub *StakehubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakehub.Contract.StakehubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakehub *StakehubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakehub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakehub *StakehubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakehub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakehub *StakehubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakehub.Contract.contract.Transact(opts, method, params...)
}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Stakehub *StakehubCaller) DEADADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "DEAD_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Stakehub *StakehubSession) DEADADDRESS() (common.Address, error) {
	return _Stakehub.Contract.DEADADDRESS(&_Stakehub.CallOpts)
}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Stakehub *StakehubCallerSession) DEADADDRESS() (common.Address, error) {
	return _Stakehub.Contract.DEADADDRESS(&_Stakehub.CallOpts)
}

// INITLOCKAMOUNT is a free data retrieval call binding the contract method 0xae0f336d.
//
// Solidity: function INIT_LOCK_AMOUNT() view returns(uint256)
func (_Stakehub *StakehubCaller) INITLOCKAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "INIT_LOCK_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITLOCKAMOUNT is a free data retrieval call binding the contract method 0xae0f336d.
//
// Solidity: function INIT_LOCK_AMOUNT() view returns(uint256)
func (_Stakehub *StakehubSession) INITLOCKAMOUNT() (*big.Int, error) {
	return _Stakehub.Contract.INITLOCKAMOUNT(&_Stakehub.CallOpts)
}

// INITLOCKAMOUNT is a free data retrieval call binding the contract method 0xae0f336d.
//
// Solidity: function INIT_LOCK_AMOUNT() view returns(uint256)
func (_Stakehub *StakehubCallerSession) INITLOCKAMOUNT() (*big.Int, error) {
	return _Stakehub.Contract.INITLOCKAMOUNT(&_Stakehub.CallOpts)
}

// REDELEGATEFEERATEBASE is a free data retrieval call binding the contract method 0xd115a206.
//
// Solidity: function REDELEGATE_FEE_RATE_BASE() view returns(uint256)
func (_Stakehub *StakehubCaller) REDELEGATEFEERATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "REDELEGATE_FEE_RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REDELEGATEFEERATEBASE is a free data retrieval call binding the contract method 0xd115a206.
//
// Solidity: function REDELEGATE_FEE_RATE_BASE() view returns(uint256)
func (_Stakehub *StakehubSession) REDELEGATEFEERATEBASE() (*big.Int, error) {
	return _Stakehub.Contract.REDELEGATEFEERATEBASE(&_Stakehub.CallOpts)
}

// REDELEGATEFEERATEBASE is a free data retrieval call binding the contract method 0xd115a206.
//
// Solidity: function REDELEGATE_FEE_RATE_BASE() view returns(uint256)
func (_Stakehub *StakehubCallerSession) REDELEGATEFEERATEBASE() (*big.Int, error) {
	return _Stakehub.Contract.REDELEGATEFEERATEBASE(&_Stakehub.CallOpts)
}

// AssetProtector is a free data retrieval call binding the contract method 0xde88700b.
//
// Solidity: function assetProtector() view returns(address)
func (_Stakehub *StakehubCaller) AssetProtector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "assetProtector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetProtector is a free data retrieval call binding the contract method 0xde88700b.
//
// Solidity: function assetProtector() view returns(address)
func (_Stakehub *StakehubSession) AssetProtector() (common.Address, error) {
	return _Stakehub.Contract.AssetProtector(&_Stakehub.CallOpts)
}

// AssetProtector is a free data retrieval call binding the contract method 0xde88700b.
//
// Solidity: function assetProtector() view returns(address)
func (_Stakehub *StakehubCallerSession) AssetProtector() (common.Address, error) {
	return _Stakehub.Contract.AssetProtector(&_Stakehub.CallOpts)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Stakehub *StakehubCaller) BlackList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "blackList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Stakehub *StakehubSession) BlackList(arg0 common.Address) (bool, error) {
	return _Stakehub.Contract.BlackList(&_Stakehub.CallOpts, arg0)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Stakehub *StakehubCallerSession) BlackList(arg0 common.Address) (bool, error) {
	return _Stakehub.Contract.BlackList(&_Stakehub.CallOpts, arg0)
}

// DowntimeJailTime is a free data retrieval call binding the contract method 0x76e7d6d6.
//
// Solidity: function downtimeJailTime() view returns(uint256)
func (_Stakehub *StakehubCaller) DowntimeJailTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "downtimeJailTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DowntimeJailTime is a free data retrieval call binding the contract method 0x76e7d6d6.
//
// Solidity: function downtimeJailTime() view returns(uint256)
func (_Stakehub *StakehubSession) DowntimeJailTime() (*big.Int, error) {
	return _Stakehub.Contract.DowntimeJailTime(&_Stakehub.CallOpts)
}

// DowntimeJailTime is a free data retrieval call binding the contract method 0x76e7d6d6.
//
// Solidity: function downtimeJailTime() view returns(uint256)
func (_Stakehub *StakehubCallerSession) DowntimeJailTime() (*big.Int, error) {
	return _Stakehub.Contract.DowntimeJailTime(&_Stakehub.CallOpts)
}

// DowntimeSlashAmount is a free data retrieval call binding the contract method 0xd8ca511f.
//
// Solidity: function downtimeSlashAmount() view returns(uint256)
func (_Stakehub *StakehubCaller) DowntimeSlashAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "downtimeSlashAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DowntimeSlashAmount is a free data retrieval call binding the contract method 0xd8ca511f.
//
// Solidity: function downtimeSlashAmount() view returns(uint256)
func (_Stakehub *StakehubSession) DowntimeSlashAmount() (*big.Int, error) {
	return _Stakehub.Contract.DowntimeSlashAmount(&_Stakehub.CallOpts)
}

// DowntimeSlashAmount is a free data retrieval call binding the contract method 0xd8ca511f.
//
// Solidity: function downtimeSlashAmount() view returns(uint256)
func (_Stakehub *StakehubCallerSession) DowntimeSlashAmount() (*big.Int, error) {
	return _Stakehub.Contract.DowntimeSlashAmount(&_Stakehub.CallOpts)
}

// FelonyJailTime is a free data retrieval call binding the contract method 0xf1f74d84.
//
// Solidity: function felonyJailTime() view returns(uint256)
func (_Stakehub *StakehubCaller) FelonyJailTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "felonyJailTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FelonyJailTime is a free data retrieval call binding the contract method 0xf1f74d84.
//
// Solidity: function felonyJailTime() view returns(uint256)
func (_Stakehub *StakehubSession) FelonyJailTime() (*big.Int, error) {
	return _Stakehub.Contract.FelonyJailTime(&_Stakehub.CallOpts)
}

// FelonyJailTime is a free data retrieval call binding the contract method 0xf1f74d84.
//
// Solidity: function felonyJailTime() view returns(uint256)
func (_Stakehub *StakehubCallerSession) FelonyJailTime() (*big.Int, error) {
	return _Stakehub.Contract.FelonyJailTime(&_Stakehub.CallOpts)
}

// FelonySlashAmount is a free data retrieval call binding the contract method 0xbdceadf3.
//
// Solidity: function felonySlashAmount() view returns(uint256)
func (_Stakehub *StakehubCaller) FelonySlashAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "felonySlashAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FelonySlashAmount is a free data retrieval call binding the contract method 0xbdceadf3.
//
// Solidity: function felonySlashAmount() view returns(uint256)
func (_Stakehub *StakehubSession) FelonySlashAmount() (*big.Int, error) {
	return _Stakehub.Contract.FelonySlashAmount(&_Stakehub.CallOpts)
}

// FelonySlashAmount is a free data retrieval call binding the contract method 0xbdceadf3.
//
// Solidity: function felonySlashAmount() view returns(uint256)
func (_Stakehub *StakehubCallerSession) FelonySlashAmount() (*big.Int, error) {
	return _Stakehub.Contract.FelonySlashAmount(&_Stakehub.CallOpts)
}

// GetOperatorAddressByConsensusAddress is a free data retrieval call binding the contract method 0xcf06248c.
//
// Solidity: function getOperatorAddressByConsensusAddress(address consensusAddress) view returns(address)
func (_Stakehub *StakehubCaller) GetOperatorAddressByConsensusAddress(opts *bind.CallOpts, consensusAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "getOperatorAddressByConsensusAddress", consensusAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorAddressByConsensusAddress is a free data retrieval call binding the contract method 0xcf06248c.
//
// Solidity: function getOperatorAddressByConsensusAddress(address consensusAddress) view returns(address)
func (_Stakehub *StakehubSession) GetOperatorAddressByConsensusAddress(consensusAddress common.Address) (common.Address, error) {
	return _Stakehub.Contract.GetOperatorAddressByConsensusAddress(&_Stakehub.CallOpts, consensusAddress)
}

// GetOperatorAddressByConsensusAddress is a free data retrieval call binding the contract method 0xcf06248c.
//
// Solidity: function getOperatorAddressByConsensusAddress(address consensusAddress) view returns(address)
func (_Stakehub *StakehubCallerSession) GetOperatorAddressByConsensusAddress(consensusAddress common.Address) (common.Address, error) {
	return _Stakehub.Contract.GetOperatorAddressByConsensusAddress(&_Stakehub.CallOpts, consensusAddress)
}

// GetOperatorAddressByVoteAddress is a free data retrieval call binding the contract method 0x7adcf7e8.
//
// Solidity: function getOperatorAddressByVoteAddress(bytes voteAddress) view returns(address)
func (_Stakehub *StakehubCaller) GetOperatorAddressByVoteAddress(opts *bind.CallOpts, voteAddress []byte) (common.Address, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "getOperatorAddressByVoteAddress", voteAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorAddressByVoteAddress is a free data retrieval call binding the contract method 0x7adcf7e8.
//
// Solidity: function getOperatorAddressByVoteAddress(bytes voteAddress) view returns(address)
func (_Stakehub *StakehubSession) GetOperatorAddressByVoteAddress(voteAddress []byte) (common.Address, error) {
	return _Stakehub.Contract.GetOperatorAddressByVoteAddress(&_Stakehub.CallOpts, voteAddress)
}

// GetOperatorAddressByVoteAddress is a free data retrieval call binding the contract method 0x7adcf7e8.
//
// Solidity: function getOperatorAddressByVoteAddress(bytes voteAddress) view returns(address)
func (_Stakehub *StakehubCallerSession) GetOperatorAddressByVoteAddress(voteAddress []byte) (common.Address, error) {
	return _Stakehub.Contract.GetOperatorAddressByVoteAddress(&_Stakehub.CallOpts, voteAddress)
}

// GetValidatorBasicInfo is a free data retrieval call binding the contract method 0xcbb04d9d.
//
// Solidity: function getValidatorBasicInfo(address operatorAddress) view returns(address consensusAddress, address creditContract, bytes voteAddress, bool jailed, uint256 jailUntil)
func (_Stakehub *StakehubCaller) GetValidatorBasicInfo(opts *bind.CallOpts, operatorAddress common.Address) (struct {
	ConsensusAddress common.Address
	CreditContract   common.Address
	VoteAddress      []byte
	Jailed           bool
	JailUntil        *big.Int
}, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "getValidatorBasicInfo", operatorAddress)

	outstruct := new(struct {
		ConsensusAddress common.Address
		CreditContract   common.Address
		VoteAddress      []byte
		Jailed           bool
		JailUntil        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CreditContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VoteAddress = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Jailed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.JailUntil = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorBasicInfo is a free data retrieval call binding the contract method 0xcbb04d9d.
//
// Solidity: function getValidatorBasicInfo(address operatorAddress) view returns(address consensusAddress, address creditContract, bytes voteAddress, bool jailed, uint256 jailUntil)
func (_Stakehub *StakehubSession) GetValidatorBasicInfo(operatorAddress common.Address) (struct {
	ConsensusAddress common.Address
	CreditContract   common.Address
	VoteAddress      []byte
	Jailed           bool
	JailUntil        *big.Int
}, error) {
	return _Stakehub.Contract.GetValidatorBasicInfo(&_Stakehub.CallOpts, operatorAddress)
}

// GetValidatorBasicInfo is a free data retrieval call binding the contract method 0xcbb04d9d.
//
// Solidity: function getValidatorBasicInfo(address operatorAddress) view returns(address consensusAddress, address creditContract, bytes voteAddress, bool jailed, uint256 jailUntil)
func (_Stakehub *StakehubCallerSession) GetValidatorBasicInfo(operatorAddress common.Address) (struct {
	ConsensusAddress common.Address
	CreditContract   common.Address
	VoteAddress      []byte
	Jailed           bool
	JailUntil        *big.Int
}, error) {
	return _Stakehub.Contract.GetValidatorBasicInfo(&_Stakehub.CallOpts, operatorAddress)
}

// GetValidatorCommission is a free data retrieval call binding the contract method 0x6ec01b27.
//
// Solidity: function getValidatorCommission(address operatorAddress) view returns((uint64,uint64,uint64))
func (_Stakehub *StakehubCaller) GetValidatorCommission(opts *bind.CallOpts, operatorAddress common.Address) (StakeHubCommission, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "getValidatorCommission", operatorAddress)

	if err != nil {
		return *new(StakeHubCommission), err
	}

	out0 := *abi.ConvertType(out[0], new(StakeHubCommission)).(*StakeHubCommission)

	return out0, err

}

// GetValidatorCommission is a free data retrieval call binding the contract method 0x6ec01b27.
//
// Solidity: function getValidatorCommission(address operatorAddress) view returns((uint64,uint64,uint64))
func (_Stakehub *StakehubSession) GetValidatorCommission(operatorAddress common.Address) (StakeHubCommission, error) {
	return _Stakehub.Contract.GetValidatorCommission(&_Stakehub.CallOpts, operatorAddress)
}

// GetValidatorCommission is a free data retrieval call binding the contract method 0x6ec01b27.
//
// Solidity: function getValidatorCommission(address operatorAddress) view returns((uint64,uint64,uint64))
func (_Stakehub *StakehubCallerSession) GetValidatorCommission(operatorAddress common.Address) (StakeHubCommission, error) {
	return _Stakehub.Contract.GetValidatorCommission(&_Stakehub.CallOpts, operatorAddress)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address operatorAddress) view returns((string,string,string,string))
func (_Stakehub *StakehubCaller) GetValidatorDescription(opts *bind.CallOpts, operatorAddress common.Address) (StakeHubDescription, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "getValidatorDescription", operatorAddress)

	if err != nil {
		return *new(StakeHubDescription), err
	}

	out0 := *abi.ConvertType(out[0], new(StakeHubDescription)).(*StakeHubDescription)

	return out0, err

}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address operatorAddress) view returns((string,string,string,string))
func (_Stakehub *StakehubSession) GetValidatorDescription(operatorAddress common.Address) (StakeHubDescription, error) {
	return _Stakehub.Contract.GetValidatorDescription(&_Stakehub.CallOpts, operatorAddress)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address operatorAddress) view returns((string,string,string,string))
func (_Stakehub *StakehubCallerSession) GetValidatorDescription(operatorAddress common.Address) (StakeHubDescription, error) {
	return _Stakehub.Contract.GetValidatorDescription(&_Stakehub.CallOpts, operatorAddress)
}

// GetValidatorElectionInfo is a free data retrieval call binding the contract method 0x63a036b5.
//
// Solidity: function getValidatorElectionInfo(uint256 offset, uint256 limit) view returns(address[] consensusAddrs, uint256[] votingPowers, bytes[] voteAddrs, uint256 totalLength)
func (_Stakehub *StakehubCaller) GetValidatorElectionInfo(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	ConsensusAddrs []common.Address
	VotingPowers   []*big.Int
	VoteAddrs      [][]byte
	TotalLength    *big.Int
}, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "getValidatorElectionInfo", offset, limit)

	outstruct := new(struct {
		ConsensusAddrs []common.Address
		VotingPowers   []*big.Int
		VoteAddrs      [][]byte
		TotalLength    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.VotingPowers = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.VoteAddrs = *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)
	outstruct.TotalLength = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorElectionInfo is a free data retrieval call binding the contract method 0x63a036b5.
//
// Solidity: function getValidatorElectionInfo(uint256 offset, uint256 limit) view returns(address[] consensusAddrs, uint256[] votingPowers, bytes[] voteAddrs, uint256 totalLength)
func (_Stakehub *StakehubSession) GetValidatorElectionInfo(offset *big.Int, limit *big.Int) (struct {
	ConsensusAddrs []common.Address
	VotingPowers   []*big.Int
	VoteAddrs      [][]byte
	TotalLength    *big.Int
}, error) {
	return _Stakehub.Contract.GetValidatorElectionInfo(&_Stakehub.CallOpts, offset, limit)
}

// GetValidatorElectionInfo is a free data retrieval call binding the contract method 0x63a036b5.
//
// Solidity: function getValidatorElectionInfo(uint256 offset, uint256 limit) view returns(address[] consensusAddrs, uint256[] votingPowers, bytes[] voteAddrs, uint256 totalLength)
func (_Stakehub *StakehubCallerSession) GetValidatorElectionInfo(offset *big.Int, limit *big.Int) (struct {
	ConsensusAddrs []common.Address
	VotingPowers   []*big.Int
	VoteAddrs      [][]byte
	TotalLength    *big.Int
}, error) {
	return _Stakehub.Contract.GetValidatorElectionInfo(&_Stakehub.CallOpts, offset, limit)
}

// GetValidatorRewardRecord is a free data retrieval call binding the contract method 0xf80a3402.
//
// Solidity: function getValidatorRewardRecord(address operatorAddress, uint256 dayIndex) view returns(uint256)
func (_Stakehub *StakehubCaller) GetValidatorRewardRecord(opts *bind.CallOpts, operatorAddress common.Address, dayIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "getValidatorRewardRecord", operatorAddress, dayIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorRewardRecord is a free data retrieval call binding the contract method 0xf80a3402.
//
// Solidity: function getValidatorRewardRecord(address operatorAddress, uint256 dayIndex) view returns(uint256)
func (_Stakehub *StakehubSession) GetValidatorRewardRecord(operatorAddress common.Address, dayIndex *big.Int) (*big.Int, error) {
	return _Stakehub.Contract.GetValidatorRewardRecord(&_Stakehub.CallOpts, operatorAddress, dayIndex)
}

// GetValidatorRewardRecord is a free data retrieval call binding the contract method 0xf80a3402.
//
// Solidity: function getValidatorRewardRecord(address operatorAddress, uint256 dayIndex) view returns(uint256)
func (_Stakehub *StakehubCallerSession) GetValidatorRewardRecord(operatorAddress common.Address, dayIndex *big.Int) (*big.Int, error) {
	return _Stakehub.Contract.GetValidatorRewardRecord(&_Stakehub.CallOpts, operatorAddress, dayIndex)
}

// GetValidatorTotalPooledBNBRecord is a free data retrieval call binding the contract method 0x8cd22b22.
//
// Solidity: function getValidatorTotalPooledBNBRecord(address operatorAddress, uint256 dayIndex) view returns(uint256)
func (_Stakehub *StakehubCaller) GetValidatorTotalPooledBNBRecord(opts *bind.CallOpts, operatorAddress common.Address, dayIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "getValidatorTotalPooledBNBRecord", operatorAddress, dayIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorTotalPooledBNBRecord is a free data retrieval call binding the contract method 0x8cd22b22.
//
// Solidity: function getValidatorTotalPooledBNBRecord(address operatorAddress, uint256 dayIndex) view returns(uint256)
func (_Stakehub *StakehubSession) GetValidatorTotalPooledBNBRecord(operatorAddress common.Address, dayIndex *big.Int) (*big.Int, error) {
	return _Stakehub.Contract.GetValidatorTotalPooledBNBRecord(&_Stakehub.CallOpts, operatorAddress, dayIndex)
}

// GetValidatorTotalPooledBNBRecord is a free data retrieval call binding the contract method 0x8cd22b22.
//
// Solidity: function getValidatorTotalPooledBNBRecord(address operatorAddress, uint256 dayIndex) view returns(uint256)
func (_Stakehub *StakehubCallerSession) GetValidatorTotalPooledBNBRecord(operatorAddress common.Address, dayIndex *big.Int) (*big.Int, error) {
	return _Stakehub.Contract.GetValidatorTotalPooledBNBRecord(&_Stakehub.CallOpts, operatorAddress, dayIndex)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Stakehub *StakehubCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Stakehub *StakehubSession) IsPaused() (bool, error) {
	return _Stakehub.Contract.IsPaused(&_Stakehub.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Stakehub *StakehubCallerSession) IsPaused() (bool, error) {
	return _Stakehub.Contract.IsPaused(&_Stakehub.CallOpts)
}

// MaxElectedValidators is a free data retrieval call binding the contract method 0xc473318f.
//
// Solidity: function maxElectedValidators() view returns(uint256)
func (_Stakehub *StakehubCaller) MaxElectedValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "maxElectedValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxElectedValidators is a free data retrieval call binding the contract method 0xc473318f.
//
// Solidity: function maxElectedValidators() view returns(uint256)
func (_Stakehub *StakehubSession) MaxElectedValidators() (*big.Int, error) {
	return _Stakehub.Contract.MaxElectedValidators(&_Stakehub.CallOpts)
}

// MaxElectedValidators is a free data retrieval call binding the contract method 0xc473318f.
//
// Solidity: function maxElectedValidators() view returns(uint256)
func (_Stakehub *StakehubCallerSession) MaxElectedValidators() (*big.Int, error) {
	return _Stakehub.Contract.MaxElectedValidators(&_Stakehub.CallOpts)
}

// MinDelegationBNBChange is a free data retrieval call binding the contract method 0x38409988.
//
// Solidity: function minDelegationBNBChange() view returns(uint256)
func (_Stakehub *StakehubCaller) MinDelegationBNBChange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "minDelegationBNBChange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegationBNBChange is a free data retrieval call binding the contract method 0x38409988.
//
// Solidity: function minDelegationBNBChange() view returns(uint256)
func (_Stakehub *StakehubSession) MinDelegationBNBChange() (*big.Int, error) {
	return _Stakehub.Contract.MinDelegationBNBChange(&_Stakehub.CallOpts)
}

// MinDelegationBNBChange is a free data retrieval call binding the contract method 0x38409988.
//
// Solidity: function minDelegationBNBChange() view returns(uint256)
func (_Stakehub *StakehubCallerSession) MinDelegationBNBChange() (*big.Int, error) {
	return _Stakehub.Contract.MinDelegationBNBChange(&_Stakehub.CallOpts)
}

// MinSelfDelegationBNB is a free data retrieval call binding the contract method 0x0661806e.
//
// Solidity: function minSelfDelegationBNB() view returns(uint256)
func (_Stakehub *StakehubCaller) MinSelfDelegationBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "minSelfDelegationBNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSelfDelegationBNB is a free data retrieval call binding the contract method 0x0661806e.
//
// Solidity: function minSelfDelegationBNB() view returns(uint256)
func (_Stakehub *StakehubSession) MinSelfDelegationBNB() (*big.Int, error) {
	return _Stakehub.Contract.MinSelfDelegationBNB(&_Stakehub.CallOpts)
}

// MinSelfDelegationBNB is a free data retrieval call binding the contract method 0x0661806e.
//
// Solidity: function minSelfDelegationBNB() view returns(uint256)
func (_Stakehub *StakehubCallerSession) MinSelfDelegationBNB() (*big.Int, error) {
	return _Stakehub.Contract.MinSelfDelegationBNB(&_Stakehub.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Stakehub *StakehubCaller) NumOfJailed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "numOfJailed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Stakehub *StakehubSession) NumOfJailed() (*big.Int, error) {
	return _Stakehub.Contract.NumOfJailed(&_Stakehub.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Stakehub *StakehubCallerSession) NumOfJailed() (*big.Int, error) {
	return _Stakehub.Contract.NumOfJailed(&_Stakehub.CallOpts)
}

// RedelegateFeeRate is a free data retrieval call binding the contract method 0xe992aaf5.
//
// Solidity: function redelegateFeeRate() view returns(uint256)
func (_Stakehub *StakehubCaller) RedelegateFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "redelegateFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedelegateFeeRate is a free data retrieval call binding the contract method 0xe992aaf5.
//
// Solidity: function redelegateFeeRate() view returns(uint256)
func (_Stakehub *StakehubSession) RedelegateFeeRate() (*big.Int, error) {
	return _Stakehub.Contract.RedelegateFeeRate(&_Stakehub.CallOpts)
}

// RedelegateFeeRate is a free data retrieval call binding the contract method 0xe992aaf5.
//
// Solidity: function redelegateFeeRate() view returns(uint256)
func (_Stakehub *StakehubCallerSession) RedelegateFeeRate() (*big.Int, error) {
	return _Stakehub.Contract.RedelegateFeeRate(&_Stakehub.CallOpts)
}

// TransferGasLimit is a free data retrieval call binding the contract method 0xe8f67c3b.
//
// Solidity: function transferGasLimit() view returns(uint256)
func (_Stakehub *StakehubCaller) TransferGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "transferGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferGasLimit is a free data retrieval call binding the contract method 0xe8f67c3b.
//
// Solidity: function transferGasLimit() view returns(uint256)
func (_Stakehub *StakehubSession) TransferGasLimit() (*big.Int, error) {
	return _Stakehub.Contract.TransferGasLimit(&_Stakehub.CallOpts)
}

// TransferGasLimit is a free data retrieval call binding the contract method 0xe8f67c3b.
//
// Solidity: function transferGasLimit() view returns(uint256)
func (_Stakehub *StakehubCallerSession) TransferGasLimit() (*big.Int, error) {
	return _Stakehub.Contract.TransferGasLimit(&_Stakehub.CallOpts)
}

// UnbondPeriod is a free data retrieval call binding the contract method 0xfc0c5ff1.
//
// Solidity: function unbondPeriod() view returns(uint256)
func (_Stakehub *StakehubCaller) UnbondPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakehub.contract.Call(opts, &out, "unbondPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondPeriod is a free data retrieval call binding the contract method 0xfc0c5ff1.
//
// Solidity: function unbondPeriod() view returns(uint256)
func (_Stakehub *StakehubSession) UnbondPeriod() (*big.Int, error) {
	return _Stakehub.Contract.UnbondPeriod(&_Stakehub.CallOpts)
}

// UnbondPeriod is a free data retrieval call binding the contract method 0xfc0c5ff1.
//
// Solidity: function unbondPeriod() view returns(uint256)
func (_Stakehub *StakehubCallerSession) UnbondPeriod() (*big.Int, error) {
	return _Stakehub.Contract.UnbondPeriod(&_Stakehub.CallOpts)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Stakehub *StakehubTransactor) AddToBlackList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "addToBlackList", account)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Stakehub *StakehubSession) AddToBlackList(account common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.AddToBlackList(&_Stakehub.TransactOpts, account)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Stakehub *StakehubTransactorSession) AddToBlackList(account common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.AddToBlackList(&_Stakehub.TransactOpts, account)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address operatorAddress, uint256 requestNumber) returns()
func (_Stakehub *StakehubTransactor) Claim(opts *bind.TransactOpts, operatorAddress common.Address, requestNumber *big.Int) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "claim", operatorAddress, requestNumber)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address operatorAddress, uint256 requestNumber) returns()
func (_Stakehub *StakehubSession) Claim(operatorAddress common.Address, requestNumber *big.Int) (*types.Transaction, error) {
	return _Stakehub.Contract.Claim(&_Stakehub.TransactOpts, operatorAddress, requestNumber)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address operatorAddress, uint256 requestNumber) returns()
func (_Stakehub *StakehubTransactorSession) Claim(operatorAddress common.Address, requestNumber *big.Int) (*types.Transaction, error) {
	return _Stakehub.Contract.Claim(&_Stakehub.TransactOpts, operatorAddress, requestNumber)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x64028fbd.
//
// Solidity: function createValidator(address consensusAddress, bytes voteAddress, bytes blsProof, (uint64,uint64,uint64) commission, (string,string,string,string) description) payable returns()
func (_Stakehub *StakehubTransactor) CreateValidator(opts *bind.TransactOpts, consensusAddress common.Address, voteAddress []byte, blsProof []byte, commission StakeHubCommission, description StakeHubDescription) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "createValidator", consensusAddress, voteAddress, blsProof, commission, description)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x64028fbd.
//
// Solidity: function createValidator(address consensusAddress, bytes voteAddress, bytes blsProof, (uint64,uint64,uint64) commission, (string,string,string,string) description) payable returns()
func (_Stakehub *StakehubSession) CreateValidator(consensusAddress common.Address, voteAddress []byte, blsProof []byte, commission StakeHubCommission, description StakeHubDescription) (*types.Transaction, error) {
	return _Stakehub.Contract.CreateValidator(&_Stakehub.TransactOpts, consensusAddress, voteAddress, blsProof, commission, description)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x64028fbd.
//
// Solidity: function createValidator(address consensusAddress, bytes voteAddress, bytes blsProof, (uint64,uint64,uint64) commission, (string,string,string,string) description) payable returns()
func (_Stakehub *StakehubTransactorSession) CreateValidator(consensusAddress common.Address, voteAddress []byte, blsProof []byte, commission StakeHubCommission, description StakeHubDescription) (*types.Transaction, error) {
	return _Stakehub.Contract.CreateValidator(&_Stakehub.TransactOpts, consensusAddress, voteAddress, blsProof, commission, description)
}

// Delegate is a paid mutator transaction binding the contract method 0x982ef0a7.
//
// Solidity: function delegate(address operatorAddress, bool delegateVotePower) payable returns()
func (_Stakehub *StakehubTransactor) Delegate(opts *bind.TransactOpts, operatorAddress common.Address, delegateVotePower bool) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "delegate", operatorAddress, delegateVotePower)
}

// Delegate is a paid mutator transaction binding the contract method 0x982ef0a7.
//
// Solidity: function delegate(address operatorAddress, bool delegateVotePower) payable returns()
func (_Stakehub *StakehubSession) Delegate(operatorAddress common.Address, delegateVotePower bool) (*types.Transaction, error) {
	return _Stakehub.Contract.Delegate(&_Stakehub.TransactOpts, operatorAddress, delegateVotePower)
}

// Delegate is a paid mutator transaction binding the contract method 0x982ef0a7.
//
// Solidity: function delegate(address operatorAddress, bool delegateVotePower) payable returns()
func (_Stakehub *StakehubTransactorSession) Delegate(operatorAddress common.Address, delegateVotePower bool) (*types.Transaction, error) {
	return _Stakehub.Contract.Delegate(&_Stakehub.TransactOpts, operatorAddress, delegateVotePower)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x092193ab.
//
// Solidity: function distributeReward(address consensusAddress) payable returns()
func (_Stakehub *StakehubTransactor) DistributeReward(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "distributeReward", consensusAddress)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x092193ab.
//
// Solidity: function distributeReward(address consensusAddress) payable returns()
func (_Stakehub *StakehubSession) DistributeReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.DistributeReward(&_Stakehub.TransactOpts, consensusAddress)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x092193ab.
//
// Solidity: function distributeReward(address consensusAddress) payable returns()
func (_Stakehub *StakehubTransactorSession) DistributeReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.DistributeReward(&_Stakehub.TransactOpts, consensusAddress)
}

// DoubleSignSlash is a paid mutator transaction binding the contract method 0xc38fbec8.
//
// Solidity: function doubleSignSlash(address consensusAddress) returns()
func (_Stakehub *StakehubTransactor) DoubleSignSlash(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "doubleSignSlash", consensusAddress)
}

// DoubleSignSlash is a paid mutator transaction binding the contract method 0xc38fbec8.
//
// Solidity: function doubleSignSlash(address consensusAddress) returns()
func (_Stakehub *StakehubSession) DoubleSignSlash(consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.DoubleSignSlash(&_Stakehub.TransactOpts, consensusAddress)
}

// DoubleSignSlash is a paid mutator transaction binding the contract method 0xc38fbec8.
//
// Solidity: function doubleSignSlash(address consensusAddress) returns()
func (_Stakehub *StakehubTransactorSession) DoubleSignSlash(consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.DoubleSignSlash(&_Stakehub.TransactOpts, consensusAddress)
}

// DowntimeSlash is a paid mutator transaction binding the contract method 0x75cc7d89.
//
// Solidity: function downtimeSlash(address consensusAddress) returns()
func (_Stakehub *StakehubTransactor) DowntimeSlash(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "downtimeSlash", consensusAddress)
}

// DowntimeSlash is a paid mutator transaction binding the contract method 0x75cc7d89.
//
// Solidity: function downtimeSlash(address consensusAddress) returns()
func (_Stakehub *StakehubSession) DowntimeSlash(consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.DowntimeSlash(&_Stakehub.TransactOpts, consensusAddress)
}

// DowntimeSlash is a paid mutator transaction binding the contract method 0x75cc7d89.
//
// Solidity: function downtimeSlash(address consensusAddress) returns()
func (_Stakehub *StakehubTransactorSession) DowntimeSlash(consensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.DowntimeSlash(&_Stakehub.TransactOpts, consensusAddress)
}

// EditCommissionRate is a paid mutator transaction binding the contract method 0x5e7cc1c9.
//
// Solidity: function editCommissionRate(uint64 commissionRate) returns()
func (_Stakehub *StakehubTransactor) EditCommissionRate(opts *bind.TransactOpts, commissionRate uint64) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "editCommissionRate", commissionRate)
}

// EditCommissionRate is a paid mutator transaction binding the contract method 0x5e7cc1c9.
//
// Solidity: function editCommissionRate(uint64 commissionRate) returns()
func (_Stakehub *StakehubSession) EditCommissionRate(commissionRate uint64) (*types.Transaction, error) {
	return _Stakehub.Contract.EditCommissionRate(&_Stakehub.TransactOpts, commissionRate)
}

// EditCommissionRate is a paid mutator transaction binding the contract method 0x5e7cc1c9.
//
// Solidity: function editCommissionRate(uint64 commissionRate) returns()
func (_Stakehub *StakehubTransactorSession) EditCommissionRate(commissionRate uint64) (*types.Transaction, error) {
	return _Stakehub.Contract.EditCommissionRate(&_Stakehub.TransactOpts, commissionRate)
}

// EditConsensusAddress is a paid mutator transaction binding the contract method 0x45211bfd.
//
// Solidity: function editConsensusAddress(address newConsensusAddress) returns()
func (_Stakehub *StakehubTransactor) EditConsensusAddress(opts *bind.TransactOpts, newConsensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "editConsensusAddress", newConsensusAddress)
}

// EditConsensusAddress is a paid mutator transaction binding the contract method 0x45211bfd.
//
// Solidity: function editConsensusAddress(address newConsensusAddress) returns()
func (_Stakehub *StakehubSession) EditConsensusAddress(newConsensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.EditConsensusAddress(&_Stakehub.TransactOpts, newConsensusAddress)
}

// EditConsensusAddress is a paid mutator transaction binding the contract method 0x45211bfd.
//
// Solidity: function editConsensusAddress(address newConsensusAddress) returns()
func (_Stakehub *StakehubTransactorSession) EditConsensusAddress(newConsensusAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.EditConsensusAddress(&_Stakehub.TransactOpts, newConsensusAddress)
}

// EditDescription is a paid mutator transaction binding the contract method 0xd6ca429d.
//
// Solidity: function editDescription((string,string,string,string) description) returns()
func (_Stakehub *StakehubTransactor) EditDescription(opts *bind.TransactOpts, description StakeHubDescription) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "editDescription", description)
}

// EditDescription is a paid mutator transaction binding the contract method 0xd6ca429d.
//
// Solidity: function editDescription((string,string,string,string) description) returns()
func (_Stakehub *StakehubSession) EditDescription(description StakeHubDescription) (*types.Transaction, error) {
	return _Stakehub.Contract.EditDescription(&_Stakehub.TransactOpts, description)
}

// EditDescription is a paid mutator transaction binding the contract method 0xd6ca429d.
//
// Solidity: function editDescription((string,string,string,string) description) returns()
func (_Stakehub *StakehubTransactorSession) EditDescription(description StakeHubDescription) (*types.Transaction, error) {
	return _Stakehub.Contract.EditDescription(&_Stakehub.TransactOpts, description)
}

// EditVoteAddress is a paid mutator transaction binding the contract method 0xfb50b31f.
//
// Solidity: function editVoteAddress(bytes newVoteAddress, bytes blsProof) returns()
func (_Stakehub *StakehubTransactor) EditVoteAddress(opts *bind.TransactOpts, newVoteAddress []byte, blsProof []byte) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "editVoteAddress", newVoteAddress, blsProof)
}

// EditVoteAddress is a paid mutator transaction binding the contract method 0xfb50b31f.
//
// Solidity: function editVoteAddress(bytes newVoteAddress, bytes blsProof) returns()
func (_Stakehub *StakehubSession) EditVoteAddress(newVoteAddress []byte, blsProof []byte) (*types.Transaction, error) {
	return _Stakehub.Contract.EditVoteAddress(&_Stakehub.TransactOpts, newVoteAddress, blsProof)
}

// EditVoteAddress is a paid mutator transaction binding the contract method 0xfb50b31f.
//
// Solidity: function editVoteAddress(bytes newVoteAddress, bytes blsProof) returns()
func (_Stakehub *StakehubTransactorSession) EditVoteAddress(newVoteAddress []byte, blsProof []byte) (*types.Transaction, error) {
	return _Stakehub.Contract.EditVoteAddress(&_Stakehub.TransactOpts, newVoteAddress, blsProof)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Stakehub *StakehubTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Stakehub *StakehubSession) Initialize() (*types.Transaction, error) {
	return _Stakehub.Contract.Initialize(&_Stakehub.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Stakehub *StakehubTransactorSession) Initialize() (*types.Transaction, error) {
	return _Stakehub.Contract.Initialize(&_Stakehub.TransactOpts)
}

// MaliciousVoteSlash is a paid mutator transaction binding the contract method 0x0e9fbf51.
//
// Solidity: function maliciousVoteSlash(bytes _voteAddr) returns()
func (_Stakehub *StakehubTransactor) MaliciousVoteSlash(opts *bind.TransactOpts, _voteAddr []byte) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "maliciousVoteSlash", _voteAddr)
}

// MaliciousVoteSlash is a paid mutator transaction binding the contract method 0x0e9fbf51.
//
// Solidity: function maliciousVoteSlash(bytes _voteAddr) returns()
func (_Stakehub *StakehubSession) MaliciousVoteSlash(_voteAddr []byte) (*types.Transaction, error) {
	return _Stakehub.Contract.MaliciousVoteSlash(&_Stakehub.TransactOpts, _voteAddr)
}

// MaliciousVoteSlash is a paid mutator transaction binding the contract method 0x0e9fbf51.
//
// Solidity: function maliciousVoteSlash(bytes _voteAddr) returns()
func (_Stakehub *StakehubTransactorSession) MaliciousVoteSlash(_voteAddr []byte) (*types.Transaction, error) {
	return _Stakehub.Contract.MaliciousVoteSlash(&_Stakehub.TransactOpts, _voteAddr)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stakehub *StakehubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stakehub *StakehubSession) Pause() (*types.Transaction, error) {
	return _Stakehub.Contract.Pause(&_Stakehub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stakehub *StakehubTransactorSession) Pause() (*types.Transaction, error) {
	return _Stakehub.Contract.Pause(&_Stakehub.TransactOpts)
}

// Redelegate is a paid mutator transaction binding the contract method 0x59491871.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 shares, bool delegateVotePower) returns()
func (_Stakehub *StakehubTransactor) Redelegate(opts *bind.TransactOpts, srcValidator common.Address, dstValidator common.Address, shares *big.Int, delegateVotePower bool) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "redelegate", srcValidator, dstValidator, shares, delegateVotePower)
}

// Redelegate is a paid mutator transaction binding the contract method 0x59491871.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 shares, bool delegateVotePower) returns()
func (_Stakehub *StakehubSession) Redelegate(srcValidator common.Address, dstValidator common.Address, shares *big.Int, delegateVotePower bool) (*types.Transaction, error) {
	return _Stakehub.Contract.Redelegate(&_Stakehub.TransactOpts, srcValidator, dstValidator, shares, delegateVotePower)
}

// Redelegate is a paid mutator transaction binding the contract method 0x59491871.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 shares, bool delegateVotePower) returns()
func (_Stakehub *StakehubTransactorSession) Redelegate(srcValidator common.Address, dstValidator common.Address, shares *big.Int, delegateVotePower bool) (*types.Transaction, error) {
	return _Stakehub.Contract.Redelegate(&_Stakehub.TransactOpts, srcValidator, dstValidator, shares, delegateVotePower)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Stakehub *StakehubTransactor) RemoveFromBlackList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "removeFromBlackList", account)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Stakehub *StakehubSession) RemoveFromBlackList(account common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.RemoveFromBlackList(&_Stakehub.TransactOpts, account)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Stakehub *StakehubTransactorSession) RemoveFromBlackList(account common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.RemoveFromBlackList(&_Stakehub.TransactOpts, account)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Stakehub *StakehubTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Stakehub *StakehubSession) Resume() (*types.Transaction, error) {
	return _Stakehub.Contract.Resume(&_Stakehub.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Stakehub *StakehubTransactorSession) Resume() (*types.Transaction, error) {
	return _Stakehub.Contract.Resume(&_Stakehub.TransactOpts)
}

// SyncGovToken is a paid mutator transaction binding the contract method 0xbaa7199e.
//
// Solidity: function syncGovToken(address[] operatorAddresses, address account) returns()
func (_Stakehub *StakehubTransactor) SyncGovToken(opts *bind.TransactOpts, operatorAddresses []common.Address, account common.Address) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "syncGovToken", operatorAddresses, account)
}

// SyncGovToken is a paid mutator transaction binding the contract method 0xbaa7199e.
//
// Solidity: function syncGovToken(address[] operatorAddresses, address account) returns()
func (_Stakehub *StakehubSession) SyncGovToken(operatorAddresses []common.Address, account common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.SyncGovToken(&_Stakehub.TransactOpts, operatorAddresses, account)
}

// SyncGovToken is a paid mutator transaction binding the contract method 0xbaa7199e.
//
// Solidity: function syncGovToken(address[] operatorAddresses, address account) returns()
func (_Stakehub *StakehubTransactorSession) SyncGovToken(operatorAddresses []common.Address, account common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.SyncGovToken(&_Stakehub.TransactOpts, operatorAddresses, account)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address operatorAddress, uint256 shares) returns()
func (_Stakehub *StakehubTransactor) Undelegate(opts *bind.TransactOpts, operatorAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "undelegate", operatorAddress, shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address operatorAddress, uint256 shares) returns()
func (_Stakehub *StakehubSession) Undelegate(operatorAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Stakehub.Contract.Undelegate(&_Stakehub.TransactOpts, operatorAddress, shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address operatorAddress, uint256 shares) returns()
func (_Stakehub *StakehubTransactorSession) Undelegate(operatorAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Stakehub.Contract.Undelegate(&_Stakehub.TransactOpts, operatorAddress, shares)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address operatorAddress) returns()
func (_Stakehub *StakehubTransactor) Unjail(opts *bind.TransactOpts, operatorAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "unjail", operatorAddress)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address operatorAddress) returns()
func (_Stakehub *StakehubSession) Unjail(operatorAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.Unjail(&_Stakehub.TransactOpts, operatorAddress)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address operatorAddress) returns()
func (_Stakehub *StakehubTransactorSession) Unjail(operatorAddress common.Address) (*types.Transaction, error) {
	return _Stakehub.Contract.Unjail(&_Stakehub.TransactOpts, operatorAddress)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Stakehub *StakehubTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Stakehub.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Stakehub *StakehubSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Stakehub.Contract.UpdateParam(&_Stakehub.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Stakehub *StakehubTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Stakehub.Contract.UpdateParam(&_Stakehub.TransactOpts, key, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakehub *StakehubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakehub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakehub *StakehubSession) Receive() (*types.Transaction, error) {
	return _Stakehub.Contract.Receive(&_Stakehub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakehub *StakehubTransactorSession) Receive() (*types.Transaction, error) {
	return _Stakehub.Contract.Receive(&_Stakehub.TransactOpts)
}

// StakehubClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Stakehub contract.
type StakehubClaimedIterator struct {
	Event *StakehubClaimed // Event containing the contract specifics and raw log

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
func (it *StakehubClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubClaimed)
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
		it.Event = new(StakehubClaimed)
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
func (it *StakehubClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubClaimed represents a Claimed event raised by the Stakehub contract.
type StakehubClaimed struct {
	OperatorAddress common.Address
	Delegator       common.Address
	BnbAmount       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) FilterClaimed(opts *bind.FilterOpts, operatorAddress []common.Address, delegator []common.Address) (*StakehubClaimedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "Claimed", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakehubClaimedIterator{contract: _Stakehub.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *StakehubClaimed, operatorAddress []common.Address, delegator []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "Claimed", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubClaimed)
				if err := _Stakehub.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed operatorAddress, address indexed delegator, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) ParseClaimed(log types.Log) (*StakehubClaimed, error) {
	event := new(StakehubClaimed)
	if err := _Stakehub.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubCommissionRateEditedIterator is returned from FilterCommissionRateEdited and is used to iterate over the raw logs and unpacked data for CommissionRateEdited events raised by the Stakehub contract.
type StakehubCommissionRateEditedIterator struct {
	Event *StakehubCommissionRateEdited // Event containing the contract specifics and raw log

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
func (it *StakehubCommissionRateEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubCommissionRateEdited)
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
		it.Event = new(StakehubCommissionRateEdited)
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
func (it *StakehubCommissionRateEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubCommissionRateEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubCommissionRateEdited represents a CommissionRateEdited event raised by the Stakehub contract.
type StakehubCommissionRateEdited struct {
	OperatorAddress common.Address
	CommissionRate  uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateEdited is a free log retrieval operation binding the contract event 0x78cdd96edf59e09cfd4d26ef6ef6c92d166effe6a40970c54821206d541932cb.
//
// Solidity: event CommissionRateEdited(address indexed operatorAddress, uint64 commissionRate)
func (_Stakehub *StakehubFilterer) FilterCommissionRateEdited(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubCommissionRateEditedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "CommissionRateEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubCommissionRateEditedIterator{contract: _Stakehub.contract, event: "CommissionRateEdited", logs: logs, sub: sub}, nil
}

// WatchCommissionRateEdited is a free log subscription operation binding the contract event 0x78cdd96edf59e09cfd4d26ef6ef6c92d166effe6a40970c54821206d541932cb.
//
// Solidity: event CommissionRateEdited(address indexed operatorAddress, uint64 commissionRate)
func (_Stakehub *StakehubFilterer) WatchCommissionRateEdited(opts *bind.WatchOpts, sink chan<- *StakehubCommissionRateEdited, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "CommissionRateEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubCommissionRateEdited)
				if err := _Stakehub.contract.UnpackLog(event, "CommissionRateEdited", log); err != nil {
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

// ParseCommissionRateEdited is a log parse operation binding the contract event 0x78cdd96edf59e09cfd4d26ef6ef6c92d166effe6a40970c54821206d541932cb.
//
// Solidity: event CommissionRateEdited(address indexed operatorAddress, uint64 commissionRate)
func (_Stakehub *StakehubFilterer) ParseCommissionRateEdited(log types.Log) (*StakehubCommissionRateEdited, error) {
	event := new(StakehubCommissionRateEdited)
	if err := _Stakehub.contract.UnpackLog(event, "CommissionRateEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubConsensusAddressEditedIterator is returned from FilterConsensusAddressEdited and is used to iterate over the raw logs and unpacked data for ConsensusAddressEdited events raised by the Stakehub contract.
type StakehubConsensusAddressEditedIterator struct {
	Event *StakehubConsensusAddressEdited // Event containing the contract specifics and raw log

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
func (it *StakehubConsensusAddressEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubConsensusAddressEdited)
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
		it.Event = new(StakehubConsensusAddressEdited)
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
func (it *StakehubConsensusAddressEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubConsensusAddressEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubConsensusAddressEdited represents a ConsensusAddressEdited event raised by the Stakehub contract.
type StakehubConsensusAddressEdited struct {
	OperatorAddress     common.Address
	NewConsensusAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterConsensusAddressEdited is a free log retrieval operation binding the contract event 0x6e4e747ca35203f16401c69805c7dd52fff67ef60b0ebc5c7fe16890530f2235.
//
// Solidity: event ConsensusAddressEdited(address indexed operatorAddress, address indexed newConsensusAddress)
func (_Stakehub *StakehubFilterer) FilterConsensusAddressEdited(opts *bind.FilterOpts, operatorAddress []common.Address, newConsensusAddress []common.Address) (*StakehubConsensusAddressEditedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var newConsensusAddressRule []interface{}
	for _, newConsensusAddressItem := range newConsensusAddress {
		newConsensusAddressRule = append(newConsensusAddressRule, newConsensusAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "ConsensusAddressEdited", operatorAddressRule, newConsensusAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubConsensusAddressEditedIterator{contract: _Stakehub.contract, event: "ConsensusAddressEdited", logs: logs, sub: sub}, nil
}

// WatchConsensusAddressEdited is a free log subscription operation binding the contract event 0x6e4e747ca35203f16401c69805c7dd52fff67ef60b0ebc5c7fe16890530f2235.
//
// Solidity: event ConsensusAddressEdited(address indexed operatorAddress, address indexed newConsensusAddress)
func (_Stakehub *StakehubFilterer) WatchConsensusAddressEdited(opts *bind.WatchOpts, sink chan<- *StakehubConsensusAddressEdited, operatorAddress []common.Address, newConsensusAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var newConsensusAddressRule []interface{}
	for _, newConsensusAddressItem := range newConsensusAddress {
		newConsensusAddressRule = append(newConsensusAddressRule, newConsensusAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "ConsensusAddressEdited", operatorAddressRule, newConsensusAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubConsensusAddressEdited)
				if err := _Stakehub.contract.UnpackLog(event, "ConsensusAddressEdited", log); err != nil {
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

// ParseConsensusAddressEdited is a log parse operation binding the contract event 0x6e4e747ca35203f16401c69805c7dd52fff67ef60b0ebc5c7fe16890530f2235.
//
// Solidity: event ConsensusAddressEdited(address indexed operatorAddress, address indexed newConsensusAddress)
func (_Stakehub *StakehubFilterer) ParseConsensusAddressEdited(log types.Log) (*StakehubConsensusAddressEdited, error) {
	event := new(StakehubConsensusAddressEdited)
	if err := _Stakehub.contract.UnpackLog(event, "ConsensusAddressEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the Stakehub contract.
type StakehubDelegatedIterator struct {
	Event *StakehubDelegated // Event containing the contract specifics and raw log

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
func (it *StakehubDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubDelegated)
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
		it.Event = new(StakehubDelegated)
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
func (it *StakehubDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubDelegated represents a Delegated event raised by the Stakehub contract.
type StakehubDelegated struct {
	OperatorAddress common.Address
	Delegator       common.Address
	Shares          *big.Int
	BnbAmount       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) FilterDelegated(opts *bind.FilterOpts, operatorAddress []common.Address, delegator []common.Address) (*StakehubDelegatedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "Delegated", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakehubDelegatedIterator{contract: _Stakehub.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *StakehubDelegated, operatorAddress []common.Address, delegator []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "Delegated", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubDelegated)
				if err := _Stakehub.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) ParseDelegated(log types.Log) (*StakehubDelegated, error) {
	event := new(StakehubDelegated)
	if err := _Stakehub.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubDescriptionEditedIterator is returned from FilterDescriptionEdited and is used to iterate over the raw logs and unpacked data for DescriptionEdited events raised by the Stakehub contract.
type StakehubDescriptionEditedIterator struct {
	Event *StakehubDescriptionEdited // Event containing the contract specifics and raw log

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
func (it *StakehubDescriptionEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubDescriptionEdited)
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
		it.Event = new(StakehubDescriptionEdited)
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
func (it *StakehubDescriptionEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubDescriptionEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubDescriptionEdited represents a DescriptionEdited event raised by the Stakehub contract.
type StakehubDescriptionEdited struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDescriptionEdited is a free log retrieval operation binding the contract event 0x85d6366b336ade7f106987ec7a8eac1e8799e508aeab045a39d2f63e0dc969d9.
//
// Solidity: event DescriptionEdited(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) FilterDescriptionEdited(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubDescriptionEditedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "DescriptionEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubDescriptionEditedIterator{contract: _Stakehub.contract, event: "DescriptionEdited", logs: logs, sub: sub}, nil
}

// WatchDescriptionEdited is a free log subscription operation binding the contract event 0x85d6366b336ade7f106987ec7a8eac1e8799e508aeab045a39d2f63e0dc969d9.
//
// Solidity: event DescriptionEdited(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) WatchDescriptionEdited(opts *bind.WatchOpts, sink chan<- *StakehubDescriptionEdited, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "DescriptionEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubDescriptionEdited)
				if err := _Stakehub.contract.UnpackLog(event, "DescriptionEdited", log); err != nil {
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

// ParseDescriptionEdited is a log parse operation binding the contract event 0x85d6366b336ade7f106987ec7a8eac1e8799e508aeab045a39d2f63e0dc969d9.
//
// Solidity: event DescriptionEdited(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) ParseDescriptionEdited(log types.Log) (*StakehubDescriptionEdited, error) {
	event := new(StakehubDescriptionEdited)
	if err := _Stakehub.contract.UnpackLog(event, "DescriptionEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Stakehub contract.
type StakehubInitializedIterator struct {
	Event *StakehubInitialized // Event containing the contract specifics and raw log

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
func (it *StakehubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubInitialized)
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
		it.Event = new(StakehubInitialized)
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
func (it *StakehubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubInitialized represents a Initialized event raised by the Stakehub contract.
type StakehubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stakehub *StakehubFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakehubInitializedIterator, error) {

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakehubInitializedIterator{contract: _Stakehub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stakehub *StakehubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakehubInitialized) (event.Subscription, error) {

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubInitialized)
				if err := _Stakehub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Stakehub *StakehubFilterer) ParseInitialized(log types.Log) (*StakehubInitialized, error) {
	event := new(StakehubInitialized)
	if err := _Stakehub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Stakehub contract.
type StakehubParamChangeIterator struct {
	Event *StakehubParamChange // Event containing the contract specifics and raw log

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
func (it *StakehubParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubParamChange)
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
		it.Event = new(StakehubParamChange)
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
func (it *StakehubParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubParamChange represents a ParamChange event raised by the Stakehub contract.
type StakehubParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Stakehub *StakehubFilterer) FilterParamChange(opts *bind.FilterOpts) (*StakehubParamChangeIterator, error) {

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return &StakehubParamChangeIterator{contract: _Stakehub.contract, event: "ParamChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Stakehub *StakehubFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *StakehubParamChange) (event.Subscription, error) {

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubParamChange)
				if err := _Stakehub.contract.UnpackLog(event, "ParamChange", log); err != nil {
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
func (_Stakehub *StakehubFilterer) ParseParamChange(log types.Log) (*StakehubParamChange, error) {
	event := new(StakehubParamChange)
	if err := _Stakehub.contract.UnpackLog(event, "ParamChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Stakehub contract.
type StakehubPausedIterator struct {
	Event *StakehubPaused // Event containing the contract specifics and raw log

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
func (it *StakehubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubPaused)
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
		it.Event = new(StakehubPaused)
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
func (it *StakehubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubPaused represents a Paused event raised by the Stakehub contract.
type StakehubPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Stakehub *StakehubFilterer) FilterPaused(opts *bind.FilterOpts) (*StakehubPausedIterator, error) {

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakehubPausedIterator{contract: _Stakehub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Stakehub *StakehubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakehubPaused) (event.Subscription, error) {

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubPaused)
				if err := _Stakehub.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Stakehub *StakehubFilterer) ParsePaused(log types.Log) (*StakehubPaused, error) {
	event := new(StakehubPaused)
	if err := _Stakehub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the Stakehub contract.
type StakehubRedelegatedIterator struct {
	Event *StakehubRedelegated // Event containing the contract specifics and raw log

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
func (it *StakehubRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubRedelegated)
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
		it.Event = new(StakehubRedelegated)
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
func (it *StakehubRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubRedelegated represents a Redelegated event raised by the Stakehub contract.
type StakehubRedelegated struct {
	SrcValidator common.Address
	DstValidator common.Address
	Delegator    common.Address
	OldShares    *big.Int
	NewShares    *big.Int
	BnbAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed srcValidator, address indexed dstValidator, address indexed delegator, uint256 oldShares, uint256 newShares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) FilterRedelegated(opts *bind.FilterOpts, srcValidator []common.Address, dstValidator []common.Address, delegator []common.Address) (*StakehubRedelegatedIterator, error) {

	var srcValidatorRule []interface{}
	for _, srcValidatorItem := range srcValidator {
		srcValidatorRule = append(srcValidatorRule, srcValidatorItem)
	}
	var dstValidatorRule []interface{}
	for _, dstValidatorItem := range dstValidator {
		dstValidatorRule = append(dstValidatorRule, dstValidatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "Redelegated", srcValidatorRule, dstValidatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakehubRedelegatedIterator{contract: _Stakehub.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed srcValidator, address indexed dstValidator, address indexed delegator, uint256 oldShares, uint256 newShares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *StakehubRedelegated, srcValidator []common.Address, dstValidator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var srcValidatorRule []interface{}
	for _, srcValidatorItem := range srcValidator {
		srcValidatorRule = append(srcValidatorRule, srcValidatorItem)
	}
	var dstValidatorRule []interface{}
	for _, dstValidatorItem := range dstValidator {
		dstValidatorRule = append(dstValidatorRule, dstValidatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "Redelegated", srcValidatorRule, dstValidatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubRedelegated)
				if err := _Stakehub.contract.UnpackLog(event, "Redelegated", log); err != nil {
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

// ParseRedelegated is a log parse operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed srcValidator, address indexed dstValidator, address indexed delegator, uint256 oldShares, uint256 newShares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) ParseRedelegated(log types.Log) (*StakehubRedelegated, error) {
	event := new(StakehubRedelegated)
	if err := _Stakehub.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Stakehub contract.
type StakehubResumedIterator struct {
	Event *StakehubResumed // Event containing the contract specifics and raw log

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
func (it *StakehubResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubResumed)
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
		it.Event = new(StakehubResumed)
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
func (it *StakehubResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubResumed represents a Resumed event raised by the Stakehub contract.
type StakehubResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Stakehub *StakehubFilterer) FilterResumed(opts *bind.FilterOpts) (*StakehubResumedIterator, error) {

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &StakehubResumedIterator{contract: _Stakehub.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Stakehub *StakehubFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *StakehubResumed) (event.Subscription, error) {

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubResumed)
				if err := _Stakehub.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_Stakehub *StakehubFilterer) ParseResumed(log types.Log) (*StakehubResumed, error) {
	event := new(StakehubResumed)
	if err := _Stakehub.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubRewardDistributeFailedIterator is returned from FilterRewardDistributeFailed and is used to iterate over the raw logs and unpacked data for RewardDistributeFailed events raised by the Stakehub contract.
type StakehubRewardDistributeFailedIterator struct {
	Event *StakehubRewardDistributeFailed // Event containing the contract specifics and raw log

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
func (it *StakehubRewardDistributeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubRewardDistributeFailed)
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
		it.Event = new(StakehubRewardDistributeFailed)
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
func (it *StakehubRewardDistributeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubRewardDistributeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubRewardDistributeFailed represents a RewardDistributeFailed event raised by the Stakehub contract.
type StakehubRewardDistributeFailed struct {
	OperatorAddress common.Address
	FailReason      []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributeFailed is a free log retrieval operation binding the contract event 0xfc8bff675087dd2da069cc3fb517b9ed001e19750c0865241a5542dba1ba170d.
//
// Solidity: event RewardDistributeFailed(address indexed operatorAddress, bytes failReason)
func (_Stakehub *StakehubFilterer) FilterRewardDistributeFailed(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubRewardDistributeFailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "RewardDistributeFailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubRewardDistributeFailedIterator{contract: _Stakehub.contract, event: "RewardDistributeFailed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributeFailed is a free log subscription operation binding the contract event 0xfc8bff675087dd2da069cc3fb517b9ed001e19750c0865241a5542dba1ba170d.
//
// Solidity: event RewardDistributeFailed(address indexed operatorAddress, bytes failReason)
func (_Stakehub *StakehubFilterer) WatchRewardDistributeFailed(opts *bind.WatchOpts, sink chan<- *StakehubRewardDistributeFailed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "RewardDistributeFailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubRewardDistributeFailed)
				if err := _Stakehub.contract.UnpackLog(event, "RewardDistributeFailed", log); err != nil {
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

// ParseRewardDistributeFailed is a log parse operation binding the contract event 0xfc8bff675087dd2da069cc3fb517b9ed001e19750c0865241a5542dba1ba170d.
//
// Solidity: event RewardDistributeFailed(address indexed operatorAddress, bytes failReason)
func (_Stakehub *StakehubFilterer) ParseRewardDistributeFailed(log types.Log) (*StakehubRewardDistributeFailed, error) {
	event := new(StakehubRewardDistributeFailed)
	if err := _Stakehub.contract.UnpackLog(event, "RewardDistributeFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubRewardDistributedIterator is returned from FilterRewardDistributed and is used to iterate over the raw logs and unpacked data for RewardDistributed events raised by the Stakehub contract.
type StakehubRewardDistributedIterator struct {
	Event *StakehubRewardDistributed // Event containing the contract specifics and raw log

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
func (it *StakehubRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubRewardDistributed)
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
		it.Event = new(StakehubRewardDistributed)
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
func (it *StakehubRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubRewardDistributed represents a RewardDistributed event raised by the Stakehub contract.
type StakehubRewardDistributed struct {
	OperatorAddress common.Address
	Reward          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributed is a free log retrieval operation binding the contract event 0xe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6.
//
// Solidity: event RewardDistributed(address indexed operatorAddress, uint256 reward)
func (_Stakehub *StakehubFilterer) FilterRewardDistributed(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubRewardDistributedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "RewardDistributed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubRewardDistributedIterator{contract: _Stakehub.contract, event: "RewardDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributed is a free log subscription operation binding the contract event 0xe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6.
//
// Solidity: event RewardDistributed(address indexed operatorAddress, uint256 reward)
func (_Stakehub *StakehubFilterer) WatchRewardDistributed(opts *bind.WatchOpts, sink chan<- *StakehubRewardDistributed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "RewardDistributed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubRewardDistributed)
				if err := _Stakehub.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
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

// ParseRewardDistributed is a log parse operation binding the contract event 0xe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6.
//
// Solidity: event RewardDistributed(address indexed operatorAddress, uint256 reward)
func (_Stakehub *StakehubFilterer) ParseRewardDistributed(log types.Log) (*StakehubRewardDistributed, error) {
	event := new(StakehubRewardDistributed)
	if err := _Stakehub.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Stakehub contract.
type StakehubUndelegatedIterator struct {
	Event *StakehubUndelegated // Event containing the contract specifics and raw log

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
func (it *StakehubUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubUndelegated)
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
		it.Event = new(StakehubUndelegated)
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
func (it *StakehubUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubUndelegated represents a Undelegated event raised by the Stakehub contract.
type StakehubUndelegated struct {
	OperatorAddress common.Address
	Delegator       common.Address
	Shares          *big.Int
	BnbAmount       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802.
//
// Solidity: event Undelegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) FilterUndelegated(opts *bind.FilterOpts, operatorAddress []common.Address, delegator []common.Address) (*StakehubUndelegatedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "Undelegated", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakehubUndelegatedIterator{contract: _Stakehub.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802.
//
// Solidity: event Undelegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *StakehubUndelegated, operatorAddress []common.Address, delegator []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "Undelegated", operatorAddressRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubUndelegated)
				if err := _Stakehub.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802.
//
// Solidity: event Undelegated(address indexed operatorAddress, address indexed delegator, uint256 shares, uint256 bnbAmount)
func (_Stakehub *StakehubFilterer) ParseUndelegated(log types.Log) (*StakehubUndelegated, error) {
	event := new(StakehubUndelegated)
	if err := _Stakehub.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubValidatorCreatedIterator is returned from FilterValidatorCreated and is used to iterate over the raw logs and unpacked data for ValidatorCreated events raised by the Stakehub contract.
type StakehubValidatorCreatedIterator struct {
	Event *StakehubValidatorCreated // Event containing the contract specifics and raw log

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
func (it *StakehubValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubValidatorCreated)
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
		it.Event = new(StakehubValidatorCreated)
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
func (it *StakehubValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubValidatorCreated represents a ValidatorCreated event raised by the Stakehub contract.
type StakehubValidatorCreated struct {
	ConsensusAddress common.Address
	OperatorAddress  common.Address
	CreditContract   common.Address
	VoteAddress      []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorCreated is a free log retrieval operation binding the contract event 0xaecd9fb95e79c75a3a1de93362c6be5fe6ab65770d8614be583884161cd8228d.
//
// Solidity: event ValidatorCreated(address indexed consensusAddress, address indexed operatorAddress, address indexed creditContract, bytes voteAddress)
func (_Stakehub *StakehubFilterer) FilterValidatorCreated(opts *bind.FilterOpts, consensusAddress []common.Address, operatorAddress []common.Address, creditContract []common.Address) (*StakehubValidatorCreatedIterator, error) {

	var consensusAddressRule []interface{}
	for _, consensusAddressItem := range consensusAddress {
		consensusAddressRule = append(consensusAddressRule, consensusAddressItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var creditContractRule []interface{}
	for _, creditContractItem := range creditContract {
		creditContractRule = append(creditContractRule, creditContractItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "ValidatorCreated", consensusAddressRule, operatorAddressRule, creditContractRule)
	if err != nil {
		return nil, err
	}
	return &StakehubValidatorCreatedIterator{contract: _Stakehub.contract, event: "ValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchValidatorCreated is a free log subscription operation binding the contract event 0xaecd9fb95e79c75a3a1de93362c6be5fe6ab65770d8614be583884161cd8228d.
//
// Solidity: event ValidatorCreated(address indexed consensusAddress, address indexed operatorAddress, address indexed creditContract, bytes voteAddress)
func (_Stakehub *StakehubFilterer) WatchValidatorCreated(opts *bind.WatchOpts, sink chan<- *StakehubValidatorCreated, consensusAddress []common.Address, operatorAddress []common.Address, creditContract []common.Address) (event.Subscription, error) {

	var consensusAddressRule []interface{}
	for _, consensusAddressItem := range consensusAddress {
		consensusAddressRule = append(consensusAddressRule, consensusAddressItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var creditContractRule []interface{}
	for _, creditContractItem := range creditContract {
		creditContractRule = append(creditContractRule, creditContractItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "ValidatorCreated", consensusAddressRule, operatorAddressRule, creditContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubValidatorCreated)
				if err := _Stakehub.contract.UnpackLog(event, "ValidatorCreated", log); err != nil {
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

// ParseValidatorCreated is a log parse operation binding the contract event 0xaecd9fb95e79c75a3a1de93362c6be5fe6ab65770d8614be583884161cd8228d.
//
// Solidity: event ValidatorCreated(address indexed consensusAddress, address indexed operatorAddress, address indexed creditContract, bytes voteAddress)
func (_Stakehub *StakehubFilterer) ParseValidatorCreated(log types.Log) (*StakehubValidatorCreated, error) {
	event := new(StakehubValidatorCreated)
	if err := _Stakehub.contract.UnpackLog(event, "ValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubValidatorEmptyJailedIterator is returned from FilterValidatorEmptyJailed and is used to iterate over the raw logs and unpacked data for ValidatorEmptyJailed events raised by the Stakehub contract.
type StakehubValidatorEmptyJailedIterator struct {
	Event *StakehubValidatorEmptyJailed // Event containing the contract specifics and raw log

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
func (it *StakehubValidatorEmptyJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubValidatorEmptyJailed)
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
		it.Event = new(StakehubValidatorEmptyJailed)
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
func (it *StakehubValidatorEmptyJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubValidatorEmptyJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubValidatorEmptyJailed represents a ValidatorEmptyJailed event raised by the Stakehub contract.
type StakehubValidatorEmptyJailed struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorEmptyJailed is a free log retrieval operation binding the contract event 0x2afdc18061ac21cff7d9f11527ab9c8dec6fabd4edf6f894ed634bebd6a20d45.
//
// Solidity: event ValidatorEmptyJailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) FilterValidatorEmptyJailed(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubValidatorEmptyJailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "ValidatorEmptyJailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubValidatorEmptyJailedIterator{contract: _Stakehub.contract, event: "ValidatorEmptyJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorEmptyJailed is a free log subscription operation binding the contract event 0x2afdc18061ac21cff7d9f11527ab9c8dec6fabd4edf6f894ed634bebd6a20d45.
//
// Solidity: event ValidatorEmptyJailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) WatchValidatorEmptyJailed(opts *bind.WatchOpts, sink chan<- *StakehubValidatorEmptyJailed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "ValidatorEmptyJailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubValidatorEmptyJailed)
				if err := _Stakehub.contract.UnpackLog(event, "ValidatorEmptyJailed", log); err != nil {
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

// ParseValidatorEmptyJailed is a log parse operation binding the contract event 0x2afdc18061ac21cff7d9f11527ab9c8dec6fabd4edf6f894ed634bebd6a20d45.
//
// Solidity: event ValidatorEmptyJailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) ParseValidatorEmptyJailed(log types.Log) (*StakehubValidatorEmptyJailed, error) {
	event := new(StakehubValidatorEmptyJailed)
	if err := _Stakehub.contract.UnpackLog(event, "ValidatorEmptyJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the Stakehub contract.
type StakehubValidatorJailedIterator struct {
	Event *StakehubValidatorJailed // Event containing the contract specifics and raw log

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
func (it *StakehubValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubValidatorJailed)
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
		it.Event = new(StakehubValidatorJailed)
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
func (it *StakehubValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubValidatorJailed represents a ValidatorJailed event raised by the Stakehub contract.
type StakehubValidatorJailed struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0x4905ac32602da3fb8b4b7b00c285e5fc4c6c2308cc908b4a1e4e9625a29c90a3.
//
// Solidity: event ValidatorJailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) FilterValidatorJailed(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubValidatorJailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "ValidatorJailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubValidatorJailedIterator{contract: _Stakehub.contract, event: "ValidatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0x4905ac32602da3fb8b4b7b00c285e5fc4c6c2308cc908b4a1e4e9625a29c90a3.
//
// Solidity: event ValidatorJailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *StakehubValidatorJailed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "ValidatorJailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubValidatorJailed)
				if err := _Stakehub.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
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

// ParseValidatorJailed is a log parse operation binding the contract event 0x4905ac32602da3fb8b4b7b00c285e5fc4c6c2308cc908b4a1e4e9625a29c90a3.
//
// Solidity: event ValidatorJailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) ParseValidatorJailed(log types.Log) (*StakehubValidatorJailed, error) {
	event := new(StakehubValidatorJailed)
	if err := _Stakehub.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the Stakehub contract.
type StakehubValidatorSlashedIterator struct {
	Event *StakehubValidatorSlashed // Event containing the contract specifics and raw log

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
func (it *StakehubValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubValidatorSlashed)
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
		it.Event = new(StakehubValidatorSlashed)
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
func (it *StakehubValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubValidatorSlashed represents a ValidatorSlashed event raised by the Stakehub contract.
type StakehubValidatorSlashed struct {
	OperatorAddress common.Address
	JailUntil       *big.Int
	SlashAmount     *big.Int
	SlashType       uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0x6e9a2ee7aee95665e3a774a212eb11441b217e3e4656ab9563793094689aabb2.
//
// Solidity: event ValidatorSlashed(address indexed operatorAddress, uint256 jailUntil, uint256 slashAmount, uint8 slashType)
func (_Stakehub *StakehubFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubValidatorSlashedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "ValidatorSlashed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubValidatorSlashedIterator{contract: _Stakehub.contract, event: "ValidatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0x6e9a2ee7aee95665e3a774a212eb11441b217e3e4656ab9563793094689aabb2.
//
// Solidity: event ValidatorSlashed(address indexed operatorAddress, uint256 jailUntil, uint256 slashAmount, uint8 slashType)
func (_Stakehub *StakehubFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *StakehubValidatorSlashed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "ValidatorSlashed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubValidatorSlashed)
				if err := _Stakehub.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
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

// ParseValidatorSlashed is a log parse operation binding the contract event 0x6e9a2ee7aee95665e3a774a212eb11441b217e3e4656ab9563793094689aabb2.
//
// Solidity: event ValidatorSlashed(address indexed operatorAddress, uint256 jailUntil, uint256 slashAmount, uint8 slashType)
func (_Stakehub *StakehubFilterer) ParseValidatorSlashed(log types.Log) (*StakehubValidatorSlashed, error) {
	event := new(StakehubValidatorSlashed)
	if err := _Stakehub.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubValidatorUnjailedIterator is returned from FilterValidatorUnjailed and is used to iterate over the raw logs and unpacked data for ValidatorUnjailed events raised by the Stakehub contract.
type StakehubValidatorUnjailedIterator struct {
	Event *StakehubValidatorUnjailed // Event containing the contract specifics and raw log

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
func (it *StakehubValidatorUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubValidatorUnjailed)
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
		it.Event = new(StakehubValidatorUnjailed)
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
func (it *StakehubValidatorUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubValidatorUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubValidatorUnjailed represents a ValidatorUnjailed event raised by the Stakehub contract.
type StakehubValidatorUnjailed struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnjailed is a free log retrieval operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) FilterValidatorUnjailed(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubValidatorUnjailedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "ValidatorUnjailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubValidatorUnjailedIterator{contract: _Stakehub.contract, event: "ValidatorUnjailed", logs: logs, sub: sub}, nil
}

// WatchValidatorUnjailed is a free log subscription operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) WatchValidatorUnjailed(opts *bind.WatchOpts, sink chan<- *StakehubValidatorUnjailed, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "ValidatorUnjailed", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubValidatorUnjailed)
				if err := _Stakehub.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
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

// ParseValidatorUnjailed is a log parse operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address indexed operatorAddress)
func (_Stakehub *StakehubFilterer) ParseValidatorUnjailed(log types.Log) (*StakehubValidatorUnjailed, error) {
	event := new(StakehubValidatorUnjailed)
	if err := _Stakehub.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakehubVoteAddressEditedIterator is returned from FilterVoteAddressEdited and is used to iterate over the raw logs and unpacked data for VoteAddressEdited events raised by the Stakehub contract.
type StakehubVoteAddressEditedIterator struct {
	Event *StakehubVoteAddressEdited // Event containing the contract specifics and raw log

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
func (it *StakehubVoteAddressEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakehubVoteAddressEdited)
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
		it.Event = new(StakehubVoteAddressEdited)
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
func (it *StakehubVoteAddressEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakehubVoteAddressEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakehubVoteAddressEdited represents a VoteAddressEdited event raised by the Stakehub contract.
type StakehubVoteAddressEdited struct {
	OperatorAddress common.Address
	NewVoteAddress  []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVoteAddressEdited is a free log retrieval operation binding the contract event 0x783156582145bd0ff7924fae6953ba054cf1233eb60739a200ddb10de068ff0d.
//
// Solidity: event VoteAddressEdited(address indexed operatorAddress, bytes newVoteAddress)
func (_Stakehub *StakehubFilterer) FilterVoteAddressEdited(opts *bind.FilterOpts, operatorAddress []common.Address) (*StakehubVoteAddressEditedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.FilterLogs(opts, "VoteAddressEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakehubVoteAddressEditedIterator{contract: _Stakehub.contract, event: "VoteAddressEdited", logs: logs, sub: sub}, nil
}

// WatchVoteAddressEdited is a free log subscription operation binding the contract event 0x783156582145bd0ff7924fae6953ba054cf1233eb60739a200ddb10de068ff0d.
//
// Solidity: event VoteAddressEdited(address indexed operatorAddress, bytes newVoteAddress)
func (_Stakehub *StakehubFilterer) WatchVoteAddressEdited(opts *bind.WatchOpts, sink chan<- *StakehubVoteAddressEdited, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Stakehub.contract.WatchLogs(opts, "VoteAddressEdited", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakehubVoteAddressEdited)
				if err := _Stakehub.contract.UnpackLog(event, "VoteAddressEdited", log); err != nil {
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

// ParseVoteAddressEdited is a log parse operation binding the contract event 0x783156582145bd0ff7924fae6953ba054cf1233eb60739a200ddb10de068ff0d.
//
// Solidity: event VoteAddressEdited(address indexed operatorAddress, bytes newVoteAddress)
func (_Stakehub *StakehubFilterer) ParseVoteAddressEdited(log types.Log) (*StakehubVoteAddressEdited, error) {
	event := new(StakehubVoteAddressEdited)
	if err := _Stakehub.contract.UnpackLog(event, "VoteAddressEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
