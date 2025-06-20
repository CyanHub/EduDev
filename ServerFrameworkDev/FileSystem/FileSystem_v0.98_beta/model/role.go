package model

import (
	"FileSystem/global"
)

type Role struct {
	global.GSModel
	ParentID    uint64       `gorm:"index"`
	Permissions []Permission `gorm:"many2many:role_permissions;comment:角色权限"`
	Name        string       `gorm:"unique;size:50;comment:角色名称"`
	Description string       `gorm:"size:255;comment:角色描述"`
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

// TableName 指定表名为 roles
func (Role) TableName() string {
	return "roles"
}
