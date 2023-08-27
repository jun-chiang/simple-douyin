package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/jun-chiang/simple-douyin/userservice/api/routers"
	"github.com/jun-chiang/simple-douyin/userservice/infrastructure/persistence/database/migrations"
)

func main() {
	// 执行数据迁移，创建数据表
	migrations.Migration()
	srvAddr := fmt.Sprintf("%s:%d", "172.17.160.75", 8091)
	srv := server.Default(server.WithHostPorts(srvAddr))
	routers.InitRouter(srv.Engine)
	srv.Spin()

}
