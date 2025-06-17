package model

import "time"

// Review 评论模型
type Review struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null"`      // 用户ID
	BoardingID  uint      `json:"boarding_id" gorm:"not null"` // 预订ID
	Rating      int       `json:"rating" gorm:"not null"`       // 评分（1-5）
	Content     string    `json:"content" gorm:"type:text"`     // 评论内容
	Images      []string  `json:"images" gorm:"-"`              // 图片URL列表
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ReviewResponse 评论响应模型
type ReviewResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	BoardingID  uint      `json:"boarding_id"`
	Rating      int       `json:"rating"`
	Content     string    `json:"content"`
	Images      []string  `json:"images,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateReviewRequest 创建评论请求
type CreateReviewRequest struct {
	UserID     uint     `json:"user_id" binding:"required"`
	BoardingID uint     `json:"boarding_id" binding:"required"`
	Rating     int      `json:"rating" binding:"required,min=1,max=5"`
	Content    string   `json:"content"`
	Images     []string `json:"images,omitempty"`
}

// UpdateReviewRequest 更新评论请求
type UpdateReviewRequest struct {
	Rating  *int      `json:"rating" binding:"omitempty,min=1,max=5"`
	Content *string   `json:"content"`
	Images  *[]string `json:"images,omitempty"`
}

// ToResponse 将Review模型转换为响应模型
func (r *Review) ToResponse() ReviewResponse {
	return ReviewResponse{
		ID:         r.ID,
		UserID:     r.UserID,
		BoardingID: r.BoardingID,
		Rating:     r.Rating,
		Content:    r.Content,
		Images:     r.Images,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

// ReviewSummary 评论汇总
type ReviewSummary struct {
	TotalReviews int     `json:"total_reviews"`
	AverageRating float64 `json:"average_rating"`
	RatingCounts  map[int]int `json:"rating_counts"` // 各评分的数量统计
}