package models

type User struct {
	ID              int64  `gorm:"id"`
	UserName        string `gorm:"user_name"`
	Password        string `gorm:"password"`
	Avatar          string `gorm:"avatar"`
	BackgroundImage string `gorm:"background_image"`
	Signature       string `gorm:"signature"`
}
