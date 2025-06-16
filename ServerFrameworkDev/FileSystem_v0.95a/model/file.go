package model

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID         uint64 `gorm:"primaryKey;type:bigint unsigned"`
	Name       string `gorm:"size:255;not null"`
	Path       string `gorm:"size:500;not null"`
	Size       int64  `gorm:"not null"`
	Type       string `gorm:"size:50"`
	UploaderID uint64 `gorm:"type:bigint unsigned;not null"`
	IsPublic   bool   `gorm:"type:tinyint(1);default:0"`
	ExpireTime *time.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

type FilePermission struct {
	gorm.Model
	FileID     uint64 `gorm:"index"`
	UserID     uint64 `gorm:"index"`
	Permission string `gorm:"size:50;index"`
}
