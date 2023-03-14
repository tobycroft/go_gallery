package UserInfoModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "lc_user_info"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_find_byUid(uid interface{}) gorose.Data {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_insert(uid, tag_id, birthday, marrige_date, baby_gender, baby_birthday interface{}) bool {
	db := self.Db.Table(Table)
	data := map[string]interface{}{
		"uid":           uid,
		"tag_id":        tag_id,
		"birthday":      birthday,
		"marrige_date":  marrige_date,
		"baby_gender":   baby_gender,
		"baby_birthday": baby_birthday,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_insert_manual(data interface{}) bool {
	db := self.Db.Table(Table)
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find_byUid(uid interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_byCoupleName(couple_name interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"couple_name": couple_name,
	}
	db.Where(where)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(id interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_update(uid, data interface{}) bool {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select_manual_byTagId(tag_id, where_key, start_date, end_date interface{}) []gorose.Data {
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
func Api_update_print(id, is_print, print_date interface{}) bool {
	db := tuuz.Db().Table(Table)
	db.Where("id", id)
	data := map[string]interface{}{
		"is_print":   is_print,
		"print_date": print_date,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
func Api_update_receive_byUid(uid, is_receive, receive_date interface{}) bool {
	db := tuuz.Db().Table(Table)
	db.Where("uid", uid)
	data := map[string]interface{}{
		"is_receive":   is_receive,
		"receive_date": receive_date,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
