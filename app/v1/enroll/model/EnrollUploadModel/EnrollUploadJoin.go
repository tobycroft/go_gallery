package EnrollUploadModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/app/v1/enroll/model/EnrollModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

func Api_joinEnroll_paginator_byTagId(tag_id any, limit, page int) gorose.Paginate {
	db := tuuz.Db().Table(Table + " a")
	db.Fields("a.attachment", "a.title", "b.tag_id", "b.name")
	db.LeftJoin(EnrollModel.Table+" b", "a.enroll_id=b.id")
	db.Where("tag_id", tag_id)
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Paginator()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return gorose.Paginate{}
	} else {
		return ret
	}
}
