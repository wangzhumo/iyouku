package controllers

import (
	"regexp"
	"strconv"
	"strings"

	"com.wangzhumo.iyouku/models"
	beego "github.com/beego/beego/v2/server/web"
)

// UserController Operations about Users
type UserController struct {
	beego.Controller
}

// SaveRegister 用户注册功能
// @router /register/save [get]
func (uc *UserController) SaveRegister() {
	// 定义接受的参数
	var (
		mobile   string
		password string
		err      error
	)
	// 获取参数
	mobile = uc.GetString("mobile")
	password = uc.GetString("password")

	// 验证手机号是否合法
	// empty
	if mobile == "" {
		uc.Data["json"] = ErrorResp(4001, PhoneEmpty)
		_ = uc.ServeJSON()
	}
	// error format
	matchPhone, err := regexp.MatchString(`^1(2|3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !matchPhone {
		uc.Data["json"] = ErrorResp(4001, PhoneFormatError)
		_ = uc.ServeJSON()
	}

	// empty psd
	if password == "" {
		uc.Data["json"] = ErrorResp(4003, PasswordEmpty)
		_ = uc.ServeJSON()
	}

	// check phone register status
	registerStatus := models.IsPhoneRegister(mobile)
	if registerStatus {
		uc.Data["json"] = ErrorResp(4005, PhoneAlreadyRegister)
		_ = uc.ServeJSON()
	} else {
		// 否则开始注册流程
		err = models.UserSave(mobile, Md5Psd(password))
		if err == nil {
			uc.Data["json"] = SucceedResp(0, RegisterSucceed, nil, 0)
			_ = uc.ServeJSON()
		} else {
			uc.Data["json"] = ErrorResp(5000, RegisterFail)
			_ = uc.ServeJSON()
		}
	}
}

// Login  用户登陆功能
// @router /login/do [get]
func (uc *UserController) Login() {
	// 定义接受的参数
	var (
		mobile   string
		password string
	)
	// 获取参数
	mobile = uc.GetString("mobile")
	password = uc.GetString("password")

	// 验证手机号是否合法
	// empty
	if mobile == "" {
		uc.Data["json"] = ErrorResp(4001, PhoneEmpty)
		_ = uc.ServeJSON()
	}

	// empty psd
	if password == "" {
		uc.Data["json"] = ErrorResp(4003, PasswordEmpty)
		_ = uc.ServeJSON()
	}

	// 查询用户
	uid, nick := models.UserLogin(mobile, Md5Psd(password))
	if uid > 0 {
		uc.Data["json"] = SucceedResp(0, LoginSucceed,
			map[string]interface{}{
				"id":   uid,
				"name": nick,
			},
			1)
		_ = uc.ServeJSON()
	} else {
		uc.Data["json"] = ErrorResp(5001, LoginFail)
		_ = uc.ServeJSON()
	}
}

// SendData 数据结构 - 给channel使用
type SendData struct {
	UserId    int
	MessageId int64
}

// SendPushMessage 发送推送消息
// @router /send/message [post]
func (uc *UserController) SendPushMessage() {
	// 评论内容
	content := uc.GetString("content")
	// 用户
	uids := uc.GetString("uids")

	// empty check
	if len(content) == 0 {
		uc.Data["json"] = ErrorResp(4001, NoContent)
		_ = uc.ServeJSON()
	}
	if len(uids) == 0 {
		uc.Data["json"] = ErrorResp(2001, RequireUids)
		_ = uc.ServeJSON()
	}

	// 保存数据
	messageId, err := models.SendMessage(content)
	if err != nil {
		uc.Data["json"] = ErrorResp(2001, SendMessageError)
		_ = uc.ServeJSON()
	} else {
		// 设置到指定的用户上去
		splitUid := strings.Split(uids, ",")
		// 使用协程来处理
		length := len(splitUid)
		// 创建sendData使用的channel
		sendChan := make(chan SendData, length)
		// 创建关闭的Channel
		closeChan := make(chan bool, length)

		// 发送UID 以及 MessageId 的消息
		go func() {
			var data SendData
			for _, uid := range splitUid {
				userId, _ := strconv.Atoi(uid)
				data.UserId = userId
				data.MessageId = messageId
				// 发送出去
				sendChan <- data
			}
			close(sendChan)
		}()

		// 执行发送过来的任务
		for i := 0; i < 5; i++ {
			// 意思是起了5个协程
			go sendMessageWithGo(sendChan, closeChan)
		}

		// 关闭
		for i := 0; i < 5; i++ {
			<-closeChan
		}
		close(closeChan)

		// 直接返回即可
		uc.Data["json"] = SucceedResp(0, RequestOk, nil, 1)
		_ = uc.ServeJSON()
	}

}

// 发送RabbitMQ消息
func sendMessageWithGo(uData chan SendData, cc chan bool) {
	for sendData := range uData {
		//_, _ = models.SendMessageToUser(messageId, userId)
		models.SendMQMessageToUser(sendData.MessageId, sendData.UserId)
	}
	cc <- true
}
