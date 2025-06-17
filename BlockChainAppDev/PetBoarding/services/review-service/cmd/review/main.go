package main

import (
	"log"

	"github.com/cyanhub/petboarding/services/review-service/internal/handler"
	"github.com/cyanhub/petboarding/services/review-service/internal/repository"
	"github.com/cyanhub/petboarding/services/review-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 初始化依赖
	reviewRepo := repository.NewReviewRepository()
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewService)

	// 注册路由
	reviewHandler.RegisterRoutes(r)

	// 添加健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "review-service",
		})
	})

	// 启动HTTP服务器
	log.Println("评论服务启动在 :8084 端口")
	if err := r.Run(":8084"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}