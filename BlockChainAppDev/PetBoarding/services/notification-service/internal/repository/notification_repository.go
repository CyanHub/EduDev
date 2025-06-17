package repository

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/CyanHub/EduDev/BlockChainAppDev/PetBoarding/services/notification-service/internal/model"
)

// NotificationRepository 定义通知仓库接口
type NotificationRepository interface {
	Create(notification *model.Notification) error
	GetByID(id string) (*model.Notification, error)
	GetByUserID(userID string) ([]*model.Notification, error)
	GetUnreadByUserID(userID string) ([]*model.Notification, error)
	GetAll() ([]*model.Notification, error)
	Update(notification *model.Notification) error
	MarkAsRead(id string) error
	MarkAllAsRead(userID string) error
	Delete(id string) error
}

// notificationRepository 实现通知仓库接口
type notificationRepository struct {
	notifications map[string]*model.Notification
	mutex         sync.RWMutex
	nextID        int
}

// NewNotificationRepository 创建通知仓库实例
func NewNotificationRepository() NotificationRepository {
	return &notificationRepository{
		notifications: make(map[string]*model.Notification),
		nextID:        1,
	}
}

// Create 创建通知
func (r *notificationRepository) Create(notification *model.Notification) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 生成ID
	notification.ID = generateID(r.nextID)
	r.nextID++

	// 设置时间戳
	now := time.Now()
	notification.CreatedAt = now
	notification.UpdatedAt = now

	// 存储通知
	r.notifications[notification.ID] = notification

	return nil
}

// GetByID 根据ID获取通知
func (r *notificationRepository) GetByID(id string) (*model.Notification, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	notification, exists := r.notifications[id]
	if !exists {
		return nil, errors.New("notification not found")
	}

	return notification, nil
}

// GetByUserID 获取用户的所有通知
func (r *notificationRepository) GetByUserID(userID string) ([]*model.Notification, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var userNotifications []*model.Notification

	for _, notification := range r.notifications {
		if notification.UserID == userID {
			userNotifications = append(userNotifications, notification)
		}
	}

	return userNotifications, nil
}

// GetUnreadByUserID 获取用户的未读通知
func (r *notificationRepository) GetUnreadByUserID(userID string) ([]*model.Notification, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var unreadNotifications []*model.Notification

	for _, notification := range r.notifications {
		if notification.UserID == userID && !notification.Read {
			unreadNotifications = append(unreadNotifications, notification)
		}
	}

	return unreadNotifications, nil
}

// GetAll 获取所有通知
func (r *notificationRepository) GetAll() ([]*model.Notification, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	notifications := make([]*model.Notification, 0, len(r.notifications))
	for _, notification := range r.notifications {
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

// Update 更新通知
func (r *notificationRepository) Update(notification *model.Notification) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.notifications[notification.ID]
	if !exists {
		return errors.New("notification not found")
	}

	// 更新时间戳
	notification.UpdatedAt = time.Now()

	// 更新通知
	r.notifications[notification.ID] = notification

	return nil
}

// MarkAsRead 将通知标记为已读
func (r *notificationRepository) MarkAsRead(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	notification, exists := r.notifications[id]
	if !exists {
		return errors.New("notification not found")
	}

	notification.Read = true
	notification.UpdatedAt = time.Now()

	return nil
}

// MarkAllAsRead 将用户的所有通知标记为已读
func (r *notificationRepository) MarkAllAsRead(userID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, notification := range r.notifications {
		if notification.UserID == userID {
			notification.Read = true
			notification.UpdatedAt = time.Now()
		}
	}

	return nil
}

// Delete 删除通知
func (r *notificationRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.notifications[id]
	if !exists {
		return errors.New("notification not found")
	}

	delete(r.notifications, id)

	return nil
}

// generateID 生成ID
func generateID(id int) string {
	return "notification-" + fmt.Sprintf("%d", id)
}