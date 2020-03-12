package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 定义全局client
var (
	client sarama.SyncProducer
)

// Init 实例化kafka
func Init(addr []string) (err error) {
	config := sarama.NewConfig()
	//tailf包使⽤
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 连接kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		return
	}
	return
}

// SendMsgToKafka 发送信息到kafka
func SendMsgToKafka(topic, str string) (err error) {
	// 构造⼀个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(str)
	// defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return
}
