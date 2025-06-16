package service

import (
	"FileSystem/global"
	"FileSystem/model"
	"FileSystem/model/request"
	"errors"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

func (u *UserService) GetUserInfo(d uint64) (any, error) {
	var user model.User
	if err := global.DB.Where("id = ?", d).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, global.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *UserService) UsernameExists(username string) bool {
	var count int64
	global.DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

var UserServiceApp = new(UserService)

func (u *UserService) Login(req request.UserLoginRequest) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, global.ErrUserNotFound
		}
		return nil, err
	}

	// 修改为bcrypt密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, global.ErrPasswordIncorrect
	}
	return &user, nil
}

// 添加用户注册逻辑（实验三中的密码加密）
func (u *UserService) Register(req request.UserRegisterRequest) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

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
		global.Logger.Error("数据库写入失败",
			zap.Error(err),
			zap.Any("user", user))
		return nil, err
	}

	global.Logger.Info("用户注册成功",
		zap.Uint64("id", user.ID),
		zap.String("username", user.Username))
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

func (s *UserService) CreateUser(user *model.User) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// 分配默认角色
		var defaultRole model.Role
		if err := tx.Where("name = ?", "普通用户").First(&defaultRole).Error; err != nil {
			return err
		}

		return tx.Model(user).Association("Roles").Append(&defaultRole)
	})
}

func (s *UserService) CheckPermission(userID uint64, permissionCode string) bool {
	var count int64
	global.DB.Model(&model.User{}).
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN role_permissions ON role_permissions.role_id = user_roles.role_id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("users.id = ? AND permissions.code = ?", userID, permissionCode).
		Count(&count)
	return count > 0
}

// 在UserService结构体下添加
func (u *UserService) IsAdmin(userID uint64) bool {
	var user model.User
	if err := global.DB.Select("role_id").First(&user, userID).Error; err != nil {
		return false
	}
	return user.RoleId == 1 // 假设1是管理员角色ID
}

func (u *UserService) GetUserByID(userID uint64) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, global.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// 分页查询用户列表
func (s *UserService) GetUserList(page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := global.DB.Model(&model.User{})
	db.Count(&total)

	if err := db.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
