// Package routers
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"com.wangzhumo.iyouku/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 用户相关的 - 注册，登录
	beego.Include(&controllers.UserController{})
	// 视频 信息，列表
	beego.Include(&controllers.VideoController{})
	// 获取地区/类型信息
	beego.Include(&controllers.BasicController{})
	// 评论相关
	beego.Include(&controllers.CommentController{})
	// 排行榜
	beego.Include(&controllers.TopController{})
	// 其他 elasticsearch 脚本
	beego.Include(&controllers.OtherController{})
}
