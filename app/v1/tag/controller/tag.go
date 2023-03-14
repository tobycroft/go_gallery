package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/tag/model/TagModel"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Any("list", tag_list)
}

func tag_list(c *gin.Context) {
	datas := TagModel.Api_select()
	RET.Success(c, 0, datas, nil)
}
