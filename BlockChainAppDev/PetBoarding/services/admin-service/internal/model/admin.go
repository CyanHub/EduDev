package model

import (
	"time"
)

// Admin 表示管理员实体
type Admin struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"-" gorm:"column:password_hash"` // 不在JSON中暴露密码
	Email     string    `json:"email" gorm:"unique"`
	Role      string    `json:"role"` // 角色：SUPER_ADMIN, ADMIN, MODERATOR
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// AdminResponse 表示管理员响应
type AdminResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateAdminRequest 表示创建管理员请求
type CreateAdminRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required"`
}

// UpdateAdminRequest 表示更新管理员请求
type UpdateAdminRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Role     string `json:"role,omitempty"`
}

// ChangePasswordRequest 表示更改密码请求
type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
}

// LoginRequest 表示登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 表示登录响应
type LoginResponse struct {
	Token  string         `json:"token"`
	Admin  AdminResponse  `json:"admin"`
}

// ToResponse 将管理员实体转换为响应
func (a *Admin) ToResponse() AdminResponse {
	return AdminResponse{
		ID:        a.ID,
		Username:  a.Username,
		Email:     a.Email,
		Role:      a.Role,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

// SystemStats 表示系统统计信息
type SystemStats struct {
	TotalUsers     int `json:"totalUsers"`
	TotalPets      int `json:"totalPets"`
	TotalBoardings int `json:"totalBoardings"`
	TotalReviews   int `json:"totalReviews"`
	ActiveBookings int `json:"activeBookings"`
	Revenue        float64 `json:"revenue"`
	AverageRating  float64 `json:"averageRating"`
}

// DashboardData 表示管理员仪表板数据
type DashboardData struct {
	Stats           SystemStats `json:"stats"`
	RecentBookings  interface{} `json:"recentBookings"`
	RecentReviews   interface{} `json:"recentReviews"`
	RevenueByMonth  map[string]float64 `json:"revenueByMonth"`
	BookingsByMonth map[string]int `json:"bookingsByMonth"`
}