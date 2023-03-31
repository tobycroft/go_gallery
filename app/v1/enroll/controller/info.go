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
	route.Any("edit", enroll_edit)
	route.Any("list", enroll_list)

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
	if e.Api_insert(uid, tag_id, age, tag_group_id, name, email, gender, cert, school_name, school_name_show, phone, province, city, district, address) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func enroll_edit(c *gin.Context) {
	uid := c.GetHeader("uid")
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
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
	if e.Api_update(id, uid, tag_id, age, tag_group_id, name, email, gender, cert, school_name, school_name_show, phone, province, city, district, address) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func enroll_list(c *gin.Context) {
	uid := c.GetHeader("uid")
	datas := EnrollModel.Api_select(uid, nil, nil, nil)
	RET.Success(c, 0, datas, nil)
}
