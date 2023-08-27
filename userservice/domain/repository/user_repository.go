package repository

import (
	dbModels "github.com/jun-chiang/simple-douyin/userservice/infrastructure/persistence/models"
)

type UserRepository interface {
	// 创建用户
	CreateUser(user *dbModels.User) (int64, error)

	// 根据用户名字查询用户
	QueryUserByName(userName string) (*dbModels.User, error)
}
