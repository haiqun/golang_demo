package main

import "fmt"

import "time"

// goruntine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后也会创建一个主goruntine去执行
func main() {
	for i := 0; i < 100; i++ {
		go hello(i) //开启一个单独的 goruntine 去执行hello函数
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main 函数结束之后，由main函数启动的goruntine也结束了
}
