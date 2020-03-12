package main

import (
	"fmt"
	"golang_demo/logagent_hq/kafka"
	"golang_demo/logagent_hq/tailflog"
	"golang_demo/logagent_hq/conf"
	"time"
	"gopkg.in/ini.v1"
)

// 声明一个全局变量去加载配置文件
var (
	cfg = new(conf.AppConf)
)

func run() {
	// 3 读取日志
	for {
		select {
		case line := <-tailflog.ReadChan():
			// 4 推送kafka
			kafka.SendMsgToKafka(cfg.KafkaConfig.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
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
	// 2 打开日志文件，准备手机日志
	err = tailflog.Init(cfg.TaillogConfig.Filename)
	if err != nil {
		fmt.Printf("init taillog failed,err:%v\n", err)
		return
	}
	run()
}
