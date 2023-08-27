package routers

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/jun-chiang/simple-douyin/userservice/api/controllers"
)

func InitRouter(r *route.Engine) {
	// 限定relationGroup的作用域
	{
		relationGoup := r.Group("/douyin/user")
		relationGoup.POST("/register", controllers.RegisterController)
	}
}
