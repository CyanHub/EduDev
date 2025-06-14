package service

import (
	"FileSystem/global"
	"FileSystem/model"
	"FileSystem/model/request"

	"gorm.io/gorm"
)

var RoleServiceApp = new(RoleService)

type RoleService struct{}

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

// DeleteRolePermissions 删除角色权限
func (r *RoleService) DeleteRolePermissions(tx *gorm.DB, roleID uint64) error {
	return tx.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error
}

// AddRolePermissions 添加角色权限
func (r *RoleService) AddRolePermissions(tx *gorm.DB, roleID uint64, permissions []string) error {
	var rolePermissions []model.RolePermission
	for _, permission := range permissions {
		rolePermissions = append(rolePermissions, model.RolePermission{
			RoleID:     roleID,
			Permission: permission,
		})
	}
	return tx.Create(&rolePermissions).Error
}

// CheckPermission 检查权限
func (r *RoleService) CheckPermission(userID uint64, permission string) (bool, error) {
	var count int64
	err := global.DB.Model(&model.UserRole{}).
		Joins("JOIN role_permissions ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ? AND role_permissions.permission = ?", userID, permission).
		Count(&count).Error

	return count > 0, err
}
