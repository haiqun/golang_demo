package tailflog

import (
	"fmt"
	"golang_demo/logagent_hq/etcd"
	"time"
)

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogConf
	newChan chan []*etcd.LogConf // 同于热更新的配置通道
	tskMap map[string]*TailTask
}


func Init(LogConf []*etcd.LogConf)  {
	tskMgr = &tailLogMgr{
		 logEntry:LogConf,
		 newChan:make(chan []*etcd.LogConf), //无缓存通道
		 tskMap:make(map[string]*TailTask),
	}
	for _,conf := range LogConf {
		NewTailTask(conf.Path,conf.Topic)
	}
	// 监控是是否有新的配置进来
	go tskMgr.run()
}

func NewConfChan()( c chan<- []*etcd.LogConf)  {
	return tskMgr.newChan
}

// 监控自己的 NewChan ,有新的配置过来更新
func (t *tailLogMgr)run()  {
	for {
		select {
		case nconf := <-t.newChan:
			fmt.Println("有新的配置进来了",nconf)
		default:
			time.Sleep(time.Second)
		}
	}
}