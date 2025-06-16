package model

import (
	"FileSystem/global"

	"gorm.io/gorm"
)

type Role struct {
	global.GSModel
	Name        string       `gorm:"unique;size:50"`
	Description string       `gorm:"size:255"`
	// 添加父角色字段
	ParentID  uint64 `gorm:"index"` 
	// 添加权限关联字段
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	global.GSModel
	Code        string `gorm:"unique;size:50"` // 权限代码 如file:upload
	Description string `gorm:"size:255"`
}

type RolePermission struct {
	RoleID       uint64 `gorm:"primaryKey"`
	PermissionID uint64 `gorm:"primaryKey"` // 修正字段名
}

func (Role) TableName() string {
	return "role"
}

type UserRole struct {
	gorm.Model
	UserID uint64 `gorm:"index"`
	RoleID uint64 `gorm:"index"`
}

// type RolePermission struct {
// 	gorm.Model
// 	RoleID     uint64 `gorm:"index"`
// 	Permission string `gorm:"index"`
// }
