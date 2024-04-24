package main

import (
	"demo1/initialize"
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/ants/v2"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	defer ants.Release()

	r := gin.Default()
	initialize.InitRouter(r)
	initialize.InitRedis()

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	r.Run(":23543")

}
