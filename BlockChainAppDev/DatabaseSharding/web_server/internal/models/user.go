package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"gorm_._model"`
	UserId     int64  `gorm:"column:user_id" json:"user_id" json:"user_id,omitempty"`
	UserName   string `gorm:"column:user_name" json:"user_name" json:"user_name,omitempty"`
	UserMobile string `gorm:"column:user_mobile" json:"user_mobile,omitempty"`
	UserEmail  string `gorm:"column:user_email" json:"user_email,omitempty"`
	UserPwd    string `gorm:"column:user_pwd" json:"user_pwd,omitempty"`
	CreatedAt  string `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt  string `gorm:"column:updated_at" json:"updated_at,omitempty"`
	//DeletedAt  string `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (u *User) TableName() string {
	return "ts_user"
}
