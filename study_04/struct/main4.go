package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

// 遇到的问题
func main() {
	i := 10
	fmt.Printf("%T\n", i)
	// 声明一个变量，类型等于int32的变量 m ，他的值是10
	// 方法1
	// var m int32
	// m = 10
	// 方法2
	// var m int32 = 10
	// 方法3
	// m := int32(10) // 强制转换类型
	// 方法4
	var m = int32(10)
	fmt.Printf("%T %v\n", m, m)
	// 结构体初始化
	// f1
	var p1 person
	p1.name = "lufei"
	fmt.Println(p1)
	// f2 没有指定初始值的字段的值就是该字段类型的零值
	p2 := person{
		name: "lufei1",
	}
	fmt.Println(p2)
	// f3
	// 必须初始化结构体的所有字段
	// 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 该方式不能和键值初始化方式混用。
	p3 := person{
		"lufei3",
		19,
		"gz",
	}
	fmt.Println(p3)
	p4 := newPerson("lufei4", 19, "广州")
	fmt.Println(p4)

}

// 为什么要构造函数 ？

func newPerson(name string, age int, city string) person {
	return person{
		name: name,
		age:  age,
		city: city,
	}
}
