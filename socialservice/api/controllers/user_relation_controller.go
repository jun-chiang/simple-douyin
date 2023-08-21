package controllers

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jun-chiang/simple-douyin/socialservice/application/service"
	"github.com/jun-chiang/simple-douyin/socialservice/infrastructure/persistence/database"
)

type DouyinRelationFollowListRequest struct {
	Token  string `json:"token"`   // 用户鉴权token
	UserID string `json:"user_id"` // 用户id
}

func GetFollowListByUserId(c context.Context, ctx *app.RequestContext) {
	// 接收json参数
	var req DouyinRelationFollowListRequest
	if err := ctx.Bind(&req); err != nil {
		panic(err)
	}
	// 类型转换
	userId, err := strconv.ParseInt(req.UserID, 0, 64)
	if err != nil {
		panic(err)
	}
	// 获取应用层服务
	userRelationService, err := service.NewUserRelationService(&database.UserRelationRepository{})
	if err != nil {
		panic(err)
	}
	// 调用服务获取数据
	users, err := userRelationService.GetFollowListByUserId(userId)
	// 控制返回数据
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, nil)
	} else {
		ctx.JSON(consts.StatusOK, users)
	}
}
