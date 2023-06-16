package WechatOrderModel

import (
	"github.com/tobycroft/gorose-pro"
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

func Api_update_status(order_id, status, status_msg any) bool {
	db := tuuz.Db().Table(Table)
	db.Where("order_id", order_id)
	data := map[string]any{
		"status":     status,
		"status_msg": status_msg,
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

func Api_find_orderId(order_id any) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("order_id", order_id)
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
