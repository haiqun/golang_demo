package main

import (
	"fmt"
	"sync"
)

// channel 循环
var ch1 chan int
var ch2 chan int
var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	ch1 = make(chan int, 100) // 这个值可以设置得小一点 ，比如手 50 因为有人一直在消费
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2() {
	defer wg.Done()
	ch2 = make(chan int, 100)
	for {
		i, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- i * i
	}
	close(ch2)
}

func f3() {
	defer wg.Done()
	ch2 = make(chan int, 100)
	for x := range ch1 { // 通道没有值的时候，队列会一直等待
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	wg.Add(2)
	f1()
	// f2()
	f3()
	wg.Wait()
	for {
		i, ok := <-ch2
		if !ok { // 没有值的时候，获取到的是默认值
			break
		}
		fmt.Println(i)
	}
}
