package EnrolllikeModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "g_enroll_like"

func Api_count_byEnrollId(enroll_id any) int64 {
	db := tuuz.Db().Table(Table)
	db.Where("enroll_id", enroll_id)
	ret, err := db.Counts()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_count(uid, enroll_id any) int64 {
	db := tuuz.Db().Table(Table)
	db.Where("enroll_id", enroll_id)
	db.Where("uid", uid)
	ret, err := db.Counts()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_count_today(uid any) int64 {
	db := tuuz.Db().Table(Table)
	db.Where("uid", uid)
	db.Where("date>current_date()")
	ret, err := db.Counts()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_insert(uid, enroll_id any) bool {
	db := tuuz.Db().Table(Table)
	data := map[string]any{
		"uid":       uid,
		"enroll_id": enroll_id,
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
