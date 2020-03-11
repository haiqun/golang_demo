package main

import "fmt"

// 给自定义的类型加方法
// 不能给别的包的类型添加方法，只能给自己的包添加方法，所以这里不能用 int 才定义个 自定义类型
type myint int

func (m myint) hello() {
	fmt.Println("hello")
}

func main() {
	m := myint(100)
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	m.hello()
}
