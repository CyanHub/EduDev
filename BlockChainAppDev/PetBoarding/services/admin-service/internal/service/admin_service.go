package service

import (
	"errors"
	"time"

	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/admin-service/internal/model"
	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/admin-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// AdminService 定义管理员服务接口
type AdminService interface {
	CreateAdmin(req *model.CreateAdminRequest) (*model.AdminResponse, error)
	GetAdminByID(id string) (*model.AdminResponse, error)
	GetAllAdmins() ([]*model.AdminResponse, error)
	UpdateAdmin(id string, req *model.UpdateAdminRequest) (*model.AdminResponse, error)
	DeleteAdmin(id string) error
	ChangePassword(id string, req *model.ChangePasswordRequest) error
	Login(req *model.LoginRequest) (*model.LoginResponse, error)
	GetDashboardData() (*model.DashboardData, error)
}

// adminService 实现管理员服务接口
type adminService struct {
	adminRepo repository.AdminRepository
	// 这里可以添加其他服务的客户端，用于获取统计数据
}

// NewAdminService 创建管理员服务实例
func NewAdminService(adminRepo repository.AdminRepository) AdminService {
	return &adminService{
		adminRepo: adminRepo,
	}
}

// CreateAdmin 创建管理员
func (s *adminService) CreateAdmin(req *model.CreateAdminRequest) (*model.AdminResponse, error) {
	// 验证请求
	if req.Username == "" {
		return nil, errors.New("username is required")
	}
	if req.Password == "" {
		return nil, errors.New("password is required")
	}
	if req.Email == "" {
		return nil, errors.New("email is required")
	}
	if req.Role == "" {
		return nil, errors.New("role is required")
	}

	// 检查用户名是否已存在
	_, err := s.adminRepo.GetByUsername(req.Username)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	// 检查邮箱是否已存在
	_, err = s.adminRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建管理员实体
	admin := &model.Admin{
		Username:  req.Username,
		Password:  string(hashedPassword),
		Email:     req.Email,
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 保存管理员
	err = s.adminRepo.Create(admin)
	if err != nil {
		return nil, err
	}

	// 返回响应
	response := admin.ToResponse()
	return &response, nil
}

// GetAdminByID 根据ID获取管理员
func (s *adminService) GetAdminByID(id string) (*model.AdminResponse, error) {
	// 获取管理员
	admin, err := s.adminRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 返回响应
	response := admin.ToResponse()
	return &response, nil
}

// GetAllAdmins 获取所有管理员
func (s *adminService) GetAllAdmins() ([]*model.AdminResponse, error) {
	// 获取所有管理员
	admins, err := s.adminRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// 转换为响应
	responses := make([]*model.AdminResponse, len(admins))
	for i, admin := range admins {
		response := admin.ToResponse()
		responses[i] = &response
	}

	return responses, nil
}

// UpdateAdmin 更新管理员
func (s *adminService) UpdateAdmin(id string, req *model.UpdateAdminRequest) (*model.AdminResponse, error) {
	// 获取管理员
	admin, err := s.adminRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 更新字段
	if req.Username != "" {
		// 检查用户名是否已存在
		existingAdmin, err := s.adminRepo.GetByUsername(req.Username)
		if err == nil && existingAdmin.ID != id {
			return nil, errors.New("username already exists")
		}
		admin.Username = req.Username
	}

	if req.Email != "" {
		// 检查邮箱是否已存在
		existingAdmin, err := s.adminRepo.GetByEmail(req.Email)
		if err == nil && existingAdmin.ID != id {
			return nil, errors.New("email already exists")
		}
		admin.Email = req.Email
	}

	if req.Role != "" {
		admin.Role = req.Role
	}

	// 更新管理员
	err = s.adminRepo.Update(admin)
	if err != nil {
		return nil, err
	}

	// 返回响应
	response := admin.ToResponse()
	return &response, nil
}

// DeleteAdmin 删除管理员
func (s *adminService) DeleteAdmin(id string) error {
	// 删除管理员
	return s.adminRepo.Delete(id)
}

// ChangePassword 更改管理员密码
func (s *adminService) ChangePassword(id string, req *model.ChangePasswordRequest) error {
	// 获取管理员
	admin, err := s.adminRepo.GetByID(id)
	if err != nil {
		return err
	}

	// 验证当前密码
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.CurrentPassword))
	if err != nil {
		return errors.New("current password is incorrect")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	return s.adminRepo.ChangePassword(id, string(hashedPassword))
}

// Login 管理员登录
func (s *adminService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	// 获取管理员
	admin, err := s.adminRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 生成令牌（实际项目中应使用JWT或其他认证机制）
	token := "sample-token-" + admin.ID // 这只是一个示例，实际应用中应使用JWT

	// 返回响应
	response := &model.LoginResponse{
		Token: token,
		Admin: admin.ToResponse(),
	}

	return response, nil
}

// GetDashboardData 获取仪表板数据
func (s *adminService) GetDashboardData() (*model.DashboardData, error) {
	// 在实际应用中，这里应该调用其他服务的API获取数据
	// 这里只是返回一些示例数据

	// 系统统计信息
	stats := model.SystemStats{
		TotalUsers:     100,
		TotalPets:      150,
		TotalBoardings: 80,
		TotalReviews:   60,
		ActiveBookings: 20,
		Revenue:        5000.0,
		AverageRating:  4.5,
	}

	// 按月收入
	revenueByMonth := map[string]float64{
		"2023-01": 1000.0,
		"2023-02": 1200.0,
		"2023-03": 1500.0,
		"2023-04": 1300.0,
		"2023-05": 1800.0,
		"2023-06": 2000.0,
	}

	// 按月预订
	bookingsByMonth := map[string]int{
		"2023-01": 10,
		"2023-02": 12,
		"2023-03": 15,
		"2023-04": 13,
		"2023-05": 18,
		"2023-06": 20,
	}

	// 仪表板数据
	dashboardData := &model.DashboardData{
		Stats:           stats,
		RecentBookings:  []string{"booking1", "booking2", "booking3"}, // 示例数据
		RecentReviews:   []string{"review1", "review2", "review3"},   // 示例数据
		RevenueByMonth:  revenueByMonth,
		BookingsByMonth: bookingsByMonth,
	}

	return dashboardData, nil
}