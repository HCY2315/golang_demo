package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

type Database struct {
	Type       string
	Host       string
	Port       string
	User       string
	Passwd     string
	Table_name string
}

var DatabaseSetting = &Database{}

var Cfg *ini.File

// 初始化配置文件
func Setup() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Setting is error:%v", err)
	}
	mapTo("Database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo is error: %v", err)
	}
}
