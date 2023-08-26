package models

type UserRelationDomain struct {
	ID         int64 `json:"id" gorm:"primary_key;auto_increment"` // 主键id
	FollowFrom int64 `json:"follow_from"`                          // 关注的发起方
	FollowTo   int64 `json:"follow_to"`
}
