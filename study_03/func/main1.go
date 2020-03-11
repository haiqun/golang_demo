package main

import "fmt"

// 函数
func f1() {
	fmt.Println("hello world")
}

func f2(x string) {
	fmt.Println("hello", x)

}

// 一个必填一个选填的 参数类型相同的函数

func f3(x int, y ...int) (ret int) {
	sum := 0
	for _, v := range y {
		sum += v
	}
	fmt.Println(x, y)
	ret = sum
	return ret
}

func main() {
	f1()
	s := f3(100, 2, 2, 3, 4, 6)
	fmt.Println(s)
}
