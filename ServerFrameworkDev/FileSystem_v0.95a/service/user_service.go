package service

import (
    "FileSystem/global"
    "FileSystem/model"
)


// GetRoleByID 根据角色 ID 获取角色信息
func (u *UserService) GetRoleByID(roleID uint64) (*model.Role, error) {
    var role model.Role
    err := global.DB.Where("id = ?", roleID).First(&role).Error
    if err != nil {
        return nil, err
    }
    return &role, nil
}
