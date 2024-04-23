package main

import (
	"demo1/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initialize.InitRouter(r)
	initialize.InitRedis()

	r.Run(":23543")

}
