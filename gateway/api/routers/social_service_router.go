package routers

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/jun-chiang/simple-douyin/common/consts"
	"github.com/jun-chiang/simple-douyin/common/helper"
)

func InitSocialServiceGateway(r *route.Engine) {

	proxy := helper.NewReverseProxy(consts.SocialService)

	r.Any(consts.BasicURI+"/relation/*path", proxy.ServeHTTP)
}
