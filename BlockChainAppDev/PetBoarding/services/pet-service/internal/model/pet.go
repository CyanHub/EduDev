package model

import "time"

// Pet 宠物模型
type Pet struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:100;not null"`
	Type        string    `json:"type" gorm:"size:50;not null"` // 例如：猫、狗、兔子等
	Breed       string    `json:"breed" gorm:"size:100"`        // 品种
	Age         int       `json:"age"`                           // 年龄（月）
	Gender      string    `json:"gender" gorm:"size:10"`         // 性别：公/母
	Weight      float64   `json:"weight"`                        // 体重（kg）
	Description string    `json:"description" gorm:"type:text"`  // 描述
	UserID      uint      `json:"user_id" gorm:"not null"`      // 所有者ID
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PetResponse 宠物响应模型
type PetResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Breed       string    `json:"breed"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	Weight      float64   `json:"weight"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreatePetRequest 创建宠物请求
type CreatePetRequest struct {
	Name        string  `json:"name" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Breed       string  `json:"breed"`
	Age         int     `json:"age" binding:"required,gte=0"`
	Gender      string  `json:"gender" binding:"required"`
	Weight      float64 `json:"weight" binding:"required,gt=0"`
	Description string  `json:"description"`
	UserID      uint    `json:"user_id" binding:"required"`
}

// UpdatePetRequest 更新宠物请求
type UpdatePetRequest struct {
	Name        *string  `json:"name"`
	Type        *string  `json:"type"`
	Breed       *string  `json:"breed"`
	Age         *int     `json:"age" binding:"omitempty,gte=0"`
	Gender      *string  `json:"gender"`
	Weight      *float64 `json:"weight" binding:"omitempty,gt=0"`
	Description *string  `json:"description"`
}

// ToResponse 将Pet模型转换为响应模型
func (p *Pet) ToResponse() PetResponse {
	return PetResponse{
		ID:          p.ID,
		Name:        p.Name,
		Type:        p.Type,
		Breed:       p.Breed,
		Age:         p.Age,
		Gender:      p.Gender,
		Weight:      p.Weight,
		Description: p.Description,
		UserID:      p.UserID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}