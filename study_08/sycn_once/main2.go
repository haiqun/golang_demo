package main

import (
	"fmt"
	"sync"
)

// sycn.Once

var ch1 chan int
var ch2 chan int
var wg sync.WaitGroup
var o sync.Once

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 chan int, ch2 chan int) {
	defer wg.Done()
	for {
		v, ok := <-ch1
		if !ok { // 当ch1 close 关闭，返回才是false
			break
		}
		ch2 <- v * v
	}
	o.Do(func() {
		close(ch2)
	})
}

func main() {
	ch1 = make(chan int, 100)
	ch2 = make(chan int, 100)
	wg.Add(3)
	// 一个生产者
	go f1(ch1)
	// 两个消费者
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()
	for x := range ch2 {
		fmt.Println(x)
	}

}
