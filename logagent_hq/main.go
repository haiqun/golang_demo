package main

import (
	"fmt"
	"golang_demo/logagent_hq/conf"
	"golang_demo/logagent_hq/etcd"
	"golang_demo/logagent_hq/kafka"
	"golang_demo/logagent_hq/tailflog"
	"golang_demo/logagent_hq/utils"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

// 声明一个全局变量去加载配置文件
var (
	cfg = new(conf.AppConf)
)

func run() {
	// 3 读取日志
	time.Sleep(time.Second*15)
	println("over")
	//for {
	//	select {
	//	case line := <-tailflog.ReadChan():
	//		// 4 推送kafka
	//		kafka.SendMsgToKafka(cfg.KafkaConfig.Topic, line.Text)
	//	default:
	//		time.Sleep(time.Second)
	//	}
	//}
}

// 程序入口
func main() {
	// 0 去读配置文件 - 结构体映射方面  
	err := ini.MapTo(cfg, "./conf/config.ini")
	// 读取单条配置信息 cfg, err := ini.Load("my.ini")
	// 读取单条配置信息  cfg.Section("").Key("app_mode").String()
	if err != nil {
		fmt.Printf("Fail config.ini to read file: %v", err)
		return 
	}
	fmt.Println(cfg.KafkaConfig.Address)
	// 1 初始化kafka链接test
	err = kafka.Init([]string{cfg.KafkaConfig.Address},cfg.KafkaConfig.MaxchanNum)
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")

	// 初始化etcd
	err = etcd.Init(cfg.EtcdConfig.Endpoints,time.Duration(cfg.EtcdConfig.Timeout) * time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")
	// 从etcd -》 获取配置项 -》 要读哪里的路径，写到哪个topic
	ip_str,err  := utils.GetOutboundIP()
	if err != nil {
		panic("GetOutboundIP failed err")
	}
	ek :=  fmt.Sprintf(cfg.EtcdConfig.Key,ip_str)
	fmt.Println(ek)
	LogConf,err := etcd.GetConf(ek)
	if err != nil {
		fmt.Printf("etcd GetConf failed,err:%v\n", err)
	}
	// 打印监控的文件与推送的topic的名称
	//go etcd.WatchConf()
	for _, i2 := range LogConf {
		fmt.Printf("监控的配置，文件路径：%s，推送的topic名 %s :%v\n",i2.Path,i2.Topic)
	}
	// 1.初始化 - 开启接受变更配置的通道
	tailflog.Init(LogConf)
	// 2.获取一个对外暴露的通道=》为上面监控的赋值
	newChanConf := tailflog.NewConfChan()
	// 然后监控改变的变量，然后推送给这个通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(ek,newChanConf)
	wg.Wait()

}
