package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/bep20"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/bep20Test"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/bscgovernor"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/crosschain"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/govtoken"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/stakehub"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/staking"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/tokenhub"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/tokenmanager"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/tokenrecoverportal"
	"github.com/bnb-chain/node-deploy/migrate_validator_tool/abi/validatorset"
)

const (
	ValidatorSetContractAddr  = "0x0000000000000000000000000000000000001000"
	TokenHubContractAddr      = "0x0000000000000000000000000000000000001004"
	GovHubContractAddr        = "0x0000000000000000000000000000000000001007"
	TokenManagerContractAddr  = "0x0000000000000000000000000000000000001008"
	StakeHubContractAddr      = "0x0000000000000000000000000000000000002002"
	CrossChainContractAddr    = "0x0000000000000000000000000000000000002000"
	StakingContractAddr       = "0x0000000000000000000000000000000000002001"
	BSCGovernorContractAddr   = "0x0000000000000000000000000000000000002004"
	GovTokenAddress           = "0x0000000000000000000000000000000000002005"
	TokenRecoveryContractAddr = "0x0000000000000000000000000000000000003000"
)

var (
	getChannelPermissionChannelID = flag.Int("get_channel_permission_channel", 0, "get channel permission")
	getChannelPermissionAddr      = flag.String("get_channel_permission_addr", "", "get channel permission")

	getSyncFee = flag.Bool("get_sync_fee", false, "")

	getValidatorElection = flag.Bool("get_validator_election", false, "")
	getValidatorSet      = flag.Bool("get_validator_set", false, "")
	getWorkingValidators = flag.Bool("get_working_validators", false, "")

	secretFile = flag.String("secret", "", "secret file path")
	password   = flag.String("password", "", "password")

	transferAmount = flag.String("amount", "0", "transfer amount")
	to             = flag.String("to", "", "transfer to address")
	gasPrice       = flag.String("gas_price", "10000000000", "transfer gasPrice")

	ccTo           = flag.String("cc_to", "", "cross chain transfer to address")
	ccContractFlag = flag.String("cc_contract", "", "cross chain transfer contract token")

	bscPrivKey              = flag.String("priv_key", "", "BSC validator operator private key")
	bscEndpoint             = flag.String("bsc_endpoint", "", "BSC RPC endpoint")
	consensusAddr           = flag.String("consensus_addr", "", "BSC validator operator consensus address")
	voteAddr                = flag.String("vote_addr", "", "BSC validator operator address")
	voteBLSProof            = flag.String("vote_bls_proof", "", "BSC validator operator bls proof")
	commissionRate          = flag.Int("commission_rate", 0, "BSC validator operator commission rate")
	commissionMaxRate       = flag.Int("commission_max_rate", 0, "BSC validator operator commission max rate")
	commissionMaxChangeRate = flag.Int("commission_max_change_rate", 0, "BSC validator operator commission max change rate")
	moniker                 = flag.String("moniker", "", "BSC validator operator moniker")
	identity                = flag.String("identity", "", "BSC validator operator identity")
	website                 = flag.String("website", "", "BSC validator operator website")
	details                 = flag.String("details", "", "BSC validator operator details")
	delegation              = flag.String("delegation", "0", "BSC validator operator delegation")

	tokenRecoverMerkleRoot    = flag.String("token_recover_merkle_root", "", "token recover merkle root")
	tokenRecoverProcter       = flag.String("token_recover_procter", "", "token recover procter")
	tokenRecoverApprover      = flag.String("token_recover_approver", "", "token recover approver")
	delegatorVoteOperatorAddr = flag.String("delegator_vote_operator_addr", "", "delegator vote operator address")

	checkProposalID                 = flag.String("check_proposal_id", "", "check proposal id")
	checkTokenRecoverContractStatus = flag.Bool("check_token_recover_contract_status", false, "check token recover contract status")

	recoverSymbol            = flag.String("recover_symbol", "", "recover symbol")
	recoverAmount            = flag.String("recover_amount", "", "recover amount")
	recoverPubKey            = flag.String("recover_pub_key", "", "recover pub key")
	recoverOwnerSignature    = flag.String("recover_owner_signature", "", "recover owner signature")
	recoverApprovalSignature = flag.String("recover_approval_signature", "", "recover approval signature")
	recoverProofs            = flag.String("recover_proofs", "", "recover proofs")
	recoverContract          = flag.String("recover_contract", "", "recover contract")

	deployContractFlag        = flag.String("deploy_contract", "", "deploy contract bytecode")
	tokenBindSymbolFlag       = flag.String("token_bind_symbol", "", "token bind symbol")
	tokenBindContractAddrFlag = flag.String("token_bind_contract_addr", "", "token bind contract address")
	tokenBindAmountFlag       = flag.String("token_bind_amount", "", "token bind amount")

	checkBEP20BalanceAddrFlag = flag.String("check_bep20_balance_addr", "", "check bep20 balance address")
	checkBEP20BalanceAccFlag  = flag.String("check_bep20_balance_acc", "", "check bep20 balance acc")

	mirrorTokenContractFlag          = flag.String("mirror_token_contract", "", "mirror token contract addr")
	getBEP2SymbolByContractFlag      = flag.String("get_bep2_symbol_by_contract", "", "get bep2 symbol by contract")
	syncContractAddrFlag             = flag.String("sync_contract_addr", "", "sync contract address")
	mintMoreMirrorTokenFlag          = flag.String("mint_more_mirror_token", "", "mint more mirror token")
	mirrorTokenTransferOutFlag       = flag.String("mirror_token_transfer_out", "", "mirror token transfer out")
	mirrorTokenTransferOutAmountFlag = flag.String("mirror_token_transfer_out_amount", "", "mirror token transfer out amount")

	getUndelegatedAmountFlag = flag.String("get_undelegated_amount", "", "get undelegated amount")

	claimTokenFlag = flag.Bool("claim_token", false, "claim token")
)

func main() {
	flag.Parse()

	if *secretFile != "" {
		convertGethPrivateKey()
		return
	}

	ethClient, err := ethclient.Dial(*bscEndpoint)
	if err != nil {
		panic(err)
	}

	if *transferAmount != "0" && len(*to) > 0 {
		from, err := FromHexKey(*bscPrivKey)
		if err != nil {
			panic(err)
		}
		toAddr := common.HexToAddress(*to)
		amount, _ := new(big.Int).SetString(*transferAmount, 10)
		gasPrice, _ := new(big.Int).SetString(*gasPrice, 10)
		fmt.Println("transfer amount:", amount, "from:", from.Addr, "to:", toAddr.String(), "gasPrice:", gasPrice)
		err = nativeTransfer(ethClient, &from, toAddr, amount, gasPrice)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	stakeHubContract, err := stakehub.NewStakehub(common.HexToAddress(StakeHubContractAddr), ethClient)
	if err != nil {
		panic(err)
	}

	validatorContract, err := validatorset.NewValidatorset(common.HexToAddress(ValidatorSetContractAddr), ethClient)
	if err != nil {
		panic(err)
	}

	governorContract, err := bscgovernor.NewBscgovernor(common.HexToAddress(BSCGovernorContractAddr), ethClient)
	if err != nil {
		panic(err)
	}
	govTokenContract, err := govtoken.NewGovtoken(common.HexToAddress(GovTokenAddress), ethClient)
	if err != nil {
		panic(err)
	}
	tokenRecoverContract, err := tokenrecoverportal.NewTokenrecoverportal(common.HexToAddress(TokenRecoveryContractAddr), ethClient)
	if err != nil {
		panic(err)
	}
	tokenHubContract, err := tokenhub.NewTokenhub(common.HexToAddress(TokenHubContractAddr), ethClient)
	if err != nil {
		panic(err)
	}
	tokenManagerContract, err := tokenmanager.NewTokenmanager(common.HexToAddress(TokenManagerContractAddr), ethClient)
	if err != nil {
		panic(err)
	}
	stakingContract, err := staking.NewStaking(common.HexToAddress(StakingContractAddr), ethClient)
	if err != nil {
		panic(err)
	}

	if *getValidatorElection {
		getElectionInfo(stakeHubContract)
		return
	}
	if *getValidatorSet {
		getCurValidatorSet(validatorContract)
		return
	}
	if *getWorkingValidators {
		getWorkingValidatorSet(validatorContract)
		return
	}
	if *getChannelPermissionAddr != "" {
		c, err := crosschain.NewCrosschain(common.HexToAddress(CrossChainContractAddr), ethClient)
		if err != nil {
			panic(err)
		}
		getChannelPermissionFromContract(
			c, common.HexToAddress(*getChannelPermissionAddr), *getChannelPermissionChannelID,
		)
		return
	}
	if *checkTokenRecoverContractStatus {
		err = checkTokenRecoverStatus(tokenRecoverContract)
		if err != nil {
			panic(err)
		}
		return
	}
	if *getSyncFee {
		getSyncFeeFromContract(tokenManagerContract)
		return
	}

	if len(*checkProposalID) > 0 {
		proposalId, _ := new(big.Int).SetString(*checkProposalID, 10)
		err = checkProposal(governorContract, proposalId)
		if err != nil {
			panic(err)
		}
		return
	}

	if len(*checkBEP20BalanceAccFlag) > 0 && len(*checkBEP20BalanceAddrFlag) > 0 {
		user := common.HexToAddress(*checkBEP20BalanceAccFlag)
		bep20Addr := common.HexToAddress(*checkBEP20BalanceAddrFlag)
		bep20Contract, err := bep20.NewBep20(bep20Addr, ethClient)
		if err != nil {
			panic(err)
		}
		balance, err := bep20Contract.BalanceOf(nil, user)
		if err != nil {
			panic(err)
		}
		fmt.Printf("balance of [%s]: %s", user, balance)
		return
	}

	if len(*getBEP2SymbolByContractFlag) > 0 {
		c := common.HexToAddress(*getBEP2SymbolByContractFlag)
		pending, err := tokenManagerContract.MirrorPendingRecord(nil, c)
		if err != nil {
			panic(err)
		}
		if pending {
			fmt.Println("token mirror is pending")
			return
		}
		bep2Symbol, err := tokenHubContract.GetBep2SymbolByContractAddr(nil, c)
		if err != nil {
			panic(err)
		}

		fmt.Println("BEP2 symbol:", string(bep2Symbol[:]))
		return
	}

	if len(*getUndelegatedAmountFlag) > 0 {
		amt, err := stakingContract.GetUndelegated(nil, common.HexToAddress(*getUndelegatedAmountFlag))
		if err != nil {
			panic(err)
		}
		fmt.Println("undelegated amount:", amt)
		return
	}

	acc, err := FromHexKey(*bscPrivKey)
	if err != nil {
		panic(err)
	}

	if *claimTokenFlag {
		err = claimToken(tokenHubContract, ethClient, &acc)
		if err != nil {
			panic(err)
		}
		return
	}

	if len(*mirrorTokenContractFlag) > 0 {
		c := common.HexToAddress(*mirrorTokenContractFlag)
		err = mirrorToken(ethClient, acc, c, tokenManagerContract, tokenHubContract)
		if err != nil {
			panic(err)
		}
		return
	}

	if len(*mintMoreMirrorTokenFlag) > 0 {
		contract, err := bep20Test.NewBep20Test(common.HexToAddress(*mintMoreMirrorTokenFlag), ethClient)
		if err != nil {
			panic(err)
		}
		txOpt, err := acc.BuildTransactOpts(context.Background(), ethClient, big.NewInt(0),
			nil, 3000000)
		if err != nil {
			panic(err)
		}
		amt, _ := new(big.Int).SetString("100000000000000000000000000", 10)
		tx, err := contract.MintMore(txOpt, amt)
		if err != nil {
			panic(err)
		}
		fmt.Println("mint more tx:", tx.Hash().String())
		r, err := bind.WaitMined(context.Background(), ethClient, tx)
		if err != nil {
			panic(err)
		}
		if r.Status != 1 {
			panic("mint more bep20 token tx execution failed")
		}
		return
	}

	if len(*syncContractAddrFlag) > 0 {
		err = syncToken(ethClient, acc, common.HexToAddress(*syncContractAddrFlag), tokenManagerContract)
		if err != nil {
			panic(err)
		}
		return
	}

	if len(*deployContractFlag) > 0 {
		f, err := os.Open(*deployContractFlag)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		bytecodeStr, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		bytecode, err := hex.DecodeString(string(bytecodeStr))
		if err != nil {
			panic(err)
		}
		_, contractAddr, err := deployContract(ethClient, acc, bytecode)
		if err != nil {
			panic(err)
		}
		fmt.Println("contract address:", contractAddr)
		return
	}

	if len(*tokenBindContractAddrFlag) > 0 && len(*tokenBindAmountFlag) > 0 && len(*tokenBindSymbolFlag) > 0 {
		bep20Addr := common.HexToAddress(*tokenBindContractAddrFlag)
		amt, _ := new(big.Int).SetString(*tokenBindAmountFlag, 10)
		err = approveBind(ethClient, acc, bep20Addr, tokenManagerContract, *tokenBindSymbolFlag, amt)
		if err != nil {
			panic(err)
		}
		return
	}

	if *recoverSymbol != "" && *recoverAmount != "" && *recoverPubKey != "" && *recoverApprovalSignature != "" &&
		*recoverOwnerSignature != "" && *recoverProofs != "" {
		amt, _ := new(big.Int).SetString(*recoverAmount, 10)
		pubkey := hexutil.MustDecode(*recoverPubKey)
		ownerSig := hexutil.MustDecode(*recoverOwnerSignature)
		approvalSig := hexutil.MustDecode(*recoverApprovalSignature)
		proofStr := strings.Split(*recoverProofs, ",")
		proofs := make([][32]byte, 0, len(proofStr))
		for _, p := range proofStr {
			var buf [32]byte
			copy(buf[:], hexutil.MustDecode(p))
			proofs = append(proofs, buf)
		}
		fmt.Println("recover symbol:", *recoverSymbol, "amount:", amt, "pubkey:", pubkey, "ownerSig:", ownerSig, "approvalSig:", approvalSig, "proofs:", proofs)
		err = tokenRecover(tokenHubContract, tokenRecoverContract, ethClient, &acc,
			*recoverSymbol, amt, pubkey, ownerSig, approvalSig, proofs)
		if err != nil {
			panic(err)
		}
		return
	}

	if *transferAmount != "0" && len(*ccTo) > 0 {
		c, err := tokenhub.NewTokenhub(common.HexToAddress(TokenHubContractAddr), ethClient)
		if err != nil {
			panic(err)
		}

		recipient := common.HexToAddress(*ccTo)
		fmt.Println("transfer amount:", *transferAmount, "from:", acc.Addr, "to:", recipient.String())
		gasLimit := uint64(3000000)
		amt, _ := new(big.Int).SetString(*transferAmount, 10)
		relayerFee := big.NewInt(1e18) // relayer fee
		bnbAmt := new(big.Int).Add(amt, relayerFee)
		ccContract := common.HexToAddress("0x0000000000000000000000000000000000000000")
		if len(*ccContractFlag) > 0 {
			ccContract = common.HexToAddress(*ccContractFlag)
			fmt.Println("ccContract:", ccContract.String(), "amt:", amt)
			bep20Contract, err := bep20.NewBep20(ccContract, ethClient)
			if err != nil {
				panic(err)
			}
			txOpt, err := acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, nil, gasLimit)
			if err != nil {
				panic(err)
			}
			tx, err := bep20Contract.Approve(txOpt, common.HexToAddress(TokenHubContractAddr), amt)
			if err != nil {
				panic(err)
			}
			fmt.Println("approve tx:", tx.Hash().String())
			r, err := bind.WaitMined(context.Background(), ethClient, tx)
			if err != nil {
				panic(err)
			}
			if r.Status != 1 {
				panic("approve tx failed")
			}

			bnbAmt = relayerFee
		}

		txOpt, err := acc.BuildTransactOpts(context.Background(), ethClient, bnbAmt, nil, gasLimit)
		if err != nil {
			panic(err)
		}

		tx, err := c.TransferOut(txOpt, ccContract,
			recipient, amt, uint64(time.Now().Add(3*time.Minute).Unix()))
		if err != nil {
			panic(err)
		}
		println("transfer out tx:", tx.Hash().String())
		r, err := bind.WaitMined(context.Background(), ethClient, tx)
		if err != nil {
			panic(err)
		}
		if r.Status != 1 {
			panic("tx failed")
		}
		return
	}

	if *tokenRecoverMerkleRoot != "" && *tokenRecoverProcter != "" && *tokenRecoverApprover != "" {
		setupTokenRecoveryContract(&acc, ethClient,
			tokenRecoverContract, governorContract,
			govTokenContract,
			stakeHubContract, validatorContract,
			*delegatorVoteOperatorAddr,
			*tokenRecoverMerkleRoot, common.HexToAddress(*tokenRecoverProcter), common.HexToAddress(*tokenRecoverApprover))
		return
	}

	delegation, _ := new(big.Int).SetString(*delegation, 10)
	balance, err := ethClient.BalanceAt(context.Background(), acc.Addr, nil)
	if err != nil {
		panic(err)
	}

	if balance.Cmp(delegation) < 0 {
		panic(fmt.Sprintln(acc.Addr, "insufficient balance:", balance, delegation))
	}

	gasLimit := uint64(3000000)
	txOpt, err := acc.BuildTransactOpts(context.Background(), ethClient, delegation, nil, gasLimit)
	if err != nil {
		panic(err)
	}
	voteAddrBytes, err := hexutil.Decode(*voteAddr)
	if err != nil {
		panic(err)
	}
	voteBLSProofBytes, err := hexutil.Decode(*voteBLSProof)
	if err != nil {
		panic(err)
	}
	consAddr := common.HexToAddress(*consensusAddr)

	tx, err := stakeHubContract.CreateValidator(txOpt, consAddr,
		voteAddrBytes, voteBLSProofBytes, stakehub.StakeHubCommission{
			Rate:          uint64(*commissionRate),
			MaxRate:       uint64(*commissionMaxRate),
			MaxChangeRate: uint64(*commissionMaxChangeRate),
		},
		stakehub.StakeHubDescription{
			Moniker:  *moniker,
			Identity: *identity,
			Website:  *website,
			Details:  *details,
		},
	)

	if err != nil {
		panic(err)
	}
	println("create validator tx:", tx.Hash().String())
	r, err := bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		panic(err)
	}
	if r.Status != 1 {
		panic("tx failed")
	}
}

type ExtAcc struct {
	Key  *ecdsa.PrivateKey
	Addr common.Address
}

func FromHexKey(hexkey string) (ExtAcc, error) {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return ExtAcc{}, err
	}
	pubKey := key.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("publicKey is not of type *ecdsa.PublicKey")
		return ExtAcc{}, err
	}
	addr := crypto.PubkeyToAddress(*pubKeyECDSA)
	return ExtAcc{key, addr}, nil
}

func (acc *ExtAcc) BuildTransactOpts(
	ctx context.Context, client *ethclient.Client,
	value *big.Int, gasPrice *big.Int, gasLimit uint64) (*bind.TransactOpts, error) {
	var err error
	if gasPrice == nil {
		gasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}
	}

	nonce, err := client.PendingNonceAt(context.Background(), acc.Addr)
	if err != nil {
		return nil, err
	}
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	//
	transactOpts, err := bind.NewKeyedTransactorWithChainID(acc.Key, chainId)
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = value
	transactOpts.GasLimit = gasLimit
	transactOpts.GasPrice = gasPrice
	return transactOpts, nil
}

func nativeTransfer(rpcClient *ethclient.Client, from *ExtAcc, to common.Address, amount *big.Int, gasPrice *big.Int) error {
	fromAddress := from.Addr
	nonce, err := rpcClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	suggestGasPrice, err := rpcClient.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	if suggestGasPrice.Cmp(gasPrice) > 0 {
		gasPrice = suggestGasPrice
	}

	chainID, err := rpcClient.NetworkID(context.Background())
	if err != nil {
		return err
	}

	gasLimit := uint64(3000000)

	tx := ethtypes.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), from.Key)
	if err != nil {
		return err
	}

	balance, err := rpcClient.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return err
	}
	fmt.Println("bsc holder balance:", balance)
	err = rpcClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	r, err := bind.WaitMined(context.Background(), rpcClient, signedTx)
	if err != nil {
		return err
	}

	if r.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("tx failed to be executed on chain")
	}

	return nil
}

func convertGethPrivateKey() {
	keyjson, err := os.ReadFile(*secretFile)
	if err != nil {
		fmt.Println("read secret error:", err)
		return
	}
	key, err := keystore.DecryptKey(keyjson, *password)
	if err != nil {
		fmt.Println("decrypt key error:", err)
		return
	}

	fmt.Println(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)))
}

func getElectionInfo(contract *stakehub.Stakehub) error {
	maxElection, err := contract.MaxElectedValidators(nil)
	if err != nil {
		panic(err)
	}
	validators, err := contract.GetValidatorElectionInfo(nil, big.NewInt(0), maxElection)
	if err != nil {
		panic(err)
	}
	fmt.Printf("MaxElectedValidators: %+v\n", maxElection.Uint64())
	fmt.Printf("%+v\n", validators)
	return nil
}

func getCurValidatorSet(contract *validatorset.Validatorset) error {
	address, err := contract.GetValidators(nil)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(address); i++ {
		validators, err := contract.CurrentValidatorSet(nil, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", validators)
	}

	return nil
}

func getWorkingValidatorSet(contract *validatorset.Validatorset) error {
	address, _, err := contract.GetMiningValidators(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(address)
	return nil
}

func getChannelPermissionFromContract(contract *crosschain.Crosschain, addr common.Address, channelID int) {
	permission, err := contract.RegisteredContractChannelMap(nil, addr, uint8(channelID))
	if err != nil {
		panic(err)
	}
	fmt.Printf("channel permission for [%s]: %v\n", addr, permission)
}

func setupTokenRecoveryContract(acc *ExtAcc, ethClient *ethclient.Client,
	tokenRecoverContract *tokenrecoverportal.Tokenrecoverportal,
	govContract *bscgovernor.Bscgovernor, govTokenContract *govtoken.Govtoken,
	stakeHubContract *stakehub.Stakehub,
	validatorSetContract *validatorset.Validatorset,
	delegatorVoteOperatorAddr string,
	merkleRoot string, procter common.Address, approver common.Address) {

	validatorOperatorAddr := common.HexToAddress(delegatorVoteOperatorAddr)

	gasLimit := uint64(3000000)
	gasPrice := big.NewInt(100 * 1e9)
	amt := new(big.Int).Mul(big.NewInt(100000000), big.NewInt(1e18)) // 100000000 BNB
	txOpt, err := acc.BuildTransactOpts(context.Background(), ethClient, amt, gasPrice, gasLimit)
	if err != nil {
		panic(err)
	}
	tx, err := stakeHubContract.Delegate(txOpt, validatorOperatorAddr, false)
	if err != nil {
		panic(err)
	}
	println("delegate tx:", tx.Hash().String())
	r, err := bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		panic(err)
	}
	if r.Status != 1 {
		panic("delegate tx failed")
	}
	votingPower, err := govTokenContract.GetVotes(nil, acc.Addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("votingPower:", votingPower, "acc", acc.Addr)

	govTokenBalance, err := govTokenContract.BalanceOf(nil, acc.Addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("govTokenBalance:", govTokenBalance, "acc", acc.Addr)
	txOpt, err = acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, gasPrice, gasLimit)
	if err != nil {
		panic(err)
	}
	tx, err = govTokenContract.Delegate(txOpt, acc.Addr)
	if err != nil {
		panic(err)
	}
	println("delegateVote tx:", tx.Hash().String())
	r, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		panic(err)
	}
	if r.Status != 1 {
		panic("delegateVote tx failed")
	}

	targets := []common.Address{
		common.HexToAddress(GovHubContractAddr),
		common.HexToAddress(GovHubContractAddr),
		common.HexToAddress(GovHubContractAddr)}
	values := []*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	callData := [][]byte{}

	merkleRootBytes, err := hexutil.Decode(merkleRoot)
	if err != nil {
		panic(err)
	}

	callData1, err := packGovCallData("merkleRoot", merkleRootBytes, common.HexToAddress(TokenRecoveryContractAddr))
	if err != nil {
		panic(err)
	}
	callData = append(callData, callData1)
	fmt.Println("update merkleRoot callData:", hexutil.Encode(callData1))

	callData2, err := packGovCallData("approvalAddress", approver[:], common.HexToAddress(TokenRecoveryContractAddr))
	if err != nil {
		panic(err)
	}
	callData = append(callData, callData2)
	fmt.Println("update approvalAddress callData:", hexutil.Encode(callData2))

	callData3, err := packGovCallData("assetProtector", procter[:], common.HexToAddress(TokenRecoveryContractAddr))
	if err != nil {
		panic(err)
	}
	callData = append(callData, callData3)
	fmt.Println("update assetProtector callData:", hexutil.Encode(callData3))

	detail := fmt.Sprintf("initializeTokenRecoveryContract-%d", rand.Int63())
	detailHash := crypto.Keccak256Hash([]byte(detail))

	txOpt, err = acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, gasPrice, gasLimit)
	if err != nil {
		panic(err)
	}
	tx, err = govContract.Propose(txOpt, targets, values, callData, detail)
	if err != nil {
		panic(err)
	}
	println("propose tx:", tx.Hash().String())
	r, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		panic(err)
	}
	if r.Status != 1 {
		panic("propose tx failed")
	}
	event, err := govContract.ParseProposalCreated(*r.Logs[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("proposal id:", event.ProposalId)

	// vote
	fmt.Println("wait 1 min to start voting")
	time.Sleep(70 * time.Second)
	txOpt, err = acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, gasPrice, gasLimit)
	if err != nil {
		panic(err)
	}
	voteTx, err := govContract.CastVote(txOpt, event.ProposalId, 1)
	if err != nil {
		panic(err)
	}
	println("vote tx:", voteTx.Hash().String())
	r, err = bind.WaitMined(context.Background(), ethClient, voteTx)
	if err != nil {
		panic(err)
	}
	if r.Status != 1 {
		panic("vote tx failed")
	}

	// queue
	fmt.Println("wait 2 min for voting end")
	time.Sleep(130 * time.Second)
	txOpt, err = acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, gasPrice, gasLimit)
	if err != nil {
		panic(err)
	}
	tx, err = govContract.Queue(txOpt, targets, values, callData, detailHash)
	if err != nil {
		panic(err)
	}
	println("queue tx:", tx.Hash().String())
	r, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		panic(err)
	}
	if r.Status != 1 {
		panic("queue tx failed")
	}

	// execute
	fmt.Println("wait 1 min for execute tx")
	time.Sleep(70 * time.Second)
	txOpt, err = acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, gasPrice, gasLimit)
	if err != nil {
		panic(err)
	}
	tx, err = govContract.Execute(txOpt, targets, values, callData, detailHash)
	if err != nil {
		panic(err)
	}
	println("execute tx:", tx.Hash().String())
	r, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		panic(err)
	}
	if r.Status != 1 {
		panic("execute tx failed")
	}
}

func packGovCallData(key string, value []byte, target common.Address) ([]byte, error) {
	stringTy, _ := abi.NewType("string", "string", nil)
	bytesTy, _ := abi.NewType("bytes", "bytes", nil)
	addressTy, _ := abi.NewType("address", "address", nil)
	method := abi.NewMethod("updateParam", "updateParam", abi.Function, "external", false, false,
		[]abi.Argument{
			{Name: "key", Type: stringTy, Indexed: false},
			{Name: "value", Type: bytesTy, Indexed: false},
			{Name: "target", Type: addressTy, Indexed: false},
		}, nil)

	methodID := method.ID
	arguments := abi.Arguments{
		{
			Type: stringTy,
		},
		{
			Type: bytesTy,
		},
		{
			Type: addressTy,
		},
	}
	encodedArguments, err := arguments.Pack(key, value, target)
	if err != nil {
		return nil, err
	}

	return append(methodID, encodedArguments...), nil
}

func checkProposal(contract *bscgovernor.Bscgovernor, proposalID *big.Int) error {
	proposal, err := contract.Proposals(nil, proposalID)
	if err != nil {
		return err
	}
	fmt.Printf("proposal: %+v\n", proposal)
	return nil
}

func checkTokenRecoverStatus(contract *tokenrecoverportal.Tokenrecoverportal) error {
	merkleRoot, err := contract.MerkleRoot(nil)
	if err != nil {
		return err
	}
	fmt.Println("merkleRoot:", merkleRoot)
	merkleRootAlreadyInit, err := contract.MerkleRootAlreadyInit(nil)
	if err != nil {
		return err
	}
	fmt.Println("merkleRootAlreadyInit:", merkleRootAlreadyInit)
	approvalAddress, err := contract.ApprovalAddress(nil)
	if err != nil {
		return err
	}
	fmt.Println("approvalAddress:", approvalAddress)
	assetProtectorAddress, err := contract.AssetProtector(nil)
	if err != nil {
		return err
	}
	fmt.Println("assetProtectorAddress:", assetProtectorAddress)

	return nil
}

func tokenRecover(
	tokenHubContract *tokenhub.Tokenhub,
	contract *tokenrecoverportal.Tokenrecoverportal, ethClient *ethclient.Client, acc *ExtAcc,
	symbol string, amt *big.Int, ownerPubKey []byte, ownerSignature []byte, approvalSignature []byte, merkleProof [][32]byte) error {

	gasLimit := uint64(3000000)
	txOpt, err := acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, nil, gasLimit)
	if err != nil {
		return err
	}
	var symbolBytes [32]byte
	copy(symbolBytes[:], symbol)
	tx, err := contract.Recover(txOpt, symbolBytes, amt, ownerPubKey, ownerSignature, approvalSignature, merkleProof)
	if err != nil {
		return err
	}
	println("token recover tx:", tx.Hash().String())
	r, err := bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		return err
	}
	if r.Status != 1 {
		return errors.New("tx failed")
	}

	fmt.Println("wait 1 min for token recover tx")
	time.Sleep(time.Minute)

	txOpt, err = acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, nil, gasLimit)
	if err != nil {
		return err
	}

	tokenContract := common.HexToAddress("0x0000000000000000000000000000000000000000")
	if symbol != "BNB" {
		tokenContract = common.HexToAddress(*recoverContract)
	}
	tx, err = tokenHubContract.WithdrawUnlockedToken(txOpt, tokenContract, acc.Addr)
	if err != nil {
		return err
	}
	println("WithdrawUnlockedToken tx:", tx.Hash().String())
	r, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		return err
	}
	if r.Status != 1 {
		return errors.New("WithdrawUnlockedToken tx failed")
	}

	if symbol == "BNB" {
		amt, err := ethClient.BalanceAt(context.Background(), acc.Addr, nil)
		if err != nil {
			return err
		}
		fmt.Printf("bnb balance of [%s]: %s\n", acc.Addr.Hex(), amt)
	} else {
		bep20Contract, err := bep20.NewBep20(tokenContract, ethClient)
		if err != nil {
			return err
		}

		amt, err := bep20Contract.BalanceOf(nil, acc.Addr)
		if err != nil {
			return err
		}

		fmt.Printf("bep20 [%s] balance of [%s]: %s\n", symbol, acc.Addr.Hex(), amt)
	}

	return nil
}

func getSyncFeeFromContract(contract *tokenmanager.Tokenmanager) {
	fee, err := contract.SyncFee(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sync fee : %v\n", fee)
	mirrorFee, err := contract.MirrorFee(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("mirror fee : %v\n", mirrorFee)
}

func signTransaction(fromEO ExtAcc, value *big.Int, data []byte, nonce uint64, gasLimit uint64, gasPrice, chainId *big.Int) (*ethtypes.Transaction, common.Hash, error) {
	tx := ethtypes.NewContractCreation(nonce, value, gasLimit, gasPrice, data)
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainId), fromEO.Key)
	if err != nil {
		return nil, common.Hash{}, err
	}

	return signedTx, signedTx.Hash(), nil
}

func deployContract(rpcClient *ethclient.Client, account ExtAcc, contractCode []byte) (common.Hash, common.Address, error) {
	fromAddress := account.Addr
	nonce, err := rpcClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.Hash{}, common.Address{}, err
	}

	gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
	if err != nil {
		return common.Hash{}, common.Address{}, err
	}

	chainID, err := rpcClient.NetworkID(context.Background())
	if err != nil {
		return common.Hash{}, common.Address{}, err
	}

	gasLimit := uint64(3000000)
	signedTx, hash, err := signTransaction(account, big.NewInt(0),
		contractCode, nonce, gasLimit, gasPrice, chainID)
	if err != nil {
		return common.Hash{}, common.Address{}, err
	}

	err = rpcClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, common.Address{}, err
	}

	found := false
	var (
		r1 *ethtypes.Receipt
	)
	for i := 0; i < 60; i++ {
		r1, err = rpcClient.TransactionReceipt(context.Background(), hash)
		if r1 != nil && err == nil {
			found = true
			break
		}
		time.Sleep(3 * time.Second)
	}
	if !found {
		return common.Hash{}, common.Address{}, errors.New("tx failed to be verified on chain or timeout")
	}
	if r1.Status != ethtypes.ReceiptStatusSuccessful {
		return common.Hash{}, common.Address{}, errors.New("tx failed to be executed on chain")
	}
	contractAddr := r1.ContractAddress

	codes, err := rpcClient.CodeAt(context.Background(), contractAddr, nil)
	if err != nil {
		return common.Hash{}, common.Address{}, err
	}

	if len(codes) == 0 {
		return common.Hash{}, common.Address{}, errors.New("contract code is not empty")
	}

	return signedTx.Hash(), contractAddr, nil
}

func approveBind(rpcClient *ethclient.Client, account ExtAcc, contract common.Address,
	tmContract *tokenmanager.Tokenmanager, symbol string, amt *big.Int) error {
	bep20Contract, err := bep20.NewBep20(contract, rpcClient)
	if err != nil {
		return err
	}
	txOpt, err := account.BuildTransactOpts(context.Background(), rpcClient, big.NewInt(0),
		nil, 3000000)
	if err != nil {
		return err
	}
	approveTx, err := bep20Contract.Approve(txOpt, common.HexToAddress(TokenManagerContractAddr), amt)
	if err != nil {
		return err
	}
	fmt.Println("approve bep20 tx hash:", approveTx.Hash())
	approveReceipt, err := bind.WaitMined(context.Background(), rpcClient, approveTx)
	if err != nil {
		return err
	}
	if approveReceipt.Status != 1 {
		return fmt.Errorf("bep20 approve tx execution failed")
	}

	txOpt, err = account.BuildTransactOpts(context.Background(), rpcClient, big.NewInt(2100000000000000),
		nil, 3000000)
	if err != nil {
		return err
	}
	approvalBindTx, err := tmContract.ApproveBind(txOpt, contract, symbol)
	fmt.Println("approve bind tx hash:", approvalBindTx.Hash())
	if err != nil {
		return err
	}
	approvalBindReceipt, err := bind.WaitMined(context.Background(), rpcClient, approvalBindTx)
	if err != nil {
		return err
	}
	if approvalBindReceipt.Status != 1 {
		return fmt.Errorf("approve bind tx execution failed")
	}

	return nil
}

func syncToken(rpcClient *ethclient.Client, account ExtAcc, contract common.Address,
	tmContract *tokenmanager.Tokenmanager) error {
	relayFee := big.NewInt(2100000000000000)
	syncFee := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10))

	txOpt, err := account.BuildTransactOpts(context.Background(), rpcClient, new(big.Int).Add(relayFee, syncFee),
		nil, 3000000)
	if err != nil {
		return err
	}
	tx, err := tmContract.Sync(txOpt, contract, uint64(time.Now().Add(60*time.Minute).Unix()))
	if err != nil {
		return err
	}

	fmt.Println("sync bep20 token tx hash:", tx.Hash())
	approveReceipt, err := bind.WaitMined(context.Background(), rpcClient, tx)
	if err != nil {
		return err
	}
	if approveReceipt.Status != 1 {
		return fmt.Errorf("sync bep20 token tx execution failed")
	}

	return nil
}

func mirrorToken(rpcClient *ethclient.Client, account ExtAcc, contract common.Address,
	tmContract *tokenmanager.Tokenmanager, tokenHubContract *tokenhub.Tokenhub) error {

	relayFee := big.NewInt(2100000000000000)
	mirrorFee := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10))
	txOpt, err := account.BuildTransactOpts(context.Background(), rpcClient, new(big.Int).Add(relayFee, mirrorFee),
		nil, 3000000)
	if err != nil {
		return err
	}

	mirrorTx, err := tmContract.Mirror(txOpt, contract, uint64(time.Now().Add(1000*time.Second).Unix()))
	fmt.Println("mirror tx hash:", mirrorTx.Hash())
	if err != nil {
		return err
	}
	mirrorReceipt, err := bind.WaitMined(context.Background(), rpcClient, mirrorTx)
	if err != nil {
		return err
	}
	if mirrorReceipt.Status != 1 {
		return fmt.Errorf("mirror tx execution failed")
	}

	return nil
}

func claimToken(tokenHubContract *tokenhub.Tokenhub,
	ethClient *ethclient.Client, acc *ExtAcc) error {
	gasLimit := uint64(3000000)
	tokenContract := common.HexToAddress("0x0000000000000000000000000000000000000000")
	txOpt, err := acc.BuildTransactOpts(context.Background(), ethClient, common.Big0, nil, gasLimit)
	if err != nil {
		return err
	}
	info, err := tokenHubContract.LockInfoMap(nil, tokenContract, acc.Addr)
	if err != nil {
		return err
	}
	fmt.Printf("lock info: %+v\n", info)
	tx, err := tokenHubContract.WithdrawUnlockedToken(txOpt, tokenContract, acc.Addr)
	if err != nil {
		return err
	}
	println("WithdrawUnlockedToken tx:", tx.Hash().String())
	r, err := bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		return err
	}
	if r.Status != 1 {
		return errors.New("WithdrawUnlockedToken tx failed")
	}

	amt, err := ethClient.BalanceAt(context.Background(), acc.Addr, nil)
	if err != nil {
		return err
	}
	fmt.Printf("bnb balance of [%s]: %s\n", acc.Addr.Hex(), amt)

	return nil
}
