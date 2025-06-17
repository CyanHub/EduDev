package service

import (
	"errors"
	"time"

	"github.com/cyanhub/petboarding/services/user-service/internal/model"
	"github.com/cyanhub/petboarding/services/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService 定义用户服务接口
type UserService interface {
	Register(req model.CreateUserRequest) (*model.UserResponse, error)
	Login(req model.LoginRequest) (*model.LoginResponse, error)
	GetUserByID(id uint) (*model.UserResponse, error)
	GetAllUsers(page, pageSize int) ([]model.UserResponse, int64, error)
	UpdateUser(id uint, req model.UpdateUserRequest) (*model.UserResponse, error)
	DeleteUser(id uint) error
}

// userService 实现UserService接口
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register 注册新用户
func (s *userService) Register(req model.CreateUserRequest) (*model.UserResponse, error) {
	// 检查用户名是否已存在
	_, err := s.userRepo.GetByUsername(req.Username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	_, err = s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("邮箱已存在")
	}

	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建用户
	now := time.Now()
	user := &model.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		FullName:  req.FullName,
		Phone:     req.Phone,
		Address:   req.Address,
		Role:      "user", // 默认角色
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

// Login 用户登录
func (s *userService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	// 根据用户名获取用户
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成JWT令牌（简化版，实际应用中应该使用JWT库）
	token := "mock-jwt-token-" + user.Username

	return &model.LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	}, nil
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(id uint) (*model.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

// GetAllUsers 获取所有用户（分页）
func (s *userService) GetAllUsers(page, pageSize int) ([]model.UserResponse, int64, error) {
	users, total, err := s.userRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	return responses, total, nil
}

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(id uint, req model.UpdateUserRequest) (*model.UserResponse, error) {
	// 获取现有用户
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 更新字段（如果提供了）
	if req.Username != "" && req.Username != user.Username {
		// 检查新用户名是否已存在
		_, err := s.userRepo.GetByUsername(req.Username)
		if err == nil {
			return nil, errors.New("用户名已存在")
		}
		user.Username = req.Username
	}

	if req.Email != "" && req.Email != user.Email {
		// 检查新邮箱是否已存在
		_, err := s.userRepo.GetByEmail(req.Email)
		if err == nil {
			return nil, errors.New("邮箱已存在")
		}
		user.Email = req.Email
	}

	if req.Password != "" {
		// 哈希新密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("密码加密失败")
		}
		user.Password = string(hashedPassword)
	}

	if req.FullName != "" {
		user.FullName = req.FullName
	}

	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if req.Address != "" {
		user.Address = req.Address
	}

	user.UpdatedAt = time.Now()

	// 保存更新
	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}