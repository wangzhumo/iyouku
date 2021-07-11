package redisClient

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
	"time"
)

// PoolConnect 获取Redis的链接
func PoolConnect() redis.Conn {
	s, _ := beego.AppConfig.String("redispsd")
	host, _ := beego.AppConfig.String("redisdp")
	redisPsd := redis.DialPassword(s)

	redisPool := &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 100 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host, redisPsd)
		},
	}
	return redisPool.Get()
}
