package TagModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "lc_tag"

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(Table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_byType(study_type, tag_type interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	if study_type != nil {
		db.Where("study_type", study_type)
	}
	if tag_type != nil {
		db.Where("tag_type", tag_type)
	}
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
