package v1

import (
	"demo1/model"
	"demo1/service"
	"demo1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
)

func (a *ApiV1) Erc20Transfer(c *gin.Context) {
	var erc20Transfer model.Erc20TransferArgs
	err := c.ShouldBindJSON(&erc20Transfer)

	amount, err := strconv.ParseInt(erc20Transfer.Amount, 10, 64)
	amountBig := big.NewInt(amount)

	if err != nil {
		utils.FailWithMsg(c, "args err")
		return
	}

	txHash, err := service.Erc20Transfer(erc20Transfer.PrivateKey, erc20Transfer.ToAddress, amountBig)
	if err != nil {
		utils.FailWithMsg(c, "inner err")
		return
	}
	utils.OkWithData(c, gin.H{"txHash": txHash})
}

func (a *ApiV1) ParseTransferByBlockHigh(c *gin.Context) {
	highS := c.Param("high")
	highI, err := strconv.ParseInt(highS, 10, 64)
	if err != nil {
		utils.FailWithMsg(c, "args err")
		return
	}
	logTransfers, err := service.Erc20LogTransfer(highI)
	if err != nil {
		fmt.Printf("%+v\n", err)
		utils.FailWithMsg(c, "inner err")
		return
	}
	utils.OkWithData(c, logTransfers)
}
