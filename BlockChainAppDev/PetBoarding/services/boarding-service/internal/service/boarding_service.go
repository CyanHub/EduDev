package service

import (
	"errors"
	"time"

	"github.com/cyanhub/petboarding/services/boarding-service/internal/model"
	"github.com/cyanhub/petboarding/services/boarding-service/internal/repository"
)

// BoardingService 预订服务接口
type BoardingService interface {
	CreateBoarding(req model.CreateBoardingRequest) (*model.BoardingResponse, error)
	GetBoardingByID(id uint) (*model.BoardingResponse, error)
	GetBoardingsByUserID(userID uint) ([]*model.BoardingResponse, error)
	GetBoardingsByPetID(petID uint) ([]*model.BoardingResponse, error)
	GetAllBoardings(page, pageSize int) ([]*model.BoardingResponse, int64, error)
	UpdateBoarding(id uint, req model.UpdateBoardingRequest) (*model.BoardingResponse, error)
	UpdateBoardingStatus(id uint, req model.UpdateBoardingStatusRequest) (*model.BoardingResponse, error)
	DeleteBoarding(id uint) error
	CheckAvailability(startDate, endDate time.Time) (bool, error)
	GetServicePrices() []model.ServicePrice
}

// boardingService 预订服务实现
type boardingService struct {
	boardingRepo repository.BoardingRepository
}

// NewBoardingService 创建预订服务实例
func NewBoardingService(boardingRepo repository.BoardingRepository) BoardingService {
	return &boardingService{
		boardingRepo: boardingRepo,
	}
}

// CreateBoarding 创建预订
func (s *boardingService) CreateBoarding(req model.CreateBoardingRequest) (*model.BoardingResponse, error) {
	// 验证请求数据
	if req.UserID == 0 {
		return nil, errors.New("用户ID不能为空")
	}
	if req.PetID == 0 {
		return nil, errors.New("宠物ID不能为空")
	}
	if req.StartDate.IsZero() || req.EndDate.IsZero() {
		return nil, errors.New("开始日期和结束日期不能为空")
	}
	if req.StartDate.After(req.EndDate) {
		return nil, errors.New("开始日期不能晚于结束日期")
	}
	if req.StartDate.Before(time.Now()) {
		return nil, errors.New("开始日期不能早于当前日期")
	}

	// 检查服务类型是否有效
	servicePrice, err := model.GetServicePrice(req.ServiceType)
	if err != nil {
		return nil, errors.New("无效的服务类型")
	}

	// 检查日期是否可用
	available, err := s.CheckAvailability(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, errors.New("所选日期已被预订，请选择其他日期")
	}

	// 计算总价
	duration := int(req.EndDate.Sub(req.StartDate).Hours() / 24)
	if duration < 1 {
		duration = 1
	}
	totalPrice := servicePrice * float64(duration)

	// 创建预订对象
	boarding := &model.Boarding{
		UserID:      req.UserID,
		PetID:       req.PetID,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Status:      model.StatusPending,
		Notes:       req.Notes,
		TotalPrice:  totalPrice,
		ServiceType: req.ServiceType,
	}

	// 保存预订
	if err := s.boardingRepo.Create(boarding); err != nil {
		return nil, err
	}

	// 返回响应
	response := boarding.ToResponse()
	return &response, nil
}

// GetBoardingByID 根据ID获取预订
func (s *boardingService) GetBoardingByID(id uint) (*model.BoardingResponse, error) {
	// 获取预订
	boarding, err := s.boardingRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 返回响应
	response := boarding.ToResponse()
	return &response, nil
}

// GetBoardingsByUserID 获取用户的所有预订
func (s *boardingService) GetBoardingsByUserID(userID uint) ([]*model.BoardingResponse, error) {
	// 获取预订列表
	boardings, err := s.boardingRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	responses := make([]*model.BoardingResponse, len(boardings))
	for i, boarding := range boardings {
		resp := boarding.ToResponse()
		responses[i] = &resp
	}

	return responses, nil
}

// GetBoardingsByPetID 获取宠物的所有预订
func (s *boardingService) GetBoardingsByPetID(petID uint) ([]*model.BoardingResponse, error) {
	// 获取预订列表
	boardings, err := s.boardingRepo.GetByPetID(petID)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	responses := make([]*model.BoardingResponse, len(boardings))
	for i, boarding := range boardings {
		resp := boarding.ToResponse()
		responses[i] = &resp
	}

	return responses, nil
}

// GetAllBoardings 获取所有预订（分页）
func (s *boardingService) GetAllBoardings(page, pageSize int) ([]*model.BoardingResponse, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 获取预订列表
	boardings, total, err := s.boardingRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	responses := make([]*model.BoardingResponse, len(boardings))
	for i, boarding := range boardings {
		resp := boarding.ToResponse()
		responses[i] = &resp
	}

	return responses, total, nil
}

// UpdateBoarding 更新预订
func (s *boardingService) UpdateBoarding(id uint, req model.UpdateBoardingRequest) (*model.BoardingResponse, error) {
	// 获取现有预订
	boarding, err := s.boardingRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 检查是否可以更新
	if boarding.Status == model.StatusCancelled || boarding.Status == model.StatusCompleted {
		return nil, errors.New("已取消或已完成的预订不能修改")
	}

	// 更新字段
	if req.StartDate != nil && req.EndDate != nil {
		// 如果同时更新开始和结束日期，需要验证
		if req.StartDate.After(*req.EndDate) {
			return nil, errors.New("开始日期不能晚于结束日期")
		}
		if req.StartDate.Before(time.Now()) {
			return nil, errors.New("开始日期不能早于当前日期")
		}

		// 检查新日期是否可用（排除当前预订）
		available, err := s.checkAvailabilityExcluding(*req.StartDate, *req.EndDate, id)
		if err != nil {
			return nil, err
		}
		if !available {
			return nil, errors.New("所选日期已被预订，请选择其他日期")
		}

		boarding.StartDate = *req.StartDate
		boarding.EndDate = *req.EndDate

		// 重新计算价格
		servicePrice, _ := model.GetServicePrice(boarding.ServiceType)
		duration := int(boarding.EndDate.Sub(boarding.StartDate).Hours() / 24)
		if duration < 1 {
			duration = 1
		}
		boarding.TotalPrice = servicePrice * float64(duration)
	} else if req.StartDate != nil {
		// 只更新开始日期
		if req.StartDate.After(boarding.EndDate) {
			return nil, errors.New("开始日期不能晚于结束日期")
		}
		if req.StartDate.Before(time.Now()) {
			return nil, errors.New("开始日期不能早于当前日期")
		}

		// 检查新日期是否可用
		available, err := s.checkAvailabilityExcluding(*req.StartDate, boarding.EndDate, id)
		if err != nil {
			return nil, err
		}
		if !available {
			return nil, errors.New("所选日期已被预订，请选择其他日期")
		}

		boarding.StartDate = *req.StartDate

		// 重新计算价格
		servicePrice, _ := model.GetServicePrice(boarding.ServiceType)
		duration := int(boarding.EndDate.Sub(boarding.StartDate).Hours() / 24)
		if duration < 1 {
			duration = 1
		}
		boarding.TotalPrice = servicePrice * float64(duration)
	} else if req.EndDate != nil {
		// 只更新结束日期
		if boarding.StartDate.After(*req.EndDate) {
			return nil, errors.New("开始日期不能晚于结束日期")
		}

		// 检查新日期是否可用
		available, err := s.checkAvailabilityExcluding(boarding.StartDate, *req.EndDate, id)
		if err != nil {
			return nil, err
		}
		if !available {
			return nil, errors.New("所选日期已被预订，请选择其他日期")
		}

		boarding.EndDate = *req.EndDate

		// 重新计算价格
		servicePrice, _ := model.GetServicePrice(boarding.ServiceType)
		duration := int(boarding.EndDate.Sub(boarding.StartDate).Hours() / 24)
		if duration < 1 {
			duration = 1
		}
		boarding.TotalPrice = servicePrice * float64(duration)
	}

	if req.Status != nil {
		boarding.Status = *req.Status
	}

	if req.Notes != nil {
		boarding.Notes = *req.Notes
	}

	if req.ServiceType != nil {
		// 检查服务类型是否有效
		servicePrice, err := model.GetServicePrice(*req.ServiceType)
		if err != nil {
			return nil, errors.New("无效的服务类型")
		}

		boarding.ServiceType = *req.ServiceType

		// 重新计算价格
		duration := int(boarding.EndDate.Sub(boarding.StartDate).Hours() / 24)
		if duration < 1 {
			duration = 1
		}
		boarding.TotalPrice = servicePrice * float64(duration)
	}

	// 保存更新
	if err := s.boardingRepo.Update(boarding); err != nil {
		return nil, err
	}

	// 返回响应
	response := boarding.ToResponse()
	return &response, nil
}

// UpdateBoardingStatus 更新预订状态
func (s *boardingService) UpdateBoardingStatus(id uint, req model.UpdateBoardingStatusRequest) (*model.BoardingResponse, error) {
	// 获取现有预订
	boarding, err := s.boardingRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 验证状态转换
	if !isValidStatusTransition(boarding.Status, req.Status) {
		return nil, errors.New("无效的状态转换")
	}

	// 更新状态
	boarding.Status = req.Status

	// 保存更新
	if err := s.boardingRepo.Update(boarding); err != nil {
		return nil, err
	}

	// 返回响应
	response := boarding.ToResponse()
	return &response, nil
}

// DeleteBoarding 删除预订
func (s *boardingService) DeleteBoarding(id uint) error {
	// 获取现有预订
	boarding, err := s.boardingRepo.GetByID(id)
	if err != nil {
		return err
	}

	// 只允许删除待确认或已取消的预订
	if boarding.Status != model.StatusPending && boarding.Status != model.StatusCancelled {
		return errors.New("只能删除待确认或已取消的预订")
	}

	return s.boardingRepo.Delete(id)
}

// CheckAvailability 检查日期是否可用
func (s *boardingService) CheckAvailability(startDate, endDate time.Time) (bool, error) {
	// 获取指定日期范围内的预订
	existingBookings, err := s.boardingRepo.GetByDateRange(startDate, endDate)
	if err != nil {
		return false, err
	}

	// 检查是否有冲突的预订
	for _, booking := range existingBookings {
		// 只考虑已确认的预订
		if booking.Status == model.StatusConfirmed {
			return false, nil
		}
	}

	return true, nil
}

// checkAvailabilityExcluding 检查日期是否可用（排除指定ID的预订）
func (s *boardingService) checkAvailabilityExcluding(startDate, endDate time.Time, excludeID uint) (bool, error) {
	// 获取指定日期范围内的预订
	existingBookings, err := s.boardingRepo.GetByDateRange(startDate, endDate)
	if err != nil {
		return false, err
	}

	// 检查是否有冲突的预订（排除指定ID）
	for _, booking := range existingBookings {
		if booking.ID != excludeID && booking.Status == model.StatusConfirmed {
			return false, nil
		}
	}

	return true, nil
}

// GetServicePrices 获取服务价格列表
func (s *boardingService) GetServicePrices() []model.ServicePrice {
	return model.GetServicePrices()
}

// isValidStatusTransition 检查状态转换是否有效
func isValidStatusTransition(from, to model.BoardingStatus) bool {
	switch from {
	case model.StatusPending:
		// 待确认 -> 已确认、已取消
		return to == model.StatusConfirmed || to == model.StatusCancelled
	case model.StatusConfirmed:
		// 已确认 -> 已完成、已取消
		return to == model.StatusCompleted || to == model.StatusCancelled
	case model.StatusCancelled:
		// 已取消 -> 不能转换到其他状态
		return false
	case model.StatusCompleted:
		// 已完成 -> 不能转换到其他状态
		return false
	default:
		return false
	}
}