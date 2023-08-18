package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FromUserID uint   `gorm:"column:from_user_id" json:"from_user_id"`
	ToUserID   uint   `gorm:"column:to_user_id" json:"to_user_id"`
	Content    string `gorm:"column:content" json:"content"`
	CreateTime uint   `json:"create_time"`
}
