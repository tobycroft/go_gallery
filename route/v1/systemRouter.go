package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/system/controller"
)

func SystemRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.InfoController(route.Group("info"))
}
