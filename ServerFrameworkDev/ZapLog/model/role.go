package model

import "github.com/CyanHub/EduDev/global"

type Role struct {
	global.GSModel
	Name     string 
	ParentId uint64
}

func (Role) TableName() string {
	return "role"
}
