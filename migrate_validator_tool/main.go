package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	abi "github.com/bnb-chain/node-deploy/migrate_validator_tool/abi"
)

const (
	StakeHubContractAddr = "0x0000000000000000000000000000000000002002"
)

var (
	secretFile = flag.String("secret", "", "secret file path")
	password   = flag.String("password", "", "password")

	transferAmount = flag.String("amount", "0", "transfer amount")
	to             = flag.String("to", "", "transfer to address")

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

	if *transferAmount != "0" {
		from, err := FromHexKey(*bscPrivKey)
		if err != nil {
			panic(err)
		}
		toAddr := common.HexToAddress(*to)
		amount, _ := new(big.Int).SetString(*transferAmount, 10)
		err = nativeTransfer(ethClient, &from, toAddr, amount)
		if err != nil {
			panic(err)
		}
		return
	}

	stakeHubContract, err := abi.NewStakehub(common.HexToAddress(StakeHubContractAddr), ethClient)
	if err != nil {
		panic(err)
	}

	acc, err := FromHexKey(*bscPrivKey)
	if err != nil {
		panic(err)
	}
	delegation, _ := new(big.Int).SetString(*delegation, 10)
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
		voteAddrBytes, voteBLSProofBytes, abi.StakeHubCommission{
			Rate:          uint64(*commissionRate),
			MaxRate:       uint64(*commissionMaxRate),
			MaxChangeRate: uint64(*commissionMaxChangeRate),
		},
		abi.StakeHubDescription{
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

func nativeTransfer(rpcClient *ethclient.Client, from *ExtAcc, to common.Address, amount *big.Int) error {
	fromAddress := from.Addr
	nonce, err := rpcClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
	if err != nil {
		return err
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
