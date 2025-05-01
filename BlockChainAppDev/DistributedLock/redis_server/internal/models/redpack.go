package models

import (
	"gorm.io/gorm"
	"time"
)

type Redpack struct {
	gorm.Model
	//RedpackId int64  `gorm:"column:redpack_id" json:"redpack_id"`
	Amount int `gorm:"column:amount" json:"amount"`
	//UserId    int64  `gorm:"column:user_id" json:"user_id"`
	Status    string `gorm:"column:status" json:"status"`
	Num       int    `gorm:"column:num" json:"num"`               // 红包数量
	ValidTime int    `gorm:"column:valid_time" json:"valid_time"` // 有效期
	ProNum    int    `gorm:"column:pro_num" json:"pro_num"`       // 已领取数量
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *Redpack) TableName() string {
	return "ts_redpack"
}
