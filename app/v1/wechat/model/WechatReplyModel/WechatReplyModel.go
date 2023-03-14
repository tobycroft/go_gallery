package WechatReplyModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "wechat_reply"

func Api_find_byWord(word interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("word", word)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
