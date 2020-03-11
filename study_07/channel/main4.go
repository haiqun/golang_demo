package main

import (
	"fmt"
	"sync"
)

// 单项通道

var wg sync.WaitGroup

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for x := range ch1 {
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	wg.Add(2)
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	f1(ch1)
	f2(ch1, ch2)
	wg.Wait()
	for {
		x, ok := <-ch2
		if !ok {
			fmt.Println("ch2的值以及取完了")
			break
		}
		fmt.Println(x)
	}

	ch3 := make(chan int, 2)
	ch3 <- 10
	ch3 <- 10
	// ch3 <- 10 

}
