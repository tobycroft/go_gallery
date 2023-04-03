package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/enroll/model/EnrollModel"
	"main.go/app/v1/facility/model/FacilityUserModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func FacilityController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(func(c *gin.Context) {
		uid := c.GetHeader("uid")
		fuser := FacilityUserModel.Api_find_byUid(uid)
		if len(fuser) < 1 {
			RET.Fail(c, 403, nil, "你不是机构管理员")
			return
		}
	})

	route.Any("data", facility_data)
	route.Any("list", facility_list)

}

func facility_data(c *gin.Context) {
	facility_name, ok := Input.Post("facility_name", c, false)
	if !ok {
		return
	}

	all_num := EnrollModel.Api_count_bySchoolName(facility_name, nil)
	payed_num := EnrollModel.Api_count_bySchoolName(facility_name, true)
	RET.Success(c, 0, map[string]any{
		"all_num":   all_num,
		"payed_num": payed_num,
	}, nil)
}

func facility_list(c *gin.Context) {

}
