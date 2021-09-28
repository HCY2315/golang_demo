package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

//设置
type Logger struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// [logger]
// LogSavePath = logs/
// LogSaveName = go-logging
// LogFileExt = log
// TimeFormat = 20060102

//重命名，取内存变量
var LoggerSetting = &Logger{}

var cfg *ini.File

//获取app.ini配置文件中的相关信息
func Setup() {
	var err error
	//从INI数据源加载和解析。
	//参数可以是文件名与字符串类型混合，也可以是[]字节中的原始数据。
	//如果列表包含不存在的文件，则返回错误。
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Setting Setup: failed to parse 'conf/api.ini':%v", err)
	}
	// "Logger" 为 app.ini 文件中的名字"
	mapTo("Logger", LoggerSetting)
}

func mapTo(section string, v interface{}) {
	//节假定命名节存在，如果不存在，则返回零值。
	//MapTo将节映射到给定的结构。
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.Mapto is err:%v", err)
	}
}
