package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID uint64 `gorm:"index"`
	RoleID uint64 `gorm:"index"`
}

type RolePermission struct {
	gorm.Model
	RoleID     uint64 `gorm:"index"`
	Permission string `gorm:"index"`
}
