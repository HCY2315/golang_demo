package main

import (
	"fmt"
	"go-logging/pkg/logging"
	"go-logging/setting"
)

//初始化配置
func init() {
	setting.Setup()
	logging.Setup()
}

func main() {
	fmt.Println(setting.LoggerSetting.TimeFormat)
}
