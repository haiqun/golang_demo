package main

import (
	"golang_demo/rabbitmq"
	"time"
)

// 订阅模式生产
func main()  {
	q := rabbitmq.NewRabbitMQPubSub("newProject")
	t := time.NewTicker(time.Second)
	for  {
		select {
			case <-t.C :
				q.PublishPub("订阅模式生产消息"+time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}