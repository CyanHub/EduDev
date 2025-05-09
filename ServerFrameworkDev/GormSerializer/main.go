package main

import (
	"fmt"
	_ "net/http/pprof"
	"time"

	// "ServerFramework/global"
	// "ServerFramework/initialize"

	"github.com/robfig/cron/v3"
)

// main 是程序的入口函数，负责初始化项目的各项配置和服务。
func main() {
	// 调用 initialize 包中的 MustConfig 函数，确保项目配置加载成功。
	// 若配置加载失败，该函数会触发 panic 终止程序。
	// initialize.MustConfig()

	// 调用 initialize 包中的 MustLoadZap 函数，确保 Zap 日志库初始化成功。
	// 若初始化失败，该函数会触发 panic 终止程序。
	// initialize.MustLoadZap()

	// 调用 initialize 包中的 MustInitDB 函数，确保数据库初始化成功。
	// 若数据库初始化失败，该函数会触发 panic 终止程序。
	// initialize.MustInitDB()

	// 调用 initialize 包中的 AutoMigrate 函数，根据模型定义自动迁移数据库表结构。
	// 传入 global 包中的 DB 数据库连接实例。
	// initialize.AutoMigrate(global.DB)

	// 注释掉的代码，调用 initialize 包中的 MustInitRedis 函数，用于确保 Redis 初始化成功。
	// 若 Redis 初始化失败，该函数会触发 panic 终止程序。
	//initialize.MustInitRedis()

	// 调用 initialize 包中的 MustRunWindowServer 函数，确保 Windows 服务器启动成功。
	// 若服务器启动失败，该函数会触发 panic 终止程序。
	// initialize.MustRunWindowServer()
	// TickerUse()
	CronUse()
}

func PrintTime() error {
	fmt.Println("当前时间是：", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func CronUse() {
	c := cron.New(cron.WithSeconds())
	c.Start()

	_, err := c.AddFunc("*/2 * * * * *", func() {
		PrintTime()
	})
	if err != nil {
		panic(fmt.Sprintf("添加定时任务失败: %v", err))
	}
	time.Sleep(10 * time.Second)
}

// func PrintTime() /*error*/ {
// 	fmt.Println("当前时间是：", time.Now().Format("2006-01-02 15:04:05"))
// 	// return nil
// }

// func CronUse() {
// 	c := cron.New(cron.WithSeconds())
// 	c.Start()

// 	_, err := c.AddFunc("*/2 * * * * *", PrintTime)
// 	if err != nil {
// 		panic(fmt.Sprintf("添加定时任务失败: %v", err))
// 	}
// 	time.Sleep(10 * time.Second)
// }

func TickerUse() {
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("当前时间是：", time.Now().Format("2006-01-02 15:04:05"))
				break
			case <-quit:
				fmt.Println("定时任务退出")

				break
			}
		}
	}()
	time.Sleep(10 * time.Second)
	quit <- struct{}{}
	time.Sleep(5 * time.Second)
}
