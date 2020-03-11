package main

import "fmt"

// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
func main() {
	// var m1 map[string]int // 没有初始化，没有在内存中开辟空间
	// fmt.Println(m1 == nil)
	m1 := make(map[string]int, 10) // 初始化的时候要估算好 map容量，避免执行过程动态扩容
	m1["中国"] = 5000
	m1["埃及"] = 2000
	fmt.Println(m1)
	// 获取不存在的key时
	fmt.Println(m1["美国"]) // 不存在的时候，拿到默认类型的值，这里拿到 0
	// 约定成俗 用ok 接受布尔值
	v, ok := m1["中国"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 删除
	delete(m1, "埃及")
	fmt.Println(m1)
	// 删除不存在的值 =》 删除不存在 ，就不操作
	delete(m1, "美国")
	fmt.Println(m1)

}
