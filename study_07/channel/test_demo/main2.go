package main

import "fmt"

var ch chan int

func f2()  {
	ret := <-ch
	fmt.Print(ret)
}

func main()  {
	// 必须要有人接受才能，使用的无缓冲通道
	ch = make(chan int )
	go f2()
	ch <- 10
	fmt.Println("over")
}