package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/jun-chiang/simple-douyin/gateway/api/routers"
)

func main() {
	srv := server.Default(server.WithHostPorts("192.168.209.159:9999"))

	// 初始化SocialService的反向代理

	routers.InitSocialServiceGateway(srv.Engine)

	srv.Spin()
}
