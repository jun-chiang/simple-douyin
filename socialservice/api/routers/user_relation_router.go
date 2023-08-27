package routers

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/jun-chiang/simple-douyin/socialservice/api/controllers"
)

func InitRouter(r *route.Engine) {
	// 限定relationGroup的作用域
	{
		relationGoup := r.Group("/douyin/relation")
		relationGoup.GET("/follow/list", controllers.GetFollowListByUserId)
		relationGoup.GET("/follower/list", controllers.GetFollowerListByUserId)
	}
}
