package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/facility/model/FacilityUserModel"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func UserController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("me", user_me)

}

func user_me(c *gin.Context) {
	uid := c.GetHeader("uid")
	user := UserModel.Api_find(uid)
	FacilityUserModel.Api_update_uid(uid, user["phone"])
	fuser := FacilityUserModel.Api_find_byUid(uid)
	if len(fuser) < 1 {
		RET.Fail(c, 401, nil, "你不是机构管理员")
		return
	}
	RET.Success(c, 0, fuser, nil)
}
