package main

import "fmt"

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 10
}

func f3(x []rune) {
	for _, v := range x {
		fmt.Println(v)
	}
}

func f4(x func() int) int {
	ret := x()
	return ret
}

func ff(x, y int) int {
	return x + y
}

// 函数的参数是个函数，函数的返回值也是一个函数
func f5(x func() int) func(x, y int) int {
	return ff
}

// 函数类型
func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	c := f4(f2)
	fmt.Println(c)

	d := f5(f2)

	fmt.Printf("%T\n", d) // func(int, int) int

}
