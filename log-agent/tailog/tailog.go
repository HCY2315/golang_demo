package tailog

import (
	"context"
	"fmt"
	"log-agent/kafka"

	"github.com/hpcloud/tail"
)

// TailTask: 一个日志收集任务的实例
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail

	// 实现退出t.run
	ctx        context.Context
	cannleFunc context.CancelFunc
}

var tailobj *TailTask

func NewTailTask(path, topic string) (tailobj *TailTask) {
	ctx, cannel := context.WithCancel(context.Background())
	tailobj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cannleFunc: cannel,
	}
	tailobj.Init() //根据路径打开对应的日志
	return tailobj
}

func (t *TailTask) Init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	go t.run()
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tailtask:%s, 结束了\n", t.topic)
			return
		case line := <-t.instance.Lines: //从tailobj的通道中一行一行的读取日志
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
