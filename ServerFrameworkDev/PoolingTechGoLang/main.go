package main

import (
	"ServerFramework/global"
	"ServerFramework/initialize"
)

func main() {
	initialize.MustConfig()           // 初始化配置文件
	initialize.MustInitRedis()        // 初始化Redis
	initialize.MustLoadZap()          // 初始化日志
	initialize.RegisterSerializer()   // 注册序列化器
	initialize.MustLoadGorm()         // 初始化数据库
	initialize.AutoMigrate(global.DB) // 自动迁移数据库
	initialize.MustCasbin()           // 初始化Casbin
	initialize.MustRunWindowServer()  // 初始化窗口服务器
}

