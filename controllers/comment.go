package controllers

import (
	"com.wangzhumo.iyouku/common"
	"com.wangzhumo.iyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

type CommentController struct {
	beego.Controller
}

// CommentInfo 返回结构需要调整
type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userInfo"`
}

// GetCommentList 获取评论列表
// @router /comment/list [*]
func (cc *CommentController) GetCommentList() {
	// 剧集id
	episodesId, _ := cc.GetInt("episodesId")
	// 分页信息
	limit, _ := cc.GetInt("limit")
	offset, _ := cc.GetInt("offset")

	// empty
	if episodesId == 0 {
		cc.Data["json"] = ErrorResp(4001, NoEpisodesID)
		_ = cc.ServeJSON()
	}
	if limit == 0 {
		limit = 10
	}

	// get data
	count, comments, err := models.GetCommentsByEpisodesId(episodesId, offset, limit)
	if err != nil {
		cc.Data["json"] = ErrorResp(4004, VideoEpisodesCommentError)
		_ = cc.ServeJSON()
	} else {
		var commentInfos []CommentInfo
		var commentInfo CommentInfo
		// 获取评论列表中用户的信息
		for _, comment := range comments {
			// 评论
			commentInfo.Id = comment.Id
			commentInfo.Content = comment.Content
			commentInfo.AddTime = comment.AddTime
			commentInfo.AddTimeTitle = common.DateFormat(comment.AddTime)
			commentInfo.PraiseCount = comment.PraiseCount
			commentInfo.UserId = comment.UserId
			commentInfo.Stamp = comment.Stamp

			// 获取用户信息
			commentInfo.UserInfo, _ = models.GetCacheUserInfo(comment.UserId)
			commentInfos = append(commentInfos, commentInfo)
		}
		cc.Data["json"] = SucceedResp(0, RequestOk, commentInfos, count)
		_ = cc.ServeJSON()
	}
}

// GetGoCommentList 获取评论列表 使用协程
// @router /comment/ex/list [*]
func (cc *CommentController) GetGoCommentList() {
	// 剧集id
	episodesId, _ := cc.GetInt("episodesId")
	// 分页信息
	limit, _ := cc.GetInt("limit")
	offset, _ := cc.GetInt("offset")

	// empty
	if episodesId == 0 {
		cc.Data["json"] = ErrorResp(4001, NoEpisodesID)
		_ = cc.ServeJSON()
	}
	if limit == 0 {
		limit = 10
	}

	// get data
	count, comments, err := models.GetCommentsByEpisodesId(episodesId, offset, limit)
	if err != nil {
		cc.Data["json"] = ErrorResp(4004, VideoEpisodesCommentError)
		_ = cc.ServeJSON()
	} else {
		var commentInfos []CommentInfo
		var commentInfo CommentInfo
		// 获取评论列表中用户的信息
		// 10 是因为 limit = 10
		uidChan := make(chan int, 10)
		closeChan := make(chan bool, 5)
		taskChan := make(chan models.UserInfo, 10)

		// uid 发送至 uidChannel
		go func() {
			for _, comment := range comments {
				uidChan <- comment.UserId
			}
			close(uidChan)
		}()

		// 发起用户信息获取任务
		for i := 0; i < 10; i++ {
			go chanGetUserInfo(uidChan, closeChan, taskChan)
		}

		// 判断任务是否完成
		go func() {
			for i := 0; i < 10; i++ {
				<-closeChan
				// 阻塞
			}
			close(taskChan)
			close(closeChan)
		}()

		// 获取taskChan中的用户数据
		userInfoMap := make(map[int]models.UserInfo)

		// channel在for循环中，有一个特性  当channel关闭时，会跳出这个循环
		for c := range taskChan {
			userInfoMap[c.Id] = c
		}

		// 获取评论列表中用户的信息
		for _, comment := range comments {
			// 评论
			commentInfo.Id = comment.Id
			commentInfo.Content = comment.Content
			commentInfo.AddTime = comment.AddTime
			commentInfo.AddTimeTitle = common.DateFormat(comment.AddTime)
			commentInfo.PraiseCount = comment.PraiseCount
			commentInfo.UserId = comment.UserId
			commentInfo.Stamp = comment.Stamp

			// 获取用户信息
			commentInfo.UserInfo, _ = userInfoMap[comment.UserId]
			commentInfos = append(commentInfos, commentInfo)
		}

		cc.Data["json"] = SucceedResp(0, RequestOk, commentInfos, count)
		_ = cc.ServeJSON()
	}
}

// Channel 获取用户信息
func chanGetUserInfo(uidChan chan int, closeChan chan bool, taskChan chan models.UserInfo) {
	for uid := range uidChan {
		userInfo, err := models.GetCacheUserInfo(uid)
		// 发送回去
		if err == nil {
			taskChan <- userInfo
		}
	}
	closeChan <- true
}

// InsertComment 写入一条评论
// @router /comment/save [*]
func (cc *CommentController) InsertComment() {
	// 评论内容
	content := cc.GetString("content")
	// 用户
	uid, _ := cc.GetInt("uid")
	episodesId, _ := cc.GetInt("episodesId")
	videoId, _ := cc.GetInt("videoId")

	// empty check
	if len(content) == 0 {
		cc.Data["json"] = ErrorResp(4001, NoContent)
		_ = cc.ServeJSON()
	}
	if uid == 0 {
		cc.Data["json"] = ErrorResp(2001, RequireLogin)
		_ = cc.ServeJSON()
	}
	if episodesId == 0 {
		cc.Data["json"] = ErrorResp(4001, NoEpisodesID)
		_ = cc.ServeJSON()
	}
	if videoId == 0 {
		cc.Data["json"] = ErrorResp(4001, NoVideoID)
		_ = cc.ServeJSON()
	}

	// 保存到数据库
	err := models.SaveComment(content, uid, episodesId, videoId)
	if err != nil {
		cc.Data["json"] = ErrorResp(5000, CommentInsertError)
		_ = cc.ServeJSON()
	} else {
		cc.Data["json"] = SucceedResp(0, RequestOk, nil, 1)
		_ = cc.ServeJSON()
	}
}

// InsertCommentEx 写入一条评论
// @router /comment/ex/save [*]
func (cc *CommentController) InsertCommentEx() {
	// 评论内容
	content := cc.GetString("content")
	// 用户
	uid, _ := cc.GetInt("uid")
	episodesId, _ := cc.GetInt("episodesId")
	videoId, _ := cc.GetInt("videoId")

	// empty check
	if len(content) == 0 {
		cc.Data["json"] = ErrorResp(4001, NoContent)
		_ = cc.ServeJSON()
	}
	if uid == 0 {
		cc.Data["json"] = ErrorResp(2001, RequireLogin)
		_ = cc.ServeJSON()
	}
	if episodesId == 0 {
		cc.Data["json"] = ErrorResp(4001, NoEpisodesID)
		_ = cc.ServeJSON()
	}
	if videoId == 0 {
		cc.Data["json"] = ErrorResp(4001, NoVideoID)
		_ = cc.ServeJSON()
	}

	// 保存到数据库
	err := models.SaveCommentEx(content, uid, episodesId, videoId)
	if err != nil {
		cc.Data["json"] = ErrorResp(5000, CommentInsertError)
		_ = cc.ServeJSON()
	} else {
		cc.Data["json"] = SucceedResp(0, RequestOk, nil, 1)
		_ = cc.ServeJSON()
	}
}
