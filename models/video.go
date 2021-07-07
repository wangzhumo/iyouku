package models

import "github.com/beego/beego/v2/client/orm"

type VideoDate struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Imgh          string
	Imgv          string
	EpisodesCount int
	IsEnd         int
}

func init() {
	orm.RegisterModel(new(VideoDate))
}

func (u *VideoDate) TableName() string {
	return "video"
}

// GetHotListByChannelID 通过ChannelID获取热播的数据
func GetHotListByChannelID(channelId int) (int64, []VideoDate, error) {
	newOrm := orm.NewOrm()
	var videos []VideoDate
	count, err := newOrm.Raw("SELECT id, title, sub_title, "+
		"channel_id, add_time, imgh, imgv ,episodes_count,is_end "+
		"FROM video where is_hot = 1 AND status = 1 AND channel_id =? "+
		"ORDER BY episodes_update DESC LIMIT 10", channelId).QueryRows(&videos)
	return count, videos, err
}

// GetRecommendByRegionID 通过regionIdD获取推荐视频
func GetRecommendByRegionID(regionId int, channelId int) (int64, []VideoDate, error) {
	newOrm := orm.NewOrm()
	var videos []VideoDate
	count, err := newOrm.Raw("SELECT id, title, sub_title, "+
		"channel_id, add_time, imgh, imgv ,episodes_count,is_end "+
		"FROM video where status = 1 AND is_recommend = 1 AND channel_id =? AND region_id =? "+
		"ORDER BY episodes_update DESC LIMIT 10", channelId, regionId).QueryRows(&videos)
	return count, videos, err
}

// GetRecommendByTypeID 通过typeId获取推荐视频
func GetRecommendByTypeID(typeId int, channelId int) (int64, []VideoDate, error) {
	newOrm := orm.NewOrm()
	var videos []VideoDate
	count, err := newOrm.Raw("SELECT id, title, sub_title, "+
		"channel_id, add_time, imgh, imgv ,episodes_count,is_end "+
		"FROM video where status = 1 AND is_recommend = 1 AND channel_id =? AND type_id =? "+
		"ORDER BY episodes_update DESC LIMIT 10", channelId, typeId).QueryRows(&videos)
	return count, videos, err
}
