package controllers

import (
	"encoding/json"

	"com.wangzhumo.iyouku/common"
	"com.wangzhumo.iyouku/models"
	esclient "com.wangzhumo.iyouku/services/es"
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
		vc.Data["json"] = ErrorResp(4001, NoChannelID)
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
		vc.Data["json"] = ErrorResp(4001, NoChannelID)
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

// ChannelRecommendByRegion 根据频道地区获取推荐视频
// @router /channel/recommend/region [get]
func (vc *VideoController) ChannelRecommendByRegion() {
	regionId, _ := vc.GetInt("regionId")
	channelId, _ := vc.GetInt("channelId")
	// empty check
	if channelId == 0 {
		vc.Data["json"] = ErrorResp(4001, NoChannelID)
		_ = vc.ServeJSON()
	}
	if regionId == 0 {
		vc.Data["json"] = ErrorResp(4001, NoRegionID)
		_ = vc.ServeJSON()
	}
	// get data
	count, videos, err := models.GetRecommendByRegionID(regionId, channelId)
	// 获取到数据，可以进行返回
	if err != nil {
		vc.Data["json"] = ErrorResp(4004, VideoChannelRecommendError)
		_ = vc.ServeJSON()
	} else {
		vc.Data["json"] = SucceedResp(0, RequestOk, videos, count)
		_ = vc.ServeJSON()
	}
}

// ChannelRecommendByType 根据频道地区获取推荐视频
// @router /channel/recommend/type [get]
func (vc *VideoController) ChannelRecommendByType() {
	typeId, _ := vc.GetInt("typeId")
	channelId, _ := vc.GetInt("channelId")
	// empty check
	if channelId == 0 {
		vc.Data["json"] = ErrorResp(4001, NoChannelID)
		_ = vc.ServeJSON()
	}
	if typeId == 0 {
		vc.Data["json"] = ErrorResp(4001, NoTypeID)
		_ = vc.ServeJSON()
	}

	// get data
	count, videos, err := models.GetRecommendByTypeID(typeId, channelId)
	// 获取到数据，可以进行返回
	if err != nil {
		vc.Data["json"] = ErrorResp(4004, VideoChannelTypeError)
		_ = vc.ServeJSON()
	} else {
		vc.Data["json"] = SucceedResp(0, RequestOk, videos, count)
		_ = vc.ServeJSON()
	}
}

// ChannelVideoByParams 根据参数获取推荐视频
// @router /channel/video [get]
func (vc *VideoController) ChannelVideoByParams() {
	// 频道信息
	channelId, _ := vc.GetInt("channelId")
	// 地区
	regionId, _ := vc.GetInt("regionId")
	// 类型参数
	typeId, _ := vc.GetInt("typeId")
	// 状态
	status := vc.GetString("end")
	// 排序
	sort := vc.GetString("sort")
	// 分页信息
	limit, _ := vc.GetInt("limit")
	offset, _ := vc.GetInt("offset")

	// empty check 【必传的参数】
	if channelId == 0 {
		vc.Data["json"] = ErrorResp(4001, NoChannelID)
		_ = vc.ServeJSON()
	}

	// 可选的参数,但是需要指定默认值
	if limit == 0 {
		limit = 12
	}

	count, videos, err := models.GetChannelVideoList(channelId, typeId, regionId, status, sort, offset, limit)
	// 获取到数据，可以进行返回
	if err != nil {
		vc.Data["json"] = ErrorResp(4004, ChannelVideoError)
		_ = vc.ServeJSON()
	} else {
		vc.Data["json"] = SucceedResp(0, RequestOk, videos, count)
		_ = vc.ServeJSON()
	}
}

// GetVideoInfo 获取视频详情
// @router /video/info [*]
func (vc *VideoController) GetVideoInfo() {
	videoId, _ := vc.GetInt("videoId")
	// empty check
	if videoId == 0 {
		vc.Data["json"] = ErrorResp(4001, NoVideoID)
		_ = vc.ServeJSON()
	}

	videoInfo, err := models.GetCacheVideoInfo(videoId)
	if err != nil {
		vc.Data["json"] = ErrorResp(4004, ChannelVideoError)
		_ = vc.ServeJSON()
	} else {
		vc.Data["json"] = SucceedResp(0, RequestOk, videoInfo, 1)
		_ = vc.ServeJSON()
	}
}

// GetVideoEpisode 获取视频剧集详情
// @router /video/episode [get]
func (vc *VideoController) GetVideoEpisode() {
	videoId, _ := vc.GetInt("videoId")
	// empty check
	if videoId == 0 {
		vc.Data["json"] = ErrorResp(4001, NoVideoID)
		_ = vc.ServeJSON()
	}

	count, videoEpisodes, err := models.GetCacheVideoEpisodes(videoId)
	if err != nil {
		vc.Data["json"] = ErrorResp(4004, VideoEpisodesError)
		_ = vc.ServeJSON()
	} else {
		vc.Data["json"] = SucceedResp(0, RequestOk, videoEpisodes, count)
		_ = vc.ServeJSON()
	}
}

// SearchVideoByKeyword 搜索视频
// @router /video/search [*]
func (vc *VideoController) SearchVideoByKeyword() {
	keyword := vc.GetString("keyword")
	limit, _ := vc.GetInt("limit")
	offset, _ := vc.GetInt("offset")
	if keyword == "" {
		vc.Data["json"] = ErrorResp(4001, NoKeyword)
		_ = vc.ServeJSON()
	}

	// 可选的参数,但是需要指定默认值
	if limit == 0 {
		limit = 10
	}

	querySort := []map[string]string{
		map[string]string{
			"id": "desc",
		},
	}

	queryParams := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"term": map[string]interface{}{
					"title": keyword,
				},
			},
		},
	}
	hd, err := esclient.EsSearch(common.EsIndexName, queryParams, offset, limit, querySort)
	if err != nil {
		vc.Data["json"] = ErrorResp(4004, VideoSearchError)
		_ = vc.ServeJSON()
	} else {
		// 获取数据
		total := hd.Total.Value
		var data []models.Video

		for _, v := range hd.Hits {
			var temp models.Video
			err = json.Unmarshal([]byte(v.Source), &temp)
			if err == nil {
				data = append(data, temp)
			}
		}

		vc.Data["json"] = SucceedResp(0, RequestOk, data, int64(total))
		_ = vc.ServeJSON()
	}
}
