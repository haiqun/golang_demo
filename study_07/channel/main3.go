package main

import "fmt"

// 关闭通道
func main() {
	var ch1 chan int
	ch1 = make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	// for x := range ch1 {
	// 	fmt.Println(x)
	// }
	<-ch1
	<-ch1
	x, ok := <-ch1
	// 返回值的 ok 为false
	if !ok {
		fmt.Println("没有值可以获取了")
	}
	fmt.Println(x, ok) // 可以直接获取，但是是通道的默认值，这里是 0
	fmt.Println(x, ok)
	fmt.Println(x, ok)
}
