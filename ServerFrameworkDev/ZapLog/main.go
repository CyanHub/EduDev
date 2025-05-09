package main

import (
	_ "net/http/pprof"

	"github.com/CyanHub/EduDev/global"
	"github.com/CyanHub/EduDev/initialize"
)

// main 是程序的入口函数，负责初始化系统的各项配置和服务，然后启动服务。
func main() {
	// 调用 initialize 包中的 MustConfig 函数，确保系统配置正确加载。
	// 若配置加载失败，该函数会终止程序运行。
	initialize.MustConfig()

	// 调用 initialize 包中的 MustLoadZap 函数，确保 Zap 日志系统正确初始化。
	// 若初始化失败，该函数会终止程序运行。
	initialize.MustLoadZap()

	// 调用 initialize 包中的 MustInitDB 函数，确保数据库连接正确初始化。
	// 若数据库连接初始化失败，该函数会终止程序运行。
	initialize.MustInitDB()

	// 调用 initialize 包中的 AutoMigrate 函数，根据模型定义自动迁移数据库表结构。
	// global.DB 是全局数据库连接对象。
	initialize.AutoMigrate(global.DB)

	// 注释掉的代码，若需要初始化 Redis 服务，可取消注释。
	// 调用 initialize 包中的 MustInitRedis 函数，确保 Redis 连接正确初始化。
	// 若 Redis 连接初始化失败，该函数会终止程序运行。
	//initialize.MustInitRedis()

	// 调用 initialize 包中的 MustRunWindowServer 函数，确保 Windows 服务正确启动。
	// 若服务启动失败，该函数会终止程序运行。
	initialize.MustRunWindowServer()
}
