package tailog

import (
	"fmt"
	"log-agent/etcd"
	"time"
)

// tailTask 管理者
type tailLogMgr struct {
	logEntry  []*etcd.LogEntry
	tskMap    map[string]*TailTask
	newConfCh chan []*etcd.LogEntry
}

var tskMgr *tailLogMgr

func Initl(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:  logEntryConf,
		tskMap:    make(map[string]*TailTask, 16),
		newConfCh: make(chan []*etcd.LogEntry), //无缓冲的通道
	}
	for _, conf := range logEntryConf {
		// NewTailTask 实现多线程运行每个任务
		tailobj := NewTailTask(conf.Path, conf.Topic)
		mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
		tskMgr.tskMap[mk] = tailobj
	}
	go tskMgr.run()
}

// 监听自己的newConfChan，有了新的配置做对应的操作
// 1、配置新增
// 2、配置删除
// 3、配置变更
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfCh:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 新增
					tailobj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailobj
				}
			}
			// 找出原来 t.taskmap有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry {
				isDelete := true
				for _, c2 := range newConf {
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[mk].cannleFunc()
					delete(t.tskMap, mk)
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数,向外暴露tskMgr 的 newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfCh
}
