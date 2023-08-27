package models

import (
	commonModels "github.com/jun-chiang/simple-douyin/common/models"
)

type RegisterRequest struct {
	Password string `json:"password"  query:"password"` // 密码，最长32个字符
	Username string `json:"username" query:"username"`  // 注册用户名，最长32个字符
}

type RegisterResponse struct {
	commonModels.BasicResponse
	Token  string `json:"token"`   // 用户鉴权token
	UserID int64  `json:"user_id"` // 用户id
}
