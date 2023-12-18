package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/moshushu/gin-blog/models"
	"github.com/moshushu/gin-blog/pkg/logging"
	"github.com/moshushu/gin-blog/pkg/setting"
	"github.com/moshushu/gin-blog/routers"
)

func main() {
	// 1、endless 方案
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("Actul pid is %d", syscall.Getpid())
	}

	log.Println("Listen Server Port", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	// 2、http.Server - Shutdown() 方案
	// router := routers.InitRouter()
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// go func() {
	// 	if err := s.ListenAndServe(); err != nil {
	// 		log.Printf("Listen: %s\n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// if err := s.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }

	// log.Println("Server exiting")
}
