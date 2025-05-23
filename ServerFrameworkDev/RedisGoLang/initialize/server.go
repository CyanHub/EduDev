package initialize

import (
	"fmt"
	"net/http"
	"runtime"

	"ServerFramework/global"
	"ServerFramework/initialize/task"
	"ServerFramework/router"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// MustRunWindowServer 启动服务
// 启动服务
// @Description: 启动服务
func MustRunWindowServer() {
	global.Cron = cron.New()
	global.Cron.Start()

	engine := gin.Default()
	userGroup := router.UserGroup{}
	userGroup.InitUserRouters(engine)

	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	address := fmt.Sprintf(":%d", global.CONFIG.App.Port)
	fmt.Println("启动服务器，监听端口：", address)
	global.Cron = cron.New(cron.WithSeconds())
	global.Cron.Start()
	task.AddClerOperationRecord(global.Cron)

	go func() {
		pprofAddress := ":6060" // 或者其他你想要的端口
		fmt.Println("启动 pprof 服务，监听端口：", pprofAddress)
		if err := http.ListenAndServe(pprofAddress, nil); err != nil {
			fmt.Println("pprof 服务启动失败:", err)
		}
	}()
	if err := engine.Run(address); err != nil {
		panic(err)
	}

}
