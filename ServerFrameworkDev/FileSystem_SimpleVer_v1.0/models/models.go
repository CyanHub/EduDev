package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Roles    []Role `gorm:"many2many:user_roles;"`
	Files    []File `gorm:"foreignKey:OwnerID"`
}

type Role struct {
	gorm.Model
	Name        string       `gorm:"unique;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Code string `gorm:"unique;not null"`
}

type FileAccess struct {
	gorm.Model
	FileID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
}

// UserRole 定义用户与角色的关联模型
type UserRole struct {
	UserID    uint           `gorm:"index;not null"`
	RoleID    uint           `gorm:"index;not null"`
}

type File struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Path     string `gorm:"not null"`
	OwnerID  uint   `gorm:"not null"`
	Owner    User   `gorm:"foreignKey:OwnerID"`
	IsPublic bool   `gorm:"default:false"`
	Size     int64  `gorm:"not null"`
	Accesses []FileAccess `gorm:"foreignKey:FileID"`
}
