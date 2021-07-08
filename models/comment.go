package models

import "github.com/beego/beego/v2/client/orm"

type Comment struct {
	Id           int
	Content      string
	AddTime      int64
	AddTimeTitle string
	UserId       int
	Stamp        int
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
