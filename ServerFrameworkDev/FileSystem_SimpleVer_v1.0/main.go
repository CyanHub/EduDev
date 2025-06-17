package main

import (
	"ServerFrameworkDev/FileSystem/config"
	"ServerFrameworkDev/FileSystem/handlers"
	"ServerFrameworkDev/FileSystem/logging"
	"ServerFrameworkDev/FileSystem/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 自定义 CORS 中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域名进行跨域调用
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		// 处理 OPTIONS 请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化日志
	logging.Init()
	defer logging.Logger.Sync()

	// 连接数据库
	dsn := cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + strconv.Itoa(cfg.Database.Port) + ")/" + cfg.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Logger.Fatal("连接数据库失败", zap.Error(err))
	}

	// 自动迁移数据库
	if err := db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.File{}); err != nil {
		logging.Logger.Fatal("迁移数据库失败", zap.Error(err))
	}

	// 初始化Gin路由
	r := gin.Default()

	// 添加 CORS 中间件
	r.Use(corsMiddleware())

	// 添加数据库实例到上下文
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// 注册路由
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.POST("/upload", handlers.UploadFile)
	r.GET("/download/:id", handlers.DownloadFile)
	r.POST("/roles", handlers.CreateRole)
	r.POST("/permissions", handlers.CreatePermission)

	// 启动服务器
	if err := r.Run(":" + strconv.Itoa(cfg.Server.Port)); err != nil {
		logging.Logger.Fatal("服务启动失败", zap.Error(err))
	}
}
