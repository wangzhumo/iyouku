package models

import (
	"com.wangzhumo.iyouku/common"
	rabbitmqClient "com.wangzhumo.iyouku/services/rabbitmq"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Comment struct {
	Id           int
	Content      string
	AddTime      int64
	AddTimeTitle string `orm:"-"`
	UserId       int
	Stamp        int
	Status       int
	PraiseCount  int
	EpisodesId   int
	VideoId      int
}

func init() {
	orm.RegisterModel(new(Comment))
}

func (u *Comment) TableName() string {
	return "comment"
}

// GetCommentsByEpisodesId 获取剧集的评论列表
func GetCommentsByEpisodesId(episodesId int, offset int, limit int) (int64, []Comment, error) {
	newOrm := orm.NewOrm()
	var comments []Comment
	queryRows, _ := newOrm.Raw("SELECT id FROM comment WHERE "+
		"status=1 AND episodes_id?", episodesId).QueryRows(&comments)
	_, err := newOrm.Raw("SELECT id, content, add_time, user_id, "+
		"stamp, praise_count, episodes_id FROM comment "+
		"WHERE status=1 AND episodes_id=? "+
		"ORDER BY add_time DESC  LIMIT ?,?", episodesId, offset, limit).QueryRows(&comments)

	return queryRows, comments, err
}

// SaveComment 保存一条评论
func SaveComment(context string, userId int, episodesId int, videoId int) (err error) {
	newOrm := orm.NewOrm()
	unixTime := time.Now().Unix()
	var comment = Comment{
		Content:    context,
		UserId:     userId,
		EpisodesId: episodesId,
		VideoId:    videoId,
		Stamp:      0,
		Status:     1,
		AddTime:    unixTime}

	_, err = newOrm.Insert(&comment)

	if err == nil {
		// 修改视频的总评论数
		newOrm.Raw("UPDATE video SET comment=comment+1 WHERE id=?", videoId).Exec()
		// 修改剧集评论数
		newOrm.Raw("UPDATE video_episodes SET comment=comment+1 WHERE id=?", episodesId).Exec()
		// 更新Redis - 通过MQ来实现(简单模式下传递)
		videoObj := map[string]int{
			"VideoId": videoId,
		}
		marshal, _ := json.Marshal(videoObj)
		// 发送给mq
		_ = rabbitmqClient.Publish("", common.TopQueue, string(marshal))
	}
	return
}

// SaveCommentCount 更新评论数
func SaveCommentCount(episodesId int, videoId int) (err error) {
	newOrm := orm.NewOrm()

	// 修改视频的总评论数
	newOrm.Raw("UPDATE video SET comment=comment+1 WHERE id=?", videoId).Exec()
	// 修改剧集评论数
	newOrm.Raw("UPDATE video_episodes SET comment=comment+1 WHERE id=?", episodesId).Exec()
	// 更新Redis - 通过MQ来实现(简单模式下传递)
	videoObj := map[string]int{
		"VideoId": videoId,
	}
	marshal, err := json.Marshal(videoObj)
	// 发送给mq
	err = rabbitmqClient.Publish("", common.TopQueue, string(marshal))
	return
}

// SaveCommentEx 保存一条评论  - 追加评论（浏览量啊,热度值啊）
func SaveCommentEx(context string, userId int, episodesId int, videoId int) (err error) {
	newOrm := orm.NewOrm()
	unixTime := time.Now().Unix()
	var comment = Comment{
		Content:    context,
		UserId:     userId,
		EpisodesId: episodesId,
		VideoId:    videoId,
		Stamp:      0,
		Status:     1,
		AddTime:    unixTime}

	_, err = newOrm.Insert(&comment)

	if err == nil {
		// 修改视频的总评论数
		newOrm.Raw("UPDATE video SET comment=comment+1 WHERE id=?", videoId).Exec()
		// 修改剧集评论数
		newOrm.Raw("UPDATE video_episodes SET comment=comment+1 WHERE id=?", episodesId).Exec()
		// 更新Redis - 通过MQ来实现(简单模式下传递)
		videoObj := map[string]int{
			"VideoId": videoId,
		}
		marshal, _ := json.Marshal(videoObj)
		// 发送给mq
		_ = rabbitmqClient.Publish("", common.TopQueue, string(marshal))

		// 追加的mq
		videoExObj := map[string]int{
			"VideoId":    videoId,
			"EpisodesId": episodesId,
		}
		videoExJson, _ := json.Marshal(videoExObj)
		// 发送到死信队列
		_ = rabbitmqClient.PublishDlx(common.CommentQueue, string(videoExJson))
	}
	return
}
