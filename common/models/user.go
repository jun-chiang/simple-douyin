package models

// User对象在多个服务中都有使用，所以提取到Common模块来

// 用于返回给前端的对象
type UserVO struct {
	Avatar          string `json:"avatar"`                               // 用户头像
	BackgroundImage string `json:"background_image"`                     // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`                       // 喜欢数
	FollowCount     int64  `json:"follow_count"`                         // 关注总数
	FollowerCount   int64  `json:"follower_count"`                       // 粉丝总数
	ID              int64  `json:"id" gorm:"primary_key;auto_increment"` // 用户id
	IsFollow        bool   `json:"is_follow"`                            // true-已关注，false-未关注
	Name            string `json:"name"`                                 // 用户名称
	Signature       string `json:"signature"`                            // 个人简介
	TotalFavorited  string `json:"total_favorited"`                      // 获赞数量
	WorkCount       int64  `json:"work_count"`                           // 作品数
}

// 数据库对象
type User struct {
	UserVO
	// 不序列化敏感字段
	Password string `json:"-"`
	GormBase
}
