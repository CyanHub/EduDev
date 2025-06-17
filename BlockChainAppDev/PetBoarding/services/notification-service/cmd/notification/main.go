package main

import (
	"log"
	"net/http"

	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/notification-service/internal/handler"
	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/notification-service/internal/repository"
	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/notification-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin路由
	router := gin.Default()

	// 初始化仓库
	notificationRepo := repository.NewNotificationRepository()

	// 初始化服务
	notificationService := service.NewNotificationService(notificationRepo)

	// 初始化处理器
	notificationHandler := handler.NewNotificationHandler(notificationService)

	// 注册路由
	notificationHandler.RegisterRoutes(router)

	// 健康检查端点
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"service": "notification-service",
		})
	})

	// 启动服务器
	log.Println("Notification Service starting on port 8085...")
	if err := router.Run(":8085"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}