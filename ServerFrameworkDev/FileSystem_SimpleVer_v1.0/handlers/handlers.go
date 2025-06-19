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
	var roleModels []models.Role // 修改为 []models.Role 类型
	var roles []string
	// 从数据库获取用户角色
	if err := db.Model(&user).Association("Roles").Find(&roleModels); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户角色失败"})
		return
	}
	// 提取角色名称
	for _, role := range roleModels {
		roles = append(roles, role.Name)
	}

	token, err := auth.GenerateToken(user.ID, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取令牌失败"})
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

	// 获取 isPublic 参数
	isPublicStr := c.PostForm("isPublic")
	isPublic := false
	if isPublicStr == "true" {
		isPublic = true
	}

	fileRecord := models.File{
		Name:    file.Filename,
		Path:    filePath,
		OwnerID: claims.UserID,
		IsPublic: isPublic,
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

// ========================管理员功能添加，后端接口实现============================
// CheckAdmin 检查用户是否为管理者
func CheckAdmin(c *gin.Context) {
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

	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Preload("Roles.Permissions").First(&user, claims.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
		return
	}

	isAdmin := false
	for _, role := range user.Roles {
		if role.Name == "管理者" {
			isAdmin = true
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"isAdmin": isAdmin})
}

// GetUsers 获取所有用户
func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetFiles 获取文件列表
func GetFiles(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var files []models.File
    // 预加载 Owner 信息
    if err := db.Preload("Owner").Find(&files).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文件列表失败"})
        return
    }
	logging.Logger.Info("获取文件列表成功", zap.Int("fileCount", len(files)))
    c.JSON(http.StatusOK, files)
}

// GetRoles 获取所有角色
func GetRoles(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var roles []models.Role
	if err := db.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取角色列表失败"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// GetPermissions 获取所有权限
func GetPermissions(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var permissions []models.Permission
	if err := db.Find(&permissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取权限列表失败"})
		return
	}
	c.JSON(http.StatusOK, permissions)
}

// AssignRoleToUser 分配角色给用户
func AssignRoleToUser(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id" validate:"required"`
		RoleID uint `json:"role_id" validate:"required"`
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
	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "开始事务失败"})
		return
	}

	// 检查 user_roles 表中是否已存在该组合
	var count int64
	if err := tx.Model(&models.UserRole{}).Where("user_id = ? AND role_id = ?", input.UserID, input.RoleID).Count(&count).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "检查角色分配失败"})
		return
	}
	if count > 0 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "该用户已拥有此角色"})
		return
	}

	if err := tx.Exec("INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)", input.UserID, input.RoleID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "分配角色失败"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "提交事务失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "角色分配成功"})
}

// AssignPermissionToRole 分配权限给角色
func AssignPermissionToRole(c *gin.Context) {
	var input struct {
		RoleID       uint `json:"role_id" validate:"required"`
		PermissionID uint `json:"permission_id" validate:"required"`
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
	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "开始事务失败"})
		return
	}

	if err := tx.Exec("INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?)", input.RoleID, input.PermissionID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "分配权限失败"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "提交事务失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "权限分配成功"})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "开始事务失败"})
		return
	}

	if err := tx.Where("id = ?", userId).Delete(&models.User{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除用户失败"})
		return
	}

	if err := tx.Exec("DELETE FROM user_roles WHERE user_id = ?", userId).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除用户关联角色失败"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "提交事务失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}

// DeleteRole 删除角色
func DeleteRole(c *gin.Context) {
	roleId := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "开始事务失败"})
		return
	}

	if err := tx.Where("id = ?", roleId).Delete(&models.Role{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除角色失败"})
		return
	}

	if err := tx.Exec("DELETE FROM user_roles WHERE role_id = ?", roleId).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除角色关联用户失败"})
		return
	}

	if err := tx.Exec("DELETE FROM role_permissions WHERE role_id = ?", roleId).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除角色关联权限失败"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "提交事务失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "角色删除成功"})
}

// DeletePermission 删除权限
func DeletePermission(c *gin.Context) {
	permissionId := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "开始事务失败"})
		return
	}

	if err := tx.Where("id = ?", permissionId).Delete(&models.Permission{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除权限失败"})
		return
	}

	if err := tx.Exec("DELETE FROM role_permissions WHERE permission_id = ?", permissionId).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除权限关联角色失败"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "提交事务失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "权限删除成功"})
}


// =========管理员添加，添加新的处理函数，同时需要进行权限检查，确保只有管理者能调用该接口。==========
// AddAdmin 添加管理者
func AddAdmin(c *gin.Context) {
    // 检查用户是否为管理者
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

    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Preload("Roles.Permissions").First(&user, claims.UserID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
        return
    }

    isAdmin := false
    for _, role := range user.Roles {
        if role.Name == "管理者" {
            isAdmin = true
            break
        }
    }

    if !isAdmin {
        c.JSON(http.StatusForbidden, gin.H{"error": "只有管理者可以添加管理者"})
        return
    }

    // 绑定请求数据
    var input UserRegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validate.Struct(input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 创建用户
    userToAdd := models.User{
        Username: input.Username,
        Password: auth.HashPassword(input.Password),
        Email:    input.Email,
    }

    if err := db.Create(&userToAdd).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
        return
    }

    // 获取管理者角色 ID
    var adminRole models.Role
    if err := db.Where("name = ?", "管理者").First(&adminRole).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取管理者角色失败"})
        return
    }

    // 绑定用户与管理者角色
    if err := db.Exec("INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)", userToAdd.ID, adminRole.ID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "分配管理者角色失败"})
        return
    }

    logging.LogAccess(userToAdd.ID, "admin_register")
    c.JSON(http.StatusCreated, gin.H{"message": "管理者添加成功"})
}

// ================管理员调用密码重置，该接口仅允许管理者调用================
// ResetPassword 重置用户密码
func ResetPassword(c *gin.Context) {
    // 检查用户是否为管理者
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

    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Preload("Roles.Permissions").First(&user, claims.UserID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
        return
    }

    isAdmin := false
    for _, role := range user.Roles {
        if role.Name == "管理者" {
            isAdmin = true
            break
        }
    }

    if !isAdmin {
        c.JSON(http.StatusForbidden, gin.H{"error": "只有管理者可以重置密码"})
        return
    }

    // 绑定请求数据
    var input struct {
        Username    string `json:"username" validate:"required"`
        NewPassword string `json:"new_password" validate:"required,min=8,containsany=!@#$%^&*"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validate.Struct(input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 查找要重置密码的用户
    var targetUser models.User
    if err := db.Where("username = ?", input.Username).First(&targetUser).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询失败"})
        return
    }

    // 更新用户密码
    targetUser.Password = auth.HashPassword(input.NewPassword)
    if err := db.Save(&targetUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "更新密码失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "密码重置成功"})
}

// CheckFileName 检查文件名是否已存在
func CheckFileName(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    name := c.Query("name")

    var count int64
    db.Model(&models.File{}).Where("name = ?", name).Count(&count)

    c.JSON(http.StatusOK, gin.H{"exists": count > 0})
}
