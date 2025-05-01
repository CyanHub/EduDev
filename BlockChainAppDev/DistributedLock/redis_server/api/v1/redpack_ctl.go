package v1

import (
	"BlockChainDev/redis_server/internal/logics"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetRedpacks 获取红包列表
func GetRedpacks(c *gin.Context) {
	redpackLogic := logics.RedpackLgc{}
	redpacks, err := redpackLogic.GetAllRedpacks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, redpacks)
}

// CreateRedpack 创建红包
func CreateRedpack(c *gin.Context) {
	var redpackRequest struct {
		Amount int `json:"amount"`
		Num    int `json:"num"`
	}
	if err := c.ShouldBindJSON(&redpackRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redpackLogic := logics.RedpackLgc{}
	redpack, err := redpackLogic.CreateRedpack(redpackRequest.Amount, redpackRequest.Num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, redpack)
}

// GrabRedpack 抢红包
func GrabRedpack(c *gin.Context) {
	userIdStr := c.Query("user_id")
	redpackIdStr := c.Query("redpack_id")

	if userIdStr == "" || redpackIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and redpack_id are required"})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	redpackId, err := strconv.ParseInt(redpackIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid redpack_id"})
		return
	}

	redpackLogic := logics.RedpackLgc{}
	record, err := redpackLogic.GrabRedpack(userId, redpackId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

// GetRedpackRecord 获取用户的红包记录
func GetRedpackRecord(c *gin.Context) {
	userIdStr := c.Query("user_id")
	if userIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	recordLogic := logics.RedpackRecordLgc{}
	records, err := recordLogic.GetRedpackRecordsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}
