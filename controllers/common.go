package controllers

import (
	"crypto/md5"
	"encoding/hex"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

// BaseResponse 返回的结构体
type BaseResponse struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"item"`
	Count int64       `json:"count"`
}

// SucceedResp 正确的返回
func SucceedResp(code int, msg interface{}, items interface{}, count int64) (json *BaseResponse) {
	json = &BaseResponse{
		Code:  code,
		Msg:   msg,
		Items: items,
		Count: count,
	}
	return
}

// ErrorResp 错误Response返回
func ErrorResp(code int, msg interface{}) (json *BaseResponse) {
	json = &BaseResponse{
		Code: code,
		Msg:  msg,
	}
	return
}

// Md5Psd 加密密码
func Md5Psd(psd string) string {
	h := md5.New()
	s, err := beego.AppConfig.String("md5code")
	if err != nil {
		s = "emptyMd5code"
	}

	h.Write([]byte(psd + s))
	return hex.EncodeToString(h.Sum(nil))
}

// DateFormat 格式化时间
func DateFormat(date int64) string {
	unix := time.Unix(date, 0)
	return unix.Format("2006-01-02")
}
