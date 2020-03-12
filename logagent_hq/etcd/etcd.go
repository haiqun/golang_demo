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
	Path string `json:"path"`
	Topic string `json:"topic"`
}

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