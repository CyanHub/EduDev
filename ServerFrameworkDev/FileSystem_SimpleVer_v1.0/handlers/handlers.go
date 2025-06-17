package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"ServerFrameworkDev/FileSystem/auth"
	"ServerFrameworkDev/FileSystem/logging"
	"ServerFrameworkDev/FileSystem/models"
)

type UserRegisterInput struct {
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
	Password string `json:"password" validate:"required,min=8,containsany=!@#$%^&*"`
	Email    string `json:"email" validate:"required,email"`
}

type UserLoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var validate = validator.New()

func RegisterUser(c *gin.Context) {
	var input UserRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	user := models.User{
		Username: input.Username,
		Password: auth.HashPassword(input.Password),
		Email:    input.Email, // 确保保存 Email
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}

	logging.LogAccess(user.ID, "register")
	c.JSON(http.StatusCreated, gin.H{"message": "用户注册成功"})
}

func LoginUser(c *gin.Context) {
	var input UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 修复 logging.Field 未定义问题，假设使用 zap 风格的日志记录
	logging.Logger.Info("已接收登录请求",
		zap.String("username", input.Username),
		zap.String("password", input.Password),
	)

	// 修复logging.Field未定义问题
	
	// 修复 ValidatePassword 未定义问题
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库加载错误"})
		return
	}

	// 验证密码
	if !auth.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码不存在"})
		return
	}

	// 修复 auth.GenerateToken 参数不匹配问题，假设需要用户角色
	var roles []string
	// 从数据库获取用户角色
	if err := db.Model(&user).Association("Roles").Find(&roles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户角色失败"})
		return
	}
	token, err := auth.GenerateToken(user.ID, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新获取令牌失败"})
		return
	}

	logging.LogAccess(user.ID, "login")
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件上传失败"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	filePath := "uploads/" + file.Filename

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	// 从 JWT 中获取用户信息
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization 请求头缺失"})
		return
	}
	claims, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌无效"})
		return
	}

	// 检查用户角色权限
	var user models.User
	if err := db.Preload("Roles.Permissions").First(&user, claims.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
		return
	}

	hasPermission := false
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			if permission.Code == "upload_file" {
				hasPermission = true
				break
			}
		}
		if hasPermission {
			break
		}
	}

	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"error": "权限访问被拒绝"})
		return
	}

	// 移除 Size 字段
	fileRecord := models.File{
		Name:    file.Filename,
		Path:    filePath,
		OwnerID: claims.UserID, // 确保设置 OwnerID
	}

	if err := db.Create(&fileRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件日志记录保存失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功", "fileId": fileRecord.ID})
}

// DownloadFile 实现文件下载和权限检查
func DownloadFile(c *gin.Context) {
	// 从 JWT 中获取用户信息
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization 请求头缺失"})
		return
	}
	claims, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌无效"})
		return
	}

	// 获取文件 ID
	fileIDStr := c.Param("id")
	fileID, err := strconv.Atoi(fileIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文件ID"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var file models.File
	if err := db.Preload("Owner").First(&file, fileID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该文件"})
		return
	}

	// 检查权限
	if file.IsPublic || file.OwnerID == claims.UserID {
		// 公共文件或文件所有者可以下载
		c.File(file.Path)
		logging.LogAccess(claims.UserID, "download file")
		return
	}

	// 检查用户角色权限
	var user models.User
	if err := db.Preload("Roles.Permissions").First(&user, claims.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
		return
	}

	hasPermission := false
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			if permission.Code == "download_file" {
				hasPermission = true
				break
			}
		}
		if hasPermission {
			break
		}
	}

	if hasPermission {
		c.File(file.Path)
		logging.LogAccess(claims.UserID, "download file")
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限下载此文件"})
	}
}

// CreateRole 创建角色
func CreateRole(c *gin.Context) {
	var input struct {
		Name string `json:"name" validate:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	role := models.Role{
		Name: input.Name,
	}

	if err := db.Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "角色已存在"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "角色创建成功", "roleId": role.ID})
}

// CreatePermission 创建权限
func CreatePermission(c *gin.Context) {
	var input struct {
		Name string `json:"name" validate:"required"`
		Code string `json:"code" validate:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	permission := models.Permission{
		Name: input.Name,
		Code: input.Code,
	}

	if err := db.Create(&permission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "权限已存在"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "权限创建成功", "permissionId": permission.ID})
}
