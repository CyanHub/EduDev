package model

import (
	"FileSystem/global"
	"time"

	"gorm.io/gorm"
)

type User struct {
	global.GSModel
	Username  string         `json:"userName" gorm:"size:50;uniqueIndex;comment:用户登录名"`
	Password  string         `json:"-" gorm:"comment:用户登录密码"`
	NickName  string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	RoleId    uint64         `json:"roleId" gorm:"default:888;comment:用户角色Id"`
	Phone     string         `json:"phone"  gorm:"comment:用户手机号"`
	Email     string         `json:"email"  gorm:"comment:用户邮箱"`
	Avatar    string         `json:"avatar" gorm:"comment:用户头像"`
	Enable    int8           `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	Balance   float64        `json:"balance" gorm:"type:float;comment:用户余额"`
	CreatedAt time.Time      `gorm:"type:datetime;autoCreateTime"`  // 明确指定类型
	UpdatedAt time.Time      `gorm:"type:datetime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`                        // 软删除字段
}

func (User) TableName() string {
	return "user"
}
