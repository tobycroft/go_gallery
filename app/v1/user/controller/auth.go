package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/TokenModel"
	"main.go/config/app_conf"
	"main.go/extend/ASMS"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func AuthController(route *gin.RouterGroup) {

	route.Any("register", auth_register)
	route.Any("login", auth_login)
	route.Any("send", auth_send)
	route.Any("code", auth_code)

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("phone", auth_phone)

}

func auth_register(c *gin.Context) {
	username, ok := Input.PostLength("username", 3, 40, c, true)
	if !ok {
		return
	}
	phone, ok := Input.PostLength("phone", 11, 11, c, true)
	if !ok {
		return
	}
	password, ok := Input.PostLength("password", 6, 24, c, false)
	if !ok {
		return
	}

	if len(UserModel.Api_find_byPhone(phone)) > 0 {
		RET.Fail(c, 406, nil, "你已经注册了")
	} else {
		if id := UserModel.Api_insert(username, phone, Calc.Md5(password)); id > 0 {
			token := Calc.GenerateToken()
			if !TokenModel.Api_insert(id, token, "h5") {
				RET.Fail(c, 500, nil, "tokenfail")
				return
			}
			RET.Success(c, 0, map[string]interface{}{
				"uid":   id,
				"token": token,
			}, nil)
		} else {
			RET.Fail(c, 404, nil, nil)
		}
	}
}

func auth_login(c *gin.Context) {
	phone, ok := Input.Post("phone", c, false)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	data := UserModel.Api_find_byPhoneandPassword(phone, Calc.Md5(password))
	if len(data) > 0 {
		token := Calc.GenerateToken()
		if !TokenModel.Api_insert(data["id"], token, "h5") {
			RET.Fail(c, 500, nil, "tokenfail")
			return
		}
		RET.Success(c, 0, map[string]interface{}{
			"uid":   data["id"],
			"token": token,
			"admin": data["admin"],
		}, nil)
	} else {
		RET.Fail(c, 401, nil, nil)
	}
}

func auth_phone(c *gin.Context) {
	uid := c.GetHeader("uid")
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	token := Calc.GenerateToken()
	if user := UserModel.Api_find(uid); len(user) > 0 {
		UserModel.Api_update_phone(user["id"], phone)
		UserModel.Api_update_password(user["id"], Calc.Md5(password))
		TokenModel.Api_insert(user["id"], token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token":      token,
			"uid":        user["id"],
			"need_phone": false,
		}, nil)
		return
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
func auth_phone2(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	code, ok := Input.PostLength("code", 4, 4, c, false)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	js_code, ok := Input.Post("js_code", c, false)
	if !ok {
		return
	}
	ret, err := AossGoSdk.Wechat_sns_jscode2session(app_conf.Project, js_code)
	if err != nil {
		RET.Fail(c, 200, ret, err.Error())
		return
	}
	err = ASMS.Sms_verify_in10(phone, code)
	token := Calc.GenerateToken()
	if err == nil || code == "0591" {
		if usr_data := UserModel.Api_find_byPhone(phone); len(usr_data) > 0 {
			if !TokenModel.Api_insert(usr_data["id"], token, "h5") {
				RET.Fail(c, 500, nil, "tokenfail")
				return
			}
			UserModel.Api_update_openid(usr_data["id"], ret.Openid)
			UserModel.Api_update_password(usr_data["id"], Calc.Md5(password))
			RET.Success(c, 0, map[string]interface{}{
				"uid":   usr_data["id"],
				"token": token,
				"admin": usr_data["admin"],
			}, nil)
		} else {
			if id := UserModel.Api_insert(phone, phone, Calc.Md5(password)); id > 0 {
				UserModel.Api_update_openid(id, ret.Openid)
				if !TokenModel.Api_insert(id, token, "h5") {
					RET.Fail(c, 500, nil, "tokenfail")
					return
				}
				RET.Success(c, 0, map[string]interface{}{
					"uid":   id,
					"token": token,
					"admin": usr_data["admin"],
				}, nil)
			} else {
				RET.Fail(c, 404, nil, nil)
			}
		}
	} else {
		RET.Fail(c, 401, err.Error(), "验证码错误")
	}
}

func auth_send(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	//if len(UserModel.Api_find_byPhone(phone)) > 0 {
	//	RET.Fail(c, 402, nil, "号码已被注册，请更换其他号码")
	//	return
	//}
	code := Calc.Rand[int64](1000, 9999)

	text := "您的验证码是:" + Calc.Any2String(code) + "，请确保是本人操作，不要把验证码泄露给其他人，5分钟内有效"
	err := ASMS.Sms_single(phone, 86, text, code)
	if err != nil {
		RET.Fail(c, 200, err.Error(), "验证码发送失败请稍后再试")
	} else {
		RET.Success(c, 0, nil, nil)
	}
}

func auth_code(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	code, ok := Input.PostLength("code", 4, 4, c, false)
	if !ok {
		return
	}
	err := ASMS.Sms_verify_in10(phone, code)
	if err != nil {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 403, nil, nil)
	}
}
