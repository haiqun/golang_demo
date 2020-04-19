package main

import (
	"fmt"
	"runtime"
)


func main() {

	/*
	   执行结果
	   B.defer
	   A.defer
	*/
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit() // 退出当前协程
			// 这里的程序不会被执行
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		// 这里也不会被执行
		fmt.Println("A")
	}()

	for {

	}
}