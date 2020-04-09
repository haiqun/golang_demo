package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func main() {
	file ,err := os.OpenFile("xx1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil{
		log.Out = file
	}else{
		log.Info("Failed to log to file")
	}
	// 格式化日志
	// log.Fatalf("Failed to send event %s to topic %s with key %d", event, topic, key)

	// 鼓励使用Fields结构化日志内容
	event := "event 1111"
	topic := "topic qqq"
	key := "key abc"
	log.WithFields(logrus.Fields{
		"event": event,
		"topic": topic,
		"key": key,
	}).Print("Failed to send event")

	/*
	2)固定Fields
	可以固定Fields不用每次都写
	*/
	entry := log.WithFields(logrus.Fields{
		"name": "test",
	})
	entry.Info("message1")	//输出 time="2020-03-31T17:47:31+08:00" level=info msg=message2 name=test
	entry.Info("message2")


}

