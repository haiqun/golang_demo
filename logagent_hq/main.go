package main

import (
	"fmt"
	"hq_study/logagent/kafka"
	"hq_study/logagent/tailflog"
	"time"
)

func run() {
	// 3 读取日志
	for {
		select {
		case line := <-tailflog.ReadChan():
			// 4 推送kafka
			kafka.SendMsgToKafka("web_my_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}

// 程序入口
func main() {
	// 1 初始化kafka链接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	// 2 打开日志文件，准备手机日志
	err = tailflog.Init("./my.log")
	if err != nil {
		fmt.Printf("init taillog failed,err:%v\n", err)
		return
	}
	run()
}
