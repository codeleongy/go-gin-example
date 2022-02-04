package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/leong-y/go-gin-example/models"
	"github.com/leong-y/go-gin-example/pkg/gredis"
	"github.com/leong-y/go-gin-example/pkg/logging"
	"github.com/leong-y/go-gin-example/pkg/setting"
	"github.com/leong-y/go-gin-example/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	gredis.Setup()
	logging.Setup()
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	// 1<<20也就是1*2^20=1MB
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	// endless.NewServer 返回一个初始化的 endlessServer 对象
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		/*
			在 BeforeBegin 时输出当前进程的 pid，调用 ListenAndServe 将实际“启动”服务
		*/
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
