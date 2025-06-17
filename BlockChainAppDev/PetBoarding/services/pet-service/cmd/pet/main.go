package main

import (
	"log"

	"github.com/cyanhub/petboarding/services/pet-service/internal/handler"
	"github.com/cyanhub/petboarding/services/pet-service/internal/repository"
	"github.com/cyanhub/petboarding/services/pet-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 初始化依赖
	petRepo := repository.NewPetRepository()
	petService := service.NewPetService(petRepo)
	petHandler := handler.NewPetHandler(petService)

	// 注册路由
	petHandler.RegisterRoutes(r)

	// 添加健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "pet-service",
		})
	})

	// 启动HTTP服务器
	log.Println("宠物服务启动在 :8082 端口")
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}