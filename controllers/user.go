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
			uc.Data["json"] = SucceedResp(5000, RegisterSucceed, nil, 0)
			_ = uc.ServeJSON()
		} else {
			uc.Data["json"] = ErrorResp(0, RegisterFail)
			_ = uc.ServeJSON()
		}
	}
}
