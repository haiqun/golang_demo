package main

import (
	"fmt"
	"time"
)

// chan的类型解决问题
// chan类型为一个空的结构体
var ch = make(chan *struct{}, 1)

func f1() {
	for {
		fmt.Println("TEst")
		time.Sleep(time.Millisecond * 500)
		// 当我在chan能取到值的时候就退出
		select {
		case <-ch:
			break
		default:
		}
	}
}

func main() {
	go f1()
	time.Sleep(time.Second * 5)
	ch <- &struct{}{}
	fmt.Println("结束")
}
