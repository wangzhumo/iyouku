package controllers

import (
	"com.wangzhumo.iyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

type VideoController struct {
	beego.Controller
}

// ChannelAdvert 频道获取顶部广告
// @router /channel/advert [get]
func (vc *VideoController) ChannelAdvert() {
	channelId, _ := vc.GetInt("channelId")

	// empty check
	if channelId == 0 {
		vc.Data["json"] = ErrorResp(4001, VideoNoChannelID)
		_ = vc.ServeJSON()
	}

	// query channelId
	count, adverts, err := models.GetAdvertByChannelID(channelId)
	// 获取到数据，可以进行返回
	if err != nil {
		vc.Data["json"] = ErrorResp(4004, VideoChannelError)
		_ = vc.ServeJSON()
	} else {
		vc.Data["json"] = SucceedResp(0, RequestOk, adverts, count)
		_ = vc.ServeJSON()
	}
}

// ChannelHotList 频道获取正在热播的内容
// @router /channel/hot [get]
func (vc *VideoController) ChannelHotList() {
	channelId, _ := vc.GetInt("channelId")

	// empty check
	if channelId == 0 {
		vc.Data["json"] = ErrorResp(4001, VideoNoChannelID)
		_ = vc.ServeJSON()
	}

	// query channelId
	count, videos, err := models.GetHotListByChannelID(channelId)
	// 获取到数据，可以进行返回
	if err != nil {
		vc.Data["json"] = ErrorResp(4004, VideoChannelHotError)
		_ = vc.ServeJSON()
	} else {
		vc.Data["json"] = SucceedResp(0, RequestOk, videos, count)
		_ = vc.ServeJSON()
	}
}
