package main

import "fmt"

// fmt的详解
func main() {

	n := 10
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	// 字符串
	s := "hello 路飞"

	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s) // “#”根据类型输出符合描述符号

}
