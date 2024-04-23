package main

import (
	"demo1/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routerGroupV1 := r.Group("/v1")
	{
		apiV1 := &v1.ApiV1{}
		blockGroup := routerGroupV1.Group("block")
		{
			blockGroup.GET("hash/:hash/transfer", apiV1.GetTransactionByHash)

			blockGroup.GET("high/:high/event/transfer", apiV1.ParseTransferByBlockHigh)
		}
		transGroup := routerGroupV1.Group("transaction")
		{
			transGroup.POST("/erc20/transfer", apiV1.Erc20Transfer)
		}

		r.Run(":23543")
	}
}
