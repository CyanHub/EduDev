package main

import (
	"ServerFramework/utils"
	"fmt"
	"time"
)

func main() {
	pool, err := utils.NewPool(5)
	if err != nil {
		panic(err)
	}

	//协程结束便关闭协程池 
	defer pool.Close()

	// 任务提交
	for i := 1; i <= 20; i++ {
		taskID := i
		pool.Submit(func() int {
			fmt.Println("任务", taskID, "开始执行")
			time.Sleep(time.Second)
			fmt.Println("任务", taskID, "执行完成")
			return taskID
		})
	}


	time.Sleep(10 * time.Second) // 休眠十秒等待协程执行完毕

	// initialize.MustInitDB()           // 初始化数据库
	// initialize.MustConfig()           // 初始化配置文件
	// initialize.MustInitRedis()        // 初始化Redis
	// initialize.MustLoadZap()          // 初始化日志
	// initialize.RegisterSerializer()   // 注册序列化器
	// initialize.MustLoadGorm()         // 初始化数据库
	// initialize.AutoMigrate(global.DB) // 自动迁移数据库
	// initialize.MustCasbin()           // 初始化Casbin
	// initialize.MustRunWindowServer()  // 初始化窗口服务器
}
