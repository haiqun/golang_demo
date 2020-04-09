package main

import (
	"log"
	"time"
)

/*
	log 支持三种函数类型的日志输出
	Print	// 普通
	Fatal //致命错误
	Panic //异常
*/

func main() {
	for  {
		log.Println("这是一条很普通的日志") // 正常打印
		v := "lufei"
		log.Printf("%s 真的很牛逼",v)
		//log.Fatal("这是一条会触发fatal的日志") // 直接抛出错误 ，停止执行
		time.Sleep(time.Millisecond * 500)
		log.Panicf("%s panicf 的错误来了",v)
	}
}