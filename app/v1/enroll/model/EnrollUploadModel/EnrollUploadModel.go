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

func (self *Interface) Api_update(id, uid, enroll_id, title, content, attachment, teacher_name, teacher_phone, is_original any) bool {
	db := self.Db.Table(Table)
	db.Where("id", id)
	db.Where("uid", uid)
	data := map[string]any{
		"enroll_id":     enroll_id,
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
