package main

import "fmt"

import "math/rand"

import "time"

func f1(ch chan<- int) {
	for {
		ch <- rand.Intn(10)
		time.Sleep(time.Second * 5) // 每隔5秒钟赋值一次
	}
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 100
	// <-ch1
	x, ok := <-ch1 // 没有值的时候，阻塞 并报错
	fmt.Println(x, ok)

	// for 循环取值
	ch2 := make(chan int, 10)
	go f1(ch2)
	// for {
	// 	x, ok := <-ch2 // 什么时候ok等于false channel 被关闭的时候
	// 	fmt.Println(x, ok)
	// 	time.Sleep(time.Second) // 每隔一秒钟取一次
	// }
	// 这种方式更好
	// for x := range ch2 { // 如果没有值就一直等
	// 	fmt.Println(x)
	// 	time.Sleep(time.Second) // 每隔一秒钟取一次
	// }

}
