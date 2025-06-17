package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/cyanhub/petboarding/services/review-service/internal/model"
)

// ReviewRepository 评论仓库接口
type ReviewRepository interface {
	Create(review *model.Review) error
	GetByID(id uint) (*model.Review, error)
	GetByUserID(userID uint) ([]*model.Review, error)
	GetByBoardingID(boardingID uint) ([]*model.Review, error)
	GetAll(page, pageSize int) ([]*model.Review, int64, error)
	Update(review *model.Review) error
	Delete(id uint) error
	GetSummary() (*model.ReviewSummary, error)
}

// reviewRepository 评论仓库实现
type reviewRepository struct {
	reviews map[uint]*model.Review
	mutex   sync.RWMutex
	nextID  uint
}

// NewReviewRepository 创建评论仓库实例
func NewReviewRepository() ReviewRepository {
	return &reviewRepository{
		reviews: make(map[uint]*model.Review),
		nextID:  1,
	}
}

// Create 创建评论
func (r *reviewRepository) Create(review *model.Review) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 设置ID和时间
	review.ID = r.nextID
	r.nextID++
	review.CreatedAt = time.Now()
	review.UpdatedAt = review.CreatedAt

	// 存储评论
	r.reviews[review.ID] = review
	return nil
}

// GetByID 根据ID获取评论
func (r *reviewRepository) GetByID(id uint) (*model.Review, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	review, exists := r.reviews[id]
	if !exists {
		return nil, errors.New("评论不存在")
	}

	return review, nil
}

// GetByUserID 获取用户的所有评论
func (r *reviewRepository) GetByUserID(userID uint) ([]*model.Review, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var userReviews []*model.Review
	for _, review := range r.reviews {
		if review.UserID == userID {
			userReviews = append(userReviews, review)
		}
	}

	return userReviews, nil
}

// GetByBoardingID 获取预订的所有评论
func (r *reviewRepository) GetByBoardingID(boardingID uint) ([]*model.Review, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var boardingReviews []*model.Review
	for _, review := range r.reviews {
		if review.BoardingID == boardingID {
			boardingReviews = append(boardingReviews, review)
		}
	}

	return boardingReviews, nil
}

// GetAll 获取所有评论（分页）
func (r *reviewRepository) GetAll(page, pageSize int) ([]*model.Review, int64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// 计算总数
	total := int64(len(r.reviews))

	// 计算分页
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= int(total) {
		return []*model.Review{}, total, nil
	}
	if end > int(total) {
		end = int(total)
	}

	// 提取分页数据
	reviews := make([]*model.Review, 0, end-start)
	i := 0
	for _, review := range r.reviews {
		if i >= start && i < end {
			reviews = append(reviews, review)
		}
		i++
		if i >= end {
			break
		}
	}

	return reviews, total, nil
}

// Update 更新评论
func (r *reviewRepository) Update(review *model.Review) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 检查评论是否存在
	_, exists := r.reviews[review.ID]
	if !exists {
		return errors.New("评论不存在")
	}

	// 更新时间
	review.UpdatedAt = time.Now()

	// 更新评论
	r.reviews[review.ID] = review
	return nil
}

// Delete 删除评论
func (r *reviewRepository) Delete(id uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 检查评论是否存在
	_, exists := r.reviews[id]
	if !exists {
		return errors.New("评论不存在")
	}

	// 删除评论
	delete(r.reviews, id)
	return nil
}

// GetSummary 获取评论汇总
func (r *reviewRepository) GetSummary() (*model.ReviewSummary, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	totalReviews := len(r.reviews)
	if totalReviews == 0 {
		return &model.ReviewSummary{
			TotalReviews:  0,
			AverageRating: 0,
			RatingCounts:  make(map[int]int),
		}, nil
	}

	// 计算评分统计
	sum := 0
	ratingCounts := make(map[int]int)
	for i := 1; i <= 5; i++ {
		ratingCounts[i] = 0
	}

	for _, review := range r.reviews {
		sum += review.Rating
		ratingCounts[review.Rating]++
	}

	// 计算平均评分
	averageRating := float64(sum) / float64(totalReviews)

	return &model.ReviewSummary{
		TotalReviews:  totalReviews,
		AverageRating: averageRating,
		RatingCounts:  ratingCounts,
	}, nil
}