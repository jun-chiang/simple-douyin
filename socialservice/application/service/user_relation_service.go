package service

import (
	"errors"

	commonModels "github.com/jun-chiang/simple-douyin/common/models"
	"github.com/jun-chiang/simple-douyin/socialservice/domain/repository"
)

type UserRelationService struct {
	// 仓储接口
	repo repository.UserRelationRepository
}

// 获取Service实例（依赖注入）
func NewUserRelationService(repo repository.UserRelationRepository) (service *UserRelationService, err error) {
	if repo == nil {
		err = errors.New("repo can not be empty")
		return
	}
	service = &UserRelationService{
		repo: repo,
	}
	return
}

// 获取用户的关注列表
func (service *UserRelationService) GetFollowListByUserId(userId int64) (users []*commonModels.User, err error) {
	// 从数据库读取数据
	users, err = service.repo.GetFollowListByUserId(userId)
	return
}

// 获取用户的关注列表
func (service *UserRelationService) GetFollowerListByUserId(userId int64) (users []*commonModels.User, err error) {
	// 从数据库读取数据
	users, err = service.repo.GetFollowerListByUserId(userId)
	return
}
