package main

import "fmt"

// panic rcover

func f1() {
	defer func() {
		fmt.Println("尝试链接")
		err := recover() // 这里捕抓的错误是panice定义的
		fmt.Println("11", err, "11")
	}()
	var os bool
	os = false
	if !os {
		panic("不可原谅的错误，如数据库连接有误")// 后面的程序不会被执行
	}
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func main() {
	f1()
	f2()
}
