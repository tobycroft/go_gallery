package UserModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "lc_user"

type Interface struct {
	Db gorose.IOrm
}

func Api_insert(username, phone, password interface{}) int64 {
	var i Interface
	i.Db = tuuz.Db()
	return i.Api_insert(username, phone, password)
}

func (self *Interface) Api_insert(username, phone, password interface{}) int64 {
	db := tuuz.Db().Table(Table)
	data := map[string]interface{}{
		"username": username,
		"phone":    phone,
		"password": password,
	}
	db.Data(data)
	ret, err := db.InsertGetId()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_find_byPhone(phone interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"phone": phone,
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

func Api_find_byPhoneandPassword(phone, password interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"phone":    phone,
		"password": password,
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

func Api_find_avail(id interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"id":     id,
		"active": 1,
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

func Api_find_avail_in(ids []interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("id", "in", ids)
	db.Where("active", 1)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_limit(id interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Fields("username,wx_name,wx_img,share,active,phone")
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

func Api_find_noPassword(id interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Fields("username,wx_name,wx_img,share,active,wx_id,wx_union,phone")
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

func Api_count() int64 {
	db := tuuz.Db().Table(Table)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func (self *Interface) Api_update_phone(uid, phone interface{}) bool {
	db := self.Db.Table(Table)
	where := map[string]interface{}{
		"id": uid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"phone": phone,
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

func Api_update_usernameAndNameAndWxImg(uid, username, wx_name, wx_img interface{}) bool {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"id": uid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"username": username,
		"wx_name":  wx_name,
		"wx_img":   wx_img,
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

func (self *Interface) Api_update_usernameAndNameAndWxImg(uid, username, wx_name interface{}) bool {
	db := self.Db.Table(Table)
	where := map[string]interface{}{
		"id": uid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"username": username,
		"wx_name":  wx_name,
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

func Api_update_phone(uid, phone interface{}) bool {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"id": uid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"phone": phone,
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

func Api_find_byWxId(wx_id interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"wx_id": wx_id,
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

func Api_find_byWxName(wx_name interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"wx_name": wx_name,
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

func Api_insert_more(username, phone, password, wx_id, wx_union, share interface{}) int64 {
	db := tuuz.Db().Table(Table)
	data := map[string]interface{}{
		"username": username,
		"phone":    phone,
		"password": password,
		"wx_id":    wx_id,
		"wx_union": wx_union,
		"share":    share,
	}
	db.Data(data)
	ret, err := db.InsertGetId()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_update_openid(uid, wx_id interface{}) bool {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"id": uid,
	}
	db.Where(where)
	data := map[string]interface{}{
		"wx_id": wx_id,
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

func Api_select_wxId_byId(ids, notin_ids []interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	if len(ids) > 0 {
		db.WhereIn("id", ids)
	}
	if len(notin_ids) > 0 {
		db.WhereNotIn("id", notin_ids)
	}
	db.WhereNotNull("wx_id")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_phone_byId(ids, notin_ids []interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	if len(ids) > 0 {
		db.WhereIn("id", ids)
	}
	if len(notin_ids) > 0 {
		db.WhereNotIn("id", notin_ids)
	}
	db.WhereNotNull("phone")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_in(ids []interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	db.WhereIn("id", ids)
	db.WhereNotNull("phone")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
