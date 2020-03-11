package main

import "fmt"

import "bufio"

import "os"

//获取客户端输入，带空格的
func useBufio() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输出你要打印的内容 ：")
	str, _ = reader.ReadString('\n') // 读到什么时候结束
	fmt.Println("你输入的内容是 :", str)
}

func useScan() {
	var str string
	fmt.Print("请输出你要打印的内容 ：")
	fmt.Scanln(&str)
	fmt.Println("你输入的内容是 :", str)
}

func main() {
	// useBufio()
	useScan()
}
