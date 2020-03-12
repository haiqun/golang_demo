package main

// 开启 tdcd 直接在安装目录下 执行 etcd 就启动了
// 默认端口 ：127.0.0.1:2380

/*
	设置一个值/修改  ./etcdctl --endpoints=http://127.0.0.1:2379 put test "1000"
	获取一个值  ./etcdctl --endpoints=http://127.0.0.1:2379 get  test
	删除一个值 ./etcdctl --endpoints=http://127.0.0.1:2379 del test
*/

// 搭建一个集群的案例：https://www.liwenzhou.com/posts/Go/go_etcd/

// etcd client put/get demo
// use etcd/clientv3

import (
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // 设置连接etcd的节点的地址跟端口
		DialTimeout: 5 * time.Second, // 设置连接超时时间
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close() // 关闭连接
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "q1mi", "dsb") // 1秒钟没返回就断开
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "q1mi") // 第三个参数可以设置前缀 clientv3.WithPerfix()
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

}