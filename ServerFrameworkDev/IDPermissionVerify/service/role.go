package service

import (
	"IDPVerify/global"
	"IDPVerify/model"
	"IDPVerify/model/request"
)

var RoleServiceApp = new(RoleService)

type RoleService struct {
}

func (r *RoleService) RoleList(req request.RoleListRequest) (total int64, list []*model.Role, err error) {
	db := global.DB.Model(&model.Role{})
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.ParentId != 0 {
		db = db.Where("parent_id = ?", req.ParentId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = db.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find(&list).Error
	if err != nil {
		return 0, nil, err
	}
	return total, list, nil
}

func (r *RoleService) RoleCreate(req request.RoleCreateRequest) (err error) {
	var c int64
	err = global.DB.Model(&model.Role{}).Where("name = ?", req.Name).Count(&c).Error
	if err != nil {
		return err
	}
	if c > 0 {
		return global.ErrRoleAlreadyExists
	}
	role := model.Role{
		Name:     req.Name,
		ParentId: req.ParentId,
	}
	err = global.DB.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

