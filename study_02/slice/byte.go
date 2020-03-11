package main

import "fmt"

//byte
func main() {
	// 切片类型的 byte ，实例化后都是 []uint8 ,
	// 实例化的两种方式如下
	var b2 = []byte{1, 2, 3, 4}
	fmt.Println(b2)
	fmt.Printf("%T , %#v\n", b2, b2)
	b1 := []byte("test")
	fmt.Printf("%T , %#v, %#v\n", b1, b1, string(b1))
}
