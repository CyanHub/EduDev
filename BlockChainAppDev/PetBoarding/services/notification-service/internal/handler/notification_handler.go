package handler

import (
	"net/http"

	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/notification-service/internal/model"
	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/notification-service/internal/service"
	"github.com/gin-gonic/gin"
)

// NotificationHandler 处理通知相关的HTTP请求
type NotificationHandler struct {
	notificationService service.NotificationService
}

// NewNotificationHandler 创建通知处理器实例
func NewNotificationHandler(notificationService service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

// RegisterRoutes 注册路由
func (h *NotificationHandler) RegisterRoutes(router *gin.Engine) {
	notificationGroup := router.Group("/api/notifications")
	{
		notificationGroup.POST("", h.CreateNotification)
		notificationGroup.GET("/:id", h.GetNotificationByID)
		notificationGroup.GET("/user/:userId", h.GetNotificationsByUserID)
		notificationGroup.GET("/user/:userId/unread", h.GetUnreadNotificationsByUserID)
		notificationGroup.GET("", h.GetAllNotifications)
		notificationGroup.PUT("/:id/read", h.MarkAsRead)
		notificationGroup.PUT("/user/:userId/read-all", h.MarkAllAsRead)
		notificationGroup.DELETE("/:id", h.DeleteNotification)
	}
}

// CreateNotification 创建通知
func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var req model.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.notificationService.CreateNotification(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetNotificationByID 根据ID获取通知
func (h *NotificationHandler) GetNotificationByID(c *gin.Context) {
	id := c.Param("id")

	response, err := h.notificationService.GetNotificationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetNotificationsByUserID 获取用户的所有通知
func (h *NotificationHandler) GetNotificationsByUserID(c *gin.Context) {
	userID := c.Param("userId")

	responses, err := h.notificationService.GetNotificationsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// GetUnreadNotificationsByUserID 获取用户的未读通知
func (h *NotificationHandler) GetUnreadNotificationsByUserID(c *gin.Context) {
	userID := c.Param("userId")

	responses, err := h.notificationService.GetUnreadNotificationsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// GetAllNotifications 获取所有通知
func (h *NotificationHandler) GetAllNotifications(c *gin.Context) {
	responses, err := h.notificationService.GetAllNotifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// MarkAsRead 将通知标记为已读
func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	id := c.Param("id")

	err := h.notificationService.MarkAsRead(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

// MarkAllAsRead 将用户的所有通知标记为已读
func (h *NotificationHandler) MarkAllAsRead(c *gin.Context) {
	userID := c.Param("userId")

	err := h.notificationService.MarkAllAsRead(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All notifications marked as read"})
}

// DeleteNotification 删除通知
func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	id := c.Param("id")

	err := h.notificationService.DeleteNotification(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted"})
}