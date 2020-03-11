package main

import "fmt"

// 闭包练习
func c(base int) (func(int) int, func(int) int) {
	n := 100
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		// n1 := n - i
		n -= i
		return n
	}
	return add, sub

}

func main() {
	f1, f2 := c(10)
	fmt.Println(f1(1), f2(2)) // 11, 9 因为base是公用的，每次调用都会修改
	fmt.Println(f1(3), f2(4)) // 12 ,8
	fmt.Println(f1(5), f2(6)) // 13, 7
}
