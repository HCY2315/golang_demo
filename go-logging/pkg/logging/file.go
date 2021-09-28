package logging

import (
	"fmt"
	"go-logging/setting"
	"time"
)

//获取日志存放的位置
func GetlogFilePath() string {
	return fmt.Sprintf("%s%s", setting.LoggerSetting.LogSavePath)
}

//根据时间进行日志切割
func GetlogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.LoggerSetting.LogSaveName,
		time.Now().Format(setting.LoggerSetting.TimeFormat),
		setting.LoggerSetting.LogFileExt,
	)
}

//获取所有的日志信息
func GetlogSumFileName() string {
	return fmt.Sprintf("%ssum.%s",
		setting.LoggerSetting.LogSavePath,
		setting.LoggerSetting.LogFileExt,
	)
}
