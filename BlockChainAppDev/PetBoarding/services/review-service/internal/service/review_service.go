package service

import (
	"errors"

	"github.com/cyanhub/petboarding/services/review-service/internal/model"
	"github.com/cyanhub/petboarding/services/review-service/internal/repository"
)

// ReviewService 评论服务接口
type ReviewService interface {
	CreateReview(req model.CreateReviewRequest) (*model.ReviewResponse, error)
	GetReviewByID(id uint) (*model.ReviewResponse, error)
	GetReviewsByUserID(userID uint) ([]*model.ReviewResponse, error)
	GetReviewsByBoardingID(boardingID uint) ([]*model.ReviewResponse, error)
	GetAllReviews(page, pageSize int) ([]*model.ReviewResponse, int64, error)
	UpdateReview(id uint, req model.UpdateReviewRequest) (*model.ReviewResponse, error)
	DeleteReview(id uint) error
	GetReviewSummary() (*model.ReviewSummary, error)
}

// reviewService 评论服务实现
type reviewService struct {
	reviewRepo repository.ReviewRepository
}

// NewReviewService 创建评论服务实例
func NewReviewService(reviewRepo repository.ReviewRepository) ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
	}
}

// CreateReview 创建评论
func (s *reviewService) CreateReview(req model.CreateReviewRequest) (*model.ReviewResponse, error) {
	// 验证请求数据
	if req.UserID == 0 {
		return nil, errors.New("用户ID不能为空")
	}
	if req.BoardingID == 0 {
		return nil, errors.New("预订ID不能为空")
	}
	if req.Rating < 1 || req.Rating > 5 {
		return nil, errors.New("评分必须在1-5之间")
	}

	// 创建评论对象
	review := &model.Review{
		UserID:     req.UserID,
		BoardingID: req.BoardingID,
		Rating:     req.Rating,
		Content:    req.Content,
		Images:     req.Images,
	}

	// 保存评论
	if err := s.reviewRepo.Create(review); err != nil {
		return nil, err
	}

	// 返回响应
	response := review.ToResponse()
	return &response, nil
}

// GetReviewByID 根据ID获取评论
func (s *reviewService) GetReviewByID(id uint) (*model.ReviewResponse, error) {
	// 获取评论
	review, err := s.reviewRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 返回响应
	response := review.ToResponse()
	return &response, nil
}

// GetReviewsByUserID 获取用户的所有评论
func (s *reviewService) GetReviewsByUserID(userID uint) ([]*model.ReviewResponse, error) {
	// 获取评论列表
	reviews, err := s.reviewRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	responses := make([]*model.ReviewResponse, len(reviews))
	for i, review := range reviews {
		resp := review.ToResponse()
		responses[i] = &resp
	}

	return responses, nil
}

// GetReviewsByBoardingID 获取预订的所有评论
func (s *reviewService) GetReviewsByBoardingID(boardingID uint) ([]*model.ReviewResponse, error) {
	// 获取评论列表
	reviews, err := s.reviewRepo.GetByBoardingID(boardingID)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	responses := make([]*model.ReviewResponse, len(reviews))
	for i, review := range reviews {
		resp := review.ToResponse()
		responses[i] = &resp
	}

	return responses, nil
}

// GetAllReviews 获取所有评论（分页）
func (s *reviewService) GetAllReviews(page, pageSize int) ([]*model.ReviewResponse, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 获取评论列表
	reviews, total, err := s.reviewRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	responses := make([]*model.ReviewResponse, len(reviews))
	for i, review := range reviews {
		resp := review.ToResponse()
		responses[i] = &resp
	}

	return responses, total, nil
}

// UpdateReview 更新评论
func (s *reviewService) UpdateReview(id uint, req model.UpdateReviewRequest) (*model.ReviewResponse, error) {
	// 获取现有评论
	review, err := s.reviewRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 更新字段
	if req.Rating != nil {
		if *req.Rating < 1 || *req.Rating > 5 {
			return nil, errors.New("评分必须在1-5之间")
		}
		review.Rating = *req.Rating
	}

	if req.Content != nil {
		review.Content = *req.Content
	}

	if req.Images != nil {
		review.Images = *req.Images
	}

	// 保存更新
	if err := s.reviewRepo.Update(review); err != nil {
		return nil, err
	}

	// 返回响应
	response := review.ToResponse()
	return &response, nil
}

// DeleteReview 删除评论
func (s *reviewService) DeleteReview(id uint) error {
	return s.reviewRepo.Delete(id)
}

// GetReviewSummary 获取评论汇总
func (s *reviewService) GetReviewSummary() (*model.ReviewSummary, error) {
	return s.reviewRepo.GetSummary()
}