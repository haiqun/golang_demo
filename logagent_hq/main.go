package main

import (
	"fmt"
	"golang_demo/logagent_hq/conf"
	"golang_demo/logagent_hq/etcd"
	"golang_demo/logagent_hq/kafka"
	"golang_demo/logagent_hq/tailflog"
	"gopkg.in/ini.v1"
	"time"
)

// 声明一个全局变量去加载配置文件
var (
	cfg = new(conf.AppConf)
)

func run() {
	// 3 读取日志
	println("ttt")
	time.Sleep(time.Second*30)
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
	// fmt.Println(cfg.KafkaConfig.Address)
	// 1 初始化kafka链接
	err = kafka.Init([]string{cfg.KafkaConfig.Address})
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

	// 从etcd 中获取配置项 -》 要读哪里的路径，写到哪个topic
	LogConf,err := etcd.GetConf(cfg.EtcdConfig.Key)
	if err != nil {
		fmt.Printf("etcd GetConf failed,err:%v\n", err)
	}
	//fmt.Println(LogConf)
	for _, i2 := range LogConf {
		fmt.Printf("value :%v\n",*i2)
	}
	// 监控配置项的改变，做热启动
	tailflog.Init(LogConf)
	// 2 打开日志文件，准备手机日志
	//err = tailflog.Init(cfg.TaillogConfig.Filename)
	//if err != nil {
	//	fmt.Printf("init taillog failed,err:%v\n", err)
	//	return
	//}
	run()
}
