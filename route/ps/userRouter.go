package ps

import (
	"github.com/gin-gonic/gin"
	"main.go/app/ps/user/controller"
)

func UserRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.AuthController(route.Group("auth"))
	controller.InfoController(route.Group("info"))
	controller.WechatController(route.Group("wechat"))
}
