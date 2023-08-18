package models

import (
	"time"

	"gorm.io/gorm"
)

// 基础模型字段
type GormBase struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
