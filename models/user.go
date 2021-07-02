package models

import (
	"com.wangzhumo.iyouku/common"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	UID      int64 `orm:"column(uid);pk"`
	Nick     string
	Name     string
	Password string
	Status   int
	AddTime  int64 `orm:"column(create)"`
	Mobile   string
	Avatar   string
}

func (u *User) TableName() string {
	return "ucenter"
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
	err := o.Read(&user, "mobile")
	if err == orm.ErrNoRows {
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
	uid, _ := common.GenerateUid()
	user := User{
		Mobile:   mobile,
		Password: encodePsd,
		Status:   1,
		Nick:     "",
		Name:     "",
		AddTime:  time.Now().Unix(),
		UID:      uid,
	}

	// 存入数据库
	_, err = o.Insert(&user)
	return
}

// UserLogin 用户登录查询
func UserLogin(mobile string, password string) (int64, string) {
	o := orm.NewOrm()
	var user User
	// 查询用户信息
	err := o.QueryTable("ucenter").Filter("mobile",
		mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}
	return user.UID, user.Nick
}
