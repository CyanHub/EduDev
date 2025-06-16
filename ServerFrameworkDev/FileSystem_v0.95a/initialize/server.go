// -- Active: 1735040439559@@127.0.0.1@3306@file_system
package initialize

import (
	"FileSystem/global"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// MustRunWindowServer 启动服务
// 启动服务
// @Description: 启动服务
func MustRunWindowServer(engine *gin.Engine) {
	global.Cron = cron.New()
	global.Cron.Start()

	address := fmt.Sprintf(":%d", global.CONFIG.App.Port) // 设置服务器监听的端口

	global.Cron = cron.New(cron.WithSeconds()) // 支持秒级别的定时任务
	// global.Cron.Start() // 启动定时任务
	// task.AddClerOperationRecordTask(global.Cron)

	go func() {
		pprofAddress := ":6060" // 或者其他你想要的端口
		fmt.Println("启动 pprof 服务，监听端口：", pprofAddress)
		if err := http.ListenAndServe(pprofAddress, nil); err != nil {
			fmt.Println("pprof 服务启动失败:", err)
		}
	}()

	fmt.Println("启动服务器，监听端口：", address)
	if err := engine.Run(address); err != nil {
		panic(err)
	}
}
