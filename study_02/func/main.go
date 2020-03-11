package main

import "fmt"

// 函数
// 函数的定义
// 命名返回值，就相当于在函数中声明了一个变量，如果已经命名过的话，return后面可以不写名称
func sum(x int, y int) (ret int) {
	ret = x + y
	return
}

// 没有参数没有返回值
func f11() {
	fmt.Println("没有参数没有返回值")
}

// 没有参数
func f2() string {
	return "没有参数"
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 多个返回值
func f5() (int, string) {
	return 1, "你好呀"
}

// 参数简写：当参数中连续两个参数类型一致时，只写一个类型就可以
func f6(x, y int, z, m, q string) string {
	return "参数简写"
}

// 可变长参数
func f7(q int,x ...int) {
	num := 0
	for _, y := range x {
		num += y
	}
	fmt.Println(num)
	fmt.Println("f7 q :",q)
}

// go语言没有默认参数

func main() {
	a1 := 12
	a2 := 12
	a3 := sum(a1, a2)
	fmt.Println(a3)
	_, s1 := f5()
	fmt.Println(s1)
	f7(100,2,2)
}
