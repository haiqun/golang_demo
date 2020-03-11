package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tail demo
// 采集日志
func main() {
	fileName := "./my.log" // 读取文件
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开文件 - 文件切割的时候，切割完重新打开
		Follow:    true,                                 // 是否追随 - 文件切割的时候，切割完继续读
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 在哪里割地方开始读
		MustExist: false,                                // 文件不存在，是否报错
		Poll:      true}
	// 打开文件
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line // 一行数据
		ok   bool
	)
	for {
		// 一行行的读
		line, ok = <-tails.Lines
		if !ok {
			// 没有数据就等待
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", line.Text)
	}
}
