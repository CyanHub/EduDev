package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id" json:"user_id"`
	UserName string `gorm:"column:user_name" json:"user_name"`
}

func (u *User) TableName() string {
	return "ts_user"
}
