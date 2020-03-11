package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func f1() {
	// fmt.Println("test")
	// 设置日志输出到指定的文件
	fileObj, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开一个文件报错了 ，err : %s\n", err)
	}
	log.SetOutput(fileObj)
	for {
		log.Println("test")
		time.Sleep(time.Second * 3)
	}
}

// 日志库
func main() {
	// f1()
}
