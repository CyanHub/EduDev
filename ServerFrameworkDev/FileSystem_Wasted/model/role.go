package model

import "FileSystem/global"

type Role struct {
	global.GSModel
	Name     string 
	ParentId uint64
}

func (Role) TableName() string {
	return "role"
}
