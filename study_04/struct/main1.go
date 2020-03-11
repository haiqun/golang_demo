package main

import "fmt"

//结构体 是 值类型

type person struct {
	name, gender string
}

// go语言中的函数永远是拷贝
func f1(s person) {
	s.gender = "女"
}

func f2(s *person) {
	// 接受指针，查找变量
	// (*s).gender = "女" // 根据内存地址找到原来的变量，修改原理的变量的值
	s.gender = "女" // 语法糖 ，自动根据指针找对应的变量
}

func main() {
	var n1 person
	n1.name = "lufei"
	n1.gender = "男"
	fmt.Println(n1)
	f1(n1)
	fmt.Println(n1)
	fmt.Printf("%p\n", &n1) // 指针 0xc000086000
	f2(&n1)
	fmt.Println(n1)
	// 结构体指针 1
	var p2 = new(person)
	fmt.Printf("%T\n", p2)  // *标识指针类型 ”*main.person“
	fmt.Printf("%p\n", p2)  // 这个p2的值就是一个内存地址 ”0xc00000c0e0“
	p2.name = "fei lu "     // 这有有糖变量 ，直接读取内存地址，然后赋值 等价于 *p2.name = "jjj"
	fmt.Printf("%p\n", &p2) // 这个就是p2的内存地址 0xc00000e020
	// 结构体指针 2
	// 2.1 key-value 初始化 =》 推荐用这种
	var p3 = &person{
		name:   "布鲁克",
		gender: "男",
	}
	fmt.Println(p3.name)
	fmt.Printf("%#v\n", p3)
	// 2.2 使用列表的形式初始化 , 这种方式的话，一定要按顺序 填写
	p4 := person{
		"索隆",
		"男",
	}
	// fmt.Println(p4)
	fmt.Printf("%#v\n", p4)

	// 结构体 占用连续的一块内存空间
	// fmt.Println(&p3.name)
	// fmt.Println(&p3.gender)

	fmt.Printf("%p\n", &p3.name)   // 这个就是p2的内存地址 0xc00000e020
	fmt.Printf("%p\n", &p3.gender) // 这个就是p2的内存地址 0xc00000e020

}
