package api

import (
	"log"

	"FileSystem/global"
	"FileSystem/model/request"
	"FileSystem/model/response"
	"FileSystem/service"
	"FileSystem/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func RoleList(c *gin.Context) {
	var req request.RoleListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("参数错误: ", utils.Translate(err))
		response.FailWithMessage(utils.Translate(err), c)
		return
	}
	total, list, err := service.RoleServiceApp.RoleList(req)
	if err != nil {
		log.Println("获取角色列表失败: ", err)
		response.FailWithMessage("获取角色列表失败", c)
		return
	}
	response.OkWithData(response.PageResult{
		Total:    total,
		List:     list,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}

// CreateRole 创建角色
func CreateRole(c *gin.Context) {
	var role request.RoleCreateRequest
	if err := c.ShouldBindJSON(&role); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	// 使用 validator/v10 进行输入校验
	validate := validator.New()
	if err := validate.Struct(role); err != nil {
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	// 在原有校验基础上增加
	validate.RegisterValidation("strong_pass", func(fl validator.FieldLevel) bool {
		pass := fl.Field().String()
		// 至少8位，包含大小写和特殊字符
		return utils.CheckPasswordStrength(pass)
	})

	if err := validate.Struct(role); err != nil {
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	// 使用 GORM 进行数据库操作
	if err := service.RoleServiceApp.RoleCreate(role); err != nil {
		global.Logger.Error("创建角色失败", zap.Error(err))
		response.FailWithMessage("创建角色失败", c)
		return
	}

	response.OkWithMessage("创建角色成功", c)
}

// AssignPermissions 分配权限
func AssignPermissions(c *gin.Context) {
	var assign request.AssignPermissions
	if err := c.ShouldBindJSON(&assign); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	// 使用 validator/v10 进行输入校验
	validate := validator.New()
	if err := validate.Struct(assign); err != nil {
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	// 使用 GORM 进行数据库操作，并实现事务操作
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := service.RoleServiceApp.DeleteRolePermissions(tx, assign.RoleID); err != nil {
			global.Logger.Error("删除角色权限失败", zap.Error(err))
			return err
		}

		if err := service.RoleServiceApp.AddRolePermissions(tx, assign.RoleID, assign.Permissions); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		global.Logger.Error("分配权限失败", zap.Error(err))
		response.FailWithMessage("分配权限失败", c)
		return
	}

	response.OkWithMessage("分配权限成功", c)
}

// CheckPermission 检查权限
func CheckPermission(c *gin.Context) {
	var check request.CheckPermission
	if err := c.ShouldBindJSON(&check); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	// 使用 validator/v10 进行输入校验
	validate := validator.New()
	if err := validate.Struct(check); err != nil {
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	// 使用 GORM 进行数据库操作
	hasPermission, err := service.RoleServiceApp.CheckPermission(check.UserID, check.Permission)
	if err != nil {
		global.Logger.Error("检查权限失败", zap.Error(err))
		response.FailWithMessage("检查权限失败", c)
		return
	}

	response.OkWithDetailed(gin.H{"hasPermission": hasPermission}, "检查权限成功", c)
}
