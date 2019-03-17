package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.conf")

	if err != nil {
		log.Fatalf("读取配置文件 'conf/app.conf' 失败： %v", err)
	}

}
