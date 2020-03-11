package main

import "fmt"

// 结构体
// 构造函数 ， 返回一个结构体变量的函数

type person struct {
	name string
	age  int
	city string
}

// 返回的是结构体本身还是结构指针
// 构造函数 约定成俗用new开头
// *** 重点 ：结构体比较大的时候 ，建议使用指针 ，减少程序的内部开销

func newPerson(name string, age int, city string) *person {
	return &person{
		name: name,
		age:  age,
		city: city,
	}
}

func main() {
	p1 := newPerson("lufei", 19, "广州")
	p2 := newPerson("suoluo", 20, "南京")
	fmt.Println(p1.name, p2)
}
