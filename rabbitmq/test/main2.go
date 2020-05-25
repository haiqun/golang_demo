package main

import (
	"golang_demo/rabbitmq"
)

//  work 模式 - 消费者
// work 模式跟simple模式的代码一样，支持多开，多个消费端互相负载
func main() {
	a := rabbitmq.NewRabbitMQSimple("test1")
	// 消费信息
	a.RecieveSimple()
}
