package EnrollModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "g_enroll"

type Interface struct {
	Db gorose.IOrm
}

func Api_select(uid, is_used, is_verify, is_payed interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("uid", uid)
	if is_used != nil {
		db.Where("is_used", is_used)
	}
	if is_verify != nil {
		db.Where("is_verify", is_verify)
	}
	if is_payed != nil {
		db.Where("is_payed", is_payed)
	}
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_insert(uid, tag_id, age, tag_group_id, name, email, gender, cert, school_name, phone, province, city, district, address interface{}) bool {
	db := self.Db.Table(Table)
	data := map[string]interface{}{
		"uid":          uid,
		"tag_id":       tag_id,
		"age":          age,
		"tag_group_id": tag_group_id,
		"name":         name,
		"email":        email,
		"gender":       gender,
		"cert":         cert,
		"school_name":  school_name,
		"phone":        phone,
		"province":     province,
		"city":         city,
		"district":     district,
		"address":      address,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_update(id, uid, tag_id, age, tag_group_id, name, email, gender, cert, school_name, phone, province, city, district, address interface{}) bool {
	db := self.Db.Table(Table)
	db.Where("id", id)
	db.Where("uid", uid)
	data := map[string]interface{}{
		"tag_id":       tag_id,
		"age":          age,
		"tag_group_id": tag_group_id,
		"name":         name,
		"email":        email,
		"gender":       gender,
		"cert":         cert,
		"school_name":  school_name,
		"phone":        phone,
		"province":     province,
		"city":         city,
		"district":     district,
		"address":      address,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_update_isUsed(order_id, is_used interface{}) bool {
	db := self.Db.Table(Table)
	db.Where("order_id", order_id)
	data := map[string]any{
		"is_used": is_used,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_update_isVerify(order_id, is_verify interface{}) bool {
	db := self.Db.Table(Table)
	db.Where("order_id", order_id)
	data := map[string]any{
		"is_verify": is_verify,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_update_isPayed(order_id, is_payed interface{}) bool {
	db := self.Db.Table(Table)
	db.Where("order_id", order_id)
	data := map[string]any{
		"is_payed": is_payed,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
