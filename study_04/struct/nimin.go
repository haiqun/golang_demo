package main

import "fmt"

// 结构体的匿名字段

type person struct {
	string // 没有名称的，直接将类型当名称，所以一个类型只有一个这样的名称，很少用
	int
}

func main() {
	p1 := person{
		"lufei",
		10,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	if p1.age > 10{
		fmt.Println(p1.string)
	}
}
