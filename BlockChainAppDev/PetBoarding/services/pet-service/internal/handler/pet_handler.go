package handler

import (
	"net/http"
	"strconv"

	"github.com/cyanhub/petboarding/services/pet-service/internal/model"
	"github.com/cyanhub/petboarding/services/pet-service/internal/service"
	"github.com/gin-gonic/gin"
)

// PetHandler 处理宠物相关的HTTP请求
type PetHandler struct {
	petService service.PetService
}

// NewPetHandler 创建宠物处理器实例
func NewPetHandler(petService service.PetService) *PetHandler {
	return &PetHandler{
		petService: petService,
	}
}

// RegisterRoutes 注册宠物相关的路由
func (h *PetHandler) RegisterRoutes(router *gin.Engine) {
	petGroup := router.Group("/api/v1/pets")
	{
		petGroup.POST("", h.CreatePet)
		petGroup.GET("", h.GetAllPets)
		petGroup.GET("/:id", h.GetPetByID)
		petGroup.GET("/user/:user_id", h.GetPetsByUserID)
		petGroup.PUT("/:id", h.UpdatePet)
		petGroup.DELETE("/:id", h.DeletePet)
	}
}

// CreatePet 处理创建宠物请求
func (h *PetHandler) CreatePet(c *gin.Context) {
	var req model.CreatePetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet, err := h.petService.CreatePet(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pet)
}

// GetPetByID 根据ID获取宠物
func (h *PetHandler) GetPetByID(c *gin.Context) {
	// 解析宠物ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的宠物ID"})
		return
	}

	pet, err := h.petService.GetPetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pet)
}

// GetPetsByUserID 获取用户的所有宠物
func (h *PetHandler) GetPetsByUserID(c *gin.Context) {
	// 解析用户ID
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	pets, err := h.petService.GetPetsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pets": pets})
}

// GetAllPets 获取所有宠物（分页）
func (h *PetHandler) GetAllPets(c *gin.Context) {
	// 解析分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	pets, total, err := h.petService.GetAllPets(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pets":       pets,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// UpdatePet 更新宠物信息
func (h *PetHandler) UpdatePet(c *gin.Context) {
	// 解析宠物ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的宠物ID"})
		return
	}

	var req model.UpdatePetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet, err := h.petService.UpdatePet(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pet)
}

// DeletePet 删除宠物
func (h *PetHandler) DeletePet(c *gin.Context) {
	// 解析宠物ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的宠物ID"})
		return
	}

	err = h.petService.DeletePet(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "宠物删除成功"})
}