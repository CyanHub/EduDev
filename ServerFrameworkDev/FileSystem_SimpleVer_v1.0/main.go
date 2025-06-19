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
	r.POST("/register", handlers.RegisterUser)                   // 注册用户
	r.POST("/login", handlers.LoginUser)                         // 登录用户
	r.POST("/upload", handlers.UploadFile)                       // 上传文件
	r.GET("/download/:id", handlers.DownloadFile)                // 下载文件
	r.GET("/files", handlers.GetFiles)                           // 获取文件列表
	r.GET("/check-file-name", handlers.CheckFileName)            // 检查文件名是否存在
	r.POST("/roles", handlers.CreateRole)                        // 创建角色
	r.POST("/permissions", handlers.CreatePermission)            // 创建权限
	r.GET("/check-admin", handlers.CheckAdmin)                   // 检查管理员权限
	r.GET("/users", handlers.GetUsers)                           // 获取用户列表
	r.GET("/roles", handlers.GetRoles)                           // 获取角色列表
	r.GET("/permissions", handlers.GetPermissions)               // 获取权限列表
	r.POST("/user-roles", handlers.AssignRoleToUser)             // 分配角色给用户
	r.POST("/role-permissions", handlers.AssignPermissionToRole) // 分配权限给角色
	r.DELETE("/users/:id", handlers.DeleteUser)                  // 删除用户
	r.DELETE("/roles/:id", handlers.DeleteRole)                  // 删除角色
	r.DELETE("/permissions/:id", handlers.DeletePermission)      // 删除权限
	r.POST("/add-admin", handlers.AddAdmin)                      // 新增管理者添加路由
	r.POST("/reset-password", handlers.ResetPassword)            // 重置密码接口

	// 启动服务器
	if err := r.Run(":" + strconv.Itoa(cfg.Server.Port)); err != nil {
		logging.Logger.Fatal("服务启动失败", zap.Error(err))
	}
}
