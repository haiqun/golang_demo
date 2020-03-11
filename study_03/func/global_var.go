package main

import "fmt"

// 作用域
var x = 100

func f1() {
	// 函数中查找变量的顺序
	// 1. 先在函数内部查找
	// 2. 函数内部没有找到，一直找到全局变量，全局变量没有找到报错
	// 3. 函数内部定义的变量不能被外部调用
	fmt.Println(x)
	y := "hhh"
	fmt.Println(y)
}

func main() {
	f1()
	// 语句块的作用域
	// if i := 1; i > 10 {
	// 	println("i大于10,i的值是:", i)
	// }
	// fmt.Println(i)

}
