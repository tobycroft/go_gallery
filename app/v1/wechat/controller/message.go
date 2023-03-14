package controller

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/wechat/model/WechatReplyModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/RET"
)

func MessageController(route *gin.RouterGroup) {

	route.Use(cors.Default())

	route.Any("receive", wechat_reveive)
	route.Any("reply", wechat_reveive)

}

func wechat_reveive(c *gin.Context) {
	var wm AossGoSdk.Wechat_message_ret_struct
	data, err := c.GetRawData()
	if err != nil {
		RET.Fail(c, 400, nil, nil)
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return
	}
	err = jsoniter.Unmarshal(data, &wm)
	if err != nil {
		RET.Fail(c, 400, nil, nil)
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return
	}
	switch wm.MsgType {
	case "text":
		fmt.Println("收到微信消息:", wm.Content)
		reply := WechatReplyModel.Api_find_byWord(wm.Content)
		if len(reply) < 1 {
			break
		}
		err := wechat_reply_text(wm.FromUserName, Calc.Any2String(reply["reply"]))
		if err != nil {
			RET.Fail(c, 200, nil, err.Error())
			return
		}
		break

	case "event":
		switch wm.Event {
		case "subscribe":
			reply := WechatReplyModel.Api_find_byWord(wm.EventKey)
			if len(reply) < 1 {
				reply = WechatReplyModel.Api_find_byWord(wm.Project + "_on_sub")
				if len(reply) < 1 {
					break
				}
			}

			err := wechat_reply_text(wm.FromUserName, Calc.Any2String(reply["reply"]))
			if err != nil {
				RET.Fail(c, 200, nil, err.Error())
				return
			}
			break

		case "unsubscribe":
			break

		case "SCAN":
			reply := WechatReplyModel.Api_find_byWord(wm.EventKey)
			if len(reply) < 1 {
				break
			}
			err := wechat_reply_text(wm.FromUserName, Calc.Any2String(reply["reply"]))
			if err != nil {
				RET.Fail(c, 200, nil, err.Error())
				return
			}
			break

		}

	}
}

func wechat_reply_text(FromUserName, reply_content string) error {
	var wm AossGoSdk.Wechat_message
	return wm.Set_openid(FromUserName).Set_message_text(reply_content).Send()
}
