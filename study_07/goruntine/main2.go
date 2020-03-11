package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 优雅的等待多线程 goruntine 结束

// goruntine 什么时候结束 
// goruntine 对应执行的函数结束 ，goruntine 就结束了

// main 函数执行完了，由main函数创建的那些goruntine 都结束了

// waitGroup

func f1() {
	rand.Seed(time.Now().UnixNano()) // 保证每次执行的时候都有点不一样，随机数不重复
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, 0-r2) // 用 0- 让数字变负数
	}
}

func f2(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup // 声明同步

func main() {
	// f1()
	// 可以选择在这里初始化
	// wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1) // 这里加一个任务
		go f2(i)
	}
	wg.Wait() // 等待  add的个数，跟done的个数相等，才能结束
}
