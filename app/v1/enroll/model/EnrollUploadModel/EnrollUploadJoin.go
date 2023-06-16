package EnrollUploadModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/app/v1/enroll/model/EnrollModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

func Api_joinEnroll_paginator_byTagId(tag_id any, search string, limit, page int) gorose.Paginate {
	db := tuuz.Db().Table(Table + " a")
	db.Fields("a.enroll_id", "a.attachment", "a.title", "b.tag_id", "b.name", "a.likes")
	db.LeftJoin(EnrollModel.Table+" b", "a.enroll_id=b.id")
	db.Where("tag_id", tag_id)
	db.Where("a.title", "like", search+"%")
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("a.id desc")
	ret, err := db.Paginator()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return gorose.Paginate{}
	} else {
		return ret
	}
}

func Api_joinEnroll_find_byEnrollId(enroll_id any) gorose.Data {
	db := tuuz.Db().Table(Table + " a")
	db.Fields("a.enroll_id", "a.attachment", "a.title", "b.tag_id,tag_group_id", "b.name", "a.likes")
	db.LeftJoin(EnrollModel.Table+" b", "a.enroll_id=b.id")
	db.Where("enroll_id", enroll_id)
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_joinEnroll_paginator_byTagId_orderByLikes(tag_id any, search string, limit, page int) gorose.Paginate {
	db := tuuz.Db().Table(Table + " a")
	db.Fields("a.enroll_id", "a.attachment", "a.title", "b.tag_id", "b.name", "a.likes")
	db.LeftJoin(EnrollModel.Table+" b", "a.enroll_id=b.id")
	db.Where("tag_id", tag_id)
	db.Where("a.title", "like", search+"%")
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("a.likes desc")
	ret, err := db.Paginator()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return gorose.Paginate{}
	} else {
		return ret
	}
}
