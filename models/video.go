package models

import "github.com/beego/beego/v2/client/orm"

type Video struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Imgh          string
	Imgv          string
	EpisodesCount int
	IsEnd         int
	ChannelId     int
	Status        int
	RegionId      int
	TypeId        int
	Sort          int
	EpisodesTime  int
	Comment       string
}

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

type Episodes struct {
	Id      int
	Title   string
	AddTime int64
	Num     int
	PlayUrl string
	Comment string
}

func init() {
	orm.RegisterModel(new(Video))
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

const SortUpdateTime = "episodesUpdateTime"
const SortComment = "comment"
const SortAddTime = "addTime"

// GetChannelVideoList 通过限制条件获取指定的Video列表
func GetChannelVideoList(channelId int, typeId int, regionId int, end string, sort string, offset int, limit int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	// 数据
	var params []orm.Params
	querySeter := o.QueryTable("video")
	querySeter = querySeter.Filter("channel_id", channelId)
	querySeter = querySeter.Filter("status", 1)
	if typeId > 0 {
		querySeter = querySeter.Filter("type_id", typeId)
	}
	if regionId > 0 {
		querySeter = querySeter.Filter("regionId", regionId)
	}

	if end == "n" {
		querySeter = querySeter.Filter("is_end", 0)
	} else if end == "y" {
		querySeter = querySeter.Filter("is_end", 1)
	}

	if sort == SortUpdateTime {
		querySeter = querySeter.OrderBy("episodes_update")
	} else if sort == SortComment {
		querySeter = querySeter.OrderBy("comment")
	} else if sort == SortAddTime {
		querySeter = querySeter.OrderBy("add_time")
	} else {
		querySeter = querySeter.OrderBy("add_time")
	}

	// 获取总条数
	count, _ := querySeter.Values(&params, "id", "title", "sub_title", "add_time", "imgh", "imgv", "episodes_count", "is_end")

	// 获取指定的Limit数据
	querySeter = querySeter.Limit(limit, offset)
	_, err := querySeter.Values(&params, "id", "title", "sub_title", "add_time", "imgh", "imgv", "episodes_count", "is_end")
	return count, params, err
}

// GetVideoInfo 通过videoId获取视频详情信息
func GetVideoInfo(videoID int) (VideoDate, error) {
	newOrm := orm.NewOrm()
	var video VideoDate
	err := newOrm.Raw("SELECT * FROM video where id =? ", videoID).QueryRow(&video)
	return video, err
}

// GetVideoEpisodes 通过videoId获取视频剧集
func GetVideoEpisodes(videoID int) (int64, []VideoDate, error) {
	newOrm := orm.NewOrm()
	var episodes []VideoDate
	rows, err := newOrm.Raw("SELECT id, title, add_time, num, play_url, comment FROM video_episodes where video_id = ? AND status =1 ORDER BY num ASC;", videoID).QueryRows(&episodes)
	return rows, episodes, err
}

// GetChannelTop 通过channelId获取视频的排行
func GetChannelTop(channelId int) (int64, []VideoDate, error) {
	o := orm.NewOrm()
	var videos []VideoDate

	rows, err := o.Raw("SELECT d, title, sub_title, "+
		"channel_id, add_time, imgh, imgv ,episodes_count "+
		"FROM video WHERE status=0 AND channel_id=?  "+
		"ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
	return rows, videos, err
}

// GetTypeTop 通过typeId获取视频的排行
func GetTypeTop(typeId int) (int64, []VideoDate, error) {
	o := orm.NewOrm()
	var videos []VideoDate

	rows, err := o.Raw("SELECT d, title, sub_title, "+
		"channel_id, add_time, imgh, imgv ,episodes_count "+
		"FROM video WHERE status=0 AND type_id=?  "+
		"ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
	return rows, videos, err
}
