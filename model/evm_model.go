package model

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type LogTransfer struct {
	From   common.Address `json:"from"`
	To     common.Address `json:"to"`
	Amount *big.Int       `json:"amount"`
}

type Erc20TransferArgs struct {
	PrivateKey string `json:"private_key"`
	ToAddress  string `json:"to_address"`
	Amount     string `json:"amount"`
}
