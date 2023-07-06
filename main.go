package main

import (
	_ "github.com/Sathya1099/beego/models"
	_ "github.com/Sathya1099/beego/routers"
	"github.com/beego/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"orm.log","level":7,"maxlines":0,"maxsize":0,"daily":false,"maxdays":1,"color":true}`)
	if err != nil {
		panic(err)
	}
	logs.Info("hello beego")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
