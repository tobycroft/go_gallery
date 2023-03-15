package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/enroll/model/EnrollModel"
	"main.go/app/v1/enroll/model/EnrollUploadModel"
	"main.go/app/v1/enroll/model/EnrolllikeModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func UploadController(route *gin.RouterGroup) {

	route.Any("list", upload_list)
	route.Any("list_rank", upload_list_rank)

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", upload_add)
	route.Any("edit", upload_edit)
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

	data := EnrollUploadModel.Api_find(uid, enroll_id)
	if len(data) > 0 {
		RET.Fail(c, 406, nil, "你已经上传过作品了")
		return
	}
	db := tuuz.Db()
	var eu EnrollUploadModel.Interface
	eu.Db = db
	db.Begin()
	if eu.Api_insert(uid, enroll_id, title, content, attachment, teacher_name, teacher_phone, is_original) {
		var e EnrollModel.Interface
		e.Db = db
		if !e.Api_update_isUpload(enroll_id, true) {
			db.Rollback()
			RET.Fail(c, 500, nil, "修改错误")
			return
		}
		db.Commit()
		RET.Success(c, 0, nil, nil)
	} else {
		db.Rollback()
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

func upload_list(c *gin.Context) {
	tag_id, ok := Input.PostInt64("tag_id", c)
	if !ok {
		return
	}
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	datas := EnrollUploadModel.Api_joinEnroll_paginator_byTagId(tag_id, limit, page)
	for i, datum := range datas.Data {
		datum["like"] = EnrolllikeModel.Api_count_byEnrollId(datum["enroll_id"])
		datas.Data[i] = datum
	}
	RET.Success(c, 0, datas, nil)
}

func upload_list_rank(c *gin.Context) {
	tag_id, ok := Input.PostInt64("tag_id", c)
	if !ok {
		return
	}
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	datas := EnrollUploadModel.Api_joinEnroll_paginator_byTagId_orderByLikes(tag_id, limit, page)
	RET.Success(c, 0, datas, nil)
}
