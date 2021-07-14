package controllers

import (
	"com.wangzhumo.iyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

type TopController struct {
	beego.Controller
}

// GetChannelTop 根据channel获取排行榜
// @router /channel/top [get]
func (tc *TopController) GetChannelTop() {
	// 获取channelId
	channelId, _ := tc.GetInt("channelId")
	// empty check
	if channelId == 0 {
		tc.Data["json"] = ErrorResp(4001, NoChannelID)
		_ = tc.ServeJSON()
	}

	// 获取channelId下的排行榜
	count, dates, err := models.GetCacheChannelTop(channelId)
	if err == nil {
		tc.Data["json"] = SucceedResp(0, RequestOk, dates, count)
		_ = tc.ServeJSON()
	} else {
		tc.Data["json"] = ErrorResp(4005, ChannelTopError)
		_ = tc.ServeJSON()
	}
}

// GetChannelType 根据Type获取排行榜
// @router /type/top [get]
func (tc *TopController) GetChannelType() {
	// 获取channelId
	typeId, _ := tc.GetInt("typeId")
	// empty check
	if typeId == 0 {
		tc.Data["json"] = ErrorResp(4001, NoTypeID)
		_ = tc.ServeJSON()
	}

	// 获取channelId下的排行榜
	count, dates, err := models.GetCacheTypeTop(typeId)
	if err == nil {
		tc.Data["json"] = SucceedResp(0, RequestOk, dates, count)
		_ = tc.ServeJSON()
	} else {
		tc.Data["json"] = ErrorResp(4005, TypeTopError)
		_ = tc.ServeJSON()
	}
}
