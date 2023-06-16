package WechatOrderModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "wechat_order"

func Api_insert(order_id, relative_id, amount any) bool {
	db := tuuz.Db().Table(Table)
	data := map[string]any{
		"order_id":    order_id,
		"relative_id": relative_id,
		"amount":      amount,
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

func Api_update_status(order_id, status any) bool {
	db := tuuz.Db().Table(Table)
	db.Where("order_id", order_id)
	data := map[string]any{
		"status": status,
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