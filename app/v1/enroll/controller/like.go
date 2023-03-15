package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/enroll/model/EnrolllikeModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LikeController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", like_add)
}

func like_add(c *gin.Context) {
	uid := c.GetHeader("uid")
	enroll_id, ok := Input.PostInt64("enroll_id", c)
	if !ok {
		return
	}
	like_limit := SystemParamModel.Api_find_val("like_limit")
	like_limit_day := SystemParamModel.Api_find_val("like_limit_day")
	count := EnrolllikeModel.Api_count(uid, enroll_id)
	if count > Calc.Any2Int64(like_limit) {
		RET.Fail(c, 406, nil, "超过投票次数")
		return
	}
	count_day := EnrolllikeModel.Api_count_today(uid)
	if count_day > Calc.Any2Int64(like_limit_day) {
		RET.Fail(c, 406, nil, "超过当日投票次数")
		return
	}
	if EnrolllikeModel.Api_insert(uid, enroll_id) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
