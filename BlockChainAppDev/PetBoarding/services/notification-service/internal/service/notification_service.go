package service

import (
	"errors"
	"time"

	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/notification-service/internal/model"
	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/notification-service/internal/repository"
)

// NotificationService 定义通知服务接口
type NotificationService interface {
	CreateNotification(req *model.CreateNotificationRequest) (*model.NotificationResponse, error)
	GetNotificationByID(id string) (*model.NotificationResponse, error)
	GetNotificationsByUserID(userID string) ([]*model.NotificationResponse, error)
	GetUnreadNotificationsByUserID(userID string) ([]*model.NotificationResponse, error)
	GetAllNotifications() ([]*model.NotificationResponse, error)
	MarkAsRead(id string) error
	MarkAllAsRead(userID string) error
	DeleteNotification(id string) error
}

// notificationService 实现通知服务接口
type notificationService struct {
	notificationRepo repository.NotificationRepository
}

// NewNotificationService 创建通知服务实例
func NewNotificationService(notificationRepo repository.NotificationRepository) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
	}
}

// CreateNotification 创建通知
func (s *notificationService) CreateNotification(req *model.CreateNotificationRequest) (*model.NotificationResponse, error) {
	// 验证请求
	if req.UserID == "" {
		return nil, errors.New("user ID is required")
	}
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
	if req.Content == "" {
		return nil, errors.New("content is required")
	}

	// 创建通知实体
	notification := &model.Notification{
		UserID:    req.UserID,
		Type:      req.Type,
		Title:     req.Title,
		Content:   req.Content,
		RelatedID: req.RelatedID,
		Read:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 保存通知
	err := s.notificationRepo.Create(notification)
	if err != nil {
		return nil, err
	}

	// 返回响应
	response := notification.ToResponse()
	return &response, nil
}

// GetNotificationByID 根据ID获取通知
func (s *notificationService) GetNotificationByID(id string) (*model.NotificationResponse, error) {
	// 获取通知
	notification, err := s.notificationRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 返回响应
	response := notification.ToResponse()
	return &response, nil
}

// GetNotificationsByUserID 获取用户的所有通知
func (s *notificationService) GetNotificationsByUserID(userID string) ([]*model.NotificationResponse, error) {
	// 获取通知
	notifications, err := s.notificationRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 转换为响应
	responses := make([]*model.NotificationResponse, len(notifications))
	for i, notification := range notifications {
		response := notification.ToResponse()
		responses[i] = &response
	}

	return responses, nil
}

// GetUnreadNotificationsByUserID 获取用户的未读通知
func (s *notificationService) GetUnreadNotificationsByUserID(userID string) ([]*model.NotificationResponse, error) {
	// 获取未读通知
	notifications, err := s.notificationRepo.GetUnreadByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 转换为响应
	responses := make([]*model.NotificationResponse, len(notifications))
	for i, notification := range notifications {
		response := notification.ToResponse()
		responses[i] = &response
	}

	return responses, nil
}

// GetAllNotifications 获取所有通知
func (s *notificationService) GetAllNotifications() ([]*model.NotificationResponse, error) {
	// 获取所有通知
	notifications, err := s.notificationRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// 转换为响应
	responses := make([]*model.NotificationResponse, len(notifications))
	for i, notification := range notifications {
		response := notification.ToResponse()
		responses[i] = &response
	}

	return responses, nil
}

// MarkAsRead 将通知标记为已读
func (s *notificationService) MarkAsRead(id string) error {
	// 标记为已读
	return s.notificationRepo.MarkAsRead(id)
}

// MarkAllAsRead 将用户的所有通知标记为已读
func (s *notificationService) MarkAllAsRead(userID string) error {
	// 标记所有为已读
	return s.notificationRepo.MarkAllAsRead(userID)
}

// DeleteNotification 删除通知
func (s *notificationService) DeleteNotification(id string) error {
	// 删除通知
	return s.notificationRepo.Delete(id)
}