package UserInfoModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/app/v1/user/model/UserModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

func Api_joinPushandUserInfo_select(tag_id, where_key, start_date, end_date interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)

	if tag_id != nil {
		db.Where("tag_id", tag_id)
	}
	db.Where(where_key, ">=", start_date)
	db.Where(where_key, "<", end_date)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_joinUser_paginator_byPhoneAndReceiveDate(phone string, start_date, end_date interface{}, limit, page int) gorose.Paginate {
	db := tuuz.Db().Table(Table + " as a")
	db.Fields("a.*,b.phone,b.wx_name,b.wx_img")
	db.LeftJoin(UserModel.Table+" as b", "a.uid=b.id")
	if phone != "" {
		db.Where("b.phone", "like", "%"+phone+"%")
	}
	if start_date != nil {
		db.Where("a.receive_date", ">=", start_date)
	}
	if end_date != nil {
		db.Where("a.receive_date", "<=", end_date)
	}
	db.Where("a.is_receive", "=", 1)
	db.OrderBy("a.id desc")
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Paginator()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return gorose.Paginate{}
	} else {
		return ret
	}
}
