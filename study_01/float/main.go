package main

import (
	"fmt"
	"math"
)

// 浮点数 布尔值
func main() {
	f2 := math.MaxFloat32 // 声明最大的浮点数 32
	fmt.Println(f2)
	i := 1.234234
	fmt.Println(i)
	fmt.Printf("%T\n", i) // 默认所以的小数都是float64
	f1 := float32(1.2344)
	fmt.Printf("%T\n", f1) // 显示声明 32位的浮点数
	b1 := true
	var b2 bool
	fmt.Printf("%T\n", b1)            // 显示声明 32位的浮点数
	fmt.Printf("%T value %v", b2, b2) // 布尔值的默认值是false %v可以打印所以的变量
}
