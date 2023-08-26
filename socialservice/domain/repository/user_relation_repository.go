package repository

import "github.com/jun-chiang/simple-douyin/common/models"

type UserRelationRepository interface {
	// 根据用户ID查询用户的关注列表
	GetFollowListByUserId(userId int64) ([]*models.User, error)
	// 根据用户ID查询用户的粉丝列表
	GetFollowerListByUserId(userId int64) ([]*models.User, error)
}
