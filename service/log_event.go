package service

import (
	"context"
	"demo1/contracts"
	"demo1/model"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
	"sync"
)

func Erc20LogTransfer(high int64) ([]model.LogTransfer, error) {

	block, err := ethClient.BlockByNumber(context.Background(), big.NewInt(high))
	if err != nil {
		return []model.LogTransfer{}, fmt.Errorf("get block err %w", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(contracts.Erc20MetaData.ABI))
	if err != nil {
		return []model.LogTransfer{}, fmt.Errorf("get erc20 abi err %w", err)
	}
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	logTransfers := make([]model.LogTransfer, 0)
	var lock sync.Mutex
	asyctl := make(chan struct{}, 12)
	for i := 0; i < 12; i++ {
		asyctl <- struct{}{}
	}
	wg := sync.WaitGroup{}
	fmt.Printf("transacations len: %d\n", len(block.Transactions()))
	for _, transaction := range block.Transactions() {
		transaction := transaction
		wg.Add(1)
		<-asyctl
		go func() {
			receipt, err := ethClient.TransactionReceipt(context.Background(), transaction.Hash())
			//fmt.Printf("receipt.Logs len: %d\n", len(receipt.Logs))
			if err != nil {
				fmt.Println("get receipt err %w\n", err)
			}
			for _, log := range receipt.Logs {
				switch log.Topics[0].Hex() {
				case logTransferSigHash.Hex():
					var transferEvent model.LogTransfer
					err = contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", log.Data)
					if err != nil {
						//println(err.Error())
					}
					transferEvent.From = common.HexToAddress(log.Topics[1].Hex())
					transferEvent.To = common.HexToAddress(log.Topics[2].Hex())
					lock.Lock()
					logTransfers = append(logTransfers, transferEvent)
					lock.Unlock()
				}
			}
			asyctl <- struct{}{}
			wg.Done()
		}()
	}
	wg.Wait()
	return logTransfers, nil
}
