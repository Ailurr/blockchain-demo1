package service

import (
	"context"
	"demo1/contracts"
	"demo1/model"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/panjf2000/ants/v2"
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
	wg := sync.WaitGroup{}
	pool, err := ants.NewPool(24)
	if err != nil {
		return []model.LogTransfer{}, fmt.Errorf("ants pool err %w", err)
	}
	fmt.Printf("transacations len: %d\n", len(block.Transactions()))
	for _, transaction := range block.Transactions() {
		transaction := transaction
		wg.Add(1)
		err := pool.Submit(func() {
			defer wg.Done()
			receipt, err := ethClient.TransactionReceipt(context.Background(), transaction.Hash())
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
		})
		if err != nil {
			return nil, err
		}
	}
	wg.Wait()
	return logTransfers, nil
}
