package controllers

import (
	"com.wangzhumo.iyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

type BasicController struct {
	beego.Controller
}

// ChannelRegion 获取channel的地区信息
// @router /channel/region [get]
func (c *BasicController) ChannelRegion() {

	channelId, _ := c.GetInt("channelId")
	// empty check
	if channelId == 0 {
		c.Data["json"] = ErrorResp(4001, NoChannelID)
		_ = c.ServeJSON()
	}

	// get data
	rows, regions, err := models.GetChannelRegion(channelId)
	if err != nil {
		c.Data["json"] = ErrorResp(4004, ChannelRegionError)
		_ = c.ServeJSON()
	} else {
		c.Data["json"] = SucceedResp(0, RequestOk, regions, rows)
		_ = c.ServeJSON()
	}
}

// ChannelType 获取channel的地区信息
// @router /channel/type [get]
func (c *BasicController) ChannelType() {

	channelId, _ := c.GetInt("channelId")
	// empty check
	if channelId == 0 {
		c.Data["json"] = ErrorResp(4001, NoChannelID)
		_ = c.ServeJSON()
	}

	// get data
	rows, regions, err := models.GetChannelType(channelId)
	if err != nil {
		c.Data["json"] = ErrorResp(4004, ChannelTypeError)
		_ = c.ServeJSON()
	} else {
		c.Data["json"] = SucceedResp(0, RequestOk, regions, rows)
		_ = c.ServeJSON()
	}
}
