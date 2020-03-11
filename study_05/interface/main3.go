package main

import "fmt"
// 空接口类型  - 空接口作为函数的参数  

type isNull interface{

}

type cat struct {
	name string
}

func (c cat)say()  {
	fmt.Printf("%s : 我是喵 \n",c.name)
}

func show(x interface{})  {
	fmt.Println(x)
}

func main() {
	var x1 isNull
	c1 := cat{
		name : "喵",
	}
	x1 = c1 
	fmt.Printf("%T , %s \n",x1,x1)
	// fmt.Println(x1.name) // 报错 x1.name undefined (type isNull is interface with no methods)

	str := "我是字符串"
	x1 = str
	fmt.Printf("%T , %s \n",x1,x1) // 可以赋值
	fmt.Println(x1)

	// x1.say()
	show(str)
	i := 199
	show(i)
	t := false
	show(t)
	
	

}