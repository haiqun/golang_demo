package main

import (
"golang_demo/rabbitmq"
"time"
)

//  work 模式 - 生产者
func main() {
	a := rabbitmq.NewRabbitMQSimple("test1")
	t := time.NewTicker(time.Second * 1)
	for {
		select {
		// 生产信息
		case <-t.C:
			a.PublishSimple("hello world"+ " " +time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

