package main

import "fmt"

// 自定义类型和类型别名

// type 后面跟的是类型
type myInt int    // 自定义
type youInt = int // 类型别名 , 编译完就不存在的了

func main() {
	var n myInt
	n = 100
	var b youInt
	b = 100
	fmt.Println(n, b)
	fmt.Printf("%T %T\n", n, b)

	var c rune // 32 位有符号整形 ， int32 的别名
	c = '中'

	fmt.Println(c)
	fmt.Printf("%T %c\n", c, c)

}
