package main

import "fmt"

// 复习

func main() {
	// 数组是值类型
	// a1 := [3]int{1, 2, 4}
	// f1(a1) // 参数进去之后，在函数内部修改不会影响外部
	// fmt.Println(a1)
	// a2 := a1 // 赋值之后修改还是不一会影响
	// a1[0] = 888
	// fmt.Println(a2)
	// 切片
	// var s1 []int
	// fmt.Println(s1) // 打印为 [] ,等等于nil，没有分配内存
	// s2 := make([]int, 2, 4)
	// fmt.Println(s2) // [0 0] 分配了内存
	// s3 := make([]int, 0, 4)
	// fmt.Println(s3) // []，不等于nil , 分配了内存
	// 切片的本质
	// s1 := []int{1, 3, 4}
	// s2 := s1
	// fmt.Println(s1, s2)
	// s1[0] = 100
	// fmt.Println(s1, s2)
	// // 切片初始化
	// var s3 []int // 没有空间
	// // s3[0] = 100 //这里会报错
	// s3 = append(s3, 1) // 自动初始化切片
	// fmt.Println(s3)

	// 对比数组初始化
	var a1 [3]int // 声明后就有内从空间，也有默认值
	a1[0] = 100
	fmt.Println(a1)

	// 指针
	n1 := &a1[0]
	fmt.Println(n1)
	fmt.Printf("%T\n", n1)
	fmt.Println(*n1)

	// map
	var m1 map[int]string
	m1 = make(map[int]string)
	m1[1] = "hahhah "
	fmt.Println(m1)
	// delete 空的map 还是不等于 nil ，因为它已经申请了内存空间

}

func f1(a [3]int) {
	a[0] = 100
}
