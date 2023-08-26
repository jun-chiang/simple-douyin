package helper

import (
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/jun-chiang/simple-douyin/common/models"
)

func NewReverseProxy(srvName string) *reverseproxy.ReverseProxy {
	proxy, err := reverseproxy.NewSingleHostReverseProxy("http://" + srvName)

	if err != nil {
		panic(err)
	}

	proxy.SetClient(models.GetNacosDiscoverClient().Client)
	proxy.SetDirector(func(req *protocol.Request) {
		req.SetRequestURI(string(reverseproxy.JoinURLPath(req, proxy.Target)))
		req.Header.SetHostBytes(req.URI().Host())
		req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	})
	return proxy
}
