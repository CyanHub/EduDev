package main

import (
	"ServerFramework/api"
	"ServerFramework/global"
	"ServerFramework/initialize"
	"ServerFramework/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    // 初始化配置文件
    initialize.MustConfig()

    // 初始化日志
    initialize.MustLoadZap()

    // 初始化数据库连接
    initialize.MustLoadGorm()

    // 初始化 Redis 连接
    initialize.MustInitRedis()

    // 创建 Gin 引擎
    engine := gin.Default()

    // 添加静态文件服务
    engine.Static("/static", "./static")

    // 设置模板目录
    engine.LoadHTMLGlob("templates/*")

    // 添加HTML文件的路由
    engine.GET("/online", func(c *gin.Context) {
        c.HTML(http.StatusOK, "online.html", nil)
    })

    engine.GET("/client", func(c *gin.Context) {
        c.HTML(http.StatusOK, "client.html", nil)
    })

    // 初始化用户路由
    userGroup := router.UserGroup{}
    userGroup.InitUserRouters(engine)

    // 注册 WebSocket 路由
    engine.GET("/user/ws", api.HandleWebSocket)

    // 启动 WebSocket 广播协程
    go api.BroadcastMessages()

    // 启动服务器
    address := fmt.Sprintf(":%d", global.CONFIG.App.Port)
    fmt.Println("服务启动成功，监听端口：", address)
    if err := engine.Run(address); err != nil {
        panic(err)
    }
}





// package main

// import (
// 	"fmt"
// 	_ "net/http/pprof"

// 	"ServerFramework/api"
// 	"ServerFramework/global"
// 	"ServerFramework/initialize"
// 	"ServerFramework/router"

// 	"github.com/gin-gonic/gin"
// )

// // main 是程序的入口函数，负责初始化项目的各项配置和服务。
// func main() {
// 	initialize.MustConfig()           // 初始化配置
// 	initialize.MustInitDB()           // 初始化数据库
// 	initialize.RegisterSerializer()   // 注册序列化器
// 	initialize.AutoMigrate(global.DB) // 自动迁移数据库结构, 确保数据库表结构与模型匹配

// 	engine := gin.Default() // 创建 Gin 引擎实例

// 	engine.Static("/static", "./static")   // 静态文件服务
// 	engine.LoadHTMLGlob("templates/*")     // 加载 HTML 模板

// 	engine.GET("/ws", api.HandleWebSocket) // WebSocket 路由
// 	userGroup := router.UserGroup{}
// 	userGroup.InitUserRouters(engine) // 初始化用户路由组

// 	go api.BroadcastMessages() // 启动 WebSocket 广播协程

// 	// 启动服务器
// 	address := fmt.Sprintf(":%d", global.CONFIG.App.Port)
// 	fmt.Println("服务启动成功，监听端口：", address)
// 	if err := engine.Run(address); err != nil {
// 		panic(err)
// 	}

// 	// initialize.MustLoadGorm() // 初始化 GORM 数据库连接
// 	// initialize.MustInitRedis() // 初始化 Redis 缓存
// 	// initialize.MustRunWindowServer() // 初始化窗口服务，用于处理窗口相关的操作
// 	// test.TickerUse() // 定时器使用示例
// 	// test.CronUse() // 定时任务使用示例
// 	// test.ExampleJSONSerializer() // JSON 序列化器使用示例
// 	// test.ExampleGobSerializer() // Gob 序列化器使用示例

// }
