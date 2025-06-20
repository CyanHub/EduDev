package main

import (
	"FileSystem/global"
	"FileSystem/initialize"
	"FileSystem/router"
	"context"
	"log"

	"go.uber.org/zap"
)

func main() {
	// 初始化全局上下文
	global.Context = context.Background()

	// 1. 初始化配置
	initialize.MustConfig()

	// 2. 初始化日志
	initialize.InitLogger()
	defer global.Logger.Sync()

	// 3. 初始化数据库
	initialize.MustInitDB()
	if err := initialize.AutoMigrate(global.DB); err != nil {
		global.Logger.Fatal("数据库表结构迁移失败", zap.Error(err))
	}

	// 4. 初始化 Redis
	initialize.MustInitRedis()

	// 5. 初始化默认管理员用户
	if err := initialize.InitAdmin(); err != nil {
		log.Fatalf("初始化管理员用户失败: %v", err)
	}

	// 6. 初始化路由
	r := router.InitRouter()

	// 7. 启动 Gin 服务器
	initialize.MustRunWindowServer(r)
}
