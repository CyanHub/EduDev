package model

import (
	"time"
)

// NotificationType 表示通知类型
type NotificationType string

const (
	BookingConfirmed  NotificationType = "BOOKING_CONFIRMED"
	BookingCancelled  NotificationType = "BOOKING_CANCELLED"
	BookingCompleted  NotificationType = "BOOKING_COMPLETED"
	PaymentReceived   NotificationType = "PAYMENT_RECEIVED"
	ReviewReminder    NotificationType = "REVIEW_REMINDER"
	SystemAnnouncement NotificationType = "SYSTEM_ANNOUNCEMENT"
)

// Notification 表示通知实体
type Notification struct {
	ID          string          `json:"id" gorm:"primaryKey"`
	UserID      string          `json:"userId" gorm:"index"`
	Type        NotificationType `json:"type"`
	Title       string          `json:"title"`
	Content     string          `json:"content"`
	Read        bool            `json:"read" gorm:"default:false"`
	RelatedID   string          `json:"relatedId,omitempty"` // 可能关联的预订ID、评论ID等
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

// NotificationResponse 表示通知响应
type NotificationResponse struct {
	ID          string          `json:"id"`
	UserID      string          `json:"userId"`
	Type        NotificationType `json:"type"`
	Title       string          `json:"title"`
	Content     string          `json:"content"`
	Read        bool            `json:"read"`
	RelatedID   string          `json:"relatedId,omitempty"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

// CreateNotificationRequest 表示创建通知请求
type CreateNotificationRequest struct {
	UserID      string          `json:"userId" binding:"required"`
	Type        NotificationType `json:"type" binding:"required"`
	Title       string          `json:"title" binding:"required"`
	Content     string          `json:"content" binding:"required"`
	RelatedID   string          `json:"relatedId,omitempty"`
}

// UpdateNotificationRequest 表示更新通知请求
type UpdateNotificationRequest struct {
	Read        bool            `json:"read"`
}

// ToResponse 将通知实体转换为响应
func (n *Notification) ToResponse() NotificationResponse {
	return NotificationResponse{
		ID:        n.ID,
		UserID:    n.UserID,
		Type:      n.Type,
		Title:     n.Title,
		Content:   n.Content,
		Read:      n.Read,
		RelatedID: n.RelatedID,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}