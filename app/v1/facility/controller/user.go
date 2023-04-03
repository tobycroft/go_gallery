package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/facility/model/FacilityUserModel"
	"main.go/tuuz/RET"
)

func UserController(route *gin.RouterGroup) {

	route.Any("me", user_me)

}

func user_me(c *gin.Context) {
	uid := c.GetHeader("uid")
	user := FacilityUserModel.Api_find_byUid(uid)
	if len(user) < 1 {
		RET.Fail(c, 401, nil, "你不是机构管理员")
		return
	}
	RET.Success(c, 0, user, nil)
}
