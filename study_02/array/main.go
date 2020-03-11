package main

import "fmt"

// 数组
// 存放元素的容器
// 必须指定存放元素的类型和容量 （长度）
// 数组的长度是数组类型的一部分 【重点】 a := [3]bool 不能比较  b:= [4]bool 为不相同的类型
func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T\n", a1, a2)
	// 数组初始化 默认都是零值 bool的时候是false 整型，浮点型是 0 字符串是 “”
	a1[0] = true
	fmt.Println(a1)
	// 初始化1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 初始化2  不设置长度,让他自己去数
	a10 := [...]int{1, 2, 3, 3, 4, 4, 5, 4, 66}
	fmt.Println(a10)
	// 初始化3
	a5 := [5]int{2: 10, 3: 199} // 按索引初始化，不设置为默认值
	fmt.Println(a5)

	// 数组的遍历
	a11 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
		

	fmt.Println(a11)
	for _, v := range a11 {
		// fmt.Println(v)
		for _, c := range v {
			fmt.Println(c)
		}
	}
	// 数组的值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1 //
	b2[0] = 100
	fmt.Println(b1)
	fmt.Println(b2)

	// 练习题
	b3 := [...]int{1, 3, 4, 9, 6, 7, 8, 6, 5}
	// 数组求和
	num := 0
	for _, k := range b3 {
		num += k
	}
	fmt.Println(num)
	fmt.Println()
	// 找出两个元素的合为10的下标
	for k1, v1 := range b3 {
		for i := k1 + 1; i < len(b3); i++ {
			if v1+b3[i] == 10 {
				fmt.Println(k1, i)
			}
		}
	}
	//支持的写法
	a19 := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}

	//不支持多维数组的内层使用...
	// b := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	fmt.Println(a19)

}
