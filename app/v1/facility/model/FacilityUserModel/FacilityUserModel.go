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
