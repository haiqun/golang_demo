package main

import "fmt"

// panic 合recover

func f1() {
	fmt.Println("a")
}
func f2() {
	defer func() { // recover 必须 配合defer使用
		err := recover() // 程序不会退出，继续执行
		fmt.Println(err)
		fmt.Println("释放数据库链接")
	}()
	// panic 必须用在 defer 之前
	panic("错误了呀") // 程序崩溃了
	fmt.Println("b")
}
func f3() {
	fmt.Println("c")
}
func main() {
	f1()
	f2()
	f3()
}
