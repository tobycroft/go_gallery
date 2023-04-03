package FacilityUserModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "g_facility_user"

func Api_find_byUid(uid interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("uid", uid)
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_byUidAndFacilityName(uid, facility_name interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("uid", uid)
	db.Where("facility_name", facility_name)
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_update_uid(uid, phone interface{}) bool {
	db := tuuz.Db().Table(Table)
	db.Where("phone", phone)
	db.Where("uid", 0)
	data := map[string]any{
		"uid": uid,
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
