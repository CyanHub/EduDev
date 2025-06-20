package model

import (
	"FileSystem/global"
)

type User struct {
	global.GSModel
	ID       uint64 `gorm:"primaryKey;type:bigint unsigned;comment:用户ID"`
	Username string `gorm:"size:50;not null;unique;comment:用户登录名"`
	Password string `gorm:"size:100;not null;comment:用户登录密码"`
	NickName string `gorm:"size:50;default:'绳网用户';comment:用户昵称"`
	Avatar   string `gorm:"size:255;default:'https://tse4-mm.cn.bing.net/th/id/OIP-C.kNpVdcZ_O_KT6StMKhNn5gHaHa?r=0&rs=1&pid=ImgDetMain&cb=idpwebp1&o=7&rm=3';comment:用户头像"`
	Status   int8   `gorm:"type:tinyint(1);default:1;comment:用户状态"`
	Phone    string `gorm:"size:20;not null;comment:用户手机号"`
	Email    string `gorm:"size:100;not null;comment:用户邮箱"`
	RoleId   uint64 `gorm:"type:bigint unsigned;default:2;comment:用户角色ID"`
}

func (User) TableName() string {
	return "users" // 确保与数据库表名一致
}

// type User struct {
// 	ID        uint64         `gorm:"primaryKey;type:bigint unsigned"`
// 	Username  string         `json:"userName" gorm:"size:50;uniqueIndex;comment:用户登录名"`
// 	Password  string         `json:"-" gorm:"comment:用户登录密码"`
// 	NickName  string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
// 	HeaderImg string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
// 	RoleId    uint64         `json:"roleId" gorm:"default:888;comment:用户角色Id"`
// 	Phone     string         `json:"phone"  gorm:"comment:用户手机号"`
// 	Email     string         `json:"email"  gorm:"comment:用户邮箱"`
// 	Avatar    string         `json:"avatar" gorm:"comment:用户头像"`
// 	Enable    int8           `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
// 	Balance   float64        `json:"balance" gorm:"type:float;comment:用户余额"`
// 	CreatedAt time.Time      `gorm:"type:datetime;autoCreateTime"`  // 明确指定类型
// 	UpdatedAt time.Time      `gorm:"type:datetime;autoUpdateTime"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`                        // 软删除字段
// }
