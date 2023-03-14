package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/gift/model/GiftRecordModel"
	"main.go/app/v1/user/model/UserInfoModel"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/config/app_conf"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"time"
)

func PrintController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("list", print_list)
	route.Any("url", print_url)
	route.Any("receive", print_receive)

	route.Use(func(c *gin.Context) {
		uid := c.GetHeader("uid")
		user := UserModel.Api_find_avail(uid)
		if len(user) < 1 {
			RET.Fail(c, 401, nil, nil)
			return
		}
		if user["admin"].(int64) != int64(1) {
			RET.Fail(c, 403, nil, nil)
			return
		}
	})
	route.Any("edit", print_edit)
}

func print_receive(c *gin.Context) {
	uid := c.GetHeader("uid")
	userinfo := UserInfoModel.Api_find_byUid(uid)
	if len(userinfo) > 1 {
		if userinfo["is_receive"].(int64) != 1 {
			if !UserInfoModel.Api_update_receive_byUid(uid, 1, time.Now()) {
				RET.Fail(c, 500, nil, nil)
				return
			}
			if !GiftRecordModel.Api_update_isValid(uid, 1) {
				RET.Fail(c, 500, nil, nil)
				return
			}
		}
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
func print_list(c *gin.Context) {
	phone, ok := Input.Post("phone", c, false)
	if !ok {
		return
	}
	var start_date interface{}
	date1, ok := Input.SPostDate("start_date", c)
	if ok {
		start_date = date1
	}
	var end_date interface{}
	date2, ok := Input.SPostDate("end_date", c)
	if ok {
		end_date = date2
	}
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	datas := UserInfoModel.Api_joinUser_paginator_byPhoneAndReceiveDate(phone, start_date, end_date, limit, page)
	RET.Success(c, 0, datas, nil)
}

func print_url(c *gin.Context) {
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	userinfo := UserInfoModel.Api_find(id)
	if len(userinfo) < 1 {
		RET.Fail(c, 404, nil, nil)
		return
	}
	user := UserModel.Api_find(userinfo["uid"])
	if len(user) < 1 {
		RET.Fail(c, 404, nil, nil)
		return
	}
	name1 := Calc.Any2String(user["ewx_name"])
	name2 := Calc.Any2String(userinfo["couple_name"])
	var canvas AossGoSdk.Canvas
	canvas.AddText(name1, AossGoSdk.Canvas_Posistion_CenterCenter, -20, 0)
	canvas.AddText(name2, AossGoSdk.Canvas_Posistion_CenterCenter, 20, 0)
	canvas.AddImage("https://static.familyeducation.org.cn/lc/20221220/f572a4d6d26d5b66b31aa4d817b7be52.jpeg", 0, 0)
	url, err := canvas.Get_Url(app_conf.Project, 1080, 1920, "ffffff")
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		RET.Success(c, 0, url, nil)
	}
}

func print_edit(c *gin.Context) {
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	is_pirnt, ok := Input.PostBool("is_print", c)
	if !ok {
		return
	}
	pirnt_date := time.Now()
	if UserInfoModel.Api_update_print(id, is_pirnt, pirnt_date) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
