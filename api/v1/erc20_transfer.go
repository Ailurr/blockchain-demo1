package v1

import (
	"demo1/model"
	"demo1/service"
	"demo1/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

	key := "ParseTransferByBlockHigh-" + highS
	cache, err := utils.GetCache(key)
	if err != nil {
		if err == redis.Nil {
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
			value, err := json.Marshal(logTransfers)
			utils.SetCache(key, utils.Bytes2String(value))
			utils.OkWithData(c, logTransfers)
		} else {
			utils.FailWithMsg(c, "inner err")
		}
	} else {
		res := make([]model.LogTransfer, 0)
		json.Unmarshal(cache, &res)
		utils.OkWithData(c, res)
	}
}
