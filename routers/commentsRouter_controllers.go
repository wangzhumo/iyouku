package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:UserController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SaveRegister",
			Router:           "/register/save",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
