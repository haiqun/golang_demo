package main

import "fmt"

// 流程控制
func main() {
	age := 10
	if age > 18 {
		fmt.Println("大于18岁，成年了")
	} else {
		fmt.Println("快去写作业")
	}

	if age > 35 {
		fmt.Println("认到中年")
	} else if age > 65 {
		fmt.Println("步入老年")
	} else {
		fmt.Println("努力学习")
	}

	// 作用域
	// if 特殊写法
	// age 只在if条件判断语句生效
	if age := 20; age > 19 {
		fmt.Println("大于18岁，成年了")
	} else {
		fmt.Println("努力学习")
	}
}
