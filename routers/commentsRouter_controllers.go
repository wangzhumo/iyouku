package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:BasicController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:BasicController"],
		beego.ControllerComments{
			Method:           "ChannelRegion",
			Router:           "/channel/region",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:BasicController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:BasicController"],
		beego.ControllerComments{
			Method:           "ChannelType",
			Router:           "/channel/type",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:CommentController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "GetCommentList",
			Router:           "/comment/list",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:CommentController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "InsertComment",
			Router:           "/comment/save",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:TopController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:TopController"],
		beego.ControllerComments{
			Method:           "GetChannelTop",
			Router:           "/channel/top",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:TopController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:TopController"],
		beego.ControllerComments{
			Method:           "GetChannelType",
			Router:           "/type/top",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

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

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:UserController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SendPushMessage",
			Router:           "/send/message",
			AllowHTTPMethods: []string{"post"},
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

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelRecommendByType",
			Router:           "/channel/recommend/type",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelVideoByParams",
			Router:           "/channel/video",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "GetVideoEpisode",
			Router:           "/video/episode",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"] = append(beego.GlobalControllerRouter["com.wangzhumo.iyouku/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "GetVideoInfo",
			Router:           "/video/info",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
