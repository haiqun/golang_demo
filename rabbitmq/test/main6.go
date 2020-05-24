package main

import (
	"golang_demo/rabbitmq"
	"time"
)

// 路由模式-生产者 main5.main6.main7 为一组测试
/*
 	队列名称随机产生
	队列消费只会接受最新接入的，没有接入前的数据不会被接收到
*/
func main()  {
	q := rabbitmq.NewRabbitMQRouting("routerPro","router_one")
	m := rabbitmq.NewRabbitMQRouting("routerPro","router_two")
	t := time.NewTicker(time.Second)
	t1 := time.NewTicker(time.Second)
	for  {
		select {
		case <-t.C :
			q.PublishRouting("路由模式key1生产消息"+time.Now().Format("2006-01-02 15:04:05"))
		case <-t1.C:
			m.PublishRouting("路由模式key2生产消息"+time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}