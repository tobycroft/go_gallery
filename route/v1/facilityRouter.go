package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/facility/controller"
)

func FacilityRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.UserController(route.Group("user"))
	controller.InfoController(route.Group("info"))

}
