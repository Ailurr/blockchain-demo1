package point

import (
	"context"
	"demo1/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
)

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

func Point3(high int64) {
	fmt.Println("-------------------------Point 3-----------------------------")
	block, _ := ethClient.BlockByNumber(context.Background(), big.NewInt(high))
	contractAbi, _ := abi.JSON(strings.NewReader(string(contracts.Erc20ABI)))
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	for _, transaction := range block.Transactions() {
		receipt, _ := ethClient.TransactionReceipt(context.Background(), transaction.Hash())
		for _, log := range receipt.Logs {
			switch log.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				var transferEvent LogTransfer
				_ = contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", log.Data)
				transferEvent.From = common.HexToAddress(log.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(log.Topics[2].Hex())
				fmt.Printf("Transfer: from:%s to:%s value:%s \n", transferEvent.From.Hex(), transferEvent.To.Hex(), transferEvent.Value)

			}
		}
	}
}
