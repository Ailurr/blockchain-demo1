package main

import (
	"context"
	"demo1/global"
	"demo1/initialize"
	"demo1/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/ants/v2"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer ants.Release()

	r := gin.Default()
	r.Use(middleware.LimiterMiddleWare())
	initialize.InitRouter(r)
	initialize.InitRedis()

	s := &http.Server{
		Addr:    ":23543",
		Handler: r,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen err: %s\n", err)
		}
	}()

	//pprof webserver
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	//graceful exit
	c := make(chan os.Signal)
	exitCh := make(chan struct{})
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		sig := <-c
		log.Printf("catch signal: %+v", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		err := global.Rdc.Close()
		if err != nil {
			fmt.Println("close redis fail")
		}
		exitCh <- struct{}{}
	}()
	<-exitCh
	log.Printf("server exitCh")
}
