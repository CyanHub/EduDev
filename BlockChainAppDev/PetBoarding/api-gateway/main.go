package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ServiceConfig 表示服务配置
type ServiceConfig struct {
	Name string
	URL  string
}

// APIGateway 表示API网关
type APIGateway struct {
	services map[string]*ServiceConfig
}

// NewAPIGateway 创建API网关实例
func NewAPIGateway() *APIGateway {
	return &APIGateway{
		services: make(map[string]*ServiceConfig),
	}
}

// RegisterService 注册服务
func (g *APIGateway) RegisterService(name, url string) {
	g.services[name] = &ServiceConfig{
		Name: name,
		URL:  url,
	}
	log.Printf("Registered service: %s with URL: %s", name, url)
}

// ProxyHandler 代理处理器
func (g *APIGateway) ProxyHandler(c *gin.Context) {
	// 从路径中提取服务名称
	path := c.Request.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid path"})
		return
	}

	// 服务名称是路径的第二部分，例如 /api/users -> users
	serviceName := parts[2]

	// 查找服务配置
	service, exists := g.services[serviceName]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Service '%s' not found", serviceName)})
		return
	}

	// 解析目标URL
	targetURL, err := url.Parse(service.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse service URL"})
		return
	}

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// 修改请求
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		// 移除API前缀，例如 /api/users/1 -> /users/1
		req.URL.Path = strings.Replace(req.URL.Path, "/api/"+serviceName, "", 1)
		if req.URL.Path == "" {
			req.URL.Path = "/"
		}

		// 添加请求头
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Origin-Host", targetURL.Host)
		req.Host = targetURL.Host
	}

	// 错误处理
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error: %v", err)
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(fmt.Sprintf("Service '%s' is unavailable", serviceName)))
	}

	// 修改响应
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Set("X-Gateway", "PetBoarding-API-Gateway")
		return nil
	}

	// 执行代理请求
	proxy.ServeHTTP(c.Writer, c.Request)
}

// CORSMiddleware 处理跨域请求
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// LoggingMiddleware 记录请求日志
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 请求信息
		log.Printf(
			"[API-GATEWAY] %s | %d | %s | %s | %s",
			c.Request.Method,
			c.Writer.Status(),
			c.Request.URL.Path,
			c.ClientIP(),
			latency,
		)
	}
}

func main() {
	// 创建Gin路由
	router := gin.Default()

	// 添加中间件
	router.Use(CORSMiddleware())
	router.Use(LoggingMiddleware())

	// 创建API网关
	gateway := NewAPIGateway()

	// 从环境变量获取服务URL
	userServiceURL := getEnv("USER_SERVICE_URL", "http://user-service:8081")
	petServiceURL := getEnv("PET_SERVICE_URL", "http://pet-service:8082")
	boardingServiceURL := getEnv("BOARDING_SERVICE_URL", "http://boarding-service:8083")
	reviewServiceURL := getEnv("REVIEW_SERVICE_URL", "http://review-service:8084")
	notificationServiceURL := getEnv("NOTIFICATION_SERVICE_URL", "http://notification-service:8085")
	adminServiceURL := getEnv("ADMIN_SERVICE_URL", "http://admin-service:8086")

	// 注册服务
	gateway.RegisterService("users", userServiceURL)
	gateway.RegisterService("pets", petServiceURL)
	gateway.RegisterService("boardings", boardingServiceURL)
	gateway.RegisterService("reviews", reviewServiceURL)
	gateway.RegisterService("notifications", notificationServiceURL)
	gateway.RegisterService("admin", adminServiceURL)

	// API路由
	router.Any("/api/:service/*path", gateway.ProxyHandler)
	router.Any("/api/:service", gateway.ProxyHandler)

	// 健康检查端点
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"service": "api-gateway",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// 首页
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to PetBoarding API Gateway",
			"services": []string{"users", "pets", "boardings", "reviews", "notifications", "admin"},
			"version": "1.0.0",
		})
	})

	// 启动服务器
	port := getEnv("PORT", "8080")
	log.Printf("API Gateway starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}