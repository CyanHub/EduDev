package main

import (
	"log"

	"github.com/cyanhub/petboarding/services/user-service/internal/handler"
	"github.com/cyanhub/petboarding/services/user-service/internal/repository"
	"github.com/cyanhub/petboarding/services/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 初始化依赖
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 注册路由
	userHandler.RegisterRoutes(r)

	// 添加健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "user-service",
		})
	})

	// 启动HTTP服务器
	log.Println("用户服务启动在 :8081 端口")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}