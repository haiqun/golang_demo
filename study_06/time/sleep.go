package main

import (
	"fmt"
	"time"
)

// sleep
func main() {
	n := 4 //
	n1 := time.Duration(n)
	fmt.Printf("%T\n", n1)
	fmt.Println(time.Now())
	time.Sleep(n1 * time.Second) // 休眠 4 秒钟
	fmt.Println(time.Now())
}
