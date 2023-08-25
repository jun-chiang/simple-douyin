package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/hertz-contrib/reverseproxy"
	commonConsts "github.com/jun-chiang/simple-douyin/common/consts"
)

func InitSocialServiceGateway(r *route.Engine) {
	proxy, err := reverseproxy.NewSingleHostReverseProxy("http://192.168.209.159:8091")
	if err != nil {
		panic(err)
	}

	r.Any(commonConsts.BasicURI+"/relation/*path", proxy.ServeHTTP)
}
