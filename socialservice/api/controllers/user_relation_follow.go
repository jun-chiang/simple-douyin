package controllers

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jun-chiang/simple-douyin/common/helper"
	commonModels "github.com/jun-chiang/simple-douyin/common/models"
	"github.com/jun-chiang/simple-douyin/socialservice/application/converters"
	"github.com/jun-chiang/simple-douyin/socialservice/application/service"
	"github.com/jun-chiang/simple-douyin/socialservice/infrastructure/persistence/database"
)

type DouyinRelationFollowListRequest struct {
	Token  string `json:"token"`   // 用户鉴权token
	UserID string `json:"user_id"` // 用户id
}

type DouyinRelationFollowListResponse struct {
	commonModels.BasicResponse
	UserList []*commonModels.UserVO `json:"user_list"` // 用户信息列表
}

// API Path :/douyin/relation/follow/list
func GetFollowListByUserId(c context.Context, ctx *app.RequestContext) {
	// 接收json参数
	var req DouyinRelationFollowListRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(consts.StatusOK, DouyinRelationFollowListResponse{
			BasicResponse: *helper.GetFailReponse("参数错误"),
			UserList:      nil,
		})
		return
	}
	// 类型转换
	userId, err := strconv.ParseInt(req.UserID, 0, 64)
	if err != nil {
		ctx.JSON(consts.StatusOK, DouyinRelationFollowListResponse{
			BasicResponse: *helper.GetFailReponse("user_id错误"),
			UserList:      nil,
		})
		return
	}
	// 获取应用层服务
	userRelationService, err := service.NewUserRelationService(&database.UserRelationRepository{})
	if err != nil {
		ctx.JSON(consts.StatusOK, DouyinRelationFollowListResponse{
			BasicResponse: *helper.GetFailReponse("查询失败"),
			UserList:      nil,
		})
		return
	}
	// 调用服务获取数据
	users, err := userRelationService.GetFollowListByUserId(userId)
	// 转换user为userVo
	userVoList := converters.NewUserConverter().ToUserVoList(users)

	if err != nil {
		ctx.JSON(consts.StatusOK, DouyinRelationFollowListResponse{
			BasicResponse: *helper.GetFailReponse("参数错误"),
			UserList:      nil,
		})
	} else {
		ctx.JSON(consts.StatusOK, DouyinRelationFollowListResponse{
			BasicResponse: *helper.GetSuccesReponse("success"),
			UserList:      userVoList,
		})
	}
}
