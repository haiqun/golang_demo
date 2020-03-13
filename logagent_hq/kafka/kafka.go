package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// 定义全局client
var (
	client sarama.SyncProducer
	chanMsg chan *msgBody
)

type msgBody struct {
	topic string
	msg string
}


// Init 实例化kafka
func Init(addr []string,maxChanNum int) (err error) {
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
	// 初始化 当前kafak的chan的空间
	chanMsg = make(chan *msgBody,maxChanNum)
	go SendMsgToKafka()
	return
}

func SendMsgToChan(topic,msg string)  {
	chanMsg <- &msgBody{
		topic:topic,
		msg:msg,
	}
}

// SendMsgToKafka 发送信息到kafka
func SendMsgToKafka() () {
	// 发送消息
	for {
		select {
			case data := <-chanMsg :
				// 构造⼀个消息
				msg := &sarama.ProducerMessage{}
				msg.Topic = data.topic
				msg.Value = sarama.StringEncoder(data.msg)
				_, _, err := client.SendMessage(msg)
				if err != nil {
					fmt.Println("send msg failed, err:", err)
					return
				}
		default:
			time.Sleep(time.Millisecond*50)
		}
	}
}
