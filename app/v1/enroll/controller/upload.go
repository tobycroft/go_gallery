package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/enroll/model/EnrollUploadModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func UploadController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", upload_add)
}

func upload_add(c *gin.Context) {
	uid := c.GetHeader("uid")
	enroll_id, ok := Input.PostInt64("enroll_id", c)
	if !ok {
		return
	}
	title, ok := Input.Post("title", c, true)
	if !ok {
		return
	}
	content, ok := Input.Post("content", c, true)
	if !ok {
		return
	}
	attachment, ok := Input.Post("attachment", c, true)
	if !ok {
		return
	}
	teacher_name, ok := Input.Post("teacher_name", c, true)
	if !ok {
		return
	}
	teacher_phone, ok := Input.Post("teacher_phone", c, true)
	if !ok {
		return
	}
	is_original, ok := Input.PostBool("is_original", c)
	if !ok {
		return
	}

	var eu EnrollUploadModel.Interface
	eu.Db = tuuz.Db()
	if eu.Api_insert(uid, enroll_id, title, content, attachment, teacher_name, teacher_phone, is_original) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func upload_edit(c *gin.Context) {
	uid := c.GetHeader("uid")
	enroll_id, ok := Input.PostInt64("enroll_id", c)
	if !ok {
		return
	}
	title, ok := Input.Post("title", c, true)
	if !ok {
		return
	}
	content, ok := Input.Post("content", c, true)
	if !ok {
		return
	}
	attachment, ok := Input.Post("attachment", c, true)
	if !ok {
		return
	}
	teacher_name, ok := Input.Post("teacher_name", c, true)
	if !ok {
		return
	}
	teacher_phone, ok := Input.Post("teacher_phone", c, true)
	if !ok {
		return
	}
	is_original, ok := Input.PostBool("is_original", c)
	if !ok {
		return
	}

	var eu EnrollUploadModel.Interface
	eu.Db = tuuz.Db()
	if eu.Api_update(uid, enroll_id, title, content, attachment, teacher_name, teacher_phone, is_original) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
