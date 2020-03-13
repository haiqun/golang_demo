package tailflog

import "golang_demo/logagent_hq/etcd"

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogConf
}

func Init(LogConf []*etcd.LogConf)  {
	tskMgr = &tailLogMgr{
		 logEntry:LogConf,
	}
	for _,conf := range LogConf {
		NewTailTask(conf.Path,conf.Topic)
	}
}

