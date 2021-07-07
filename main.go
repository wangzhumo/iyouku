package main

import (
	_ "com.wangzhumo.iyouku/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
)

func main() {
	goarch := runtime.GOOS
	defaultDB,_ := beego.AppConfig.String("defaultDb")
	if goarch == "darwin" {
		defaultDB, _ = beego.AppConfig.String("darwinDb")
	}else{
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

	beego.Run()
}
