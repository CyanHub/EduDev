package main

import (
	"log"
	"net/http"

	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/admin-service/internal/handler"
	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/admin-service/internal/repository"
	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/admin-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin路由
	router := gin.Default()

	// 初始化仓库
	adminRepo := repository.NewAdminRepository()

	// 初始化服务
	adminService := service.NewAdminService(adminRepo)

	// 初始化处理器
	adminHandler := handler.NewAdminHandler(adminService)

	// 注册路由
	adminHandler.RegisterRoutes(router)

	// 健康检查端点
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"service": "admin-service",
		})
	})

	// 启动服务器
	log.Println("Admin Service starting on port 8086...")
	if err := router.Run(":8086"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}