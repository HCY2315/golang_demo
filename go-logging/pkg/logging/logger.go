package logging

import (
	"fmt"
	"go-logging/pkg/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	F *os.File

	DefaultPrefix      = ""
	logPrefix          = ""
	logger             *log.Logger
	DefaultCallerDepth = 2
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func Setup() {
	var err error
	filepath := GetlogFilePath()
	filename := GetlogFileName()
	F, err = file.MustOpen(filename, filepath)
	if err != nil {
		log.Fatalf("log setup error %v", err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debuge(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

func setPrefix(level Level) {
	//调用者报告有关上函数调用的文件和行号信息
	//调用goroutine的堆栈。参数skip是堆栈帧数
	//升序，0表示调用者的调用者。（由于历史原因
	//skip的含义在调用者和调用者之间是不同的。）返回值报告
	//程序计数器、文件名和相应文件内的行号
	//打电话。如果无法恢复信息，则布尔值ok为false。
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		//Base返回path的最后一个元素。
		//在提取最后一个元素之前，将删除尾部路径分隔符。
		//如果路径为空，则Base返回“.”。
		//如果路径完全由分隔符组成，则Base返回单个分隔符。
		logPrefix = fmt.Sprintf("[%s][%s:%d],", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprint("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
