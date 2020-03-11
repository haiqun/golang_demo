package main

import (
	"fmt"
	"os"
)

// os.Args  获取命令行参数

func main() {
	data := os.Args //[]string
	fmt.Printf("%T\n", data)
	fmt.Printf("%#v\n", data)
}
