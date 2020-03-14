package etcd

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
)

var cli *clientv3.Client

func Init(address string,t time.Duration)(err error)  {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{address}, // 设置连接etcd的节点的地址跟端口
		DialTimeout: t, // 设置连接超时时间
	})
	return
}

type LogConf struct{
	Path string `json:"path"` //日志存放的路径
	Topic string `json:"topic"` // 日志要存放的kafka topic
}


// GetConf 获取配置
func GetConf(key string)(LogConf []*LogConf ,err error)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// xxx_conf
	resp, err := cli.Get(ctx, key) // 第三个参数可以设置前缀 clientv3.WithPerfix()
	cancel()
	if err != nil {
		//fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	// d:/tmp/nginx.log web.log
	// f:/xxx/redis.log redis.log
	for _, ev := range resp.Kvs {
		//fmt.Println(string(ev.Value))
		err = json.Unmarshal(ev.Value,&LogConf)
		if err != nil {
			fmt.Printf("json.Unmarshal failed:%v", err)
			return
		}
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	return
}

// 监控配置 c chan<- []*etcd.LogConf
func WatchConf(key string,ch chan<- []*LogConf )  {
	rch := cli.Watch(context.Background(), key) // 返回类型 <-chan WatchResponse
	// 尝试从通道中获取值
	for wresp := range rch {
		for _, ev := range wresp.Events {
			// 如果有改变，通知 tailLogMgr
			// 初始化一个值，如果是删除的情况下，将改变量返回出去
			var newConf []*LogConf
			if  ev.Type == clientv3.EventTypeDelete{
				
			}else{
				// 解压json 设置值
				err := json.Unmarshal(ev.Kv.Value,&newConf)
				if err != nil {
					fmt.Printf("json Unmarshal failed err:%s",err)
					continue
				}
			}
			ch<- newConf
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}