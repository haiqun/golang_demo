package main

import "fmt"

// 结构体嵌套

type addr struct {
	city    string
	provice string
}

type person struct {
	name string
	age  int
	addr addr
}

type person1 struct {
	name string
	age  int
	addr // 匿名嵌套 等效于 addr : addr
}

type person2 struct {
	name      string
	age       int
	addr      // 匿名嵌套 , 有冲突字段
	workplace //
}

type workplace struct {
	city    string
	provice string
}

type company struct {
	name string
	addr addr
}

func main() {
	p1 := person{
		name: "lufei",
		age:  19,
		addr: addr{
			city:    "广州",
			provice: "广东",
		},
	}
	c1 := company{
		name: "dido",
		addr: addr{
			city:    "上海",
			provice: "上海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.addr.provice)
	fmt.Println(c1)
	fmt.Println(c1.name, c1.addr.city)

	p2 := person1{
		name: "suolong",
		age:  20,
		addr: addr{ // 匿名嵌套体，打两个类型名称
			city:    "广州",
			provice: "广东",
		},
	}
	fmt.Println(p2)
	fmt.Println(p2.name, p2.addr.city)
	fmt.Println(p2.name, p2.city) // 匿名嵌套体，支持直接获取 , 如果匿名字段有冲突，不支持直接获取（有冲突）

	p3 := person2{
		name : "test",
		age : 22,
		addr : addr{
			city : "深圳",
			provice : "广东",
		},
		workplace : workplace{
			city : "北京",
			provice : "河北",
		},
	}
	fmt.Println(p3)
	fmt.Println(p3.addr.city)
	fmt.Println(p3.workplace.city)
}
