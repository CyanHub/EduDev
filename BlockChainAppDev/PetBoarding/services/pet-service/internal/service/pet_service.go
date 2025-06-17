package service

import (
	"errors"

	"github.com/cyanhub/petboarding/services/pet-service/internal/model"
	"github.com/cyanhub/petboarding/services/pet-service/internal/repository"
)

// PetService 宠物服务接口
type PetService interface {
	CreatePet(req model.CreatePetRequest) (*model.PetResponse, error)
	GetPetByID(id uint) (*model.PetResponse, error)
	GetPetsByUserID(userID uint) ([]*model.PetResponse, error)
	GetAllPets(page, pageSize int) ([]*model.PetResponse, int64, error)
	UpdatePet(id uint, req model.UpdatePetRequest) (*model.PetResponse, error)
	DeletePet(id uint) error
}

// petService 宠物服务实现
type petService struct {
	petRepo repository.PetRepository
}

// NewPetService 创建宠物服务实例
func NewPetService(petRepo repository.PetRepository) PetService {
	return &petService{
		petRepo: petRepo,
	}
}

// CreatePet 创建宠物
func (s *petService) CreatePet(req model.CreatePetRequest) (*model.PetResponse, error) {
	// 验证请求数据
	if req.Name == "" {
		return nil, errors.New("宠物名称不能为空")
	}
	if req.Type == "" {
		return nil, errors.New("宠物类型不能为空")
	}
	if req.Age < 0 {
		return nil, errors.New("宠物年龄不能为负数")
	}
	if req.Weight <= 0 {
		return nil, errors.New("宠物体重必须大于0")
	}
	if req.UserID == 0 {
		return nil, errors.New("必须指定宠物所有者")
	}

	// 创建宠物对象
	pet := &model.Pet{
		Name:        req.Name,
		Type:        req.Type,
		Breed:       req.Breed,
		Age:         req.Age,
		Gender:      req.Gender,
		Weight:      req.Weight,
		Description: req.Description,
		UserID:      req.UserID,
	}

	// 保存宠物
	if err := s.petRepo.Create(pet); err != nil {
		return nil, err
	}

	// 返回响应
	response := pet.ToResponse()
	return &response, nil
}

// GetPetByID 根据ID获取宠物
func (s *petService) GetPetByID(id uint) (*model.PetResponse, error) {
	// 获取宠物
	pet, err := s.petRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 返回响应
	response := pet.ToResponse()
	return &response, nil
}

// GetPetsByUserID 获取用户的所有宠物
func (s *petService) GetPetsByUserID(userID uint) ([]*model.PetResponse, error) {
	// 获取宠物列表
	pets, err := s.petRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	responses := make([]*model.PetResponse, len(pets))
	for i, pet := range pets {
		resp := pet.ToResponse()
		responses[i] = &resp
	}

	return responses, nil
}

// GetAllPets 获取所有宠物（分页）
func (s *petService) GetAllPets(page, pageSize int) ([]*model.PetResponse, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 获取宠物列表
	pets, total, err := s.petRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	responses := make([]*model.PetResponse, len(pets))
	for i, pet := range pets {
		resp := pet.ToResponse()
		responses[i] = &resp
	}

	return responses, total, nil
}

// UpdatePet 更新宠物
func (s *petService) UpdatePet(id uint, req model.UpdatePetRequest) (*model.PetResponse, error) {
	// 获取现有宠物
	pet, err := s.petRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 更新字段
	if req.Name != nil {
		pet.Name = *req.Name
	}
	if req.Type != nil {
		pet.Type = *req.Type
	}
	if req.Breed != nil {
		pet.Breed = *req.Breed
	}
	if req.Age != nil {
		pet.Age = *req.Age
	}
	if req.Gender != nil {
		pet.Gender = *req.Gender
	}
	if req.Weight != nil {
		pet.Weight = *req.Weight
	}
	if req.Description != nil {
		pet.Description = *req.Description
	}

	// 保存更新
	if err := s.petRepo.Update(pet); err != nil {
		return nil, err
	}

	// 返回响应
	response := pet.ToResponse()
	return &response, nil
}

// DeletePet 删除宠物
func (s *petService) DeletePet(id uint) error {
	return s.petRepo.Delete(id)
}