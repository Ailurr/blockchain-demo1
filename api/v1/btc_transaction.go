package v1

import (
	"demo1/service"
	"demo1/utils"
	"github.com/gin-gonic/gin"
)

func (a *ApiV1) GetTransactionByHash(c *gin.Context) {
	hash := c.Param("hash")

	parsedTrans, err := service.GetBlockTransInfo(hash)

	if err != nil {
		utils.FailWithMsg(c, "parse err")
	}
	utils.OkWithData(c, parsedTrans)
}
