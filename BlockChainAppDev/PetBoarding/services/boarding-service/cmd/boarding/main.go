package main

import (
	"log"

	"github.com/cyanhub/petboarding/services/boarding-service/internal/handler"
	"github.com/cyanhub/petboarding/services/boarding-service/internal/repository"
	"github.com/cyanhub/petboarding/services/boarding-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 初始化依赖
	boardingRepo := repository.NewBoardingRepository()
	boardingService := service.NewBoardingService(boardingRepo)
	boardingHandler := handler.NewBoardingHandler(boardingService)

	// 注册路由
	boardingHandler.RegisterRoutes(r)

	// 添加健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "boarding-service",
		})
	})

	// 启动HTTP服务器
	log.Println("预订服务启动在 :8083 端口")
	if err := r.Run(":8083"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}