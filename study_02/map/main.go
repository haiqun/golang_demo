package main

import "fmt"

// map和slice 组合
func main() {
	// 元素类型为map的切片
	// n1 := make([]int, 2, 4) // make一个切片，类型为int ，长度为 2 ，容量为4 
	s1 := make([]map[int]string, 10, 10)
	// 没有对内部的map进行初始化 符合类型的数据，make完需要实例化
	// s1[0][1] = "中国" //报错
	s1[0] = make(map[int]string, 1) // make 第一个为类型，第二个参数是长度 1 ，第三个是容量 1
	s1[0][1] = "中国"
	fmt.Println(s1[0])
	fmt.Println()
	fmt.Println(s1)
	// 值为切片的map
	a1 := make(map[string][]int)
	a1["北京"] = []int{1, 2, 3, 4, 5, 9}
	fmt.Println(a1)
}
