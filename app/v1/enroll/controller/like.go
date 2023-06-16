package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/enroll/model/EnrollUploadModel"
	"main.go/app/v1/enroll/model/EnrolllikeModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LikeController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", like_add)
	route.Any("status", like_status)
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
	if count >= Calc.Any2Int64(like_limit) {
		RET.Fail(c, 406, nil, "超过投票次数")
		return
	}
	count_day := EnrolllikeModel.Api_count_today(uid)
	if count_day >= Calc.Any2Int64(like_limit_day) {
		RET.Fail(c, 406, nil, "超过当日投票次数")
		return
	}
	var el EnrolllikeModel.Interface
	db := tuuz.Db()
	el.Db = db
	db.Begin()
	if !el.Api_insert(uid, enroll_id) {
		db.Rollback()
		RET.Fail(c, 500, nil, nil)
		return
	}
	var eu EnrollUploadModel.Interface
	eu.Db = db
	if !eu.Api_inc_like(enroll_id) {
		db.Rollback()
		RET.Fail(c, 500, nil, nil)
		return
	}
	db.Commit()
	RET.Success(c, 0, nil, nil)
}

func like_status(c *gin.Context) {
	uid := c.GetHeader("uid")
	enroll_id, ok := Input.PostInt64("enroll_id", c)
	if !ok {
		return
	}
	like_limit := SystemParamModel.Api_find_val("like_limit")
	like_limit_day := SystemParamModel.Api_find_val("like_limit_day")
	count := EnrolllikeModel.Api_count(uid, enroll_id)
	count_day := EnrolllikeModel.Api_count_today(uid)
	RET.Success(c, 0, map[string]any{
		"like_sum":       count,
		"like_today":     count_day,
		"like_limit":     like_limit,
		"like_limit_day": like_limit_day,
		"today_left":     Calc.ToDecimal(like_limit_day).Sub(Calc.ToDecimal(count_day)),
		"total_left":     Calc.ToDecimal(like_limit).Sub(Calc.ToDecimal(count)),
		"current_avail":  Calc.ToDecimal(like_limit_day).Sub(Calc.ToDecimal(count_day)).GreaterThan(decimal.Zero) && Calc.ToDecimal(like_limit).Sub(Calc.ToDecimal(count)).GreaterThanOrEqual(decimal.Zero),
	}, nil)
}
