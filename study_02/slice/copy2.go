package main

import (
	"fmt"
	"sort"
)

// copy slice
func main() {
	s2 := [...]string{"test", "wo", "中文"} // 声明一个数组
	// s3 := s2 // 开辟了新的内存空间，修改源不影响其他的
	// s3[0] = "他"
	// s2 := s1
	s1 := s2[:]
	var s3 = make([]string, 3, 10) // make ( type len cap) 为什么len等于0 不行呢？
	copy(s3, s1)
	fmt.Println( s1, s3)
	// s1[2] = "日本" // copy 过去的不会变，copy是复制了内存的空间
	// fmt.Println(s1, s2, s3)
	// 删除切片的操作 不要 中间的值
	// s1 := s2[:]
	// fmt.Printf("%p\n", s1) // 显示内存地址
	// fmt.Println(s1, len(s1), cap(s1))
	// s1 = append(s1[:1], s1[2:]...) // 操作了底层的数组，所以原来的值会变，没有改动的位置，值不变
	// fmt.Println(s1, len(s1), cap(s1))
	// s1[0] = "enene"
	// fmt.Println(s2, len(s2), cap(s2))
	// fmt.Printf("%p\n", s1)
	// fmt.Println(s3, len(s3), cap(s3))
	// 练习题
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i)) // 这里是追加 ，所以原来切片的默认值还是存在的
	}
	fmt.Println(a)

	// 对切片进行排序
	a1 := [...]int{2, 3, 5, 20, 1}
	a2 := a1[:]
	sort.Ints(a2)
	fmt.Println(a2)
}
