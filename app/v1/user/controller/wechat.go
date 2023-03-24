package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/TokenModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"time"
)

func WechatController(route *gin.RouterGroup) {

	route.Use(cors.Default())

	route.Any("testjump", wechat_testjump)
	route.Any("login", wechat_login)
	route.Any("bind", wechat_bind)
	route.Any("phone", wechat_phone)
	route.Any("scene", wechat_scene)
	route.Any("create", wechat_create)
	route.Any("create_file", wechat_create_file)
	route.Any("scheme", wechat_scheme)
	route.Any("redirect", wechat_redirect)
	route.Any("signature", wechat_signature)
	route.Any("signaturet", wechat_signaturet)

	route.Any("reply", wechat_reply)

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("authurl", wechat_authurl)
}

func wechat_scheme(c *gin.Context) {
	path, ok := Input.Post("path", c, false)
	if !ok {
		return
	}
	query, ok := Input.Post("query", c, false)
	if !ok {
		return
	}
	ret, err := AossGoSdk.Wechat_wxa_generatescheme(app_conf.Project, path, query, true, 180)
	if err != nil {
		RET.Fail(c, 200, ret, err.Error())
	} else {
		RET.Success(c, 0, ret.Openlink, nil)
	}
}

func wechat_login(c *gin.Context) {
	js_code, ok := Input.Combi("js_code", c, false)
	if !ok {
		return
	}
	ret, err := AossGoSdk.Wechat_sns_jscode2session(app_conf.Project, js_code)
	if err != nil {
		RET.Fail(c, 200, ret, err.Error())
		return
	}
	md5_pass := Calc.Md5(app_conf.Project + ret.SessionKey)
	token := Calc.GenerateToken()
	if user := UserModel.Api_find_byWxId(ret.Openid); len(user) > 0 {
		TokenModel.Api_insert(user["id"], token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token": token,
			"uid":   user["id"],
		}, nil)
		return
	}
	if id := UserModel.Api_insert_more("wx_"+ret.Openid, "wx_"+ret.Openid, md5_pass, ret.Openid, ret.Unionid, ""); id > 0 {
		token = Calc.GenerateToken()
		TokenModel.Api_insert(id, token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token": token,
			"uid":   id,
		}, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func wechat_bind(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	openid, ok := Input.Post("openid", c, false)
	if !ok {
		return
	}
	unionid, ok := Input.Post("unionid", c, false)
	if !ok {
		return
	}
	token := Calc.GenerateToken()
	if user := UserModel.Api_find_byWxId(openid); len(user) > 0 {
		TokenModel.Api_insert(user["id"], token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token": token,
			"uid":   user["id"],
		}, nil)
		return
	}
	if id := UserModel.Api_insert_more(phone, phone, Calc.Md5(password), openid, unionid, ""); id > 0 {
		token = Calc.GenerateToken()
		TokenModel.Api_insert(id, token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token": token,
			"uid":   id,
		}, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func wechat_phone(c *gin.Context) {
	uid := c.GetHeader("uid")
	code, ok := Input.Post("code", c, false)
	if !ok {
		return
	}
	ret, err := AossGoSdk.Wechat_wxa_getuserphonenumber(app_conf.Project, code)
	if err != nil {
		RET.Fail(c, 200, ret, err.Error())
		return
	}
	if err != nil {
		RET.Fail(c, 402, nil, nil)
		return
	}
	if UserModel.Api_update_phone(uid, ret.PurePhoneNumber) {
		RET.Success(c, 0, ret.PurePhoneNumber, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func wechat_create_file(c *gin.Context) {
	data, ok := Input.Combi("data", c, false)
	if !ok {
		return
	}
	url, err := AossGoSdk.Wechat_wxa_unlimited_file(app_conf.Project, data, "pages/registerInfo/registerInfo")
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		RET.Success(c, 0, url, nil)
	}
}

func wechat_create(c *gin.Context) {
	data, ok := Input.Combi("data", c, false)
	if !ok {
		return
	}
	file_url, err := AossGoSdk.Wechat_wxa_unlimited_file(app_conf.Project, data, "pages/registerInfo/registerInfo")
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		c.Redirect(302, file_url)
	}
}

func wechat_scene(c *gin.Context) {
	scene, ok := Input.Combi("scene", c, false)
	if !ok {
		return
	}

	sc, err := AossGoSdk.Wechat_wxa_scene(app_conf.Project, scene)
	if err != nil {
		RET.Fail(c, 404, nil, err.Error())
		return
	}
	RET.Success(c, 0, sc.Val, nil)
}

func wechat_redirect(c *gin.Context) {
	//ret := c.Request.URL.Query()
	////if err != nil {
	////	c.String(200, "err:"+err.Error())
	////} else {
	//c.String(200, "succ:"+ret.Encode())
	//ret2, _ := c.GetRawData()
	//c.String(200, "succ2:"+string(ret2))
	//}
	redirect, ok := Input.Get("redirect", c, false)
	if !ok {
		c.Redirect(301, "https://lc.familyeducation.org.cn/")
		return
	}
	code, ok := c.GetQuery("code")
	if !ok {
		c.Redirect(301, redirect+"?code=204&echo=授权码失败，请先关注账号&data="+code)
		return
	}
	state, ok := c.GetQuery("state")
	if !ok {
		c.Redirect(301, redirect+"?code=204&echo=用户账号获取失败，请先登录&data="+code)
		return
	}
	openid, err := AossGoSdk.Wechat_offi_openid_from_code(app_conf.Project, code)
	if err != nil {
		c.Redirect(301, redirect+"?code=200&echo="+err.Error()+"&data="+code)
		return
	}
	user := UserModel.Api_find_byPhone(state)
	if len(user) > 0 {
		UserModel.Api_update_openid(user["id"], openid)
		c.Redirect(301, redirect+"?code=0&echo=成功&data="+openid)
	} else {
		c.Redirect(301, redirect+"?code=204&echo=未找到用户，请先登录后再申请授权&data="+openid)
	}
}

func wechat_testjump(c *gin.Context) {
	redirect, ok := Input.Get("url", c, false)
	if !ok {
		c.Redirect(303, "https://gallery.familyeducation.org.cn/#/login?code=0&echo=成功&data=test")
	} else {
		RET.Success(c, 0, redirect+"?code=0&echo=成功&data=test", nil)
		//c.Redirect(303, redirect+"?code=0&echo=成功&data=test")

	}
}

func wechat_authurl(c *gin.Context) {
	uid := c.GetHeader("uid")
	user := UserModel.Api_find(uid)
	redirect, ok := Input.Post("url", c, false)
	if !ok {
		return
	}
	redirect_uri := "https://api.gallery.familyeducation.org.cn/v1/user/wechat/redirect?redirect=" + redirect
	url, err := AossGoSdk.Wechat_offi_openidUrl(app_conf.Project, redirect_uri, "code", "snsapi_base", Calc.Any2String(user["phone"]), false)
	if err != nil {
		RET.Fail(c, 200, nil, err)
	} else {
		RET.Success(c, 0, url, nil)
	}
}

func wechat_signature(c *gin.Context) {
	nonceStr := Calc.Md5(time.Now().String())
	timestamp := time.Now()
	url, ok := Input.Post("url", c, false)
	if !ok {
		return
	}
	signature, err := AossGoSdk.Wechat_ticket_signature(app_conf.Project, nonceStr, timestamp, url)
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		RET.Success(c, 0, map[string]interface{}{
			"nonceStr":  nonceStr,
			"timestamp": timestamp.Unix(),
			"signature": signature,
		}, nil)
	}
}

func wechat_signaturet(c *gin.Context) {
	nonceStr := "Wm3WZYTPz0wzccnW"
	timestamp := time.Now()
	url := "https://lc.familyeducation.org.cn"
	signature, err := AossGoSdk.Wechat_ticket_signature(app_conf.Project, nonceStr, timestamp, url)
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		RET.Success(c, 0, map[string]interface{}{
			"nonceStr":  nonceStr,
			"timestamp": timestamp.Unix(),
			"signature": signature,
		}, nil)
	}
}

func wechat_reply(c *gin.Context) {
	var wm AossGoSdk.Wechat_message
	err := wm.Set_openid("o44RD0sz-k16La_qohNjbAvGL6Zs").Set_message_text("jkasdjkhiasdkj").Send()
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		RET.Success(c, 0, map[string]interface{}{}, nil)
	}
}
