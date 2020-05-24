package main

import (
	"golang_demo/rabbitmq"
	"time"
)

// topic模式-生产者 main8 main9 main10 为一组测试
/*
 	队列名称随机产生
	队列消费只会接受最新接入的，没有接入前的数据不会被接收到

    重要  ******  生产者-key固定
	重要  ******  消费者匹配
*/
func main()  {
	q := rabbitmq.NewRabbitMQTopic("topicPro1","topic.kkk.abc")
	m := rabbitmq.NewRabbitMQTopic("topicPro1","topic_ko.ik")
	t := time.NewTicker(time.Second)
	t1 := time.NewTicker(time.Second)
	for  {
		select {
		case <-t.C :
			q.PublishTopic("topic模式key1生产消息"+time.Now().Format("2006-01-02 15:04:05"))
		case <-t1.C:
			m.PublishTopic("topic模式key2生产消息"+time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}