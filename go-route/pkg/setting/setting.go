package setting

import (
	"log"

	"github.com/go-ini/ini"
)

// Mysql Server
type Mysql struct {
	Address  string
	Port     int
	Database string
	Password string
	User     string
}

var MysqlSetting = &Mysql{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		log.Fatalf("setting Setup, fail to parse 'conf/config.ini': %v", err)
		return
	}
	mapto("Mysql", MysqlSetting)
}

func mapto(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.mapto setting is err: %v", err)
	}
}
