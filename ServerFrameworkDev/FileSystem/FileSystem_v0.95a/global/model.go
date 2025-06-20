package global

import (
	"time"
)

type GSModel struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`             // 主键ID
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"` // 更新时间
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}
