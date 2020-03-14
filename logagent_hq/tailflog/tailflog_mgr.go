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
		tailObj := NewTailTask(conf.Path,conf.Topic)
		// 记录,后续可以判断是否修改-或调整
		tskMgr.tskMap[conf.Path] = tailObj
	}
	// 监控是是否有新的配置进来
	go tskMgr.run()
}

// 这其实就是将自己的变量暴露出去，给其他人赋值
func NewConfChan()( c chan<- []*etcd.LogConf)  {
	return tskMgr.newChan
}

// 监控自己的 NewChan ,有新的配置过来更新
func (t *tailLogMgr)run()  {
	for {
		select {
		case nconf := <-t.newChan:
			fmt.Println("新的配置:",nconf)
			fmt.Println("原来的配置:",t.logEntry)
			for _,v := range nconf {
				// 配置可能新增-修改-删除 ？ 如何判断配置属于新增还是修改，所以需要在初始化的话时候，记录通道，这时候作比较
				_,ok := t.tskMap[v.Path]
				if ok {
					// 路径跟topic没有变更
					continue
				}else{
					//新增
					tailObj := NewTailTask(v.Path,v.Topic)
					t.tskMap[v.Path] = tailObj
				}
				//fmt.Println(*v)
			}
			// 处理删除
			for _,v := range t.logEntry{
				is_delete := true
				for _,item := range nconf  {
					if v.Path == item.Path && v.Topic == item.Topic {
						is_delete = false
						continue
					}
				}
				// 如果已经不存在了，那么就删除掉这个配置
				if is_delete {
					t.tskMap[v.Path].cancelFunc()
					delete(t.tskMap,v.Path)
				}
			}

		default:
			time.Sleep(time.Second)
		}
	}
}