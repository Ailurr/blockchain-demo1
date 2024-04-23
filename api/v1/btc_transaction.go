package v1

import (
	"demo1/model"
	"demo1/service"
	"demo1/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func (a *ApiV1) GetTransactionByHash(c *gin.Context) {
	hash := c.Param("hash")
	key := "GetTransactionByHash-" + hash
	cache, err := utils.GetCache(key)
	if err != nil {
		if err == redis.Nil {
			parsedTrans, err := service.GetBlockTransInfo(hash)
			if err != nil {
				utils.FailWithMsg(c, "parse err")
			}
			value, err := json.Marshal(parsedTrans)
			utils.SetCache(key, utils.Bytes2String(value))
			utils.OkWithData(c, parsedTrans)
		} else {
			utils.FailWithMsg(c, "inner err")
		}
	} else {
		res := make([]model.ParsedTrans, 0)
		json.Unmarshal(cache, &res)
		utils.OkWithData(c, res)
	}
}
