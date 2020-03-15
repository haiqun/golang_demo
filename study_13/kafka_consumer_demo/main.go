package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.3.8:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	//  因为一个topic的存储可能是分别存在不同的分区的，所以需要找到对应分区的获取内容

	partitionList, err := consumer.Partitions("m.hiii") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition : err %v\n", err)
		return
	}
	fmt.Println("分区列表",partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("m.hiii", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				//fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
				fmt.Println(string(msg.Value))
			}
		}(pc)
	}
	// 因为上面是异步去打印信息的，所以这里要等待，不能直接关了main的主goroutine
	select {

	}
}