package route

import (
	"github.com/gin-gonic/gin"
	"main.go/route/ps"
	v1 "main.go/route/v1"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	version1 := router.Group("/v1")
	{
		version1.Use(func(context *gin.Context) {
		}, gin.Recovery())
		version1.Any("/", func(context *gin.Context) {
			context.String(0, version1.BasePath())
		})
		index := version1.Group("index")
		{
			v1.IndexRouter(index)
		}
		system := version1.Group("system")
		{
			v1.SystemRouter(system)
		}
		tag := version1.Group("tag")
		{
			v1.TagRouter(tag)
		}
		user := version1.Group("user")
		{
			v1.UserRouter(user)
		}
		wechat := version1.Group("wechat")
		{
			v1.WechatRouter(wechat)
		}
		enroll := version1.Group("enroll")
		{
			v1.EnrollRouter(enroll)
		}
		facility := version1.Group("facility")
		{
			v1.FacilityRouter(facility)
		}
	}
	ps1 := router.Group("/ps")
	{
		user := ps1.Group("user")
		{
			ps.UserRouter(user)
		}
		enroll := ps1.Group("enroll")
		{
			ps.EnrollRouter(enroll)
		}
	}
}
