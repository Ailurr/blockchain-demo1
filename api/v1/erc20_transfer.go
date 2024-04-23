package v1

import (
	"demo1/model"
	"demo1/service"
	"demo1/utils"
	"github.com/gin-gonic/gin"
	"log"
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

	txHash := service.Erc20Transfer(erc20Transfer.PrivateKey, erc20Transfer.ToAddress, amountBig)
	utils.OkWithData(c, gin.H{"txHash": txHash})
}

func (a *ApiV1) ParseTransferByBlockHigh(c *gin.Context) {
	highS := c.Param("high")
	highI, err := strconv.ParseInt(highS, 10, 64)
	log.Println(highS, highI)
	if err != nil {
		utils.FailWithMsg(c, "args err")
		return
	}
	logTransfers := service.Erc20LogTransfer(highI)
	utils.OkWithData(c, logTransfers)
}