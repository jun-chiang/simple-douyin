package models

import "github.com/jun-chiang/simple-douyin/common/consts"

// 每个Response都会有的字段，提取出来，到时候嵌入到结构体里面去
type BasicResponse struct {
	StatusCode consts.StatusCode `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
}
