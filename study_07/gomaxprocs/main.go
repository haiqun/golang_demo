package main

import (
	"fmt"
	"runtime"
	"sync"
)

//GOMAXPROCS

var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("F1:%d\n", i)
	}
}

func f2() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("F2:%d\n", i)
	}
}

func main() {
	// 设置执行者个数 如果只有一个那就正常顺序
	runtime.GOMAXPROCS(4)
	// 查看cup核心数 我的辣鸡电脑只有 4个
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()
}
