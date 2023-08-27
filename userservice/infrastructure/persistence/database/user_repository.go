package database

import (
	"github.com/jun-chiang/simple-douyin/userservice/infrastructure/persistence/config"
	dbModels "github.com/jun-chiang/simple-douyin/userservice/infrastructure/persistence/models"
)

type UserRepository struct {
}

// 创建用户
func (repo *UserRepository) CreateUser(user *dbModels.User) (int64, error) {
	db := config.GetInstance().Conn
	err := db.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

// 根据用户名字查询用户
func (repo *UserRepository) QueryUserByName(userName string) (*dbModels.User, error) {
	db := config.GetInstance().Conn

	var user dbModels.User
	if err := db.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
