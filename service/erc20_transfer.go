package service

import (
	"context"
	"demo1/contracts"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Erc20Transfer(privatekey string, toAdr string, value *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		log.Fatal(err)
	}

	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}

	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//gasLimit := uint64(21000)         // in units
	//tip := big.NewInt(2000000000)     // maxPriorityFeePerGas = 2 Gwei
	//feeCap := big.NewInt(20000000000) // maxFeePerGas = 20 Gwei

	toAddress := common.HexToAddress(toAdr)

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("fail to get chain id: %w", err)
	}
	//var data []byte
	//tx := types.NewTx(&types.DynamicFeeTx{
	//	ChainID:   chainID,
	//	Nonce:     nonce,
	//	GasFeeCap: feeCap,
	//	GasTipCap: tip,
	//	Gas:       gasLimit,
	//	To:        &toAddress,
	//	Amount:     value,
	//	Data:      data,
	//})
	//
	//signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = ethClient.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Transaction hash: %s", signedTx.Hash().Hex())
	//构建参数对象
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("fail to create transfer opts: %w", err)
	}
	gasTipCap, err := ethClient.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", fmt.Errorf("fail to get gasTipCap: %w", err)
	}
	opts.GasFeeCap = big.NewInt(108694000460)
	opts.GasLimit = uint64(100000)
	opts.GasTipCap = gasTipCap

	amount, _ := new(big.Int).SetString("1000000000", 10)
	usdt, _ := contracts.NewErc20(common.HexToAddress("0xaA8E23Fb1079EA71e0a56F48a2aA51851D8433D0"), ethClient)
	tx, err := usdt.Transfer(opts, toAddress, amount)
	if err != nil {
		return "", fmt.Errorf("fail to transfer: %w", err)
	}

	return tx.Hash().Hex(), nil
}
