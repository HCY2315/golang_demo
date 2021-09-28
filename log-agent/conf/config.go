package conf

import (
	"fmt"

	"github.com/go-ini/ini"
)

type Kafka struct {
	Address  string
	Topic    string
	Chan_max int
}

var KafkaSetting = &Kafka{}

type Etcd struct {
	Address string
	Timeout int
	Key_log string
}

var EtcdSetting = &Etcd{}

// -------- unused
type Tail struct {
	Path string
}

var TailSetting = &Tail{}

var cfg *ini.File

func Init() {
	var err error
	cfg, err = ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("init load config.ini err:", err)
		return
	}
	mapto("kafka", KafkaSetting)
	mapto("etcd", EtcdSetting)
	mapto("tail", TailSetting)
}

func mapto(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Println("Cfg.Mapto is err:", err)
		return
	}
}
