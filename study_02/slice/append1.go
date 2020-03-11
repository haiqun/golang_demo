package main

import "fmt"

// append an copy
func main() {
	s1 := []int{1, 2, 3, 4, 1} // 直接声明 切片变量
	fmt.Println(s1)
	// 不能这样追加 s1[4] = 5 编译时会报错
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1, 999) // 必须用原理的切片变量接受返回值 
	fmt.Println(s1, len(s1), cap(s1)) // 少于1024 直接 容量翻倍
	s1 = append(s1, 998, 8938)        // 当原来切片的底层数组放不下的时候，会把底层数组"换"一个更大的
	fmt.Println(s1, len(s1), cap(s1)) //
	// copy 复制内存，修改原理的变量，不会影响到copy出来的数据
	s2 := []int{3, 4, 5}
	s1 = append(s1, s2...) // ... 表示拆开，分别追加
	fmt.Println(s1, len(s1), cap(s1))

}
