package main

import "fmt"

func f1(a ...interface{}) {
	fmt.Printf("value : %#v ,type : %T\n", a, a)
}

func main() {
	// s1 := []int{1, 2, 3, 4}
	// fmt.Printf("value : %#v ,type : %T\n", s1, s1)
	s2 := []interface{}{1, 2, 3, 4, 5} // 这时候 1，2，4 是接口类型的切片的值
	fmt.Printf("value : %#v ,type : %T\n", s2, s2)
	// for x := range s2 {
	// 	fmt.Println(x)
	// }
	// 将 {[]interface {}{1, 2, 3, 4, 5}} 当前一个参数 【切片的第一个元素也是切片】
	f1(s2)
	// ...是将 1，2，4 当成接口类型的切片的值，分别传进去
	f1(s2...)

}
