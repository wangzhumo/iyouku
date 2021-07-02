package main

import (
	_ "com.wangzhumo.iyouku/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defaultDB, _ := beego.AppConfig.String("defaultDb")
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return
	}
	err = orm.RegisterDataBase("default", "mysql", defaultDB)
	if err != nil {
		return
	}

	beego.Run()
}
