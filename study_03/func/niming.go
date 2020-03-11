package main

import "fmt"

// 匿名函数

func main() {
	// 函数内部不能再声明带名字的函数
	// 匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(1, 3)
	// 如果只调用一次 ，可以写成简写立即执行函数-就是函数的最后再写一个（）
	// 立即执行函数 执行完就没有了
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

}
