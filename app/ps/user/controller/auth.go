package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/ps/user/action/RetAction"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseModel/TokenModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/Net"
	"main.go/tuuz/RET"
)

func AuthController(route *gin.RouterGroup) {

	route.Any("register", auth_register)
	route.Any("login", auth_login)
	route.Any("send", auth_send)
	route.Any("code", auth_code)

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
	ret, err := Net.Post("http://api.ps.familyeducation.org.cn/v1/user/auth/code", nil, map[string]any{
		"phone":    phone,
		"password": password,
	}, nil, nil)
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
		return
	}
	RET.Success(c, 0, ret, nil)
}

type login_ret struct {
	Code int `json:"code"`
	Data struct {
		Uid   int    `json:"uid"`
		Token string `json:"token"`
		Admin int    `json:"admin"`
	} `json:"data"`
	Echo string `json:"echo"`
}

func auth_phone(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	code, ok := Input.PostLength("code", 4, 8, c, false)
	if !ok {
		return
	}
	var l login_ret
	ret, err := Net.Post("http://api.ps.familyeducation.org.cn/v1/user/auth/phone", nil, map[string]any{
		"phone": phone,
		"code":  code,
	}, nil, nil)
	err = RetAction.App_ret(ret, err, &l)
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
		return
	}
	uid :=
		RET.Success(c, 0, l.Data, nil)
}

func auth_send(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	ret, err := Net.Post("http://api.ps.familyeducation.org.cn/v1/user/auth/send", nil, map[string]any{
		"phone": phone,
	}, nil, nil)
	err = RetAction.App_ret(ret, err, nil)
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
		return
	}
	RET.Success(c, 0, nil, nil)
}

func auth_code(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	ret, err := Net.Post("http://api.ps.familyeducation.org.cn/v1/user/auth/code", nil, map[string]any{
		"phone": phone,
	}, nil, nil)
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
		return
	}
	RET.Success(c, 0, ret, nil)
}
