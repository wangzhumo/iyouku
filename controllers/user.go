package controllers

import (
	"com.wangzhumo.iyouku/models"
	beego "github.com/beego/beego/v2/server/web"
	"regexp"
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
	//matchPassword, err := regexp.MatchString(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`, password)
	//if !matchPassword || err != nil {
	//	uc.Data["json"] = ErrorResp(4004, PasswordFormatError)
	//	_ = uc.ServeJSON()
	//}

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

// SendPushMessage 发送推送消息
// @router /send/message
func (uc *UserController) SendPushMessage(){
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
	// models.SaveComment(content)

}