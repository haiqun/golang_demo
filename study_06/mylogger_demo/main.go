package main

import (
	"fmt"
	"time"

	"github.com/hq_stady/study_06/mylogger"
)

// 声明一个全局的接口变量  ， 包内都可以调用
var log mylogger.Logger

func consoloTest() {
	log = mylogger.NewconsoloLogger("debug")
	id := 111
	name := "lufei"
	log.Debug("debug 我的id是：%d ,我的名字是： %s ", id, name)
	log.Warning("WARNING")
	log.Error("teERRORst")
	log.Fatal("FATAL  我的id是：%d ,我的名字是： %s ", id, name)
}

func fileText() {
	log = mylogger.NewFileLogger("debug", "./", "text.log", 10*1024)
	// n1 := time.Now().UnixNano()
	for {
		fmt.Println(time.Now().UnixNano())
		id := 111
		name := "lufei"
		log.Debug("debug 我的id是：%d ,我的名字是： %s ", id, name)
		log.Warning("WARNING")
		log.Error("teERRORst")
		log.Fatal("FATAL  我的id是：%d ,我的名字是： %s ", id, name)

		time.Sleep(time.Second * 3)
	}
}

func main() {

	// 控制台日志输出
	// consoloTest()

	// 不要直接执行 用 go run 跑-》 有死循环
	/*
		文件切割的时候，会关闭文件目录，
		这时候并发多线程去写日志，
		可能引起文件信息获取的错误，
		这时候写日志就会失败，
		获取文件详细信息报错 ：stat text.log: use of closed file1583553418834098000
		获取文件详细信息报错 ：stat text.log.err: use of closed file1583553421836413000
	*/
	// 记录文件，日志输出
	fileText()
}
