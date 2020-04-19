package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 5; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 5; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	// 只放了一个cup去执行，变成串行
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}