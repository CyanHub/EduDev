package main

import (
	_ "net/http/pprof"

	"ServerFramework/global"
	"ServerFramework/initialize"
	"ServerFramework/test"
)

// main 是程序的入口函数，负责初始化项目的各项配置和服务。
func main() {
	initialize.MustConfig()           // 初始化配置
	initialize.MustInitDB()           // 初始化数据库
	initialize.RegisterSerializer()   // 注册序列化器
	initialize.AutoMigrate(global.DB) // 自动迁移数据库结构, 确保数据库表结构与模型匹配
	// initialize.MustLoadGorm() // 初始化 GORM 数据库连接
	// initialize.MustInitRedis() // 初始化 Redis 缓存
	// initialize.MustRunWindowServer() // 初始化窗口服务，用于处理窗口相关的操作
	// test.TickerUse() // 定时器使用示例
	// test.CronUse() // 定时任务使用示例
	// test.ExampleJSONSerializer() // JSON 序列化器使用示例
	test.ExampleGobSerializer() // Gob 序列化器使用示例
}
