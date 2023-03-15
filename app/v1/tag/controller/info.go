package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/tag/model/TagModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Any("list", tag_list)
	route.Any("get", tag_get)
}

func tag_list(c *gin.Context) {
	datas := TagModel.Api_select()
	RET.Success(c, 0, datas, nil)
}

func tag_get(c *gin.Context) {
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	data := TagModel.Api_find(id)
	RET.Success(c, 0, data, nil)
}
