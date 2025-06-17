package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/cyanhub/petboarding/services/pet-service/internal/model"
)

// PetRepository 宠物仓库接口
type PetRepository interface {
	Create(pet *model.Pet) error
	GetByID(id uint) (*model.Pet, error)
	GetByUserID(userID uint) ([]*model.Pet, error)
	GetAll(page, pageSize int) ([]*model.Pet, int64, error)
	Update(pet *model.Pet) error
	Delete(id uint) error
}

// petRepository 宠物仓库实现
type petRepository struct {
	pets  map[uint]*model.Pet
	mutex sync.RWMutex
	nextID uint
}

// NewPetRepository 创建宠物仓库实例
func NewPetRepository() PetRepository {
	return &petRepository{
		pets:   make(map[uint]*model.Pet),
		nextID: 1,
	}
}

// Create 创建宠物
func (r *petRepository) Create(pet *model.Pet) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 设置ID和时间
	pet.ID = r.nextID
	r.nextID++
	pet.CreatedAt = time.Now()
	pet.UpdatedAt = pet.CreatedAt

	// 存储宠物
	r.pets[pet.ID] = pet
	return nil
}

// GetByID 根据ID获取宠物
func (r *petRepository) GetByID(id uint) (*model.Pet, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	pet, exists := r.pets[id]
	if !exists {
		return nil, errors.New("宠物不存在")
	}

	return pet, nil
}

// GetByUserID 获取用户的所有宠物
func (r *petRepository) GetByUserID(userID uint) ([]*model.Pet, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var userPets []*model.Pet
	for _, pet := range r.pets {
		if pet.UserID == userID {
			userPets = append(userPets, pet)
		}
	}

	return userPets, nil
}

// GetAll 获取所有宠物（分页）
func (r *petRepository) GetAll(page, pageSize int) ([]*model.Pet, int64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// 计算总数
	total := int64(len(r.pets))

	// 计算分页
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= int(total) {
		return []*model.Pet{}, total, nil
	}
	if end > int(total) {
		end = int(total)
	}

	// 提取分页数据
	pets := make([]*model.Pet, 0, end-start)
	i := 0
	for _, pet := range r.pets {
		if i >= start && i < end {
			pets = append(pets, pet)
		}
		i++
		if i >= end {
			break
		}
	}

	return pets, total, nil
}

// Update 更新宠物
func (r *petRepository) Update(pet *model.Pet) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 检查宠物是否存在
	_, exists := r.pets[pet.ID]
	if !exists {
		return errors.New("宠物不存在")
	}

	// 更新时间
	pet.UpdatedAt = time.Now()

	// 更新宠物
	r.pets[pet.ID] = pet
	return nil
}

// Delete 删除宠物
func (r *petRepository) Delete(id uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 检查宠物是否存在
	_, exists := r.pets[id]
	if !exists {
		return errors.New("宠物不存在")
	}

	// 删除宠物
	delete(r.pets, id)
	return nil
}