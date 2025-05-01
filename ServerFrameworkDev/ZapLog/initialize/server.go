package initialize

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/CyanHub/EduDev/global"
	"github.com/CyanHub/EduDev/initialize/task"
	"github.com/CyanHub/EduDev/router"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func MustRunWindowServer() {
	global.Cron = cron.New()
	global.Cron.Start()
	task.ClearOperationRecord("0/1 * * * * *")

	engine := gin.Default()
	userGroup := router.UserGroup{}
	userGroup.InitUserRouters(engine)

	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	address := fmt.Sprintf(":%d", global.CONFIG.App.Port)
	fmt.Println("启动服务器，监听端口：", address)
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
