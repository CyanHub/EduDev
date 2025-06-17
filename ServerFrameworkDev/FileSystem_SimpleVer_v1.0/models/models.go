package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Roles    []Role `gorm:"many2many:user_roles;"`
	Email    string `gorm:"unique;not null"`
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

type File struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Path     string `gorm:"not null"`
	OwnerID  uint   `gorm:"not null"`
	Owner    User   `gorm:"foreignKey:OwnerID"`
	IsPublic bool   `gorm:"default:false"`
	Size     int64  `gorm:"not null"`
	Accesses []FileAccess
}

type FileAccess struct {
	gorm.Model
	FileID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
}
