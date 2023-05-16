package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/enroll/model/EnrollModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", enroll_add)
	//线下参加
	route.Any("offline", enroll_offline)
	route.Any("edit", enroll_edit)
	route.Any("list", enroll_list)
	route.Any("get", enroll_get)

}

func enroll_add(c *gin.Context) {
	uid := c.GetHeader("uid")
	tag_id, ok := Input.PostInt64("tag_id", c)
	if !ok {
		return
	}
	age, ok := Input.PostInt64("age", c)
	if !ok {
		return
	}
	tag_group_id, ok := Input.PostInt64("tag_group_id", c)
	if !ok {
		return
	}
	name, ok := Input.Post("name", c, true)
	if !ok {
		return
	}
	receiver_name := name
	email, ok := Input.Post("email", c, true)
	if !ok {
		return
	}
	gender, ok := Input.PostInt64("gender", c)
	if !ok {
		return
	}
	cert, ok := Input.Post("cert", c, true)
	if !ok {
		return
	}
	school_name, ok := Input.Post("school_name", c, true)
	if !ok {
		return
	}
	school_name_show, ok := Input.Post("school_name_show", c, true)
	if !ok {
		return
	}
	phone, ok := Input.Post("phone", c, true)
	if !ok {
		return
	}
	province, ok := Input.Post("province", c, true)
	if !ok {
		return
	}
	city, ok := Input.Post("city", c, true)
	if !ok {
		return
	}
	district, ok := Input.Post("district", c, true)
	if !ok {
		return
	}
	address, ok := Input.Post("address", c, true)
	if !ok {
		return
	}
	var e EnrollModel.Interface
	e.Db = tuuz.Db()
	e.Db.Begin()
	if len(e.Api_find_byUidAndCert(uid, cert, tag_id)) > 0 {
		e.Db.Rollback()
		RET.Fail(c, 406, nil, "一个孩子同类型活动只能参加一次")
		return
	}
	if id := e.Api_insert_ps("ps", uid, tag_id, age, tag_group_id, name, receiver_name, email, gender, cert, school_name, school_name_show, phone, province, city, district, address); id > 0 {
		e.Db.Commit()
		RET.Success(c, 0, id, nil)
	} else {
		e.Db.Rollback()
		RET.Fail(c, 500, nil, nil)
	}
}
