package main

import (
	"com.wangzhumo.iyouku/common"
	"com.wangzhumo.iyouku/models"
	rabbitmqClient "com.wangzhumo.iyouku/services/rabbitmq"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
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

	// 接受RabbitMQ中comment的Dlx消息
	rabbitmqClient.ConsumerDlx(common.CommentExchange,
		common.CommentQueue, common.CommentDlxExchange,
		common.CommentDlxQueue, 10, callback)
}

// mq的回调
func callback(msg string) {
	type Data struct {
		VideoId    int
		EpisodesId int
	}
	// 解析
	var data Data
	err := json.Unmarshal([]byte(msg), &data)

	// 直接写入数据库即可
	if err == nil {
		models.SaveCommentCount(data.EpisodesId, data.VideoId)
	}
}
