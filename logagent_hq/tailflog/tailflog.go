package tailflog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"golang_demo/logagent_hq/kafka"
)

// 收集日志
//var (
//	tailObj *tail.Tail
//)

// 日志手机任务
type TailTask struct {
	path string // 日志路径
	topic string // 推送的topce
	instance *tail.Tail // 通过tail 打开的文件话柄
	// 为了让t.run的函数能给退出
	ctx context.Context
	cancelFunc context.CancelFunc
}


// NewTailMsg tecd获取的多条配置信息
func NewTailTask(path,topic string) (tailObj *TailTask) {
	ctx,cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:path,
		topic:topic,
		ctx:ctx,
		cancelFunc:cancel,
	}
	tailObj.initTailObj()
	return
}

// Init 实例化
func (t *TailTask)initTailObj() () {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开文件 - 文件切割的时候，切割完重新打开
		Follow:    true,                                 // 是否追随 - 文件切割的时候，切割完继续读
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 在哪里割地方开始读
		MustExist: false,                                // 文件不存在，是否报错
		Poll:      true}
	// 打开文件
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Printf("tail file failed:%s",err)
	}
	// 开启一个go 去执行日志已收集与上报
	go t.run()
}

func ( t *TailTask)run()  {
	for {
		select {
		// 如果ctx 通知了要关闭收集的话，可以跳出
		case <- t.ctx.Done():
			fmt.Printf("监控文件%s退出了",t.path)
			return
		case line := <-t.instance.Lines:
			// 推送到kafka的chan =》 topce , msg
			fmt.Println(line.Text)
			kafka.SendMsgToChan(t.topic, line.Text)
		}
	}
}