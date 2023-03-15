package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/tag/model/TagGroupModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func GroupController(route *gin.RouterGroup) {

}

func group_list(c *gin.Context) {
	datas := TagGroupModel.Api_select()
	RET.Success(c, 0, datas, nil)
}

func group_get(c *gin.Context) {
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	data := TagGroupModel.Api_find(id)
	RET.Success(c, 0, data, nil)
}
