package main

import "time"
import "fmt"

// goruntine 匿名函数
func main() {
	for i := 0; i < 100; i++ {
		go func(i int) { // 不传值的闭包会往上获取数据，也可以拿到外层的i
			fmt.Println(i) // 如果不传值，也可以执行，但是会出现重复的i ，
		}(i)
	}
	fmt.Println("mian")
	time.Sleep(time.Second)
}
