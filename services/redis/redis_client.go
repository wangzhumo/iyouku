package redisClient

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
	"time"
)

// PoolConnect 获取Redis的链接
func PoolConnect() redis.Conn {
	s, _ := beego.AppConfig.String("redispsd")
	host, _ := beego.AppConfig.String("redisdb")
	redisPsd := redis.DialPassword(s)

	redisPool := &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 100 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", host, redisPsd)
			if err != nil {
				return dial, err
			}
			//if _,err := dial.Do("AUTH",redisPsd); err != nil {
			//	dial.Close()
			//	return nil, err
			//}
			return dial, err
		},
	}
	return redisPool.Get()
}
