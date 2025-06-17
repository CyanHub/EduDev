package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/cyanhub/petboarding/services/boarding-service/internal/model"
)

// BoardingRepository 预订仓库接口
type BoardingRepository interface {
	Create(boarding *model.Boarding) error
	GetByID(id uint) (*model.Boarding, error)
	GetByUserID(userID uint) ([]*model.Boarding, error)
	GetByPetID(petID uint) ([]*model.Boarding, error)
	GetAll(page, pageSize int) ([]*model.Boarding, int64, error)
	Update(boarding *model.Boarding) error
	Delete(id uint) error
	GetByDateRange(startDate, endDate time.Time) ([]*model.Boarding, error)
}

// boardingRepository 预订仓库实现
type boardingRepository struct {
	boardings map[uint]*model.Boarding
	mutex     sync.RWMutex
	nextID    uint
}

// NewBoardingRepository 创建预订仓库实例
func NewBoardingRepository() BoardingRepository {
	return &boardingRepository{
		boardings: make(map[uint]*model.Boarding),
		nextID:    1,
	}
}

// Create 创建预订
func (r *boardingRepository) Create(boarding *model.Boarding) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 设置ID和时间
	boarding.ID = r.nextID
	r.nextID++
	boarding.CreatedAt = time.Now()
	boarding.UpdatedAt = boarding.CreatedAt

	// 存储预订
	r.boardings[boarding.ID] = boarding
	return nil
}

// GetByID 根据ID获取预订
func (r *boardingRepository) GetByID(id uint) (*model.Boarding, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	boarding, exists := r.boardings[id]
	if !exists {
		return nil, errors.New("预订不存在")
	}

	return boarding, nil
}

// GetByUserID 获取用户的所有预订
func (r *boardingRepository) GetByUserID(userID uint) ([]*model.Boarding, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var userBoardings []*model.Boarding
	for _, boarding := range r.boardings {
		if boarding.UserID == userID {
			userBoardings = append(userBoardings, boarding)
		}
	}

	return userBoardings, nil
}

// GetByPetID 获取宠物的所有预订
func (r *boardingRepository) GetByPetID(petID uint) ([]*model.Boarding, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var petBoardings []*model.Boarding
	for _, boarding := range r.boardings {
		if boarding.PetID == petID {
			petBoardings = append(petBoardings, boarding)
		}
	}

	return petBoardings, nil
}

// GetAll 获取所有预订（分页）
func (r *boardingRepository) GetAll(page, pageSize int) ([]*model.Boarding, int64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// 计算总数
	total := int64(len(r.boardings))

	// 计算分页
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= int(total) {
		return []*model.Boarding{}, total, nil
	}
	if end > int(total) {
		end = int(total)
	}

	// 提取分页数据
	boardings := make([]*model.Boarding, 0, end-start)
	i := 0
	for _, boarding := range r.boardings {
		if i >= start && i < end {
			boardings = append(boardings, boarding)
		}
		i++
		if i >= end {
			break
		}
	}

	return boardings, total, nil
}

// Update 更新预订
func (r *boardingRepository) Update(boarding *model.Boarding) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 检查预订是否存在
	_, exists := r.boardings[boarding.ID]
	if !exists {
		return errors.New("预订不存在")
	}

	// 更新时间
	boarding.UpdatedAt = time.Now()

	// 更新预订
	r.boardings[boarding.ID] = boarding
	return nil
}

// Delete 删除预订
func (r *boardingRepository) Delete(id uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 检查预订是否存在
	_, exists := r.boardings[id]
	if !exists {
		return errors.New("预订不存在")
	}

	// 删除预订
	delete(r.boardings, id)
	return nil
}

// GetByDateRange 获取指定日期范围内的预订
func (r *boardingRepository) GetByDateRange(startDate, endDate time.Time) ([]*model.Boarding, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var rangeBookings []*model.Boarding
	for _, boarding := range r.boardings {
		// 检查预订日期是否与指定范围重叠
		if (boarding.StartDate.Before(endDate) || boarding.StartDate.Equal(endDate)) &&
			(boarding.EndDate.After(startDate) || boarding.EndDate.Equal(startDate)) {
			rangeBookings = append(rangeBookings, boarding)
		}
	}

	return rangeBookings, nil
}