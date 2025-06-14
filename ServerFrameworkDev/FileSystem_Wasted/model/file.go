package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name       string `gorm:"size:255;not null"`
	Path       string `gorm:"size:512;not null"`
	Size       int64  `gorm:"not null"`
	Type       string `gorm:"size:50"`
	UserID     uint64 `gorm:"not null;index"`
	UploaderID uint64 `gorm:"not null;index"`
}

type FilePermission struct {
	gorm.Model
	FileID     uint64 `gorm:"index"`
	UserID     uint64 `gorm:"index"`
	Permission string `gorm:"size:50;index"`
}
