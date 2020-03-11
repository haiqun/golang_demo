package main

import "fmt"

// defer 后面的预计延迟执行，延迟到返回的时候在再执行
func main() {
	fmt.Println("start")
	i := 0
	defer fmt.Println(i)
	defer fmt.Println("hhh")
	defer fmt.Println("aaa")
	i += 100
	fmt.Println("end")
}

