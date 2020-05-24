package main

import (
	"golang_demo/rabbitmq"
	"time"
)

func main() {
	a := rabbitmq.NewRabbitMQSimple("test1")
	go func() {
		// 消费信息
		a.RecieveSimple()
	}()
	t := time.NewTicker(time.Millisecond * 10)
	for {
		select {
		// 生产信息
		 	case <-t.C:
				a.PublishSimple("hello world"+ " " +time.Now().Format("2006-01-02 15:04:05.000"))
		}
	}
}
