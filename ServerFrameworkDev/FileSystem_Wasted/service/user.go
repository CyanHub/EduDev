package service

import (
	"FileSystem/global"
	"FileSystem/model"
	"FileSystem/model/request"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (u *UserService) Login(req request.UserLoginRequest) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return nil, global.ErrUserNotFound
	}

	// 修改为bcrypt密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, global.ErrPasswordIncorrect
	}
	return &user, nil
}

func (u *UserService) Register(req request.UserRegisterRequest) (*model.User, error) {
	// 检查用户名是否已存在
	var count int64
	if err := global.DB.Model(&model.User{}).
		Where("username = ?", req.Username).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, global.ErrUserAlreadyExists // 在注册时检查用户名是否已存在，替换了 errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 创建用户
	user := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		NickName: req.NickName,
		Email:    req.Email,
		Phone:    req.Phone,  // 添加手机号字段
		RoleId:   2,          // 默认普通用户角色
		Avatar:   req.Avatar, // 确保avatar字段正确
	}

	if err := global.DB.Create(&user).Error; err != nil {
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

func (u *UserService) UpdateUser(req request.UserUpdateRequest) (*model.User, error) {
	var user model.User
	if err := global.DB.First(&user, req.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, global.ErrUserNotFound
		}
		return nil, err
	}

	// 更新用户信息
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	if err := global.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
