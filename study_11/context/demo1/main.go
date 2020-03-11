package main

import (
	"fmt"
	"sync"
	"time"
)

// 开关的方式解决

var notify bool
var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	for {
		fmt.Println("test")
		// 休眠500毫秒
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f1()
	time.Sleep(time.Second * 5)
	notify = true
	// 退出
	wg.Wait()
}
