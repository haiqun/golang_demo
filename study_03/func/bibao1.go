package main

import "fmt"

// 闭包
// 闭包是什么？ 
// 闭包是一个函数 ，这个函数包含了他外部作用域的一个变量

// 闭包原理 
// 1. 函数可以作为返回值 
// 2. 函数内部查找变量的顺序，现在自己内部找，找不到往外层找 
func f1(f func()) {
	fmt.Println("this is f1 ")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2 ")
	fmt.Println(x + y)
}

func f3(x func(a, b int), n, m int) func() {
	tmp := func() {
		x(n, m)
	}
	return tmp
}

func main() {
	ret := f3(f2, 4, 5)
	f1(ret)
}
