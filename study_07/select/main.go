package main

import "fmt"

// select
// 哪个执行就调用哪个，如果都可以执行，那就随机

func main() {
	ch1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		// 哪个执行就调用哪个，如果都可以执行，那就随机
		select {
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:

		}

	}
}
