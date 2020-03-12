package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// watch demo
/*
	监控操作 ：
	1 . 某个值的修改，删除都会通道通知
	使用场景 ：
	1 . 配置环境，热更新

*/

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// watch key:q1mi change 监控q1mi这个字段的变化
	rch := cli.Watch(context.Background(), "q1mi") // 返回类型 <-chan WatchResponse
	// 尝试从通道中获取值
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}