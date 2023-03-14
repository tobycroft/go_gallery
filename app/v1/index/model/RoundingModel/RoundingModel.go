package RoundingModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "lc_rounding"

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
