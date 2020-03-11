package main

import "fmt"

import "sync"

var a []int
var b chan int // 需要指定通道中的类型 ， 需要开辟空间

var wg sync.WaitGroup

func noBufChannel() {
	b1 := make(chan int) // 没有缓冲区
	// b1 <- 10 // 报错，没有人接受 all goroutines are asleep - deadlock! 死锁
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 接受一个从水管中获取的值
		x := <-b1
		fmt.Println("从b1水管中获取一个值", x)
	}()
	// 没有使用缓冲区的话，如果加上goruntine就可以满足这样的条件
	b1 <- 10
	fmt.Println("这里是 将10丢到水管b1中...")
	close(b)
	wg.Wait()
}

func bufChannel() {
	// 发送值
	b = make(chan int, 10) // 带缓存区的通道初始化 10为指定可以放的值，这个值可以设置很大
	// 如果设置很大的话，建议使用指针
	fmt.Printf("%T --- %v\n", b, b)
	b <- 10 // 因为初始化有设置缓冲区，那么直接放值不会报错
	b <- 20
	fmt.Printf("%T --- %v\n", b, b)
	x := <-b
	fmt.Printf("%T --- %v\n", x, x)
	close(b)
}

// channel
func main() {
	// noBufChannel()
	// bufChannel()
}
