package model

import "time"

// BoardingStatus 预订状态枚举
type BoardingStatus string

const (
	StatusPending   BoardingStatus = "pending"   // 待确认
	StatusConfirmed BoardingStatus = "confirmed" // 已确认
	StatusCancelled BoardingStatus = "cancelled" // 已取消
	StatusCompleted BoardingStatus = "completed" // 已完成
)

// Boarding 宠物寄养预订模型
type Boarding struct {
	ID          uint          `json:"id" gorm:"primaryKey"`
	UserID      uint          `json:"user_id" gorm:"not null"`      // 用户ID
	PetID       uint          `json:"pet_id" gorm:"not null"`       // 宠物ID
	StartDate   time.Time     `json:"start_date" gorm:"not null"`   // 开始日期
	EndDate     time.Time     `json:"end_date" gorm:"not null"`     // 结束日期
	Status      BoardingStatus `json:"status" gorm:"not null"`      // 预订状态
	Notes       string        `json:"notes" gorm:"type:text"`       // 备注
	TotalPrice  float64       `json:"total_price" gorm:"not null"` // 总价格
	ServiceType string        `json:"service_type" gorm:"size:50"` // 服务类型（基础寄养、豪华寄养等）
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

// BoardingResponse 预订响应模型
type BoardingResponse struct {
	ID          uint          `json:"id"`
	UserID      uint          `json:"user_id"`
	PetID       uint          `json:"pet_id"`
	StartDate   time.Time     `json:"start_date"`
	EndDate     time.Time     `json:"end_date"`
	Status      BoardingStatus `json:"status"`
	Notes       string        `json:"notes"`
	TotalPrice  float64       `json:"total_price"`
	ServiceType string        `json:"service_type"`
	Duration    int           `json:"duration"` // 寄养天数
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

// CreateBoardingRequest 创建预订请求
type CreateBoardingRequest struct {
	UserID      uint      `json:"user_id" binding:"required"`
	PetID       uint      `json:"pet_id" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	Notes       string    `json:"notes"`
	ServiceType string    `json:"service_type" binding:"required"`
}

// UpdateBoardingRequest 更新预订请求
type UpdateBoardingRequest struct {
	StartDate   *time.Time     `json:"start_date"`
	EndDate     *time.Time     `json:"end_date"`
	Status      *BoardingStatus `json:"status"`
	Notes       *string        `json:"notes"`
	ServiceType *string        `json:"service_type"`
}

// UpdateBoardingStatusRequest 更新预订状态请求
type UpdateBoardingStatusRequest struct {
	Status BoardingStatus `json:"status" binding:"required"`
}

// ToResponse 将Boarding模型转换为响应模型
func (b *Boarding) ToResponse() BoardingResponse {
	// 计算寄养天数
	duration := int(b.EndDate.Sub(b.StartDate).Hours() / 24)
	if duration < 1 {
		duration = 1
	}

	return BoardingResponse{
		ID:          b.ID,
		UserID:      b.UserID,
		PetID:       b.PetID,
		StartDate:   b.StartDate,
		EndDate:     b.EndDate,
		Status:      b.Status,
		Notes:       b.Notes,
		TotalPrice:  b.TotalPrice,
		ServiceType: b.ServiceType,
		Duration:    duration,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}

// ServicePrice 服务价格配置
type ServicePrice struct {
	ServiceType string  `json:"service_type"`
	DailyPrice  float64 `json:"daily_price"`
	Description string  `json:"description"`
}

// GetServicePrices 获取服务价格列表
func GetServicePrices() []ServicePrice {
	return []ServicePrice{
		{
			ServiceType: "basic",
			DailyPrice:  100.0,
			Description: "基础寄养服务，包含基本食宿和每日两次遛狗",
		},
		{
			ServiceType: "premium",
			DailyPrice:  200.0,
			Description: "高级寄养服务，包含优质食宿、每日三次遛狗和基础洗护",
		},
		{
			ServiceType: "deluxe",
			DailyPrice:  300.0,
			Description: "豪华寄养服务，包含顶级食宿、专人陪伴、每日洗护和训练",
		},
	}
}

// GetServicePrice 根据服务类型获取价格
func GetServicePrice(serviceType string) (float64, error) {
	for _, price := range GetServicePrices() {
		if price.ServiceType == serviceType {
			return price.DailyPrice, nil
		}
	}
	return 0, nil
}