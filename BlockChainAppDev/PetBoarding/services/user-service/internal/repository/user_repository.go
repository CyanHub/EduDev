package repository

import (
	"errors"
	"sync"

	"github.com/cyanhub/petboarding/services/user-service/internal/model"
	"gorm.io/gorm"
)

// UserRepository 定义用户仓库接口
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll(page, pageSize int) ([]model.User, int64, error)
	Update(user *model.User) error
	Delete(id uint) error
}

// userRepository 实现UserRepository接口
type userRepository struct {
	db *gorm.DB
	// 临时使用内存存储，后续会替换为数据库
	users     map[uint]*model.User
	nextID    uint
	userMutex sync.RWMutex
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository() UserRepository {
	// 初始化一些测试数据
	users := make(map[uint]*model.User)
	users[1] = &model.User{
		ID:       1,
		Username: "user1",
		Email:    "user1@example.com",
		Password: "password1", // 实际应用中应该存储哈希值
		Role:     "user",
	}
	users[2] = &model.User{
		ID:       2,
		Username: "user2",
		Email:    "user2@example.com",
		Password: "password2",
		Role:     "user",
	}

	return &userRepository{
		users:  users,
		nextID: 3,
	}
}

// Create 创建新用户
func (r *userRepository) Create(user *model.User) error {
	r.userMutex.Lock()
	defer r.userMutex.Unlock()

	// 检查用户名是否已存在
	for _, existingUser := range r.users {
		if existingUser.Username == user.Username {
			return errors.New("用户名已存在")
		}
		if existingUser.Email == user.Email {
			return errors.New("邮箱已存在")
		}
	}

	// 设置ID并保存
	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user

	return nil
}

// GetByID 根据ID获取用户
func (r *userRepository) GetByID(id uint) (*model.User, error) {
	r.userMutex.RLock()
	defer r.userMutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("用户不存在")
	}

	return user, nil
}

// GetByUsername 根据用户名获取用户
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	r.userMutex.RLock()
	defer r.userMutex.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("用户不存在")
}

// GetByEmail 根据邮箱获取用户
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	r.userMutex.RLock()
	defer r.userMutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("用户不存在")
}

// GetAll 获取所有用户（分页）
func (r *userRepository) GetAll(page, pageSize int) ([]model.User, int64, error) {
	r.userMutex.RLock()
	defer r.userMutex.RUnlock()

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	totalCount := int64(len(r.users))
	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= int(totalCount) {
		return []model.User{}, totalCount, nil
	}

	if end > int(totalCount) {
		end = int(totalCount)
	}

	result := make([]model.User, 0, end-start)
	current := 0

	for _, user := range r.users {
		if current >= start && current < end {
			result = append(result, *user)
		}
		current++
		if current >= end {
			break
		}
	}

	return result, totalCount, nil
}

// Update 更新用户信息
func (r *userRepository) Update(user *model.User) error {
	r.userMutex.Lock()
	defer r.userMutex.Unlock()

	_, exists := r.users[user.ID]
	if !exists {
		return errors.New("用户不存在")
	}

	// 检查用户名和邮箱是否与其他用户冲突
	for id, existingUser := range r.users {
		if id == user.ID {
			continue
		}
		if existingUser.Username == user.Username {
			return errors.New("用户名已存在")
		}
		if existingUser.Email == user.Email {
			return errors.New("邮箱已存在")
		}
	}

	r.users[user.ID] = user
	return nil
}

// Delete 删除用户
func (r *userRepository) Delete(id uint) error {
	r.userMutex.Lock()
	defer r.userMutex.Unlock()

	_, exists := r.users[id]
	if !exists {
		return errors.New("用户不存在")
	}

	delete(r.users, id)
	return nil
}