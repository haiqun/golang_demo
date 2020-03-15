package kafka

import (
	//"encoding/json"
	"github.com/Shopify/sarama"
	"golang_demo/log_transfer/es"
)
import "fmt"

var consumer sarama.Consumer

// 照一个json的结构体
type logData struct {
	Data string `json:"data"`
}

func Init(address []string,topic string) (err error) {
	consumer, err = sarama.NewConsumer(address, nil)
	if err != nil {
		return
	}
	fmt.Printf("kafka connect %s success!\n",address)
	go runConsumerInfo(topic)
	return
}

func runConsumerInfo(topic string) {
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition : err %v\n", err)
		return
	}
	fmt.Printf("当前topic=>%s ,所有分区列表:%v\n",topic,partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				//fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
				//fmt.Println(string(msg.Value))
				// 发给es
				//var logd = new(logData)
				//logd := logData{Data:string(msg.Value)}
				logd := map[string]interface{}{
					"data": string(msg.Value),
				}
				// 改为丢到chan中，防止函数调函数
				info := es.MsgData{
					Index:topic,
					Mgs:logd,
				}
				es.SendToEsChan(&info)
				//err = es.SendMsg(topic,logd)
				//if err!=nil {
				//	fmt.Println("es SendMsg err:",err)
				//	continue
				//}
			}
		}(pc)
	}
	// 这里不能退出
	select {

	}
}