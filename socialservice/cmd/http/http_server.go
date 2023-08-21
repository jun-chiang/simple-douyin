package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/jun-chiang/simple-douyin/socialservice/api/routers"
)

func main() {
	// 执行数据迁移，创建数据表
	// migrations.Migration()
	srv := server.Default(server.WithHostPorts("192.168.209.159:8091"))

	routers.InitRouter(srv.Engine)

	srv.Spin()

}
