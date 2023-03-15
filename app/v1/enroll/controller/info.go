package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
)

func EnrollController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", enroll_add)

}

func enroll_add(c *gin.Context) {
	uid := c.GetHeader("uid")
	tag_id, ok := Input.PostInt64("tag", c)
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

}
