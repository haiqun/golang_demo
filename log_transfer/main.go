package main

import (
	"fmt"
	"golang_demo/log_transfer/conf"
	"golang_demo/log_transfer/kafka"
	"golang_demo/log_transfer/es"
	"gopkg.in/ini.v1"
)

func main()  {
	// 读取配置项
	cfg := new(conf.AppConf)
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err!=nil {
		panic("加载配置文件失败")
	}
	// 初始化es
	err = es.Init(cfg.EsConfig.Address)
	if err!=nil {
	 fmt.Println(" kafka init failed:",err)
	}

	// 初始化kafka - 并监控-获取数据，推送到 es
	err = kafka.Init([]string{cfg.KafkaConfig.Address},cfg.KafkaConfig.Topic)
	if err!=nil {
		fmt.Println(" kafka init failed:",err)
	}
	// 将kafka的数据推送到es
	select {
		
	}
}

