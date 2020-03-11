package main

import "fmt"

// for 循环 - golang唯一的循环结构
func main() {
	// 分几种
	// 基本格式
	// for i := 1; i < 3; i++ {
	// 	fmt.Printf("我是哈哈哈,第%v次循环\n", i)
	// }
	// 变种1
	// i := 1
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// 变种2
	// i := 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// 无限循环
	// for {
	// fmt.Printf("i")
	// }
	// fmt.Printf("我是哈哈哈")
	// for range 循环
	// s := "hello 我是海贼王路飞"
	// for i, v := range s {
	// 	fmt.Printf("%d,%c\n", i, v)
	// }

	// 联系题 打印99乘法表
	i := 1
	for i < 10 {
		j := 1
		for i >= j {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
			j++
		}
		fmt.Println()
		i++
	}
}
