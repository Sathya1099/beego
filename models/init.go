package models

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
)

var ormer orm.Ormer

func init() {
	// orm.Debug = true
	orm.RegisterModel(new(Object))
	user, _ := config.String("dbuser")
	log.Println("17", user)
	pass, _ := config.String("dbpass")
	host, _ := config.String("dbhost")
	port, _ := config.String("dbport")
	dbname, _ := config.String("dbname")
	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			user, pass, host, port, dbname,
		))
	orm.RunSyncdb("default", false, true)
	ormer = orm.NewOrm()
}
