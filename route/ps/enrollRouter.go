package ps

import (
	"github.com/gin-gonic/gin"
	"main.go/app/ps/enroll/controller"
)

func EnrollRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.InfoController(route.Group("info"))
	controller.UploadController(route.Group("upload"))
	controller.LikeController(route.Group("like"))

	controller.FacilityController(route.Group("facility"))
}
