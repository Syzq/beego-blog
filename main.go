package main

import (
	"beego-blog/pkg/setting"
	_ "beego-blog/routers"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	_ "beego-blog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	beego.Run()
}
func init() {
	sec := setting.Cfg.Section("database")

	dbType := sec.Key("TYPE").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	dbName := sec.Key("NAME").String()

	// set default database
	orm.RegisterDataBase("default", dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8%s", user, password, host, dbName, "&loc=Asia%2FShanghai"), 30)

	// create table
	orm.RunSyncdb("default", false, true)
}
