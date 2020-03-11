package main

// 整型
import "fmt"

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%o\n", i1) // 十进制转八进制 -> 文件权限时会使用
	fmt.Printf("%b\n", i1) // 十进制转二进制 1100101 => 64*1-32*1-16-8-4*1-2-1*1
	fmt.Printf("%x\n", i1) // 十进制转十六进制 -> 内存地址时会使用到

	// 查看变量的类型
	fmt.Printf("%T\n", i1) //

	// 声明一个int8类型的变量
	i4 := int8(12)          // 明确指定类型，否则默认是int类型
	fmt.Printf("%T\n %v", i4,i4) //

}
