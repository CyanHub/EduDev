package model

import (
    "gorm.io/gorm"
)

// Subject 表示课程信息的结构体。
type Subject struct {
    gorm.Model
    Name       string                 `gorm:"column:name"`
    Tags       []string               `gorm:"type:text;serializer:json"` // 使用 JSON 序列化
    Syllabus   []string               `gorm:"type:text;serializer:json"` // 使用 JSON 序列化
    Properties map[string]interface{} `gorm:"type:text;serializer:json"` // 使用 JSON 序列化
}
