package model

import "gorm.io/gorm"

// UserRole 用户角色关联模型
type UserRole struct {
	gorm.Model
	UserID uint64 `gorm:"index;comment:用户 ID"`
	RoleID uint64 `gorm:"index;comment:角色 ID"`
}
