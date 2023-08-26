package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/jun-chiang/simple-douyin/common/consts"
	"github.com/jun-chiang/simple-douyin/gateway/api/routers"
)

func main() {
	srv := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", consts.ProjectDevelopIP, 9999)))

	// 初始化SocialService的反向代理
	routers.InitSocialServiceGateway(srv.Engine)

	srv.Spin()
}
