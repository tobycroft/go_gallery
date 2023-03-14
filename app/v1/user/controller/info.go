package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/user/model/UserInfoModel"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("edit", info_edit)
	route.Any("get", info_get)
	route.Any("my", info_my)

}

func info_get(c *gin.Context) {
	uid := c.GetHeader("uid")
	data := UserInfoModel.Api_find_byUid(uid)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func info_my(c *gin.Context) {
	uid := c.GetHeader("uid")
	data := UserModel.Api_find_noPassword(uid)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func info_edit(c *gin.Context) {
	uid := c.GetHeader("uid")
	data := map[string]interface{}{}
	wx_name, ok := Input.Post("wx_name", c, true)
	if !ok {
		return
	}
	mdp := Input.NewModelPost(c)
	mdp.Xss(true)
	mdp.Fields("couple_name")
	mdp.PostString("face").
		PostInt64("tag_id").
		PostDateTime("birthday").
		PostDateTime("marrige_date").
		PostInt64("baby_gender").
		PostDateTime("baby_birthday").
		PostDateTime("pregnant_date").
		PostString("couple_name").
		PostString("province").PostString("city").PostString("district").PostString("street").PostString("address")
	data, err, msg := mdp.GetPostMap()
	if err != nil {
		RET.Fail(c, 400, err.Error(), msg)
		return
	}
	db := tuuz.Db()
	var u UserModel.Interface
	u.Db = db
	var ui UserInfoModel.Interface
	ui.Db = db

	db.Begin()

	if !u.Api_update_usernameAndNameAndWxImg(uid, wx_name, wx_name) {
		db.Rollback()
		RET.Fail(c, 500, nil, nil)
		return
	}
	data["uid"] = uid
	if len(ui.Api_find_byUid(uid)) > 0 {
		if ui.Api_update(uid, data) {
			db.Commit()
			RET.Success(c, 0, nil, nil)
			return
		}
	} else {
		if ui.Api_insert_manual(data) {
			db.Commit()
			RET.Success(c, 0, nil, nil)
			return
		}
	}
	db.Rollback()
	RET.Fail(c, 500, nil, nil)
}
