package service

import (
	"errors"

	"github.com/jun-chiang/simple-douyin/common/consts"
	"github.com/jun-chiang/simple-douyin/common/utils"
	user "github.com/jun-chiang/simple-douyin/userservice/domain/models"
	"github.com/jun-chiang/simple-douyin/userservice/domain/repository"
	dbModels "github.com/jun-chiang/simple-douyin/userservice/infrastructure/persistence/models"
)

type UserService struct {
	// 仓储接口
	repo repository.UserRepository
}

// 获取Service实例（依赖注入）
func NewUserService(repo repository.UserRepository) (service *UserService, err error) {
	if repo == nil {
		err = errors.New("repo can not be empty")
		return
	}
	service = &UserService{
		repo: repo,
	}
	return
}

// 用户注册
func (service *UserService) UserRegister(req *user.RegisterRequest) (int64, error) {
	user, err := service.repo.QueryUserByName(req.Username)
	if err != nil {
		return -1, err
	}

	if *user != (dbModels.User{}) {
		return -1, err
	}

	passWord, err := utils.Crypt(req.Password)
	user_id, err := service.repo.CreateUser(&dbModels.User{
		UserName:        req.Username,
		Password:        passWord,
		Avatar:          consts.TestAva,
		BackgroundImage: consts.TestBackground,
		Signature:       consts.TestSign,
	})
	return user_id, nil
}
