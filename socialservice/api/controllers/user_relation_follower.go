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

type DouyinRelationFollowerListRequest struct {
	Token  string `json:"token"`   // 用户鉴权token
	UserID string `json:"user_id"` // 用户id
}

type DouyinRelationFollowerListResponse struct {
	commonModels.BasicResponse
	UserList []*commonModels.UserVO `json:"user_list"` // 用户列表
}

// API Path : /douyin/relation/follower/list
func GetFollowerListByUserId(c context.Context, ctx *app.RequestContext) {
	var req DouyinRelationFollowerListRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(consts.StatusOK, DouyinRelationFollowerListResponse{
			BasicResponse: *helper.GetFailReponse("参数错误！"),
			UserList:      nil,
		})
		return
	}

	userId, err := strconv.ParseInt(req.UserID, 0, 64)
	if err != nil {
		ctx.JSON(consts.StatusOK, DouyinRelationFollowerListResponse{
			BasicResponse: *helper.GetFailReponse("user_id错误"),
			UserList:      nil,
		})
		return
	}

	userRelationService, err := service.NewUserRelationService(&database.UserRelationRepository{})
	if err != nil {
		panic(err)
	}

	users, err := userRelationService.GetFollowerListByUserId(userId)
	if err != nil {
		ctx.JSON(consts.StatusOK, DouyinRelationFollowerListResponse{
			BasicResponse: *helper.GetFailReponse("查询出错"),
			UserList:      nil,
		})
		return
	}
	userVoList := converters.NewUserConverter().ToUserVoList(users)
	ctx.JSON(consts.StatusOK, DouyinRelationFollowerListResponse{
		BasicResponse: *helper.GetSuccesReponse("success"),
		UserList:      userVoList,
	})
}
