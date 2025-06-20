package initialize

import (
	"FileSystem/global"
	"FileSystem/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// InitAdmin 初始化默认管理员用户
func InitAdmin() error {
	// 检查是否已经存在管理员用户
	var adminCount int64
	global.DB.Model(&model.User{}).
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("roles.name = ?", "admin").
		Count(&adminCount)

	if adminCount > 0 {
		return nil // 已有管理员用户，不重复创建
	}

	// 确保管理员角色存在
	var adminRole model.Role
	err := global.DB.Where("name = ?", "admin").First(&adminRole).Error
	if err != nil {
		if err == model.RecordNotFound {
			// 若角色不存在，则创建管理员角色
			adminRole = model.Role{
				Name:        "admin",
				Description: "超级管理员",
			}
			if err := global.DB.Create(&adminRole).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// 定义默认管理员用户信息
	defaultAdmin := model.User{
		Username: "super_admin",
		NickName: "超级管理员",
		Email:    "admin@root.com",
		Phone:    "10101010101",
		Avatar:   "",
	}

	// 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	defaultAdmin.Password = string(hashedPassword)

	// 开启事务创建用户和分配角色
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		if err := tx.Create(&defaultAdmin).Error; err != nil {
			return err
		}

		// 分配角色
		userRole := model.UserRole{
			UserID: defaultAdmin.ID,
			RoleID: adminRole.ID,
		}
		if err := tx.Create(&userRole).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
