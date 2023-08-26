package converters

import "github.com/jun-chiang/simple-douyin/common/models"

type UserConverter struct {
}

func NewUserConverter() *UserConverter {
	return &UserConverter{}
}

// 把user对象转换成userVo对象
func (converter *UserConverter) ToUserVo(user *models.User) models.UserVO {
	return user.UserVO
}

// 把user列表转换成userVo列表
func (converter *UserConverter) ToUserVoList(users []*models.User) []*models.UserVO {
	userVoList := make([]*models.UserVO, len(users))
	for i, v := range users {
		userVoList[i] = &v.UserVO
	}
	return userVoList
}
