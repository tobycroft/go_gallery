package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/wechat/controller"
)

func WechatRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.MessageController(route.Group("message"))

	controller.PayController(route.Group("pay"))
	controller.ApiController(route.Group("api"))

}
