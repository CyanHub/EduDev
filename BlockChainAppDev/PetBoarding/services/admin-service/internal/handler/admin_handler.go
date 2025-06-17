package handler

import (
	"net/http"

	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/admin-service/internal/model"
	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/admin-service/internal/service"
	"github.com/gin-gonic/gin"
)

// AdminHandler 处理管理员相关的HTTP请求
type AdminHandler struct {
	adminService service.AdminService
}

// NewAdminHandler 创建管理员处理器实例
func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

// RegisterRoutes 注册路由
func (h *AdminHandler) RegisterRoutes(router *gin.Engine) {
	adminGroup := router.Group("/api/admin")
	{
		// 认证相关
		adminGroup.POST("/login", h.Login)

		// 管理员管理
		adminGroup.POST("/admins", h.CreateAdmin)
		adminGroup.GET("/admins/:id", h.GetAdminByID)
		adminGroup.GET("/admins", h.GetAllAdmins)
		adminGroup.PUT("/admins/:id", h.UpdateAdmin)
		adminGroup.DELETE("/admins/:id", h.DeleteAdmin)
		adminGroup.PUT("/admins/:id/password", h.ChangePassword)

		// 仪表板
		adminGroup.GET("/dashboard", h.GetDashboardData)
	}
}

// Login 管理员登录
func (h *AdminHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.adminService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateAdmin 创建管理员
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var req model.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.adminService.CreateAdmin(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetAdminByID 根据ID获取管理员
func (h *AdminHandler) GetAdminByID(c *gin.Context) {
	id := c.Param("id")

	response, err := h.adminService.GetAdminByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllAdmins 获取所有管理员
func (h *AdminHandler) GetAllAdmins(c *gin.Context) {
	responses, err := h.adminService.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

// UpdateAdmin 更新管理员
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	id := c.Param("id")

	var req model.UpdateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.adminService.UpdateAdmin(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteAdmin 删除管理员
func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	id := c.Param("id")

	err := h.adminService.DeleteAdmin(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}

// ChangePassword 更改管理员密码
func (h *AdminHandler) ChangePassword(c *gin.Context) {
	id := c.Param("id")

	var req model.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.adminService.ChangePassword(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

// GetDashboardData 获取仪表板数据
func (h *AdminHandler) GetDashboardData(c *gin.Context) {
	data, err := h.adminService.GetDashboardData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}