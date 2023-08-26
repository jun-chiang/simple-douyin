package models

import commonModels "github.com/jun-chiang/simple-douyin/common/models"

type User struct {
	// 基础的数据库字段
	commonModels.User
	// GORM关系映射
	Following []*User `gorm:"many2many:user_relations;foreignKey:ID;joinForeignKey:FollowFrom;references:ID;joinReferences:FollowTo"`
	Followers []*User `gorm:"many2many:user_relations;foreignKey:ID;joinForeignKey:FollowTo;references:ID;joinReferences:FollowFrom"`
}
