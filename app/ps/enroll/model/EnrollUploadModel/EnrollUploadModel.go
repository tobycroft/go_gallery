package EnrollUploadModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "g_enroll_upload"

type Interface struct {
	Db gorose.IOrm
}

func Api_find(uid, enroll_id any) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("uid", uid)
	db.Where("enroll_id", enroll_id)
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_insert(uid, enroll_id, title, content, attachment, teacher_name, teacher_phone, is_original any) bool {
	db := self.Db.Table(Table)
	data := map[string]any{
		"uid":           uid,
		"enroll_id":     enroll_id,
		"title":         title,
		"content":       content,
		"attachment":    attachment,
		"teacher_name":  teacher_name,
		"teacher_phone": teacher_phone,
		"is_original":   is_original,
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

func (self *Interface) Api_update(uid, enroll_id, title, content, attachment, teacher_name, teacher_phone, is_original any) bool {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Where("enroll_id", enroll_id)
	data := map[string]any{
		"title":         title,
		"content":       content,
		"attachment":    attachment,
		"teacher_name":  teacher_name,
		"teacher_phone": teacher_phone,
		"is_original":   is_original,
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

func (self *Interface) Api_inc_like(enroll_id any) bool {
	db := self.Db.Table(Table)
	db.Where("enroll_id")
	_, err := db.Increment("likes", 1)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
