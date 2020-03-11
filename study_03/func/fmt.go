package main

import "fmt"

//fpm 函数介绍
func main() {
	// fmt.Print("test")
	// fmt.Print("---test")
	// fmt.Println()
	// fmt.Println("test")
	// fmt.Println("---test")
	// a := 1
	// b := [...]int{1}
	// fmt.Printf("%p\n", &a)
	// %T 查看类型
	// %c 字符输出
	// %p 指针
	// %t 布尔值
	// %v 值的默认格式
	// %#v 值的go语法展示
	// 获取用户输出 scan
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容：", s)
	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("姓名%s 年龄%d 班级%s", &name, &age, &class)
	println(name, age, class)

}
