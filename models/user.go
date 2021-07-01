package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id       int64
	Nick     string
	Name     string
	Password string
	Status   int
	AddTime  int64
	Mobile   string
	Avatar   string
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func init() {
	orm.RegisterModel(new(User))
}

// IsPhoneRegister 手机号是否已经注册
func IsPhoneRegister(mobile string) (status bool) {
	// 获取orm
	o := orm.NewOrm()
	user := User{Mobile: mobile}

	// 查询数据
	err := o.Read(&user, "Mobile")
	if err != orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}
	return true
}

// UserSave 创建用户
func UserSave(mobile string, encodePsd string) (err error) {
	// 获取orm
	o := orm.NewOrm()
	user := User{
		Mobile:   mobile,
		Password: encodePsd,
		Status:   1,
		Nick:     "",
		Name:     "",
		AddTime:  time.Now().Unix(),
	}

	// 存入数据库
	_, err = o.Insert(user)
	return
}
