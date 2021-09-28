package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

// 向 kafka 中写日志的操作

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer //声明kafka的生产者 client
	logDataChan chan *logData
)

func Init(add []string, MaxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//连接kafka
	client, err = sarama.NewSyncProducer(add, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return err
	}
	// 初始化全局的 chan
	logDataChan = make(chan *logData, MaxSize)

	// 开启后台的 goroutine 从通道中获取数据发往kafka
	go sendToKafka()
	return nil
}

// 给外部暴露一个函数，该函数只把日志数据发送给一个内部的channel中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

// 真正发送到Kafka的函数
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
			}
			fmt.Printf("pid: %v, offset: %v", pid, offset)
		default:
			time.Sleep(time.Microsecond * 500)
		}
	}
}
