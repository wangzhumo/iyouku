package main

import (
	"com.wangzhumo.iyouku/common"
	"com.wangzhumo.iyouku/models"
	rabbitmqClient "com.wangzhumo.iyouku/services/rabbitmq"
	redisClient "com.wangzhumo.iyouku/services/redis"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"strconv"
)

func main() {
	//不在外面的beego范围内，这里指定配置文件
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	goarch := runtime.GOOS
	defaultDB, _ := beego.AppConfig.String("defaultDb")
	if goarch == "darwin" {
		defaultDB, _ = beego.AppConfig.String("darwinDb")
	} else {
		defaultDB, _ = beego.AppConfig.String("windowsDb")
	}
	fmt.Println("defaultDB = " + defaultDB)
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return
	}
	err = orm.RegisterDataBase("default", "mysql", defaultDB)
	if err != nil {
		return
	}

	// 接受RabbitMQ中Top消息
	rabbitmqClient.Consumer("", common.TopQueue, callback)
}

// mq的回调
func callback(msg string) {
	type Data struct {
		VideoId int
	}
	// 解析
	var data Data
	err := json.Unmarshal([]byte(msg), &data)
	videoInfo, err := models.GetCacheVideoInfo(data.VideoId)
	if err == nil {
		// 更新Redis的信息
		connect := redisClient.PoolConnect()
		defer connect.Close()

		// 需要更新两个排行榜
		redisChannelKey := "video:top:channel:channelId:" + strconv.Itoa(videoInfo.ChannelId)
		redisTypeKey := "video:top:type:typeId:" + strconv.Itoa(videoInfo.TypeId)
		redisVideoKey := "video:id:" + strconv.Itoa(videoInfo.Id)

		// 直接更新即可
		// ZINCRBY key increment member
		// 自减	   key 自减的数量  自减的成员项
		// zincrby  让指定的key它的数字值，减去传入数据
		connect.Do("zincrby", redisChannelKey, 1, data.VideoId)
		// 就是说让redisChannelKey这个key中 videoID 这个成员 减去 1
		connect.Do("zincrby", redisTypeKey, 1, data.VideoId)
		_, err := connect.Do("hincrby", redisVideoKey, "Comment", 1)
		if err != nil {
			return
		}
	}

}
