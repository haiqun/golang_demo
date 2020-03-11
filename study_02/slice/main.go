package main

import "fmt"

//切片  在数组的上层做了封装，使其更灵活
/*
	1 切片指向一个底层的数组
	2 切片的长度就是它元素的个数
	3 切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
*/
func main() {
	var s1 []int // 定义了存放int类型元素的切片
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil) // nil 为没有开辟空间

	// 初始化
	s1 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	s2 = []string{"上海", "北京", "广州"}
	s3 := s1[5:] //到最后
	fmt.Printf("s3的值：%v ,len %d, cap %d \n", s3, len(s3), cap(s3))
	s1[6] = 1999
	fmt.Printf("s3的值：%v ,len %d, cap %d \n", s3, len(s3), cap(s3))

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 长度与容量
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 6, 7, 8, 8, 10}
	a2 := a1[0:4] // 1, 3, 5, 6 左闭右开
	fmt.Println(a2)
	a3 := a1[5:] //到最后
	a4 := a1[:4] //开始到底7
	a5 := a1[:]  //全部
	fmt.Println(a3, a4, a5)
	fmt.Println(a3, len(a3), cap(a3))
	// 切片是引用类型，都指向一个底层的数组 ，修改了底层数组的值，切片也会跟着变

	
}
