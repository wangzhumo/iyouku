package controllers

import (
	"strconv"

	"com.wangzhumo.iyouku/common"
	esClient "com.wangzhumo.iyouku/services/es"

	"com.wangzhumo.iyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

type OtherController struct {
	beego.Controller
}

// SendAllVideoToES 把所有视频数据导入到ES
// @router /video/send/es [*]
func (oc *OtherController) SendAllVideoToES() {
	// 获取所有的数据
	count, vidoes, err := models.GetAllVideoList()

	for _, v := range vidoes {
		// 发送到ES中去
		esBody := map[string]interface{}{
			"id":                   v.Id,
			"title":                v.Title,
			"sub_title":            v.SubTitle,
			"add_time":             v.AddTime,
			"img":                  v.Img,
			"img1":                 v.Img1,
			"episodes_count":       v.EpisodesCount,
			"episodes_update_time": v.EpisodesUpdateTime,
			"channel_id":           v.ChannelId,
			"type_id":              v.TypeId,
			"user_id":              v.UserId,
			"region_id":            v.RegionId,
			"is_end":               v.IsEnd,
			"is_hot":               v.IsHot,
			"is_recommend":         v.IsRecommend,
			"comment":              v.Comment,
		}
		esClient.EsAdd(common.EsIndexName, "video-"+strconv.Itoa(v.Id), esBody)
	}

	if err == nil {
		oc.Data["json"] = SucceedResp(0, RequestOk, nil, count)
		_ = oc.ServeJSON()
	} else {
		oc.Data["json"] = ErrorResp(5005, TypeTopError)
		_ = oc.ServeJSON()
	}
}
