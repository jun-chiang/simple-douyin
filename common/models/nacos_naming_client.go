package models

import (
	"sync"

	"github.com/jun-chiang/simple-douyin/common/consts"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NacosNamingClient struct {
	Client naming_client.INamingClient
}

var (
	initnamingCliOnce sync.Once
	namingClient      *NacosNamingClient
)

func GetNacosNamingClient() *NacosNamingClient {
	initnamingCliOnce.Do(func() {
		nacosCli, err := clients.NewNamingClient(vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId:         consts.NacosNamespaceID,
				NotLoadCacheAtStart: true,
				TimeoutMs:           5000,
				LogDir:              "/tmp/nacos/log",
				CacheDir:            "/tmp/nacos/cache",
				LogLevel:            "debug",
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig(consts.NacosIp, consts.NacosPort),
			},
		})
		if err != nil {
			panic(err)
		}
		namingClient = &NacosNamingClient{
			Client: nacosCli,
		}
	})
	return namingClient
}
