package service

import (
	"context"
	"demo1/contracts"
	"demo1/model"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Erc20LogTransfer(high int64) ([]model.LogTransfer, error) {
	logTransfers := make([]model.LogTransfer, 0)
	block, err := ethClient.BlockByNumber(context.Background(), big.NewInt(high))
	if err != nil {
		return logTransfers, fmt.Errorf("get block err %w", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(contracts.Erc20MetaData.ABI))
	if err != nil {
		return logTransfers, fmt.Errorf("get erc20 abi err %w", err)
	}
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for _, transaction := range block.Transactions() {
		receipt, err := ethClient.TransactionReceipt(context.Background(), transaction.Hash())
		if err != nil {
			return logTransfers, fmt.Errorf("get receipt err %w", err)
		}
		for _, log := range receipt.Logs {
			switch log.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				var transferEvent model.LogTransfer
				err = contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", log.Data)
				if err != nil {
					//has nil amount
				}
				transferEvent.From = common.HexToAddress(log.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(log.Topics[2].Hex())
				logTransfers = append(logTransfers, transferEvent)
			}
		}
	}

	return logTransfers, nil
}
