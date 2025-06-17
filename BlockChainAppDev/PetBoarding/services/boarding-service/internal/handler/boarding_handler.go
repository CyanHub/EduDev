package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cyanhub/petboarding/services/boarding-service/internal/model"
	"github.com/cyanhub/petboarding/services/boarding-service/internal/service"
	"github.com/gin-gonic/gin"
)

// BoardingHandler 处理预订相关的HTTP请求
type BoardingHandler struct {
	boardingService service.BoardingService
}

// NewBoardingHandler 创建预订处理器实例
func NewBoardingHandler(boardingService service.BoardingService) *BoardingHandler {
	return &BoardingHandler{
		boardingService: boardingService,
	}
}

// RegisterRoutes 注册预订相关的路由
func (h *BoardingHandler) RegisterRoutes(router *gin.Engine) {
	boardingGroup := router.Group("/api/v1/boardings")
	{
		boardingGroup.POST("", h.CreateBoarding)
		boardingGroup.GET("", h.GetAllBoardings)
		boardingGroup.GET("/:id", h.GetBoardingByID)
		boardingGroup.GET("/user/:user_id", h.GetBoardingsByUserID)
		boardingGroup.GET("/pet/:pet_id", h.GetBoardingsByPetID)
		boardingGroup.PUT("/:id", h.UpdateBoarding)
		boardingGroup.PUT("/:id/status", h.UpdateBoardingStatus)
		boardingGroup.DELETE("/:id", h.DeleteBoarding)
		boardingGroup.GET("/availability", h.CheckAvailability)
		boardingGroup.GET("/service-prices", h.GetServicePrices)
	}
}

// CreateBoarding 处理创建预订请求
func (h *BoardingHandler) CreateBoarding(c *gin.Context) {
	var req model.CreateBoardingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boarding, err := h.boardingService.CreateBoarding(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, boarding)
}

// GetBoardingByID 根据ID获取预订
func (h *BoardingHandler) GetBoardingByID(c *gin.Context) {
	// 解析预订ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的预订ID"})
		return
	}

	boarding, err := h.boardingService.GetBoardingByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, boarding)
}

// GetBoardingsByUserID 获取用户的所有预订
func (h *BoardingHandler) GetBoardingsByUserID(c *gin.Context) {
	// 解析用户ID
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	boardings, err := h.boardingService.GetBoardingsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"boardings": boardings})
}

// GetBoardingsByPetID 获取宠物的所有预订
func (h *BoardingHandler) GetBoardingsByPetID(c *gin.Context) {
	// 解析宠物ID
	petID, err := strconv.ParseUint(c.Param("pet_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的宠物ID"})
		return
	}

	boardings, err := h.boardingService.GetBoardingsByPetID(uint(petID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"boardings": boardings})
}

// GetAllBoardings 获取所有预订（分页）
func (h *BoardingHandler) GetAllBoardings(c *gin.Context) {
	// 解析分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	boardings, total, err := h.boardingService.GetAllBoardings(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"boardings":  boardings,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// UpdateBoarding 更新预订信息
func (h *BoardingHandler) UpdateBoarding(c *gin.Context) {
	// 解析预订ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的预订ID"})
		return
	}

	var req model.UpdateBoardingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boarding, err := h.boardingService.UpdateBoarding(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, boarding)
}

// UpdateBoardingStatus 更新预订状态
func (h *BoardingHandler) UpdateBoardingStatus(c *gin.Context) {
	// 解析预订ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的预订ID"})
		return
	}

	var req model.UpdateBoardingStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boarding, err := h.boardingService.UpdateBoardingStatus(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, boarding)
}

// DeleteBoarding 删除预订
func (h *BoardingHandler) DeleteBoarding(c *gin.Context) {
	// 解析预订ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的预订ID"})
		return
	}

	err = h.boardingService.DeleteBoarding(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "预订删除成功"})
}

// CheckAvailability 检查日期是否可用
func (h *BoardingHandler) CheckAvailability(c *gin.Context) {
	// 解析日期参数
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "开始日期和结束日期不能为空"})
		return
	}

	// 解析日期
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的开始日期格式，请使用YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的结束日期格式，请使用YYYY-MM-DD"})
		return
	}

	// 检查日期是否可用
	available, err := h.boardingService.CheckAvailability(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available": available})
}

// GetServicePrices 获取服务价格列表
func (h *BoardingHandler) GetServicePrices(c *gin.Context) {
	prices := h.boardingService.GetServicePrices()
	c.JSON(http.StatusOK, gin.H{"prices": prices})
}