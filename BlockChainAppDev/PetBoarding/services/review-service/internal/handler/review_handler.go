package handler

import (
	"net/http"
	"strconv"

	"github.com/cyanhub/petboarding/services/review-service/internal/model"
	"github.com/cyanhub/petboarding/services/review-service/internal/service"
	"github.com/gin-gonic/gin"
)

// ReviewHandler 处理评论相关的HTTP请求
type ReviewHandler struct {
	reviewService service.ReviewService
}

// NewReviewHandler 创建评论处理器实例
func NewReviewHandler(reviewService service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

// RegisterRoutes 注册评论相关的路由
func (h *ReviewHandler) RegisterRoutes(router *gin.Engine) {
	reviewGroup := router.Group("/api/v1/reviews")
	{
		reviewGroup.POST("", h.CreateReview)
		reviewGroup.GET("", h.GetAllReviews)
		reviewGroup.GET("/:id", h.GetReviewByID)
		reviewGroup.GET("/user/:user_id", h.GetReviewsByUserID)
		reviewGroup.GET("/boarding/:boarding_id", h.GetReviewsByBoardingID)
		reviewGroup.PUT("/:id", h.UpdateReview)
		reviewGroup.DELETE("/:id", h.DeleteReview)
		reviewGroup.GET("/summary", h.GetReviewSummary)
	}
}

// CreateReview 处理创建评论请求
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req model.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.reviewService.CreateReview(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

// GetReviewByID 根据ID获取评论
func (h *ReviewHandler) GetReviewByID(c *gin.Context) {
	// 解析评论ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID"})
		return
	}

	review, err := h.reviewService.GetReviewByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

// GetReviewsByUserID 获取用户的所有评论
func (h *ReviewHandler) GetReviewsByUserID(c *gin.Context) {
	// 解析用户ID
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	reviews, err := h.reviewService.GetReviewsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

// GetReviewsByBoardingID 获取预订的所有评论
func (h *ReviewHandler) GetReviewsByBoardingID(c *gin.Context) {
	// 解析预订ID
	boardingID, err := strconv.ParseUint(c.Param("boarding_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的预订ID"})
		return
	}

	reviews, err := h.reviewService.GetReviewsByBoardingID(uint(boardingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

// GetAllReviews 获取所有评论（分页）
func (h *ReviewHandler) GetAllReviews(c *gin.Context) {
	// 解析分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	reviews, total, err := h.reviewService.GetAllReviews(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reviews":    reviews,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// UpdateReview 更新评论信息
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	// 解析评论ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID"})
		return
	}

	var req model.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.reviewService.UpdateReview(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

// DeleteReview 删除评论
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	// 解析评论ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID"})
		return
	}

	err = h.reviewService.DeleteReview(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评论删除成功"})
}

// GetReviewSummary 获取评论汇总
func (h *ReviewHandler) GetReviewSummary(c *gin.Context) {
	summary, err := h.reviewService.GetReviewSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}