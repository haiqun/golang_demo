package main

import (
	"golang_demo/rabbitmq"
)

// 路由模式-消费者
/*
 	队列名称随机产生
	队列消费只会接受最新接入的，没有接入前的数据不会被接收到
*/
func main()  {
	q := rabbitmq.NewRabbitMQTopic("topicPro1","topic_ko.#")
	q.RecieveTopic()
}