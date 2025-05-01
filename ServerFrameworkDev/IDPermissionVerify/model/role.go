package model

import "IDPVerify/global"

type Role struct {
	global.GSModel
	Name     string `json:"name" gorm:"type:varchar(255);uniqueIndex;comment:角色名称"`
	ParentId uint64 `json:"parentId" gorm:"default:0;comment:父级角色Id"`
}

func (Role) TableName() string {
	return "role"
}
