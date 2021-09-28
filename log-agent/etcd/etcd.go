package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化ETCD的函数
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, err:", err)
		return err
	}
	return nil
}

// 从etcd中根据key获取配置项
func GetConf(key string) (LogEntryConf []*LogEntry, err error) {
	ctx, cannal := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cannal()
	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		// fmt.Println(ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &LogEntryConf)
		if err != nil {
			fmt.Println(" unmarshal etcd value failed, err:", err)
			return
		}
	}
	return
}

// etcd watch
func WatchConf(key string, newConfch chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	// 从通道中尝试取值(监视的信息)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key: %v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			// 通知 taillog.tskmgr
			// 1、先判断操作的类型
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				// 如果是删除操作，手动传递一个空的配置项
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("Unmarshal err:", err)
					continue
				}
			}
			fmt.Printf("get new conf: %v\n", newConf)
			newConfch <- newConf
		}
	}
}
