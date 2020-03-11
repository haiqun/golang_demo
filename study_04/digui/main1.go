package main

import "fmt"

// 递归 - 自己调用自己
// 适合处理相同问题，但是规模越来越小的场景
// 递归一定要有明确的退出条件

func f1(n int) int {
	if n > 1 {
		return n * f1(n-1)
	}
	return 1
}

func main() {
	n := f1(5)
	fmt.Println(n)
}
