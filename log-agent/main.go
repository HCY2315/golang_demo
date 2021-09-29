package main

import (
	"fmt"
	"log-agent/conf"
	"log-agent/etcd"
	"log-agent/kafka"
	"log-agent/tailog"
	"log-agent/utils"
	"sync"
	"time"
)

func init() {
	conf.Init()
}

func main() {
	//0、初始化配置文件
	conf.Init()

	//1、初始化 kafka 连接
	err := kafka.Init([]string{conf.KafkaSetting.Address}, conf.KafkaSetting.Chan_max)
	if err != nil {
		fmt.Println("kafka 初始化失败")
		return
	}
	fmt.Println("kafka init success")

	//2、初始化 etcd
	err = etcd.Init(conf.EtcdSetting.Address, time.Duration(conf.EtcdSetting.Timeout)*time.Second)
	if err != nil {
		fmt.Println("etcd init err:", err)
		return
	}
	fmt.Println("etcd init success")

	// 为了实现每个logagent拉取自己独有的日志配置，所以要以自己的ip地址作为区分
	ipStr, err := utils.GetOutboundIp()
	if err != nil {
		return
	}
	ercdConfKey := fmt.Sprintf(conf.EtcdSetting.Key_log, ipStr)

	// 2.1 从etcd中获取日志采集项的配置信息
	LogEntryConf, err := etcd.GetConf(ercdConfKey)
	if err != nil {
		fmt.Println("form etcd get config err:", err)
	}
	for index, value := range LogEntryConf {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// 2.2 派一个哨兵去监视日志收集项的变化（有变化及时通知我们的logagent实现热加载）
	tailog.Initl(LogEntryConf)
	newconfchan := tailog.NewConfChan() //从 taillog包中回去对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(ercdConfKey, newconfchan) //发现最新的配置信息通知会通知上面的通道
	wg.Wait()

	// 3、收集日志发送kafka
	// 3.1、循环每一个日志收集项，叉棍见tailobj
	// for _, conf := range LogEntryConf {
	// 	tailog.NewTailTask(conf.Path, conf.Topic)
	// }

	// //3、打开日志文件准备手机日志
	// err = tailog.Init(conf.TailSetting.Path)
	// if err != nil {
	// 	fmt.Println("init taillog failed, err:", err)
	// 	return
	// }
	// run()
}

// func run() {
// 	//读取日志，发送到Kafka
// 	for {
// 		select {
// 		case line := <-tailog.ReadChan():
// 			kafka.SendToKafka(conf.KafkaSetting.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second * 5)
// 		}
// 	}
// }

// func conf() {
// 	//0、添加配置项
// 	cfg, err := ini.Load("./conf/config.ini")
// 	if err != nil {
// 		fmt.Println("conf init failed, err:", err)
// 		return
// 	}

// 	//kafka 连接参数
// 	Kafka := cfg.Section("kafka")
// 	Kafkddress = Kafka.Key("address").String()
// 	KakfaTopic = Kafka.Key("topic").String()

// 	// 读取的日志位置
// 	Logpath = cfg.Section("taillog").Key("path").String()
// }
