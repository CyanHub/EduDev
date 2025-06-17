package repository

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/admin-service/internal/model"
)

// AdminRepository 定义管理员仓库接口
type AdminRepository interface {
	Create(admin *model.Admin) error
	GetByID(id string) (*model.Admin, error)
	GetByUsername(username string) (*model.Admin, error)
	GetByEmail(email string) (*model.Admin, error)
	GetAll() ([]*model.Admin, error)
	Update(admin *model.Admin) error
	Delete(id string) error
	ChangePassword(id, passwordHash string) error
}

// adminRepository 实现管理员仓库接口
type adminRepository struct {
	admins map[string]*model.Admin
	mutex  sync.RWMutex
	nextID int
}

// NewAdminRepository 创建管理员仓库实例
func NewAdminRepository() AdminRepository {
	return &adminRepository{
		admins: make(map[string]*model.Admin),
		nextID: 1,
	}
}

// Create 创建管理员
func (r *adminRepository) Create(admin *model.Admin) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 检查用户名是否已存在
	for _, existingAdmin := range r.admins {
		if existingAdmin.Username == admin.Username {
			return errors.New("username already exists")
		}
		if existingAdmin.Email == admin.Email {
			return errors.New("email already exists")
		}
	}

	// 生成ID
	admin.ID = generateID(r.nextID)
	r.nextID++

	// 设置时间戳
	now := time.Now()
	admin.CreatedAt = now
	admin.UpdatedAt = now

	// 存储管理员
	r.admins[admin.ID] = admin

	return nil
}

// GetByID 根据ID获取管理员
func (r *adminRepository) GetByID(id string) (*model.Admin, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	admin, exists := r.admins[id]
	if !exists {
		return nil, errors.New("admin not found")
	}

	return admin, nil
}

// GetByUsername 根据用户名获取管理员
func (r *adminRepository) GetByUsername(username string) (*model.Admin, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, admin := range r.admins {
		if admin.Username == username {
			return admin, nil
		}
	}

	return nil, errors.New("admin not found")
}

// GetByEmail 根据邮箱获取管理员
func (r *adminRepository) GetByEmail(email string) (*model.Admin, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, admin := range r.admins {
		if admin.Email == email {
			return admin, nil
		}
	}

	return nil, errors.New("admin not found")
}

// GetAll 获取所有管理员
func (r *adminRepository) GetAll() ([]*model.Admin, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	admins := make([]*model.Admin, 0, len(r.admins))
	for _, admin := range r.admins {
		admins = append(admins, admin)
	}

	return admins, nil
}

// Update 更新管理员
func (r *adminRepository) Update(admin *model.Admin) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 检查ID是否存在
	existingAdmin, exists := r.admins[admin.ID]
	if !exists {
		return errors.New("admin not found")
	}

	// 检查用户名是否已存在（如果更改了用户名）
	if admin.Username != existingAdmin.Username {
		for _, a := range r.admins {
			if a.ID != admin.ID && a.Username == admin.Username {
				return errors.New("username already exists")
			}
		}
	}

	// 检查邮箱是否已存在（如果更改了邮箱）
	if admin.Email != existingAdmin.Email {
		for _, a := range r.admins {
			if a.ID != admin.ID && a.Email == admin.Email {
				return errors.New("email already exists")
			}
		}
	}

	// 保留原始密码（如果没有提供新密码）
	if admin.Password == "" {
		admin.Password = existingAdmin.Password
	}

	// 更新时间戳
	admin.UpdatedAt = time.Now()
	// 保留创建时间
	admin.CreatedAt = existingAdmin.CreatedAt

	// 更新管理员
	r.admins[admin.ID] = admin

	return nil
}

// Delete 删除管理员
func (r *adminRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.admins[id]
	if !exists {
		return errors.New("admin not found")
	}

	delete(r.admins, id)

	return nil
}

// ChangePassword 更改管理员密码
func (r *adminRepository) ChangePassword(id, passwordHash string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	admin, exists := r.admins[id]
	if !exists {
		return errors.New("admin not found")
	}

	admin.Password = passwordHash
	admin.UpdatedAt = time.Now()

	return nil
}

// generateID 生成ID
func generateID(id int) string {
	return fmt.Sprintf("admin-%d", id)
}