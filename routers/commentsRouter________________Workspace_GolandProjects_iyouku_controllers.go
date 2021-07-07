package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:UserController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login/do",
			AllowHTTPMethods: []string{"get"},
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

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelAdvert",
			Router:           "/channel/advert",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelHotList",
			Router:           "/channel/hot",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelRecommendByRegion",
			Router:           "/channel/recommend/region",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
