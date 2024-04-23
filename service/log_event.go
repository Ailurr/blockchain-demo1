package service

import (
	"context"
	"demo1/contracts"
	"demo1/model"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Erc20LogTransfer(high int64) []model.LogTransfer {
	logTransfers := make([]model.LogTransfer, 0)
	block, _ := ethClient.BlockByNumber(context.Background(), big.NewInt(high))
	contractAbi, _ := abi.JSON(strings.NewReader(string(contracts.Erc20MetaData.ABI)))

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for _, transaction := range block.Transactions() {
		receipt, _ := ethClient.TransactionReceipt(context.Background(), transaction.Hash())
		for _, log := range receipt.Logs {
			switch log.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				var transferEvent model.LogTransfer
				_ = contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", log.Data)
				transferEvent.From = common.HexToAddress(log.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(log.Topics[2].Hex())
				logTransfers = append(logTransfers, transferEvent)
			}
		}
	}

	return logTransfers
}