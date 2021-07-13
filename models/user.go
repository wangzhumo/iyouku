package models

import (
	"com.wangzhumo.iyouku/common"
	redisClient "com.wangzhumo.iyouku/services/redis"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

// User self
type User struct {
	Id       int64
	Name     string
	Password string
	AddTime  int64
	Status   int
	Mobile   string
	Avatar   string
}

// UserInfo normal user info
type UserInfo struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	AddTime string `json:"addTime"`
	Avatar  string `json:"avatar"`
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
		Name:     "",
		AddTime:  time.Now().Unix(),
		Id:       uid,
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
	err := o.QueryTable("user").Filter("mobile",
		mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}
	return user.Id, user.Name
}

// GetUserInfo 查询用户信息（normal）
func GetUserInfo(userId int64) (UserInfo, error) {
	o := orm.NewOrm()
	var user UserInfo
	err := o.Raw("SELECT id,name,add_time,avatar "+
		"FROM user WHERE id=? LIMIT 1", userId).QueryRow(&user)
	return user, err
}

// GetCacheUserInfo 获取用户数据
func GetCacheUserInfo(userId int64) (userInfo UserInfo, err error) {
	var user UserInfo
	// 连接redis
	conn := redisClient.PoolConnect()
	defer conn.Close()
	// 通过key获取视频信息
	userRedisKey := "video:id:" + strconv.FormatInt(userId, 10)
	exists, err := redis.Bool(conn.Do("exists", userRedisKey))
	if exists {
		// 存在就直接获取，而不是数据库查询
		values, _ := redis.Values(conn.Do("hgetall", userRedisKey))
		err = redis.ScanStruct(values, &user)
	} else {
		user, err = GetUserInfo(userId)
		if err == nil {
			_, err := conn.Do("hmset", redis.Args{userRedisKey}.AddFlat(user)...)
			if err == nil {
				conn.Do("expire", userRedisKey, 86400)
			}
		}
	}
	return user, err
}
