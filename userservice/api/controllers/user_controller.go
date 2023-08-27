package controllers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	netConsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	commonConsts "github.com/jun-chiang/simple-douyin/common/consts"
	commonModels "github.com/jun-chiang/simple-douyin/common/models"
	"github.com/jun-chiang/simple-douyin/userservice/application/service"
	user "github.com/jun-chiang/simple-douyin/userservice/domain/models"
	"github.com/jun-chiang/simple-douyin/userservice/infrastructure/persistence/database"
)

// API Path :/douyin/user/register
func RegisterController(c context.Context, ctx *app.RequestContext) {
	// 接收json参数
	var req user.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		panic(err)
	}

	// 获取应用层服务
	userService, err := service.NewUserService(&database.UserRepository{})
	if err != nil {
		panic(err)
	}
	// 调用服务获取数据
	user_id, err := userService.UserRegister(&req)

	var statusMsg string = "success"
	if err != nil {
		statusMsg = "failure"
	}
	// 构建返回对象
	respData := user.RegisterResponse{
		BasicResponse: commonModels.BasicResponse{
			StatusCode: commonConsts.StatusOk,
			StatusMsg:  statusMsg,
		},
		UserID: user_id,
	}
	// 控制返回数据
	if err != nil {
		ctx.JSON(netConsts.StatusBadRequest, nil)
	} else {
		ctx.JSON(netConsts.StatusOK, respData)
	}
}
