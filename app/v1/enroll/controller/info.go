package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/enroll/model/EnrollModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"time"
)

func InfoController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", enroll_add)
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
	mp := Input.NewModelPost(c)
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	//mp.PostInt64("tag_id")
	mp.PostInt64("age")
	mp.PostInt64("tag_group_id")
	mp.PostString("name")
	mp.PostString("email")
	mp.PostInt64("gender")
	mp.PostString("nacertme")
	mp.PostString("school_name")
	mp.PostString("school_name_show")
	mp.PostString("phone")
	mp.PostString("province")
	mp.PostString("city")
	mp.PostString("district")
	mp.PostString("address")
	mp.PostDateTime("expect_date")
	data, err, errmsg := mp.GetPostMap()
	if err != nil {
		RET.Fail(c, 400, nil, errmsg)
		return
	}
	if mp.Has("expect_date") {
		if mp.Find("expect_date").(time.Time).Before(time.Now().AddDate(0, 0, 3)) {
			RET.Fail(c, 406, nil, "时间需要预约三天以后")
			return
		}
	}
	var e EnrollModel.Interface
	e.Db = tuuz.Db()
	if e.Api_update_auto(id, uid, data) {
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

func enroll_get(c *gin.Context) {
	uid := c.GetHeader("uid")
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	data := EnrollModel.Api_find_byUid(uid, id)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
