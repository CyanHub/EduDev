package main

import (
	"FileSystem/global"
	"FileSystem/initialize"
	"FileSystem/router"
	"os"

	"go.uber.org/zap"
)

func main() {
	// 1. 初始化配置
	initialize.MustConfig()

	// 2. 初始化日志
	initialize.MustLoadZap()
	defer global.Logger.Sync()

	// 3. 初始化数据库
	initialize.MustInitDB()
	initialize.MustLoadGorm()
	initialize.AutoMigrate(global.DB)

	// 4. 初始化Redis
	initialize.MustInitRedis()

	// 5. 初始化Casbin权限管理
	initialize.MustCasbin()

	// 6. 启动Gin服务器
	initialize.MustRunWindowServer()

	// 7. 初始化路由
	// 启动前确保创建pages目录
	if err := os.MkdirAll("pages", 0755); err != nil {
		global.Logger.Fatal("创建静态文件目录失败", zap.Error(err))
	}
	
	
	// 启动服务
	r := router.InitRouter()
	if err := r.Run(":9090"); err != nil {
		global.Logger.Fatal("服务启动失败", zap.Error(err))
	}
}
