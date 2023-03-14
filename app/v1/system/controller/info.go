package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Any("list", info_list)
}

func info_list(c *gin.Context) {
	datas := SystemParamModel.Api_select()
	RET.Success(c, 0, datas, nil)
}
