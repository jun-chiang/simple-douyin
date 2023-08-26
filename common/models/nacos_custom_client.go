package models

import (
	"sync"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/hertz-contrib/registry/nacos"
)

type NacosDiscoverClient struct {
	Client *client.Client
}

var (
	discoverClient   *NacosDiscoverClient
	initDiscoverOnce sync.Once
)

func GetNacosDiscoverClient() *NacosDiscoverClient {
	initDiscoverOnce.Do(func() {
		cli, err := client.NewClient()
		if err != nil {
			panic(err)
		}
		cli.Use(sd.Discovery(nacos.NewNacosResolver(GetNacosNamingClient().Client)))
		discoverClient = &NacosDiscoverClient{
			Client: cli,
		}
	})
	return discoverClient
}
