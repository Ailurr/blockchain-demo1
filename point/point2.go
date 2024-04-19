package point

import (
	"context"
	"demo1/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func Point2(privatekey string, toAdr string, value *big.Int) {
	fmt.Println("-------------------------Point 2-----------------------------")

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
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var data []byte
	//tx := types.NewTx(&types.DynamicFeeTx{
	//	ChainID:   chainID,
	//	Nonce:     nonce,
	//	GasFeeCap: feeCap,
	//	GasTipCap: tip,
	//	Gas:       gasLimit,
	//	To:        &toAddress,
	//	Value:     value,
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
		fmt.Println("bind.NewKeyedTransactorWithChainID error ,", err)
		return
	}
	gasTipCap, _ := ethClient.SuggestGasTipCap(context.Background())
	opts.GasFeeCap = big.NewInt(108694000460)
	opts.GasLimit = uint64(100000)
	opts.GasTipCap = gasTipCap

	amount, _ := new(big.Int).SetString("1000000000", 10)
	usdt, _ := contracts.NewUsdt(common.HexToAddress("0xaA8E23Fb1079EA71e0a56F48a2aA51851D8433D0"), ethClient)
	tx, err := usdt.Transfer(opts, toAddress, amount)
	if err != nil {
		fmt.Println("token.Transfer error ,", err)
		return
	}

	fmt.Println("transfer tx : ", tx.Hash().Hex())
}
