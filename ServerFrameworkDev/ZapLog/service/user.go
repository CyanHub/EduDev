package service

import (
	"github.com/CyanHub/EduDev/global"
	"github.com/CyanHub/EduDev/model"
	"github.com/CyanHub/EduDev/model/request"
)

var UserServiceApp = new(UserService)

type UserService struct{}

func (u *UserService) Login(req request.UserLoginRequest) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return nil, global.ErrUserNotFound
	}
	if user.Password != req.Password {
		return nil, global.ErrPasswordIncorrect
	}
	return &user, nil
}

func (u *UserService) Register(req request.UserRegisterRequest) (*model.User, error) {
	var user model.User
	var count int64
	// 检查用户是否存在
	err := global.DB.Model(&model.User{}).Where("username = ?", req.Username).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, global.ErrUserAlreadyExists
	}
	user.Username = req.Username
	user.Password = req.Password
	user.NickName = req.NickName
	err = global.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) UserList(req request.UserListRequest) (total int64, data []*model.User, err error) {
	var users []*model.User
	db := global.DB.Model(&model.User{})
	if req.Username != "" {
		db = db.Where("username = ?", req.Username)
	}
	if req.NickName != "" {
		db = db.Where("nick_name = ?", req.NickName)
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}
	if req.RoleId != 0 {
		db = db.Where("role_id = ?", req.RoleId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = db.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find(&users).Error
	if err != nil {
		return 0, nil, err
	}
	return total, users, nil
}
