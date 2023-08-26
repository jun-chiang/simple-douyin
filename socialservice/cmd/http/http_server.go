package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/jun-chiang/simple-douyin/common/consts"
	"github.com/jun-chiang/simple-douyin/common/models"
	"github.com/jun-chiang/simple-douyin/socialservice/api/routers"
)

func main() {
	// 执行数据迁移，创建数据表
	// migrations.Migration()
	srvAddr := fmt.Sprintf("%s:%d", consts.ProjectDevelopIP, 8091)
	srv := server.Default(server.WithHostPorts(srvAddr),
		server.WithRegistry(
			nacos.NewNacosRegistry(models.GetNacosNamingClient().Client),
			&registry.Info{
				ServiceName: "social-service",
				Addr:        utils.NewNetAddr("tcp", srvAddr),
				Weight:      10,
			}),
	)

	routers.InitRouter(srv.Engine)

	srv.Spin()

}
