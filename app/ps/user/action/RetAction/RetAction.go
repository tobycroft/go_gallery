package RetAction

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
)

type ret struct {
	Code int    `json:"code"`
	Echo string `json:"echo"`
}

func App_ret(body string, err error, Struct any) error {
	if err != nil {
		return err
	}
	var r ret
	err = jsoniter.UnmarshalFromString(body, &r)
	if err != nil {
		return err
	}
	if r.Code != 0 {
		return errors.New(r.Echo)
	}
	if Struct != nil {
		return jsoniter.UnmarshalFromString(body, &Struct)
	}
	return nil
}
