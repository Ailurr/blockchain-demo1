package main

import (
	"demo1/initialize"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	r := gin.Default()
	initialize.InitRouter(r)
	initialize.InitRedis()

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	r.Run(":23543")

}
