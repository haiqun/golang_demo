package main

import "fmt"

// 高阶函数复习
func f1(name string) {
	fmt.Println("你好~", name)
}

func f2(x func(string), name string) {
	x(name)
}

func f3() func(x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}

// 闭包
func f4(f func()) {
	fmt.Println("这里是f4")
	f()
}

// 把 f1 套到 f4

func f5(f func(string), m string) func() {
	return func() {
		f(m)
	}
}

func main() {
	f2(f1, "路飞")
	ret := f3()
	n := ret(156, 398)
	fmt.Println(n)

	f44 := f5(f1, "布鲁克")
	fmt.Printf("%T\n", f44)
	f4(f44)

}
