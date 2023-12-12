// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorset

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

// ValidatorsetMetaData contains all meta data concerning the Validatorset contract.
var ValidatorsetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BIND_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BURN_RATIO_SCALE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CODE_OK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_STAKE_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DUSTY_INCOMING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_FAIL_CHECK_VALIDATORS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_FAIL_DECODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_LEN_OF_VAL_MISMATCH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_RELAYFEE_TOO_LARGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERROR_UNKNOWN_PACKAGE_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXPIRE_TIME_SECOND_GAP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_TOKEN_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INCENTIVIZE_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_BURN_RATIO\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_MAINTAIN_SLASH_SCALE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_MAX_NUM_OF_MAINTAINING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_NUM_OF_CABINETS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_SYSTEM_REWARD_RATIO\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_VALIDATORSET_BYTES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"JAIL_MESSAGE_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_NUM_OF_VALIDATORS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_SYSTEM_REWARD_BALANCE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_CREDIT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYSTEM_REWARD_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYSTEM_REWARD_RATIO_SCALE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TIMELOCK_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_HUB_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_MANAGER_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_RECOVER_PORTAL_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_IN_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_OUT_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATORS_UPDATE_MESSAGE_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"alreadyInit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bscChainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnRatio\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnRatioInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canEnterMaintenance\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentValidatorSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"BBCFeeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"jailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"incoming\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentValidatorSetMap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentVoteAddrFullSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"valAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distributeFinalityReward\",\"inputs\":[{\"name\":\"valAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enterMaintenance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitMaintenance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expireTimeSecondGap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"felony\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentValidatorIndex\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIncoming\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLivingValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMiningValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkingValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"workingValidatorCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleAckPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleFailAckPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleSynPackage\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"responsePayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isCurrentValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMonitoredForMaliciousVote\",\"inputs\":[{\"name\":\"voteAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSystemRewardIncluded\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWorkingValidator\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"jailValidator\",\"inputs\":[{\"name\":\"consensusAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maintainSlashScale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxNumOfCandidates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxNumOfMaintaining\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxNumOfWorkingCandidates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"misdemeanor\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numOfCabinets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfJailed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfMaintaining\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previousBalanceOfSystemReward\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previousHeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previousVoteAddrFullSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"systemRewardRatio\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalInComing\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateValidatorSetV2\",\"inputs\":[{\"name\":\"_consensusAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_votingPowers\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"_voteAddrs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorExtraSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"enterMaintenanceHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isMaintaining\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"voteAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"batchTransfer\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"batchTransferFailed\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"batchTransferLowerFailed\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"deprecatedDeposit\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"deprecatedFinalityRewardDeposit\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"directTransfer\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"directTransferFail\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"failReasonWithStr\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"feeBurned\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"finalityRewardDeposit\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"paramChange\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"systemTransfer\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"unexpectedPackage\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"msgBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorDeposit\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorEmptyJailed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorEnterMaintenance\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorExitMaintenance\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorFelony\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorJailed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorMisdemeanor\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorSetUpdated\",\"inputs\":[],\"anonymous\":false}]",
}

// ValidatorsetABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorsetMetaData.ABI instead.
var ValidatorsetABI = ValidatorsetMetaData.ABI

// Validatorset is an auto generated Go binding around an Ethereum contract.
type Validatorset struct {
	ValidatorsetCaller     // Read-only binding to the contract
	ValidatorsetTransactor // Write-only binding to the contract
	ValidatorsetFilterer   // Log filterer for contract events
}

// ValidatorsetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsetSession struct {
	Contract     *Validatorset     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsetCallerSession struct {
	Contract *ValidatorsetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorsetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsetTransactorSession struct {
	Contract     *ValidatorsetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorsetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsetRaw struct {
	Contract *Validatorset // Generic contract binding to access the raw methods on
}

// ValidatorsetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsetCallerRaw struct {
	Contract *ValidatorsetCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsetTransactorRaw struct {
	Contract *ValidatorsetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorset creates a new instance of Validatorset, bound to a specific deployed contract.
func NewValidatorset(address common.Address, backend bind.ContractBackend) (*Validatorset, error) {
	contract, err := bindValidatorset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validatorset{ValidatorsetCaller: ValidatorsetCaller{contract: contract}, ValidatorsetTransactor: ValidatorsetTransactor{contract: contract}, ValidatorsetFilterer: ValidatorsetFilterer{contract: contract}}, nil
}

// NewValidatorsetCaller creates a new read-only instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsetCaller, error) {
	contract, err := bindValidatorset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetCaller{contract: contract}, nil
}

// NewValidatorsetTransactor creates a new write-only instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsetTransactor, error) {
	contract, err := bindValidatorset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetTransactor{contract: contract}, nil
}

// NewValidatorsetFilterer creates a new log filterer instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsetFilterer, error) {
	contract, err := bindValidatorset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetFilterer{contract: contract}, nil
}

// bindValidatorset binds a generic wrapper to an already deployed contract.
func bindValidatorset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorsetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorset *ValidatorsetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorset.Contract.ValidatorsetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorset *ValidatorsetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.Contract.ValidatorsetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorset *ValidatorsetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorset.Contract.ValidatorsetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorset *ValidatorsetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorset *ValidatorsetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorset *ValidatorsetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorset.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetSession) BINDCHANNELID() (uint8, error) {
	return _Validatorset.Contract.BINDCHANNELID(&_Validatorset.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) BINDCHANNELID() (uint8, error) {
	return _Validatorset.Contract.BINDCHANNELID(&_Validatorset.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_Validatorset *ValidatorsetCaller) BURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_Validatorset *ValidatorsetSession) BURNADDRESS() (common.Address, error) {
	return _Validatorset.Contract.BURNADDRESS(&_Validatorset.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) BURNADDRESS() (common.Address, error) {
	return _Validatorset.Contract.BURNADDRESS(&_Validatorset.CallOpts)
}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) BURNRATIOSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "BURN_RATIO_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetSession) BURNRATIOSCALE() (*big.Int, error) {
	return _Validatorset.Contract.BURNRATIOSCALE(&_Validatorset.CallOpts)
}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) BURNRATIOSCALE() (*big.Int, error) {
	return _Validatorset.Contract.BURNRATIOSCALE(&_Validatorset.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Validatorset *ValidatorsetCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Validatorset *ValidatorsetSession) CODEOK() (uint32, error) {
	return _Validatorset.Contract.CODEOK(&_Validatorset.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Validatorset *ValidatorsetCallerSession) CODEOK() (uint32, error) {
	return _Validatorset.Contract.CODEOK(&_Validatorset.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.CROSSCHAINCONTRACTADDR(&_Validatorset.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.CROSSCHAINCONTRACTADDR(&_Validatorset.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) CROSSSTAKECHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "CROSS_STAKE_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Validatorset.Contract.CROSSSTAKECHANNELID(&_Validatorset.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Validatorset.Contract.CROSSSTAKECHANNELID(&_Validatorset.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) DUSTYINCOMING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "DUSTY_INCOMING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_Validatorset *ValidatorsetSession) DUSTYINCOMING() (*big.Int, error) {
	return _Validatorset.Contract.DUSTYINCOMING(&_Validatorset.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) DUSTYINCOMING() (*big.Int, error) {
	return _Validatorset.Contract.DUSTYINCOMING(&_Validatorset.CallOpts)
}

// EPOCH is a free data retrieval call binding the contract method 0xa0dc2758.
//
// Solidity: function EPOCH() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) EPOCH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "EPOCH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EPOCH is a free data retrieval call binding the contract method 0xa0dc2758.
//
// Solidity: function EPOCH() view returns(uint256)
func (_Validatorset *ValidatorsetSession) EPOCH() (*big.Int, error) {
	return _Validatorset.Contract.EPOCH(&_Validatorset.CallOpts)
}

// EPOCH is a free data retrieval call binding the contract method 0xa0dc2758.
//
// Solidity: function EPOCH() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) EPOCH() (*big.Int, error) {
	return _Validatorset.Contract.EPOCH(&_Validatorset.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_Validatorset *ValidatorsetCaller) ERRORFAILCHECKVALIDATORS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "ERROR_FAIL_CHECK_VALIDATORS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_Validatorset *ValidatorsetSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _Validatorset.Contract.ERRORFAILCHECKVALIDATORS(&_Validatorset.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_Validatorset *ValidatorsetCallerSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _Validatorset.Contract.ERRORFAILCHECKVALIDATORS(&_Validatorset.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Validatorset *ValidatorsetCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Validatorset *ValidatorsetSession) ERRORFAILDECODE() (uint32, error) {
	return _Validatorset.Contract.ERRORFAILDECODE(&_Validatorset.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Validatorset *ValidatorsetCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Validatorset.Contract.ERRORFAILDECODE(&_Validatorset.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_Validatorset *ValidatorsetCaller) ERRORLENOFVALMISMATCH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "ERROR_LEN_OF_VAL_MISMATCH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_Validatorset *ValidatorsetSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _Validatorset.Contract.ERRORLENOFVALMISMATCH(&_Validatorset.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_Validatorset *ValidatorsetCallerSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _Validatorset.Contract.ERRORLENOFVALMISMATCH(&_Validatorset.CallOpts)
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_Validatorset *ValidatorsetCaller) ERRORRELAYFEETOOLARGE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "ERROR_RELAYFEE_TOO_LARGE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_Validatorset *ValidatorsetSession) ERRORRELAYFEETOOLARGE() (uint32, error) {
	return _Validatorset.Contract.ERRORRELAYFEETOOLARGE(&_Validatorset.CallOpts)
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_Validatorset *ValidatorsetCallerSession) ERRORRELAYFEETOOLARGE() (uint32, error) {
	return _Validatorset.Contract.ERRORRELAYFEETOOLARGE(&_Validatorset.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_Validatorset *ValidatorsetCaller) ERRORUNKNOWNPACKAGETYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "ERROR_UNKNOWN_PACKAGE_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_Validatorset *ValidatorsetSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _Validatorset.Contract.ERRORUNKNOWNPACKAGETYPE(&_Validatorset.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_Validatorset *ValidatorsetCallerSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _Validatorset.Contract.ERRORUNKNOWNPACKAGETYPE(&_Validatorset.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) EXPIRETIMESECONDGAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "EXPIRE_TIME_SECOND_GAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_Validatorset *ValidatorsetSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _Validatorset.Contract.EXPIRETIMESECONDGAP(&_Validatorset.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _Validatorset.Contract.EXPIRETIMESECONDGAP(&_Validatorset.CallOpts)
}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) GOVERNORADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "GOVERNOR_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) GOVERNORADDR() (common.Address, error) {
	return _Validatorset.Contract.GOVERNORADDR(&_Validatorset.CallOpts)
}

// GOVERNORADDR is a free data retrieval call binding the contract method 0xdf8079e9.
//
// Solidity: function GOVERNOR_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) GOVERNORADDR() (common.Address, error) {
	return _Validatorset.Contract.GOVERNORADDR(&_Validatorset.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetSession) GOVCHANNELID() (uint8, error) {
	return _Validatorset.Contract.GOVCHANNELID(&_Validatorset.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) GOVCHANNELID() (uint8, error) {
	return _Validatorset.Contract.GOVCHANNELID(&_Validatorset.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) GOVHUBADDR() (common.Address, error) {
	return _Validatorset.Contract.GOVHUBADDR(&_Validatorset.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Validatorset.Contract.GOVHUBADDR(&_Validatorset.CallOpts)
}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) GOVTOKENADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "GOV_TOKEN_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) GOVTOKENADDR() (common.Address, error) {
	return _Validatorset.Contract.GOVTOKENADDR(&_Validatorset.CallOpts)
}

// GOVTOKENADDR is a free data retrieval call binding the contract method 0x28087028.
//
// Solidity: function GOV_TOKEN_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) GOVTOKENADDR() (common.Address, error) {
	return _Validatorset.Contract.GOVTOKENADDR(&_Validatorset.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Validatorset.Contract.INCENTIVIZEADDR(&_Validatorset.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Validatorset.Contract.INCENTIVIZEADDR(&_Validatorset.CallOpts)
}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) INITBURNRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "INIT_BURN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_Validatorset *ValidatorsetSession) INITBURNRATIO() (*big.Int, error) {
	return _Validatorset.Contract.INITBURNRATIO(&_Validatorset.CallOpts)
}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) INITBURNRATIO() (*big.Int, error) {
	return _Validatorset.Contract.INITBURNRATIO(&_Validatorset.CallOpts)
}

// INITMAINTAINSLASHSCALE is a free data retrieval call binding the contract method 0xc6d33945.
//
// Solidity: function INIT_MAINTAIN_SLASH_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) INITMAINTAINSLASHSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "INIT_MAINTAIN_SLASH_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITMAINTAINSLASHSCALE is a free data retrieval call binding the contract method 0xc6d33945.
//
// Solidity: function INIT_MAINTAIN_SLASH_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetSession) INITMAINTAINSLASHSCALE() (*big.Int, error) {
	return _Validatorset.Contract.INITMAINTAINSLASHSCALE(&_Validatorset.CallOpts)
}

// INITMAINTAINSLASHSCALE is a free data retrieval call binding the contract method 0xc6d33945.
//
// Solidity: function INIT_MAINTAIN_SLASH_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) INITMAINTAINSLASHSCALE() (*big.Int, error) {
	return _Validatorset.Contract.INITMAINTAINSLASHSCALE(&_Validatorset.CallOpts)
}

// INITMAXNUMOFMAINTAINING is a free data retrieval call binding the contract method 0x9fe0f816.
//
// Solidity: function INIT_MAX_NUM_OF_MAINTAINING() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) INITMAXNUMOFMAINTAINING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "INIT_MAX_NUM_OF_MAINTAINING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITMAXNUMOFMAINTAINING is a free data retrieval call binding the contract method 0x9fe0f816.
//
// Solidity: function INIT_MAX_NUM_OF_MAINTAINING() view returns(uint256)
func (_Validatorset *ValidatorsetSession) INITMAXNUMOFMAINTAINING() (*big.Int, error) {
	return _Validatorset.Contract.INITMAXNUMOFMAINTAINING(&_Validatorset.CallOpts)
}

// INITMAXNUMOFMAINTAINING is a free data retrieval call binding the contract method 0x9fe0f816.
//
// Solidity: function INIT_MAX_NUM_OF_MAINTAINING() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) INITMAXNUMOFMAINTAINING() (*big.Int, error) {
	return _Validatorset.Contract.INITMAXNUMOFMAINTAINING(&_Validatorset.CallOpts)
}

// INITNUMOFCABINETS is a free data retrieval call binding the contract method 0xb8cf4ef1.
//
// Solidity: function INIT_NUM_OF_CABINETS() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) INITNUMOFCABINETS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "INIT_NUM_OF_CABINETS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITNUMOFCABINETS is a free data retrieval call binding the contract method 0xb8cf4ef1.
//
// Solidity: function INIT_NUM_OF_CABINETS() view returns(uint256)
func (_Validatorset *ValidatorsetSession) INITNUMOFCABINETS() (*big.Int, error) {
	return _Validatorset.Contract.INITNUMOFCABINETS(&_Validatorset.CallOpts)
}

// INITNUMOFCABINETS is a free data retrieval call binding the contract method 0xb8cf4ef1.
//
// Solidity: function INIT_NUM_OF_CABINETS() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) INITNUMOFCABINETS() (*big.Int, error) {
	return _Validatorset.Contract.INITNUMOFCABINETS(&_Validatorset.CallOpts)
}

// INITSYSTEMREWARDRATIO is a free data retrieval call binding the contract method 0xc466689d.
//
// Solidity: function INIT_SYSTEM_REWARD_RATIO() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) INITSYSTEMREWARDRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "INIT_SYSTEM_REWARD_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITSYSTEMREWARDRATIO is a free data retrieval call binding the contract method 0xc466689d.
//
// Solidity: function INIT_SYSTEM_REWARD_RATIO() view returns(uint256)
func (_Validatorset *ValidatorsetSession) INITSYSTEMREWARDRATIO() (*big.Int, error) {
	return _Validatorset.Contract.INITSYSTEMREWARDRATIO(&_Validatorset.CallOpts)
}

// INITSYSTEMREWARDRATIO is a free data retrieval call binding the contract method 0xc466689d.
//
// Solidity: function INIT_SYSTEM_REWARD_RATIO() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) INITSYSTEMREWARDRATIO() (*big.Int, error) {
	return _Validatorset.Contract.INITSYSTEMREWARDRATIO(&_Validatorset.CallOpts)
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_Validatorset *ValidatorsetCaller) INITVALIDATORSETBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "INIT_VALIDATORSET_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_Validatorset *ValidatorsetSession) INITVALIDATORSETBYTES() ([]byte, error) {
	return _Validatorset.Contract.INITVALIDATORSETBYTES(&_Validatorset.CallOpts)
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_Validatorset *ValidatorsetCallerSession) INITVALIDATORSETBYTES() ([]byte, error) {
	return _Validatorset.Contract.INITVALIDATORSETBYTES(&_Validatorset.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) JAILMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "JAIL_MESSAGE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_Validatorset *ValidatorsetSession) JAILMESSAGETYPE() (uint8, error) {
	return _Validatorset.Contract.JAILMESSAGETYPE(&_Validatorset.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) JAILMESSAGETYPE() (uint8, error) {
	return _Validatorset.Contract.JAILMESSAGETYPE(&_Validatorset.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Validatorset.Contract.LIGHTCLIENTADDR(&_Validatorset.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Validatorset.Contract.LIGHTCLIENTADDR(&_Validatorset.CallOpts)
}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) MAXNUMOFVALIDATORS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "MAX_NUM_OF_VALIDATORS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_Validatorset *ValidatorsetSession) MAXNUMOFVALIDATORS() (*big.Int, error) {
	return _Validatorset.Contract.MAXNUMOFVALIDATORS(&_Validatorset.CallOpts)
}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) MAXNUMOFVALIDATORS() (*big.Int, error) {
	return _Validatorset.Contract.MAXNUMOFVALIDATORS(&_Validatorset.CallOpts)
}

// MAXSYSTEMREWARDBALANCE is a free data retrieval call binding the contract method 0xaef198a9.
//
// Solidity: function MAX_SYSTEM_REWARD_BALANCE() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) MAXSYSTEMREWARDBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "MAX_SYSTEM_REWARD_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSYSTEMREWARDBALANCE is a free data retrieval call binding the contract method 0xaef198a9.
//
// Solidity: function MAX_SYSTEM_REWARD_BALANCE() view returns(uint256)
func (_Validatorset *ValidatorsetSession) MAXSYSTEMREWARDBALANCE() (*big.Int, error) {
	return _Validatorset.Contract.MAXSYSTEMREWARDBALANCE(&_Validatorset.CallOpts)
}

// MAXSYSTEMREWARDBALANCE is a free data retrieval call binding the contract method 0xaef198a9.
//
// Solidity: function MAX_SYSTEM_REWARD_BALANCE() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) MAXSYSTEMREWARDBALANCE() (*big.Int, error) {
	return _Validatorset.Contract.MAXSYSTEMREWARDBALANCE(&_Validatorset.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Validatorset *ValidatorsetSession) PRECISION() (*big.Int, error) {
	return _Validatorset.Contract.PRECISION(&_Validatorset.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) PRECISION() (*big.Int, error) {
	return _Validatorset.Contract.PRECISION(&_Validatorset.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.RELAYERHUBCONTRACTADDR(&_Validatorset.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.RELAYERHUBCONTRACTADDR(&_Validatorset.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetSession) SLASHCHANNELID() (uint8, error) {
	return _Validatorset.Contract.SLASHCHANNELID(&_Validatorset.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Validatorset.Contract.SLASHCHANNELID(&_Validatorset.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.SLASHCONTRACTADDR(&_Validatorset.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.SLASHCONTRACTADDR(&_Validatorset.CallOpts)
}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) STAKECREDITADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "STAKE_CREDIT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) STAKECREDITADDR() (common.Address, error) {
	return _Validatorset.Contract.STAKECREDITADDR(&_Validatorset.CallOpts)
}

// STAKECREDITADDR is a free data retrieval call binding the contract method 0x7e434d54.
//
// Solidity: function STAKE_CREDIT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) STAKECREDITADDR() (common.Address, error) {
	return _Validatorset.Contract.STAKECREDITADDR(&_Validatorset.CallOpts)
}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) STAKEHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "STAKE_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) STAKEHUBADDR() (common.Address, error) {
	return _Validatorset.Contract.STAKEHUBADDR(&_Validatorset.CallOpts)
}

// STAKEHUBADDR is a free data retrieval call binding the contract method 0xaa82dce1.
//
// Solidity: function STAKE_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) STAKEHUBADDR() (common.Address, error) {
	return _Validatorset.Contract.STAKEHUBADDR(&_Validatorset.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetSession) STAKINGCHANNELID() (uint8, error) {
	return _Validatorset.Contract.STAKINGCHANNELID(&_Validatorset.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Validatorset.Contract.STAKINGCHANNELID(&_Validatorset.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) STAKINGCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "STAKING_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.STAKINGCONTRACTADDR(&_Validatorset.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.STAKINGCONTRACTADDR(&_Validatorset.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Validatorset.Contract.SYSTEMREWARDADDR(&_Validatorset.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Validatorset.Contract.SYSTEMREWARDADDR(&_Validatorset.CallOpts)
}

// SYSTEMREWARDRATIOSCALE is a free data retrieval call binding the contract method 0x603d86d3.
//
// Solidity: function SYSTEM_REWARD_RATIO_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) SYSTEMREWARDRATIOSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "SYSTEM_REWARD_RATIO_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SYSTEMREWARDRATIOSCALE is a free data retrieval call binding the contract method 0x603d86d3.
//
// Solidity: function SYSTEM_REWARD_RATIO_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetSession) SYSTEMREWARDRATIOSCALE() (*big.Int, error) {
	return _Validatorset.Contract.SYSTEMREWARDRATIOSCALE(&_Validatorset.CallOpts)
}

// SYSTEMREWARDRATIOSCALE is a free data retrieval call binding the contract method 0x603d86d3.
//
// Solidity: function SYSTEM_REWARD_RATIO_SCALE() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) SYSTEMREWARDRATIOSCALE() (*big.Int, error) {
	return _Validatorset.Contract.SYSTEMREWARDRATIOSCALE(&_Validatorset.CallOpts)
}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) TIMELOCKADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "TIMELOCK_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) TIMELOCKADDR() (common.Address, error) {
	return _Validatorset.Contract.TIMELOCKADDR(&_Validatorset.CallOpts)
}

// TIMELOCKADDR is a free data retrieval call binding the contract method 0x51b4dce3.
//
// Solidity: function TIMELOCK_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) TIMELOCKADDR() (common.Address, error) {
	return _Validatorset.Contract.TIMELOCKADDR(&_Validatorset.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) TOKENHUBADDR() (common.Address, error) {
	return _Validatorset.Contract.TOKENHUBADDR(&_Validatorset.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Validatorset.Contract.TOKENHUBADDR(&_Validatorset.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Validatorset.Contract.TOKENMANAGERADDR(&_Validatorset.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Validatorset.Contract.TOKENMANAGERADDR(&_Validatorset.CallOpts)
}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) TOKENRECOVERPORTALADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "TOKEN_RECOVER_PORTAL_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) TOKENRECOVERPORTALADDR() (common.Address, error) {
	return _Validatorset.Contract.TOKENRECOVERPORTALADDR(&_Validatorset.CallOpts)
}

// TOKENRECOVERPORTALADDR is a free data retrieval call binding the contract method 0xaad56063.
//
// Solidity: function TOKEN_RECOVER_PORTAL_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) TOKENRECOVERPORTALADDR() (common.Address, error) {
	return _Validatorset.Contract.TOKENRECOVERPORTALADDR(&_Validatorset.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Validatorset.Contract.TRANSFERINCHANNELID(&_Validatorset.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Validatorset.Contract.TRANSFERINCHANNELID(&_Validatorset.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Validatorset.Contract.TRANSFEROUTCHANNELID(&_Validatorset.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Validatorset.Contract.TRANSFEROUTCHANNELID(&_Validatorset.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_Validatorset *ValidatorsetCaller) VALIDATORSUPDATEMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "VALIDATORS_UPDATE_MESSAGE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_Validatorset *ValidatorsetSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _Validatorset.Contract.VALIDATORSUPDATEMESSAGETYPE(&_Validatorset.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_Validatorset *ValidatorsetCallerSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _Validatorset.Contract.VALIDATORSUPDATEMESSAGETYPE(&_Validatorset.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.VALIDATORCONTRACTADDR(&_Validatorset.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Validatorset *ValidatorsetCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Validatorset.Contract.VALIDATORCONTRACTADDR(&_Validatorset.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorset *ValidatorsetCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorset *ValidatorsetSession) AlreadyInit() (bool, error) {
	return _Validatorset.Contract.AlreadyInit(&_Validatorset.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorset *ValidatorsetCallerSession) AlreadyInit() (bool, error) {
	return _Validatorset.Contract.AlreadyInit(&_Validatorset.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Validatorset *ValidatorsetCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Validatorset *ValidatorsetSession) BscChainID() (uint16, error) {
	return _Validatorset.Contract.BscChainID(&_Validatorset.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Validatorset *ValidatorsetCallerSession) BscChainID() (uint16, error) {
	return _Validatorset.Contract.BscChainID(&_Validatorset.CallOpts)
}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) BurnRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "burnRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_Validatorset *ValidatorsetSession) BurnRatio() (*big.Int, error) {
	return _Validatorset.Contract.BurnRatio(&_Validatorset.CallOpts)
}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) BurnRatio() (*big.Int, error) {
	return _Validatorset.Contract.BurnRatio(&_Validatorset.CallOpts)
}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_Validatorset *ValidatorsetCaller) BurnRatioInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "burnRatioInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_Validatorset *ValidatorsetSession) BurnRatioInitialized() (bool, error) {
	return _Validatorset.Contract.BurnRatioInitialized(&_Validatorset.CallOpts)
}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_Validatorset *ValidatorsetCallerSession) BurnRatioInitialized() (bool, error) {
	return _Validatorset.Contract.BurnRatioInitialized(&_Validatorset.CallOpts)
}

// CanEnterMaintenance is a free data retrieval call binding the contract method 0x321d398a.
//
// Solidity: function canEnterMaintenance(uint256 index) view returns(bool)
func (_Validatorset *ValidatorsetCaller) CanEnterMaintenance(opts *bind.CallOpts, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "canEnterMaintenance", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEnterMaintenance is a free data retrieval call binding the contract method 0x321d398a.
//
// Solidity: function canEnterMaintenance(uint256 index) view returns(bool)
func (_Validatorset *ValidatorsetSession) CanEnterMaintenance(index *big.Int) (bool, error) {
	return _Validatorset.Contract.CanEnterMaintenance(&_Validatorset.CallOpts, index)
}

// CanEnterMaintenance is a free data retrieval call binding the contract method 0x321d398a.
//
// Solidity: function canEnterMaintenance(uint256 index) view returns(bool)
func (_Validatorset *ValidatorsetCallerSession) CanEnterMaintenance(index *big.Int) (bool, error) {
	return _Validatorset.Contract.CanEnterMaintenance(&_Validatorset.CallOpts, index)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Validatorset *ValidatorsetCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "currentValidatorSet", arg0)

	outstruct := new(struct {
		ConsensusAddress common.Address
		FeeAddress       common.Address
		BBCFeeAddress    common.Address
		VotingPower      uint64
		Jailed           bool
		Incoming         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FeeAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BBCFeeAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.VotingPower = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Jailed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Incoming = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Validatorset *ValidatorsetSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _Validatorset.Contract.CurrentValidatorSet(&_Validatorset.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Validatorset *ValidatorsetCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _Validatorset.Contract.CurrentValidatorSet(&_Validatorset.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validatorset *ValidatorsetCaller) CurrentValidatorSetMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "currentValidatorSetMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validatorset *ValidatorsetSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _Validatorset.Contract.CurrentValidatorSetMap(&_Validatorset.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _Validatorset.Contract.CurrentValidatorSetMap(&_Validatorset.CallOpts, arg0)
}

// CurrentVoteAddrFullSet is a free data retrieval call binding the contract method 0xce910b0c.
//
// Solidity: function currentVoteAddrFullSet(uint256 ) view returns(bytes)
func (_Validatorset *ValidatorsetCaller) CurrentVoteAddrFullSet(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "currentVoteAddrFullSet", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CurrentVoteAddrFullSet is a free data retrieval call binding the contract method 0xce910b0c.
//
// Solidity: function currentVoteAddrFullSet(uint256 ) view returns(bytes)
func (_Validatorset *ValidatorsetSession) CurrentVoteAddrFullSet(arg0 *big.Int) ([]byte, error) {
	return _Validatorset.Contract.CurrentVoteAddrFullSet(&_Validatorset.CallOpts, arg0)
}

// CurrentVoteAddrFullSet is a free data retrieval call binding the contract method 0xce910b0c.
//
// Solidity: function currentVoteAddrFullSet(uint256 ) view returns(bytes)
func (_Validatorset *ValidatorsetCallerSession) CurrentVoteAddrFullSet(arg0 *big.Int) ([]byte, error) {
	return _Validatorset.Contract.CurrentVoteAddrFullSet(&_Validatorset.CallOpts, arg0)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) ExpireTimeSecondGap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "expireTimeSecondGap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_Validatorset *ValidatorsetSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _Validatorset.Contract.ExpireTimeSecondGap(&_Validatorset.CallOpts)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _Validatorset.Contract.ExpireTimeSecondGap(&_Validatorset.CallOpts)
}

// GetCurrentValidatorIndex is a free data retrieval call binding the contract method 0x8d19a410.
//
// Solidity: function getCurrentValidatorIndex(address _validator) view returns(uint256)
func (_Validatorset *ValidatorsetCaller) GetCurrentValidatorIndex(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getCurrentValidatorIndex", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentValidatorIndex is a free data retrieval call binding the contract method 0x8d19a410.
//
// Solidity: function getCurrentValidatorIndex(address _validator) view returns(uint256)
func (_Validatorset *ValidatorsetSession) GetCurrentValidatorIndex(_validator common.Address) (*big.Int, error) {
	return _Validatorset.Contract.GetCurrentValidatorIndex(&_Validatorset.CallOpts, _validator)
}

// GetCurrentValidatorIndex is a free data retrieval call binding the contract method 0x8d19a410.
//
// Solidity: function getCurrentValidatorIndex(address _validator) view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) GetCurrentValidatorIndex(_validator common.Address) (*big.Int, error) {
	return _Validatorset.Contract.GetCurrentValidatorIndex(&_Validatorset.CallOpts, _validator)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_Validatorset *ValidatorsetCaller) GetIncoming(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getIncoming", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_Validatorset *ValidatorsetSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _Validatorset.Contract.GetIncoming(&_Validatorset.CallOpts, validator)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _Validatorset.Contract.GetIncoming(&_Validatorset.CallOpts, validator)
}

// GetLivingValidators is a free data retrieval call binding the contract method 0x3b071dcc.
//
// Solidity: function getLivingValidators() view returns(address[], bytes[])
func (_Validatorset *ValidatorsetCaller) GetLivingValidators(opts *bind.CallOpts) ([]common.Address, [][]byte, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getLivingValidators")

	if err != nil {
		return *new([]common.Address), *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return out0, out1, err

}

// GetLivingValidators is a free data retrieval call binding the contract method 0x3b071dcc.
//
// Solidity: function getLivingValidators() view returns(address[], bytes[])
func (_Validatorset *ValidatorsetSession) GetLivingValidators() ([]common.Address, [][]byte, error) {
	return _Validatorset.Contract.GetLivingValidators(&_Validatorset.CallOpts)
}

// GetLivingValidators is a free data retrieval call binding the contract method 0x3b071dcc.
//
// Solidity: function getLivingValidators() view returns(address[], bytes[])
func (_Validatorset *ValidatorsetCallerSession) GetLivingValidators() ([]common.Address, [][]byte, error) {
	return _Validatorset.Contract.GetLivingValidators(&_Validatorset.CallOpts)
}

// GetMiningValidators is a free data retrieval call binding the contract method 0x4df6e0c3.
//
// Solidity: function getMiningValidators() view returns(address[], bytes[])
func (_Validatorset *ValidatorsetCaller) GetMiningValidators(opts *bind.CallOpts) ([]common.Address, [][]byte, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getMiningValidators")

	if err != nil {
		return *new([]common.Address), *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return out0, out1, err

}

// GetMiningValidators is a free data retrieval call binding the contract method 0x4df6e0c3.
//
// Solidity: function getMiningValidators() view returns(address[], bytes[])
func (_Validatorset *ValidatorsetSession) GetMiningValidators() ([]common.Address, [][]byte, error) {
	return _Validatorset.Contract.GetMiningValidators(&_Validatorset.CallOpts)
}

// GetMiningValidators is a free data retrieval call binding the contract method 0x4df6e0c3.
//
// Solidity: function getMiningValidators() view returns(address[], bytes[])
func (_Validatorset *ValidatorsetCallerSession) GetMiningValidators() ([]common.Address, [][]byte, error) {
	return _Validatorset.Contract.GetMiningValidators(&_Validatorset.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validatorset *ValidatorsetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validatorset *ValidatorsetSession) GetValidators() ([]common.Address, error) {
	return _Validatorset.Contract.GetValidators(&_Validatorset.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validatorset *ValidatorsetCallerSession) GetValidators() ([]common.Address, error) {
	return _Validatorset.Contract.GetValidators(&_Validatorset.CallOpts)
}

// GetWorkingValidatorCount is a free data retrieval call binding the contract method 0xd68fb56a.
//
// Solidity: function getWorkingValidatorCount() view returns(uint256 workingValidatorCount)
func (_Validatorset *ValidatorsetCaller) GetWorkingValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getWorkingValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWorkingValidatorCount is a free data retrieval call binding the contract method 0xd68fb56a.
//
// Solidity: function getWorkingValidatorCount() view returns(uint256 workingValidatorCount)
func (_Validatorset *ValidatorsetSession) GetWorkingValidatorCount() (*big.Int, error) {
	return _Validatorset.Contract.GetWorkingValidatorCount(&_Validatorset.CallOpts)
}

// GetWorkingValidatorCount is a free data retrieval call binding the contract method 0xd68fb56a.
//
// Solidity: function getWorkingValidatorCount() view returns(uint256 workingValidatorCount)
func (_Validatorset *ValidatorsetCallerSession) GetWorkingValidatorCount() (*big.Int, error) {
	return _Validatorset.Contract.GetWorkingValidatorCount(&_Validatorset.CallOpts)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address validator) view returns(bool)
func (_Validatorset *ValidatorsetCaller) IsCurrentValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "isCurrentValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address validator) view returns(bool)
func (_Validatorset *ValidatorsetSession) IsCurrentValidator(validator common.Address) (bool, error) {
	return _Validatorset.Contract.IsCurrentValidator(&_Validatorset.CallOpts, validator)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address validator) view returns(bool)
func (_Validatorset *ValidatorsetCallerSession) IsCurrentValidator(validator common.Address) (bool, error) {
	return _Validatorset.Contract.IsCurrentValidator(&_Validatorset.CallOpts, validator)
}

// IsMonitoredForMaliciousVote is a free data retrieval call binding the contract method 0xea321e49.
//
// Solidity: function isMonitoredForMaliciousVote(bytes voteAddr) view returns(bool)
func (_Validatorset *ValidatorsetCaller) IsMonitoredForMaliciousVote(opts *bind.CallOpts, voteAddr []byte) (bool, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "isMonitoredForMaliciousVote", voteAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMonitoredForMaliciousVote is a free data retrieval call binding the contract method 0xea321e49.
//
// Solidity: function isMonitoredForMaliciousVote(bytes voteAddr) view returns(bool)
func (_Validatorset *ValidatorsetSession) IsMonitoredForMaliciousVote(voteAddr []byte) (bool, error) {
	return _Validatorset.Contract.IsMonitoredForMaliciousVote(&_Validatorset.CallOpts, voteAddr)
}

// IsMonitoredForMaliciousVote is a free data retrieval call binding the contract method 0xea321e49.
//
// Solidity: function isMonitoredForMaliciousVote(bytes voteAddr) view returns(bool)
func (_Validatorset *ValidatorsetCallerSession) IsMonitoredForMaliciousVote(voteAddr []byte) (bool, error) {
	return _Validatorset.Contract.IsMonitoredForMaliciousVote(&_Validatorset.CallOpts, voteAddr)
}

// IsSystemRewardIncluded is a free data retrieval call binding the contract method 0x8a7beb01.
//
// Solidity: function isSystemRewardIncluded() view returns(bool)
func (_Validatorset *ValidatorsetCaller) IsSystemRewardIncluded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "isSystemRewardIncluded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSystemRewardIncluded is a free data retrieval call binding the contract method 0x8a7beb01.
//
// Solidity: function isSystemRewardIncluded() view returns(bool)
func (_Validatorset *ValidatorsetSession) IsSystemRewardIncluded() (bool, error) {
	return _Validatorset.Contract.IsSystemRewardIncluded(&_Validatorset.CallOpts)
}

// IsSystemRewardIncluded is a free data retrieval call binding the contract method 0x8a7beb01.
//
// Solidity: function isSystemRewardIncluded() view returns(bool)
func (_Validatorset *ValidatorsetCallerSession) IsSystemRewardIncluded() (bool, error) {
	return _Validatorset.Contract.IsSystemRewardIncluded(&_Validatorset.CallOpts)
}

// IsWorkingValidator is a free data retrieval call binding the contract method 0x3365af3a.
//
// Solidity: function isWorkingValidator(uint256 index) view returns(bool)
func (_Validatorset *ValidatorsetCaller) IsWorkingValidator(opts *bind.CallOpts, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "isWorkingValidator", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWorkingValidator is a free data retrieval call binding the contract method 0x3365af3a.
//
// Solidity: function isWorkingValidator(uint256 index) view returns(bool)
func (_Validatorset *ValidatorsetSession) IsWorkingValidator(index *big.Int) (bool, error) {
	return _Validatorset.Contract.IsWorkingValidator(&_Validatorset.CallOpts, index)
}

// IsWorkingValidator is a free data retrieval call binding the contract method 0x3365af3a.
//
// Solidity: function isWorkingValidator(uint256 index) view returns(bool)
func (_Validatorset *ValidatorsetCallerSession) IsWorkingValidator(index *big.Int) (bool, error) {
	return _Validatorset.Contract.IsWorkingValidator(&_Validatorset.CallOpts, index)
}

// MaintainSlashScale is a free data retrieval call binding the contract method 0x8b5ad0c9.
//
// Solidity: function maintainSlashScale() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) MaintainSlashScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "maintainSlashScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaintainSlashScale is a free data retrieval call binding the contract method 0x8b5ad0c9.
//
// Solidity: function maintainSlashScale() view returns(uint256)
func (_Validatorset *ValidatorsetSession) MaintainSlashScale() (*big.Int, error) {
	return _Validatorset.Contract.MaintainSlashScale(&_Validatorset.CallOpts)
}

// MaintainSlashScale is a free data retrieval call binding the contract method 0x8b5ad0c9.
//
// Solidity: function maintainSlashScale() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) MaintainSlashScale() (*big.Int, error) {
	return _Validatorset.Contract.MaintainSlashScale(&_Validatorset.CallOpts)
}

// MaxNumOfCandidates is a free data retrieval call binding the contract method 0xe40716a1.
//
// Solidity: function maxNumOfCandidates() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) MaxNumOfCandidates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "maxNumOfCandidates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumOfCandidates is a free data retrieval call binding the contract method 0xe40716a1.
//
// Solidity: function maxNumOfCandidates() view returns(uint256)
func (_Validatorset *ValidatorsetSession) MaxNumOfCandidates() (*big.Int, error) {
	return _Validatorset.Contract.MaxNumOfCandidates(&_Validatorset.CallOpts)
}

// MaxNumOfCandidates is a free data retrieval call binding the contract method 0xe40716a1.
//
// Solidity: function maxNumOfCandidates() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) MaxNumOfCandidates() (*big.Int, error) {
	return _Validatorset.Contract.MaxNumOfCandidates(&_Validatorset.CallOpts)
}

// MaxNumOfMaintaining is a free data retrieval call binding the contract method 0x45cf9daf.
//
// Solidity: function maxNumOfMaintaining() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) MaxNumOfMaintaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "maxNumOfMaintaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumOfMaintaining is a free data retrieval call binding the contract method 0x45cf9daf.
//
// Solidity: function maxNumOfMaintaining() view returns(uint256)
func (_Validatorset *ValidatorsetSession) MaxNumOfMaintaining() (*big.Int, error) {
	return _Validatorset.Contract.MaxNumOfMaintaining(&_Validatorset.CallOpts)
}

// MaxNumOfMaintaining is a free data retrieval call binding the contract method 0x45cf9daf.
//
// Solidity: function maxNumOfMaintaining() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) MaxNumOfMaintaining() (*big.Int, error) {
	return _Validatorset.Contract.MaxNumOfMaintaining(&_Validatorset.CallOpts)
}

// MaxNumOfWorkingCandidates is a free data retrieval call binding the contract method 0xf92eb86b.
//
// Solidity: function maxNumOfWorkingCandidates() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) MaxNumOfWorkingCandidates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "maxNumOfWorkingCandidates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumOfWorkingCandidates is a free data retrieval call binding the contract method 0xf92eb86b.
//
// Solidity: function maxNumOfWorkingCandidates() view returns(uint256)
func (_Validatorset *ValidatorsetSession) MaxNumOfWorkingCandidates() (*big.Int, error) {
	return _Validatorset.Contract.MaxNumOfWorkingCandidates(&_Validatorset.CallOpts)
}

// MaxNumOfWorkingCandidates is a free data retrieval call binding the contract method 0xf92eb86b.
//
// Solidity: function maxNumOfWorkingCandidates() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) MaxNumOfWorkingCandidates() (*big.Int, error) {
	return _Validatorset.Contract.MaxNumOfWorkingCandidates(&_Validatorset.CallOpts)
}

// NumOfCabinets is a free data retrieval call binding the contract method 0x7a84ca2a.
//
// Solidity: function numOfCabinets() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) NumOfCabinets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "numOfCabinets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfCabinets is a free data retrieval call binding the contract method 0x7a84ca2a.
//
// Solidity: function numOfCabinets() view returns(uint256)
func (_Validatorset *ValidatorsetSession) NumOfCabinets() (*big.Int, error) {
	return _Validatorset.Contract.NumOfCabinets(&_Validatorset.CallOpts)
}

// NumOfCabinets is a free data retrieval call binding the contract method 0x7a84ca2a.
//
// Solidity: function numOfCabinets() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) NumOfCabinets() (*big.Int, error) {
	return _Validatorset.Contract.NumOfCabinets(&_Validatorset.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) NumOfJailed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "numOfJailed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Validatorset *ValidatorsetSession) NumOfJailed() (*big.Int, error) {
	return _Validatorset.Contract.NumOfJailed(&_Validatorset.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) NumOfJailed() (*big.Int, error) {
	return _Validatorset.Contract.NumOfJailed(&_Validatorset.CallOpts)
}

// NumOfMaintaining is a free data retrieval call binding the contract method 0x07a56847.
//
// Solidity: function numOfMaintaining() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) NumOfMaintaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "numOfMaintaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfMaintaining is a free data retrieval call binding the contract method 0x07a56847.
//
// Solidity: function numOfMaintaining() view returns(uint256)
func (_Validatorset *ValidatorsetSession) NumOfMaintaining() (*big.Int, error) {
	return _Validatorset.Contract.NumOfMaintaining(&_Validatorset.CallOpts)
}

// NumOfMaintaining is a free data retrieval call binding the contract method 0x07a56847.
//
// Solidity: function numOfMaintaining() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) NumOfMaintaining() (*big.Int, error) {
	return _Validatorset.Contract.NumOfMaintaining(&_Validatorset.CallOpts)
}

// PreviousBalanceOfSystemReward is a free data retrieval call binding the contract method 0x88b32f11.
//
// Solidity: function previousBalanceOfSystemReward() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) PreviousBalanceOfSystemReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "previousBalanceOfSystemReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousBalanceOfSystemReward is a free data retrieval call binding the contract method 0x88b32f11.
//
// Solidity: function previousBalanceOfSystemReward() view returns(uint256)
func (_Validatorset *ValidatorsetSession) PreviousBalanceOfSystemReward() (*big.Int, error) {
	return _Validatorset.Contract.PreviousBalanceOfSystemReward(&_Validatorset.CallOpts)
}

// PreviousBalanceOfSystemReward is a free data retrieval call binding the contract method 0x88b32f11.
//
// Solidity: function previousBalanceOfSystemReward() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) PreviousBalanceOfSystemReward() (*big.Int, error) {
	return _Validatorset.Contract.PreviousBalanceOfSystemReward(&_Validatorset.CallOpts)
}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) PreviousHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "previousHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() view returns(uint256)
func (_Validatorset *ValidatorsetSession) PreviousHeight() (*big.Int, error) {
	return _Validatorset.Contract.PreviousHeight(&_Validatorset.CallOpts)
}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) PreviousHeight() (*big.Int, error) {
	return _Validatorset.Contract.PreviousHeight(&_Validatorset.CallOpts)
}

// PreviousVoteAddrFullSet is a free data retrieval call binding the contract method 0x60eba4fe.
//
// Solidity: function previousVoteAddrFullSet(uint256 ) view returns(bytes)
func (_Validatorset *ValidatorsetCaller) PreviousVoteAddrFullSet(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "previousVoteAddrFullSet", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PreviousVoteAddrFullSet is a free data retrieval call binding the contract method 0x60eba4fe.
//
// Solidity: function previousVoteAddrFullSet(uint256 ) view returns(bytes)
func (_Validatorset *ValidatorsetSession) PreviousVoteAddrFullSet(arg0 *big.Int) ([]byte, error) {
	return _Validatorset.Contract.PreviousVoteAddrFullSet(&_Validatorset.CallOpts, arg0)
}

// PreviousVoteAddrFullSet is a free data retrieval call binding the contract method 0x60eba4fe.
//
// Solidity: function previousVoteAddrFullSet(uint256 ) view returns(bytes)
func (_Validatorset *ValidatorsetCallerSession) PreviousVoteAddrFullSet(arg0 *big.Int) ([]byte, error) {
	return _Validatorset.Contract.PreviousVoteAddrFullSet(&_Validatorset.CallOpts, arg0)
}

// SystemRewardRatio is a free data retrieval call binding the contract method 0x5de1e22c.
//
// Solidity: function systemRewardRatio() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) SystemRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "systemRewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SystemRewardRatio is a free data retrieval call binding the contract method 0x5de1e22c.
//
// Solidity: function systemRewardRatio() view returns(uint256)
func (_Validatorset *ValidatorsetSession) SystemRewardRatio() (*big.Int, error) {
	return _Validatorset.Contract.SystemRewardRatio(&_Validatorset.CallOpts)
}

// SystemRewardRatio is a free data retrieval call binding the contract method 0x5de1e22c.
//
// Solidity: function systemRewardRatio() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) SystemRewardRatio() (*big.Int, error) {
	return _Validatorset.Contract.SystemRewardRatio(&_Validatorset.CallOpts)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) TotalInComing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "totalInComing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_Validatorset *ValidatorsetSession) TotalInComing() (*big.Int, error) {
	return _Validatorset.Contract.TotalInComing(&_Validatorset.CallOpts)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) TotalInComing() (*big.Int, error) {
	return _Validatorset.Contract.TotalInComing(&_Validatorset.CallOpts)
}

// ValidatorExtraSet is a free data retrieval call binding the contract method 0xfd4ad81f.
//
// Solidity: function validatorExtraSet(uint256 ) view returns(uint256 enterMaintenanceHeight, bool isMaintaining, bytes voteAddress)
func (_Validatorset *ValidatorsetCaller) ValidatorExtraSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EnterMaintenanceHeight *big.Int
	IsMaintaining          bool
	VoteAddress            []byte
}, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "validatorExtraSet", arg0)

	outstruct := new(struct {
		EnterMaintenanceHeight *big.Int
		IsMaintaining          bool
		VoteAddress            []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EnterMaintenanceHeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsMaintaining = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.VoteAddress = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ValidatorExtraSet is a free data retrieval call binding the contract method 0xfd4ad81f.
//
// Solidity: function validatorExtraSet(uint256 ) view returns(uint256 enterMaintenanceHeight, bool isMaintaining, bytes voteAddress)
func (_Validatorset *ValidatorsetSession) ValidatorExtraSet(arg0 *big.Int) (struct {
	EnterMaintenanceHeight *big.Int
	IsMaintaining          bool
	VoteAddress            []byte
}, error) {
	return _Validatorset.Contract.ValidatorExtraSet(&_Validatorset.CallOpts, arg0)
}

// ValidatorExtraSet is a free data retrieval call binding the contract method 0xfd4ad81f.
//
// Solidity: function validatorExtraSet(uint256 ) view returns(uint256 enterMaintenanceHeight, bool isMaintaining, bytes voteAddress)
func (_Validatorset *ValidatorsetCallerSession) ValidatorExtraSet(arg0 *big.Int) (struct {
	EnterMaintenanceHeight *big.Int
	IsMaintaining          bool
	VoteAddress            []byte
}, error) {
	return _Validatorset.Contract.ValidatorExtraSet(&_Validatorset.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_Validatorset *ValidatorsetTransactor) Deposit(opts *bind.TransactOpts, valAddr common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "deposit", valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_Validatorset *ValidatorsetSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.Deposit(&_Validatorset.TransactOpts, valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_Validatorset *ValidatorsetTransactorSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.Deposit(&_Validatorset.TransactOpts, valAddr)
}

// DistributeFinalityReward is a paid mutator transaction binding the contract method 0x300c3567.
//
// Solidity: function distributeFinalityReward(address[] valAddrs, uint256[] weights) returns()
func (_Validatorset *ValidatorsetTransactor) DistributeFinalityReward(opts *bind.TransactOpts, valAddrs []common.Address, weights []*big.Int) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "distributeFinalityReward", valAddrs, weights)
}

// DistributeFinalityReward is a paid mutator transaction binding the contract method 0x300c3567.
//
// Solidity: function distributeFinalityReward(address[] valAddrs, uint256[] weights) returns()
func (_Validatorset *ValidatorsetSession) DistributeFinalityReward(valAddrs []common.Address, weights []*big.Int) (*types.Transaction, error) {
	return _Validatorset.Contract.DistributeFinalityReward(&_Validatorset.TransactOpts, valAddrs, weights)
}

// DistributeFinalityReward is a paid mutator transaction binding the contract method 0x300c3567.
//
// Solidity: function distributeFinalityReward(address[] valAddrs, uint256[] weights) returns()
func (_Validatorset *ValidatorsetTransactorSession) DistributeFinalityReward(valAddrs []common.Address, weights []*big.Int) (*types.Transaction, error) {
	return _Validatorset.Contract.DistributeFinalityReward(&_Validatorset.TransactOpts, valAddrs, weights)
}

// EnterMaintenance is a paid mutator transaction binding the contract method 0x9369d7de.
//
// Solidity: function enterMaintenance() returns()
func (_Validatorset *ValidatorsetTransactor) EnterMaintenance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "enterMaintenance")
}

// EnterMaintenance is a paid mutator transaction binding the contract method 0x9369d7de.
//
// Solidity: function enterMaintenance() returns()
func (_Validatorset *ValidatorsetSession) EnterMaintenance() (*types.Transaction, error) {
	return _Validatorset.Contract.EnterMaintenance(&_Validatorset.TransactOpts)
}

// EnterMaintenance is a paid mutator transaction binding the contract method 0x9369d7de.
//
// Solidity: function enterMaintenance() returns()
func (_Validatorset *ValidatorsetTransactorSession) EnterMaintenance() (*types.Transaction, error) {
	return _Validatorset.Contract.EnterMaintenance(&_Validatorset.TransactOpts)
}

// ExitMaintenance is a paid mutator transaction binding the contract method 0x04c4fec6.
//
// Solidity: function exitMaintenance() returns()
func (_Validatorset *ValidatorsetTransactor) ExitMaintenance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "exitMaintenance")
}

// ExitMaintenance is a paid mutator transaction binding the contract method 0x04c4fec6.
//
// Solidity: function exitMaintenance() returns()
func (_Validatorset *ValidatorsetSession) ExitMaintenance() (*types.Transaction, error) {
	return _Validatorset.Contract.ExitMaintenance(&_Validatorset.TransactOpts)
}

// ExitMaintenance is a paid mutator transaction binding the contract method 0x04c4fec6.
//
// Solidity: function exitMaintenance() returns()
func (_Validatorset *ValidatorsetTransactorSession) ExitMaintenance() (*types.Transaction, error) {
	return _Validatorset.Contract.ExitMaintenance(&_Validatorset.TransactOpts)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Validatorset *ValidatorsetTransactor) Felony(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "felony", validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Validatorset *ValidatorsetSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.Felony(&_Validatorset.TransactOpts, validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Validatorset *ValidatorsetTransactorSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.Felony(&_Validatorset.TransactOpts, validator)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Validatorset *ValidatorsetTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Validatorset *ValidatorsetSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.HandleAckPackage(&_Validatorset.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Validatorset *ValidatorsetTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.HandleAckPackage(&_Validatorset.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Validatorset *ValidatorsetTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Validatorset *ValidatorsetSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.HandleFailAckPackage(&_Validatorset.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Validatorset *ValidatorsetTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.HandleFailAckPackage(&_Validatorset.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Validatorset *ValidatorsetTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "handleSynPackage", arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Validatorset *ValidatorsetSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.HandleSynPackage(&_Validatorset.TransactOpts, arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Validatorset *ValidatorsetTransactorSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.HandleSynPackage(&_Validatorset.TransactOpts, arg0, msgBytes)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Validatorset *ValidatorsetTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Validatorset *ValidatorsetSession) Init() (*types.Transaction, error) {
	return _Validatorset.Contract.Init(&_Validatorset.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Validatorset *ValidatorsetTransactorSession) Init() (*types.Transaction, error) {
	return _Validatorset.Contract.Init(&_Validatorset.TransactOpts)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorset *ValidatorsetTransactor) JailValidator(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "jailValidator", consensusAddress)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorset *ValidatorsetSession) JailValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.JailValidator(&_Validatorset.TransactOpts, consensusAddress)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorset *ValidatorsetTransactorSession) JailValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.JailValidator(&_Validatorset.TransactOpts, consensusAddress)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Validatorset *ValidatorsetTransactor) Misdemeanor(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "misdemeanor", validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Validatorset *ValidatorsetSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.Misdemeanor(&_Validatorset.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Validatorset *ValidatorsetTransactorSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.Misdemeanor(&_Validatorset.TransactOpts, validator)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Validatorset *ValidatorsetTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Validatorset *ValidatorsetSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.UpdateParam(&_Validatorset.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Validatorset *ValidatorsetTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.UpdateParam(&_Validatorset.TransactOpts, key, value)
}

// UpdateValidatorSetV2 is a paid mutator transaction binding the contract method 0x1e4c1524.
//
// Solidity: function updateValidatorSetV2(address[] _consensusAddrs, uint64[] _votingPowers, bytes[] _voteAddrs) returns()
func (_Validatorset *ValidatorsetTransactor) UpdateValidatorSetV2(opts *bind.TransactOpts, _consensusAddrs []common.Address, _votingPowers []uint64, _voteAddrs [][]byte) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "updateValidatorSetV2", _consensusAddrs, _votingPowers, _voteAddrs)
}

// UpdateValidatorSetV2 is a paid mutator transaction binding the contract method 0x1e4c1524.
//
// Solidity: function updateValidatorSetV2(address[] _consensusAddrs, uint64[] _votingPowers, bytes[] _voteAddrs) returns()
func (_Validatorset *ValidatorsetSession) UpdateValidatorSetV2(_consensusAddrs []common.Address, _votingPowers []uint64, _voteAddrs [][]byte) (*types.Transaction, error) {
	return _Validatorset.Contract.UpdateValidatorSetV2(&_Validatorset.TransactOpts, _consensusAddrs, _votingPowers, _voteAddrs)
}

// UpdateValidatorSetV2 is a paid mutator transaction binding the contract method 0x1e4c1524.
//
// Solidity: function updateValidatorSetV2(address[] _consensusAddrs, uint64[] _votingPowers, bytes[] _voteAddrs) returns()
func (_Validatorset *ValidatorsetTransactorSession) UpdateValidatorSetV2(_consensusAddrs []common.Address, _votingPowers []uint64, _voteAddrs [][]byte) (*types.Transaction, error) {
	return _Validatorset.Contract.UpdateValidatorSetV2(&_Validatorset.TransactOpts, _consensusAddrs, _votingPowers, _voteAddrs)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validatorset *ValidatorsetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validatorset *ValidatorsetSession) Receive() (*types.Transaction, error) {
	return _Validatorset.Contract.Receive(&_Validatorset.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validatorset *ValidatorsetTransactorSession) Receive() (*types.Transaction, error) {
	return _Validatorset.Contract.Receive(&_Validatorset.TransactOpts)
}

// ValidatorsetBatchTransferIterator is returned from FilterBatchTransfer and is used to iterate over the raw logs and unpacked data for BatchTransfer events raised by the Validatorset contract.
type ValidatorsetBatchTransferIterator struct {
	Event *ValidatorsetBatchTransfer // Event containing the contract specifics and raw log

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
func (it *ValidatorsetBatchTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetBatchTransfer)
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
		it.Event = new(ValidatorsetBatchTransfer)
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
func (it *ValidatorsetBatchTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetBatchTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetBatchTransfer represents a BatchTransfer event raised by the Validatorset contract.
type ValidatorsetBatchTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransfer is a free log retrieval operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterBatchTransfer(opts *bind.FilterOpts) (*ValidatorsetBatchTransferIterator, error) {

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "batchTransfer")
	if err != nil {
		return nil, err
	}
	return &ValidatorsetBatchTransferIterator{contract: _Validatorset.contract, event: "batchTransfer", logs: logs, sub: sub}, nil
}

// WatchBatchTransfer is a free log subscription operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchBatchTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorsetBatchTransfer) (event.Subscription, error) {

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "batchTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetBatchTransfer)
				if err := _Validatorset.contract.UnpackLog(event, "batchTransfer", log); err != nil {
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

// ParseBatchTransfer is a log parse operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseBatchTransfer(log types.Log) (*ValidatorsetBatchTransfer, error) {
	event := new(ValidatorsetBatchTransfer)
	if err := _Validatorset.contract.UnpackLog(event, "batchTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetBatchTransferFailedIterator is returned from FilterBatchTransferFailed and is used to iterate over the raw logs and unpacked data for BatchTransferFailed events raised by the Validatorset contract.
type ValidatorsetBatchTransferFailedIterator struct {
	Event *ValidatorsetBatchTransferFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorsetBatchTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetBatchTransferFailed)
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
		it.Event = new(ValidatorsetBatchTransferFailed)
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
func (it *ValidatorsetBatchTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetBatchTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetBatchTransferFailed represents a BatchTransferFailed event raised by the Validatorset contract.
type ValidatorsetBatchTransferFailed struct {
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferFailed is a free log retrieval operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_Validatorset *ValidatorsetFilterer) FilterBatchTransferFailed(opts *bind.FilterOpts, amount []*big.Int) (*ValidatorsetBatchTransferFailedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "batchTransferFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetBatchTransferFailedIterator{contract: _Validatorset.contract, event: "batchTransferFailed", logs: logs, sub: sub}, nil
}

// WatchBatchTransferFailed is a free log subscription operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_Validatorset *ValidatorsetFilterer) WatchBatchTransferFailed(opts *bind.WatchOpts, sink chan<- *ValidatorsetBatchTransferFailed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "batchTransferFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetBatchTransferFailed)
				if err := _Validatorset.contract.UnpackLog(event, "batchTransferFailed", log); err != nil {
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

// ParseBatchTransferFailed is a log parse operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_Validatorset *ValidatorsetFilterer) ParseBatchTransferFailed(log types.Log) (*ValidatorsetBatchTransferFailed, error) {
	event := new(ValidatorsetBatchTransferFailed)
	if err := _Validatorset.contract.UnpackLog(event, "batchTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetBatchTransferLowerFailedIterator is returned from FilterBatchTransferLowerFailed and is used to iterate over the raw logs and unpacked data for BatchTransferLowerFailed events raised by the Validatorset contract.
type ValidatorsetBatchTransferLowerFailedIterator struct {
	Event *ValidatorsetBatchTransferLowerFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorsetBatchTransferLowerFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetBatchTransferLowerFailed)
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
		it.Event = new(ValidatorsetBatchTransferLowerFailed)
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
func (it *ValidatorsetBatchTransferLowerFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetBatchTransferLowerFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetBatchTransferLowerFailed represents a BatchTransferLowerFailed event raised by the Validatorset contract.
type ValidatorsetBatchTransferLowerFailed struct {
	Amount *big.Int
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferLowerFailed is a free log retrieval operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_Validatorset *ValidatorsetFilterer) FilterBatchTransferLowerFailed(opts *bind.FilterOpts, amount []*big.Int) (*ValidatorsetBatchTransferLowerFailedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "batchTransferLowerFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetBatchTransferLowerFailedIterator{contract: _Validatorset.contract, event: "batchTransferLowerFailed", logs: logs, sub: sub}, nil
}

// WatchBatchTransferLowerFailed is a free log subscription operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_Validatorset *ValidatorsetFilterer) WatchBatchTransferLowerFailed(opts *bind.WatchOpts, sink chan<- *ValidatorsetBatchTransferLowerFailed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "batchTransferLowerFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetBatchTransferLowerFailed)
				if err := _Validatorset.contract.UnpackLog(event, "batchTransferLowerFailed", log); err != nil {
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

// ParseBatchTransferLowerFailed is a log parse operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_Validatorset *ValidatorsetFilterer) ParseBatchTransferLowerFailed(log types.Log) (*ValidatorsetBatchTransferLowerFailed, error) {
	event := new(ValidatorsetBatchTransferLowerFailed)
	if err := _Validatorset.contract.UnpackLog(event, "batchTransferLowerFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetDeprecatedDepositIterator is returned from FilterDeprecatedDeposit and is used to iterate over the raw logs and unpacked data for DeprecatedDeposit events raised by the Validatorset contract.
type ValidatorsetDeprecatedDepositIterator struct {
	Event *ValidatorsetDeprecatedDeposit // Event containing the contract specifics and raw log

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
func (it *ValidatorsetDeprecatedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetDeprecatedDeposit)
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
		it.Event = new(ValidatorsetDeprecatedDeposit)
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
func (it *ValidatorsetDeprecatedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetDeprecatedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetDeprecatedDeposit represents a DeprecatedDeposit event raised by the Validatorset contract.
type ValidatorsetDeprecatedDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedDeposit is a free log retrieval operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterDeprecatedDeposit(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetDeprecatedDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "deprecatedDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetDeprecatedDepositIterator{contract: _Validatorset.contract, event: "deprecatedDeposit", logs: logs, sub: sub}, nil
}

// WatchDeprecatedDeposit is a free log subscription operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchDeprecatedDeposit(opts *bind.WatchOpts, sink chan<- *ValidatorsetDeprecatedDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "deprecatedDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetDeprecatedDeposit)
				if err := _Validatorset.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
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

// ParseDeprecatedDeposit is a log parse operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseDeprecatedDeposit(log types.Log) (*ValidatorsetDeprecatedDeposit, error) {
	event := new(ValidatorsetDeprecatedDeposit)
	if err := _Validatorset.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetDeprecatedFinalityRewardDepositIterator is returned from FilterDeprecatedFinalityRewardDeposit and is used to iterate over the raw logs and unpacked data for DeprecatedFinalityRewardDeposit events raised by the Validatorset contract.
type ValidatorsetDeprecatedFinalityRewardDepositIterator struct {
	Event *ValidatorsetDeprecatedFinalityRewardDeposit // Event containing the contract specifics and raw log

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
func (it *ValidatorsetDeprecatedFinalityRewardDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetDeprecatedFinalityRewardDeposit)
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
		it.Event = new(ValidatorsetDeprecatedFinalityRewardDeposit)
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
func (it *ValidatorsetDeprecatedFinalityRewardDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetDeprecatedFinalityRewardDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetDeprecatedFinalityRewardDeposit represents a DeprecatedFinalityRewardDeposit event raised by the Validatorset contract.
type ValidatorsetDeprecatedFinalityRewardDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedFinalityRewardDeposit is a free log retrieval operation binding the contract event 0xb9c75cbbfde137c4281689580799ef5f52144e78858f776a5979b2b212137d85.
//
// Solidity: event deprecatedFinalityRewardDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterDeprecatedFinalityRewardDeposit(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetDeprecatedFinalityRewardDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "deprecatedFinalityRewardDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetDeprecatedFinalityRewardDepositIterator{contract: _Validatorset.contract, event: "deprecatedFinalityRewardDeposit", logs: logs, sub: sub}, nil
}

// WatchDeprecatedFinalityRewardDeposit is a free log subscription operation binding the contract event 0xb9c75cbbfde137c4281689580799ef5f52144e78858f776a5979b2b212137d85.
//
// Solidity: event deprecatedFinalityRewardDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchDeprecatedFinalityRewardDeposit(opts *bind.WatchOpts, sink chan<- *ValidatorsetDeprecatedFinalityRewardDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "deprecatedFinalityRewardDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetDeprecatedFinalityRewardDeposit)
				if err := _Validatorset.contract.UnpackLog(event, "deprecatedFinalityRewardDeposit", log); err != nil {
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

// ParseDeprecatedFinalityRewardDeposit is a log parse operation binding the contract event 0xb9c75cbbfde137c4281689580799ef5f52144e78858f776a5979b2b212137d85.
//
// Solidity: event deprecatedFinalityRewardDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseDeprecatedFinalityRewardDeposit(log types.Log) (*ValidatorsetDeprecatedFinalityRewardDeposit, error) {
	event := new(ValidatorsetDeprecatedFinalityRewardDeposit)
	if err := _Validatorset.contract.UnpackLog(event, "deprecatedFinalityRewardDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetDirectTransferIterator is returned from FilterDirectTransfer and is used to iterate over the raw logs and unpacked data for DirectTransfer events raised by the Validatorset contract.
type ValidatorsetDirectTransferIterator struct {
	Event *ValidatorsetDirectTransfer // Event containing the contract specifics and raw log

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
func (it *ValidatorsetDirectTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetDirectTransfer)
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
		it.Event = new(ValidatorsetDirectTransfer)
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
func (it *ValidatorsetDirectTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetDirectTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetDirectTransfer represents a DirectTransfer event raised by the Validatorset contract.
type ValidatorsetDirectTransfer struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransfer is a free log retrieval operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterDirectTransfer(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetDirectTransferIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "directTransfer", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetDirectTransferIterator{contract: _Validatorset.contract, event: "directTransfer", logs: logs, sub: sub}, nil
}

// WatchDirectTransfer is a free log subscription operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchDirectTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorsetDirectTransfer, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "directTransfer", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetDirectTransfer)
				if err := _Validatorset.contract.UnpackLog(event, "directTransfer", log); err != nil {
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

// ParseDirectTransfer is a log parse operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseDirectTransfer(log types.Log) (*ValidatorsetDirectTransfer, error) {
	event := new(ValidatorsetDirectTransfer)
	if err := _Validatorset.contract.UnpackLog(event, "directTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetDirectTransferFailIterator is returned from FilterDirectTransferFail and is used to iterate over the raw logs and unpacked data for DirectTransferFail events raised by the Validatorset contract.
type ValidatorsetDirectTransferFailIterator struct {
	Event *ValidatorsetDirectTransferFail // Event containing the contract specifics and raw log

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
func (it *ValidatorsetDirectTransferFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetDirectTransferFail)
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
		it.Event = new(ValidatorsetDirectTransferFail)
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
func (it *ValidatorsetDirectTransferFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetDirectTransferFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetDirectTransferFail represents a DirectTransferFail event raised by the Validatorset contract.
type ValidatorsetDirectTransferFail struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransferFail is a free log retrieval operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterDirectTransferFail(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetDirectTransferFailIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "directTransferFail", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetDirectTransferFailIterator{contract: _Validatorset.contract, event: "directTransferFail", logs: logs, sub: sub}, nil
}

// WatchDirectTransferFail is a free log subscription operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchDirectTransferFail(opts *bind.WatchOpts, sink chan<- *ValidatorsetDirectTransferFail, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "directTransferFail", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetDirectTransferFail)
				if err := _Validatorset.contract.UnpackLog(event, "directTransferFail", log); err != nil {
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

// ParseDirectTransferFail is a log parse operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseDirectTransferFail(log types.Log) (*ValidatorsetDirectTransferFail, error) {
	event := new(ValidatorsetDirectTransferFail)
	if err := _Validatorset.contract.UnpackLog(event, "directTransferFail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetFailReasonWithStrIterator is returned from FilterFailReasonWithStr and is used to iterate over the raw logs and unpacked data for FailReasonWithStr events raised by the Validatorset contract.
type ValidatorsetFailReasonWithStrIterator struct {
	Event *ValidatorsetFailReasonWithStr // Event containing the contract specifics and raw log

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
func (it *ValidatorsetFailReasonWithStrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetFailReasonWithStr)
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
		it.Event = new(ValidatorsetFailReasonWithStr)
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
func (it *ValidatorsetFailReasonWithStrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetFailReasonWithStrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetFailReasonWithStr represents a FailReasonWithStr event raised by the Validatorset contract.
type ValidatorsetFailReasonWithStr struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailReasonWithStr is a free log retrieval operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Validatorset *ValidatorsetFilterer) FilterFailReasonWithStr(opts *bind.FilterOpts) (*ValidatorsetFailReasonWithStrIterator, error) {

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return &ValidatorsetFailReasonWithStrIterator{contract: _Validatorset.contract, event: "failReasonWithStr", logs: logs, sub: sub}, nil
}

// WatchFailReasonWithStr is a free log subscription operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Validatorset *ValidatorsetFilterer) WatchFailReasonWithStr(opts *bind.WatchOpts, sink chan<- *ValidatorsetFailReasonWithStr) (event.Subscription, error) {

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetFailReasonWithStr)
				if err := _Validatorset.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
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
func (_Validatorset *ValidatorsetFilterer) ParseFailReasonWithStr(log types.Log) (*ValidatorsetFailReasonWithStr, error) {
	event := new(ValidatorsetFailReasonWithStr)
	if err := _Validatorset.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetFeeBurnedIterator is returned from FilterFeeBurned and is used to iterate over the raw logs and unpacked data for FeeBurned events raised by the Validatorset contract.
type ValidatorsetFeeBurnedIterator struct {
	Event *ValidatorsetFeeBurned // Event containing the contract specifics and raw log

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
func (it *ValidatorsetFeeBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetFeeBurned)
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
		it.Event = new(ValidatorsetFeeBurned)
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
func (it *ValidatorsetFeeBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetFeeBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetFeeBurned represents a FeeBurned event raised by the Validatorset contract.
type ValidatorsetFeeBurned struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeBurned is a free log retrieval operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterFeeBurned(opts *bind.FilterOpts) (*ValidatorsetFeeBurnedIterator, error) {

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "feeBurned")
	if err != nil {
		return nil, err
	}
	return &ValidatorsetFeeBurnedIterator{contract: _Validatorset.contract, event: "feeBurned", logs: logs, sub: sub}, nil
}

// WatchFeeBurned is a free log subscription operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchFeeBurned(opts *bind.WatchOpts, sink chan<- *ValidatorsetFeeBurned) (event.Subscription, error) {

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "feeBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetFeeBurned)
				if err := _Validatorset.contract.UnpackLog(event, "feeBurned", log); err != nil {
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

// ParseFeeBurned is a log parse operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseFeeBurned(log types.Log) (*ValidatorsetFeeBurned, error) {
	event := new(ValidatorsetFeeBurned)
	if err := _Validatorset.contract.UnpackLog(event, "feeBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetFinalityRewardDepositIterator is returned from FilterFinalityRewardDeposit and is used to iterate over the raw logs and unpacked data for FinalityRewardDeposit events raised by the Validatorset contract.
type ValidatorsetFinalityRewardDepositIterator struct {
	Event *ValidatorsetFinalityRewardDeposit // Event containing the contract specifics and raw log

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
func (it *ValidatorsetFinalityRewardDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetFinalityRewardDeposit)
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
		it.Event = new(ValidatorsetFinalityRewardDeposit)
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
func (it *ValidatorsetFinalityRewardDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetFinalityRewardDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetFinalityRewardDeposit represents a FinalityRewardDeposit event raised by the Validatorset contract.
type ValidatorsetFinalityRewardDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFinalityRewardDeposit is a free log retrieval operation binding the contract event 0xcb0aad6cf9cd03bdf6137e359f541c42f38b39f007cae8e89e88aa7d8c6617b2.
//
// Solidity: event finalityRewardDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterFinalityRewardDeposit(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetFinalityRewardDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "finalityRewardDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetFinalityRewardDepositIterator{contract: _Validatorset.contract, event: "finalityRewardDeposit", logs: logs, sub: sub}, nil
}

// WatchFinalityRewardDeposit is a free log subscription operation binding the contract event 0xcb0aad6cf9cd03bdf6137e359f541c42f38b39f007cae8e89e88aa7d8c6617b2.
//
// Solidity: event finalityRewardDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchFinalityRewardDeposit(opts *bind.WatchOpts, sink chan<- *ValidatorsetFinalityRewardDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "finalityRewardDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetFinalityRewardDeposit)
				if err := _Validatorset.contract.UnpackLog(event, "finalityRewardDeposit", log); err != nil {
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

// ParseFinalityRewardDeposit is a log parse operation binding the contract event 0xcb0aad6cf9cd03bdf6137e359f541c42f38b39f007cae8e89e88aa7d8c6617b2.
//
// Solidity: event finalityRewardDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseFinalityRewardDeposit(log types.Log) (*ValidatorsetFinalityRewardDeposit, error) {
	event := new(ValidatorsetFinalityRewardDeposit)
	if err := _Validatorset.contract.UnpackLog(event, "finalityRewardDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Validatorset contract.
type ValidatorsetParamChangeIterator struct {
	Event *ValidatorsetParamChange // Event containing the contract specifics and raw log

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
func (it *ValidatorsetParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetParamChange)
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
		it.Event = new(ValidatorsetParamChange)
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
func (it *ValidatorsetParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetParamChange represents a ParamChange event raised by the Validatorset contract.
type ValidatorsetParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Validatorset *ValidatorsetFilterer) FilterParamChange(opts *bind.FilterOpts) (*ValidatorsetParamChangeIterator, error) {

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &ValidatorsetParamChangeIterator{contract: _Validatorset.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Validatorset *ValidatorsetFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *ValidatorsetParamChange) (event.Subscription, error) {

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetParamChange)
				if err := _Validatorset.contract.UnpackLog(event, "paramChange", log); err != nil {
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
func (_Validatorset *ValidatorsetFilterer) ParseParamChange(log types.Log) (*ValidatorsetParamChange, error) {
	event := new(ValidatorsetParamChange)
	if err := _Validatorset.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetSystemTransferIterator is returned from FilterSystemTransfer and is used to iterate over the raw logs and unpacked data for SystemTransfer events raised by the Validatorset contract.
type ValidatorsetSystemTransferIterator struct {
	Event *ValidatorsetSystemTransfer // Event containing the contract specifics and raw log

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
func (it *ValidatorsetSystemTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetSystemTransfer)
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
		it.Event = new(ValidatorsetSystemTransfer)
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
func (it *ValidatorsetSystemTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetSystemTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetSystemTransfer represents a SystemTransfer event raised by the Validatorset contract.
type ValidatorsetSystemTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSystemTransfer is a free log retrieval operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterSystemTransfer(opts *bind.FilterOpts) (*ValidatorsetSystemTransferIterator, error) {

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "systemTransfer")
	if err != nil {
		return nil, err
	}
	return &ValidatorsetSystemTransferIterator{contract: _Validatorset.contract, event: "systemTransfer", logs: logs, sub: sub}, nil
}

// WatchSystemTransfer is a free log subscription operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchSystemTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorsetSystemTransfer) (event.Subscription, error) {

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "systemTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetSystemTransfer)
				if err := _Validatorset.contract.UnpackLog(event, "systemTransfer", log); err != nil {
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

// ParseSystemTransfer is a log parse operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseSystemTransfer(log types.Log) (*ValidatorsetSystemTransfer, error) {
	event := new(ValidatorsetSystemTransfer)
	if err := _Validatorset.contract.UnpackLog(event, "systemTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetUnexpectedPackageIterator is returned from FilterUnexpectedPackage and is used to iterate over the raw logs and unpacked data for UnexpectedPackage events raised by the Validatorset contract.
type ValidatorsetUnexpectedPackageIterator struct {
	Event *ValidatorsetUnexpectedPackage // Event containing the contract specifics and raw log

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
func (it *ValidatorsetUnexpectedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetUnexpectedPackage)
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
		it.Event = new(ValidatorsetUnexpectedPackage)
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
func (it *ValidatorsetUnexpectedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetUnexpectedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetUnexpectedPackage represents a UnexpectedPackage event raised by the Validatorset contract.
type ValidatorsetUnexpectedPackage struct {
	ChannelId uint8
	MsgBytes  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedPackage is a free log retrieval operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Validatorset *ValidatorsetFilterer) FilterUnexpectedPackage(opts *bind.FilterOpts) (*ValidatorsetUnexpectedPackageIterator, error) {

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return &ValidatorsetUnexpectedPackageIterator{contract: _Validatorset.contract, event: "unexpectedPackage", logs: logs, sub: sub}, nil
}

// WatchUnexpectedPackage is a free log subscription operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Validatorset *ValidatorsetFilterer) WatchUnexpectedPackage(opts *bind.WatchOpts, sink chan<- *ValidatorsetUnexpectedPackage) (event.Subscription, error) {

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetUnexpectedPackage)
				if err := _Validatorset.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
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
func (_Validatorset *ValidatorsetFilterer) ParseUnexpectedPackage(log types.Log) (*ValidatorsetUnexpectedPackage, error) {
	event := new(ValidatorsetUnexpectedPackage)
	if err := _Validatorset.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetValidatorDepositIterator is returned from FilterValidatorDeposit and is used to iterate over the raw logs and unpacked data for ValidatorDeposit events raised by the Validatorset contract.
type ValidatorsetValidatorDepositIterator struct {
	Event *ValidatorsetValidatorDeposit // Event containing the contract specifics and raw log

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
func (it *ValidatorsetValidatorDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetValidatorDeposit)
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
		it.Event = new(ValidatorsetValidatorDeposit)
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
func (it *ValidatorsetValidatorDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetValidatorDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetValidatorDeposit represents a ValidatorDeposit event raised by the Validatorset contract.
type ValidatorsetValidatorDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeposit is a free log retrieval operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterValidatorDeposit(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetValidatorDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "validatorDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetValidatorDepositIterator{contract: _Validatorset.contract, event: "validatorDeposit", logs: logs, sub: sub}, nil
}

// WatchValidatorDeposit is a free log subscription operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchValidatorDeposit(opts *bind.WatchOpts, sink chan<- *ValidatorsetValidatorDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "validatorDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetValidatorDeposit)
				if err := _Validatorset.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
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

// ParseValidatorDeposit is a log parse operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseValidatorDeposit(log types.Log) (*ValidatorsetValidatorDeposit, error) {
	event := new(ValidatorsetValidatorDeposit)
	if err := _Validatorset.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetValidatorEmptyJailedIterator is returned from FilterValidatorEmptyJailed and is used to iterate over the raw logs and unpacked data for ValidatorEmptyJailed events raised by the Validatorset contract.
type ValidatorsetValidatorEmptyJailedIterator struct {
	Event *ValidatorsetValidatorEmptyJailed // Event containing the contract specifics and raw log

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
func (it *ValidatorsetValidatorEmptyJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetValidatorEmptyJailed)
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
		it.Event = new(ValidatorsetValidatorEmptyJailed)
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
func (it *ValidatorsetValidatorEmptyJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetValidatorEmptyJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetValidatorEmptyJailed represents a ValidatorEmptyJailed event raised by the Validatorset contract.
type ValidatorsetValidatorEmptyJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorEmptyJailed is a free log retrieval operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) FilterValidatorEmptyJailed(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetValidatorEmptyJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "validatorEmptyJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetValidatorEmptyJailedIterator{contract: _Validatorset.contract, event: "validatorEmptyJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorEmptyJailed is a free log subscription operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) WatchValidatorEmptyJailed(opts *bind.WatchOpts, sink chan<- *ValidatorsetValidatorEmptyJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "validatorEmptyJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetValidatorEmptyJailed)
				if err := _Validatorset.contract.UnpackLog(event, "validatorEmptyJailed", log); err != nil {
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

// ParseValidatorEmptyJailed is a log parse operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) ParseValidatorEmptyJailed(log types.Log) (*ValidatorsetValidatorEmptyJailed, error) {
	event := new(ValidatorsetValidatorEmptyJailed)
	if err := _Validatorset.contract.UnpackLog(event, "validatorEmptyJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetValidatorEnterMaintenanceIterator is returned from FilterValidatorEnterMaintenance and is used to iterate over the raw logs and unpacked data for ValidatorEnterMaintenance events raised by the Validatorset contract.
type ValidatorsetValidatorEnterMaintenanceIterator struct {
	Event *ValidatorsetValidatorEnterMaintenance // Event containing the contract specifics and raw log

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
func (it *ValidatorsetValidatorEnterMaintenanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetValidatorEnterMaintenance)
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
		it.Event = new(ValidatorsetValidatorEnterMaintenance)
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
func (it *ValidatorsetValidatorEnterMaintenanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetValidatorEnterMaintenanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetValidatorEnterMaintenance represents a ValidatorEnterMaintenance event raised by the Validatorset contract.
type ValidatorsetValidatorEnterMaintenance struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorEnterMaintenance is a free log retrieval operation binding the contract event 0xf62981a567ec3cec866c6fa93c55bcdf841d6292d18b8d522ececa769375d82d.
//
// Solidity: event validatorEnterMaintenance(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) FilterValidatorEnterMaintenance(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetValidatorEnterMaintenanceIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "validatorEnterMaintenance", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetValidatorEnterMaintenanceIterator{contract: _Validatorset.contract, event: "validatorEnterMaintenance", logs: logs, sub: sub}, nil
}

// WatchValidatorEnterMaintenance is a free log subscription operation binding the contract event 0xf62981a567ec3cec866c6fa93c55bcdf841d6292d18b8d522ececa769375d82d.
//
// Solidity: event validatorEnterMaintenance(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) WatchValidatorEnterMaintenance(opts *bind.WatchOpts, sink chan<- *ValidatorsetValidatorEnterMaintenance, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "validatorEnterMaintenance", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetValidatorEnterMaintenance)
				if err := _Validatorset.contract.UnpackLog(event, "validatorEnterMaintenance", log); err != nil {
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

// ParseValidatorEnterMaintenance is a log parse operation binding the contract event 0xf62981a567ec3cec866c6fa93c55bcdf841d6292d18b8d522ececa769375d82d.
//
// Solidity: event validatorEnterMaintenance(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) ParseValidatorEnterMaintenance(log types.Log) (*ValidatorsetValidatorEnterMaintenance, error) {
	event := new(ValidatorsetValidatorEnterMaintenance)
	if err := _Validatorset.contract.UnpackLog(event, "validatorEnterMaintenance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetValidatorExitMaintenanceIterator is returned from FilterValidatorExitMaintenance and is used to iterate over the raw logs and unpacked data for ValidatorExitMaintenance events raised by the Validatorset contract.
type ValidatorsetValidatorExitMaintenanceIterator struct {
	Event *ValidatorsetValidatorExitMaintenance // Event containing the contract specifics and raw log

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
func (it *ValidatorsetValidatorExitMaintenanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetValidatorExitMaintenance)
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
		it.Event = new(ValidatorsetValidatorExitMaintenance)
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
func (it *ValidatorsetValidatorExitMaintenanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetValidatorExitMaintenanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetValidatorExitMaintenance represents a ValidatorExitMaintenance event raised by the Validatorset contract.
type ValidatorsetValidatorExitMaintenance struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorExitMaintenance is a free log retrieval operation binding the contract event 0xb9d38178dc641ff1817967a63c9078cbcd955a9f1fcd75e0e3636de615d44d3b.
//
// Solidity: event validatorExitMaintenance(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) FilterValidatorExitMaintenance(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetValidatorExitMaintenanceIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "validatorExitMaintenance", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetValidatorExitMaintenanceIterator{contract: _Validatorset.contract, event: "validatorExitMaintenance", logs: logs, sub: sub}, nil
}

// WatchValidatorExitMaintenance is a free log subscription operation binding the contract event 0xb9d38178dc641ff1817967a63c9078cbcd955a9f1fcd75e0e3636de615d44d3b.
//
// Solidity: event validatorExitMaintenance(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) WatchValidatorExitMaintenance(opts *bind.WatchOpts, sink chan<- *ValidatorsetValidatorExitMaintenance, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "validatorExitMaintenance", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetValidatorExitMaintenance)
				if err := _Validatorset.contract.UnpackLog(event, "validatorExitMaintenance", log); err != nil {
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

// ParseValidatorExitMaintenance is a log parse operation binding the contract event 0xb9d38178dc641ff1817967a63c9078cbcd955a9f1fcd75e0e3636de615d44d3b.
//
// Solidity: event validatorExitMaintenance(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) ParseValidatorExitMaintenance(log types.Log) (*ValidatorsetValidatorExitMaintenance, error) {
	event := new(ValidatorsetValidatorExitMaintenance)
	if err := _Validatorset.contract.UnpackLog(event, "validatorExitMaintenance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetValidatorFelonyIterator is returned from FilterValidatorFelony and is used to iterate over the raw logs and unpacked data for ValidatorFelony events raised by the Validatorset contract.
type ValidatorsetValidatorFelonyIterator struct {
	Event *ValidatorsetValidatorFelony // Event containing the contract specifics and raw log

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
func (it *ValidatorsetValidatorFelonyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetValidatorFelony)
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
		it.Event = new(ValidatorsetValidatorFelony)
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
func (it *ValidatorsetValidatorFelonyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetValidatorFelonyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetValidatorFelony represents a ValidatorFelony event raised by the Validatorset contract.
type ValidatorsetValidatorFelony struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorFelony is a free log retrieval operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterValidatorFelony(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetValidatorFelonyIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "validatorFelony", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetValidatorFelonyIterator{contract: _Validatorset.contract, event: "validatorFelony", logs: logs, sub: sub}, nil
}

// WatchValidatorFelony is a free log subscription operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchValidatorFelony(opts *bind.WatchOpts, sink chan<- *ValidatorsetValidatorFelony, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "validatorFelony", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetValidatorFelony)
				if err := _Validatorset.contract.UnpackLog(event, "validatorFelony", log); err != nil {
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

// ParseValidatorFelony is a log parse operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseValidatorFelony(log types.Log) (*ValidatorsetValidatorFelony, error) {
	event := new(ValidatorsetValidatorFelony)
	if err := _Validatorset.contract.UnpackLog(event, "validatorFelony", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the Validatorset contract.
type ValidatorsetValidatorJailedIterator struct {
	Event *ValidatorsetValidatorJailed // Event containing the contract specifics and raw log

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
func (it *ValidatorsetValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetValidatorJailed)
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
		it.Event = new(ValidatorsetValidatorJailed)
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
func (it *ValidatorsetValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetValidatorJailed represents a ValidatorJailed event raised by the Validatorset contract.
type ValidatorsetValidatorJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetValidatorJailedIterator{contract: _Validatorset.contract, event: "validatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *ValidatorsetValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetValidatorJailed)
				if err := _Validatorset.contract.UnpackLog(event, "validatorJailed", log); err != nil {
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

// ParseValidatorJailed is a log parse operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Validatorset *ValidatorsetFilterer) ParseValidatorJailed(log types.Log) (*ValidatorsetValidatorJailed, error) {
	event := new(ValidatorsetValidatorJailed)
	if err := _Validatorset.contract.UnpackLog(event, "validatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetValidatorMisdemeanorIterator is returned from FilterValidatorMisdemeanor and is used to iterate over the raw logs and unpacked data for ValidatorMisdemeanor events raised by the Validatorset contract.
type ValidatorsetValidatorMisdemeanorIterator struct {
	Event *ValidatorsetValidatorMisdemeanor // Event containing the contract specifics and raw log

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
func (it *ValidatorsetValidatorMisdemeanorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetValidatorMisdemeanor)
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
		it.Event = new(ValidatorsetValidatorMisdemeanor)
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
func (it *ValidatorsetValidatorMisdemeanorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetValidatorMisdemeanorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetValidatorMisdemeanor represents a ValidatorMisdemeanor event raised by the Validatorset contract.
type ValidatorsetValidatorMisdemeanor struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorMisdemeanor is a free log retrieval operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) FilterValidatorMisdemeanor(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsetValidatorMisdemeanorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "validatorMisdemeanor", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetValidatorMisdemeanorIterator{contract: _Validatorset.contract, event: "validatorMisdemeanor", logs: logs, sub: sub}, nil
}

// WatchValidatorMisdemeanor is a free log subscription operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) WatchValidatorMisdemeanor(opts *bind.WatchOpts, sink chan<- *ValidatorsetValidatorMisdemeanor, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "validatorMisdemeanor", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetValidatorMisdemeanor)
				if err := _Validatorset.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
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

// ParseValidatorMisdemeanor is a log parse operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_Validatorset *ValidatorsetFilterer) ParseValidatorMisdemeanor(log types.Log) (*ValidatorsetValidatorMisdemeanor, error) {
	event := new(ValidatorsetValidatorMisdemeanor)
	if err := _Validatorset.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsetValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the Validatorset contract.
type ValidatorsetValidatorSetUpdatedIterator struct {
	Event *ValidatorsetValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorsetValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetValidatorSetUpdated)
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
		it.Event = new(ValidatorsetValidatorSetUpdated)
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
func (it *ValidatorsetValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetValidatorSetUpdated represents a ValidatorSetUpdated event raised by the Validatorset contract.
type ValidatorsetValidatorSetUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Validatorset *ValidatorsetFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts) (*ValidatorsetValidatorSetUpdatedIterator, error) {

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorsetValidatorSetUpdatedIterator{contract: _Validatorset.contract, event: "validatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Validatorset *ValidatorsetFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsetValidatorSetUpdated) (event.Subscription, error) {

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetValidatorSetUpdated)
				if err := _Validatorset.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Validatorset *ValidatorsetFilterer) ParseValidatorSetUpdated(log types.Log) (*ValidatorsetValidatorSetUpdated, error) {
	event := new(ValidatorsetValidatorSetUpdated)
	if err := _Validatorset.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
