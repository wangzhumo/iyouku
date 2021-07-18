package models

import (
	"encoding/json"
	"strconv"

	redisClient "com.wangzhumo.iyouku/services/redis"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gomodule/redigo/redis"
)

type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	IsHot              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime int
	Comment            int
	UserId             int
	IsRecommend        int
}

type VideoDate struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
	Comment       int
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
		"channel_id, add_time, img, img1 ,episodes_count,is_end "+
		"FROM video where is_hot = 1 AND status = 1 AND channel_id =? "+
		"ORDER BY episodes_update_time DESC LIMIT 10", channelId).QueryRows(&videos)
	return count, videos, err
}

// GetRecommendByRegionID 通过regionIdD获取推荐视频
func GetRecommendByRegionID(regionId int, channelId int) (int64, []VideoDate, error) {
	newOrm := orm.NewOrm()
	var videos []VideoDate
	count, err := newOrm.Raw("SELECT id, title, sub_title, "+
		"channel_id, add_time, img, img1 ,episodes_count,is_end "+
		"FROM video where status = 1 AND is_recommend = 1 AND channel_id =? AND region_id =? "+
		"ORDER BY episodes_update_time DESC LIMIT 10", channelId, regionId).QueryRows(&videos)
	return count, videos, err
}

// GetRecommendByTypeID 通过typeId获取推荐视频
func GetRecommendByTypeID(typeId int, channelId int) (int64, []VideoDate, error) {
	newOrm := orm.NewOrm()
	var videos []VideoDate
	count, err := newOrm.Raw("SELECT id, title, sub_title, "+
		"channel_id, add_time, img, img1 ,episodes_count,is_end "+
		"FROM video where status = 1 AND is_recommend = 1 AND channel_id =? AND type_id =? "+
		"ORDER BY episodes_update_time DESC LIMIT 10", channelId, typeId).QueryRows(&videos)
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
		querySeter = querySeter.OrderBy("episodes_update_time")
	} else if sort == SortComment {
		querySeter = querySeter.OrderBy("comment")
	} else if sort == SortAddTime {
		querySeter = querySeter.OrderBy("add_time")
	} else {
		querySeter = querySeter.OrderBy("add_time")
	}

	// 获取总条数
	count, _ := querySeter.Values(&params, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")

	// 获取指定的Limit数据
	querySeter = querySeter.Limit(limit, offset)
	_, err := querySeter.Values(&params, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	return count, params, err
}

// GetVideoInfo 通过videoId获取视频详情信息
func GetVideoInfo(videoID int) (Video, error) {
	newOrm := orm.NewOrm()
	var video Video
	err := newOrm.Raw("SELECT * FROM video where id =? ", videoID).QueryRow(&video)
	return video, err
}

// GetCacheVideoInfo 通过videoId获取视频详情信息 - Redis获取
func GetCacheVideoInfo(videoID int) (videoInfo Video, err error) {
	// 获取视频
	var video Video
	// 获取video连接
	connect := redisClient.PoolConnect()
	// 异常关闭
	defer connect.Close()
	// 通过key获取视频信息
	videoRedisKey := "video:id:" + strconv.Itoa(videoID)

	// 判断是否存在
	videoExists, err := redis.Bool(connect.Do("exists", videoRedisKey))
	if videoExists {
		values, _ := redis.Values(connect.Do("hgetall", videoRedisKey))
		// 如果没问题，就直接发出去，否则查数据
		err = redis.ScanStruct(values, &video)
	} else {
		video, err = GetVideoInfo(videoID)
		if err == nil {
			// 把数据存入redis
			_, err := connect.Do("hmset", redis.Args{videoRedisKey}.AddFlat(video)...)
			if err == nil {
				// 这是一个约定，存入的时候，记得设置过期时间
				connect.Do("expire", videoRedisKey, 86400)
			}
		}
	}
	return video, err
}

// GetCacheVideoEpisodes 通过videoId获取视频剧集
func GetCacheVideoEpisodes(videoID int) (int64, []Episodes, error) {
	// 获取视频
	var (
		err      error
		episodes []Episodes
		len      int64
	)

	// 获取video连接
	connect := redisClient.PoolConnect()
	// 异常关闭
	defer connect.Close()
	// 通过key获取视频信息
	videoRedisKey := "video:episodes:id:" + strconv.Itoa(videoID)

	// 判断是否存在
	videoExists, err := redis.Bool(connect.Do("exists", videoRedisKey))
	if videoExists {
		len, err = redis.Int64(connect.Do("llen", videoRedisKey))
		if err == nil {
			values, _ := redis.Values(connect.Do("lrange", videoRedisKey, "0", "-1"))
			var episodesInfo Episodes
			for _, value := range values {
				err := json.Unmarshal(value.([]byte), &episodesInfo)
				if err == nil {
					episodes = append(episodes, episodesInfo)
				}
			}
		}
	} else {
		len, episodes, err = GetVideoEpisodes(videoID)
		if err == nil {
			for _, episode := range episodes {
				episodeJson, err := json.Marshal(episode)
				if err == nil {
					connect.Do("rpush", videoRedisKey, episodeJson)
				}
			}
			connect.Do("expire", videoRedisKey, 86400)
		}
	}
	return len, episodes, err
}

// GetVideoEpisodes 通过videoId获取视频剧集 - redis 缓存
func GetVideoEpisodes(videoID int) (int64, []Episodes, error) {
	newOrm := orm.NewOrm()
	var episodes []Episodes
	rows, err := newOrm.Raw("SELECT id, title, add_time, num, play_url, comment FROM video_episodes where video_id = ? AND status =1 ORDER BY num ASC;", videoID).QueryRows(&episodes)
	return rows, episodes, err
}

// GetChannelTop 通过channelId获取视频的排行
func GetChannelTop(channelId int) (int64, []VideoDate, error) {
	o := orm.NewOrm()
	var videos []VideoDate

	rows, err := o.Raw("SELECT id, title, sub_title, "+
		"channel_id, add_time, img, img1 ,episodes_count ,comment "+
		"FROM video WHERE status=1 AND channel_id=?  "+
		"ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
	return rows, videos, err
}

// GetCacheChannelTop 通过channelId获取视频的排行
func GetCacheChannelTop(channelId int) (int64, []VideoDate, error) {
	var (
		rows []VideoDate
		num  int64
		err  error
	)
	// redis
	conn := redisClient.PoolConnect()
	defer conn.Close()
	// key
	channelIdKey := "video:top:channel:channelId:" + strconv.Itoa(channelId)
	// 获取数据
	exists, err := redis.Bool(conn.Do("exists", channelIdKey))
	if exists {
		// zrevrange ->  // id value
		values, _ := redis.Values(conn.Do("zrevrange", channelIdKey, "0", "10", "WITHSCORES"))
		for index, value := range values {
			// 这一个是ID,否则是comment的数据 - score
			if index%2 == 0 {
				videoId, err := strconv.Atoi(string(value.([]byte)))
				videoInfo, err := GetCacheVideoInfo(videoId)
				if err == nil {
					// 转换数据
					var videoDataInfo VideoDate
					videoDataInfo.Id = videoInfo.Id
					videoDataInfo.Title = videoInfo.Title
					videoDataInfo.SubTitle = videoInfo.SubTitle
					videoDataInfo.AddTime = videoInfo.AddTime
					videoDataInfo.Img = videoInfo.Img
					videoDataInfo.Img1 = videoInfo.Img1
					videoDataInfo.EpisodesCount = videoInfo.EpisodesCount
					videoDataInfo.IsEnd = videoInfo.IsEnd
					videoDataInfo.Comment = videoInfo.Comment

					rows = append(rows, videoDataInfo)
					num++
				}
			}
		}

	} else {
		// mysql
		num, rows, err = GetChannelTop(channelId)
		if err == nil {
			// 存入redis
			for _, videoDate := range rows {
				conn.Do("zadd", channelIdKey, videoDate.Comment, videoDate.Id)
			}
			conn.Do("expire", channelIdKey, 30*86400)
		}
	}

	return num, rows, err
}

// GetCacheTypeTop 通过typeId获取视频的排行
func GetCacheTypeTop(typeId int) (int64, []VideoDate, error) {
	var (
		rows []VideoDate
		num  int64
		err  error
	)
	// 连接、关闭
	conn := redisClient.PoolConnect()
	defer conn.Close()

	// 是否存在
	typeIdKey := "video:top:type:typeId:" + strconv.Itoa(typeId)
	exists, err := redis.Bool(conn.Do("exists", typeIdKey))
	if exists {
		// 如果存在
		values, _ := redis.Values(conn.Do("zrevrange", typeIdKey, "0", "10", "WITHSCORES"))
		for index, value := range values {
			if index%2 == 0 {
				videoId, err := strconv.Atoi(string(value.([]byte)))
				videoInfo, err := GetCacheVideoInfo(videoId)
				if err == nil {
					// 转换数据
					var videoDataInfo VideoDate
					videoDataInfo.Id = videoInfo.Id
					videoDataInfo.Title = videoInfo.Title
					videoDataInfo.SubTitle = videoInfo.SubTitle
					videoDataInfo.AddTime = videoInfo.AddTime
					videoDataInfo.Img = videoInfo.Img
					videoDataInfo.Img1 = videoInfo.Img1
					videoDataInfo.EpisodesCount = videoInfo.EpisodesCount
					videoDataInfo.IsEnd = videoInfo.IsEnd
					videoDataInfo.Comment = videoInfo.Comment

					rows = append(rows, videoDataInfo)
					num++
				}
			}
		}
	} else {
		// mysql
		num, rows, err = GetTypeTop(typeId)
		if err == nil {
			// 存入redis
			for _, videoDate := range rows {
				conn.Do("zadd", typeIdKey, videoDate.Comment, videoDate.Id)
			}
			conn.Do("expire", typeIdKey, 30*86400)
		}
	}
	return num, rows, err
}

// GetTypeTop 通过typeId获取视频的排行
func GetTypeTop(typeId int) (int64, []VideoDate, error) {
	o := orm.NewOrm()
	var videos []VideoDate

	rows, err := o.Raw("SELECT id, title, sub_title, "+
		"channel_id, add_time, img, img1 ,episodes_count ,comment "+
		"FROM video WHERE status=1 AND type_id=?  "+
		"ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
	return rows, videos, err
}

// GetAllVideoList 获取所有的视频信息
func GetAllVideoList() (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	rows, err := o.Raw("SELECT id, title, sub_title, status, " +
		"add_time, img, img1 ,channel_id, type_id,region_id,user_id, " +
		"episodes_count ,episodes_update_time,is_end,is_hot,is_recommend, comment " +
		"FROM video").QueryRows(&videos)

	return rows, videos, err
}
