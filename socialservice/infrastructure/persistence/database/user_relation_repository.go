package database

import (
	commonModels "github.com/jun-chiang/simple-douyin/common/models"
	"github.com/jun-chiang/simple-douyin/socialservice/infrastructure/persistence/config"
	dbModels "github.com/jun-chiang/simple-douyin/socialservice/infrastructure/persistence/models"
)

type UserRelationRepository struct {
}

// 根据用户ID查询用户的关注列表
func (repo *UserRelationRepository) GetFollowListByUserId(userId int64) (users []*commonModels.User, err error) {
	db := config.GetInstance().Conn
	err = db.Model(&dbModels.User{}).
		Joins("join user_relations as ur on ur.follow_to = users.id").
		Where("ur.follow_from = ?", userId).
		Find(&users).Error
	return
}

// 根据用户ID查询用户的粉丝列表
func (repo *UserRelationRepository) GetFollowerListByUserId(userId int64) (users []*commonModels.User, err error) {
	return
}
