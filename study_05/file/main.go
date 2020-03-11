package main

import (
	"fmt"
	"os"
)

func f1() {
	fileObj, err := os.Open("./log.txt1")
	// defer fileObj.Close() // 这里有问题，如果文件不存在，obj为nil 不能调用close 方法
	if err != nil {
		fmt.Printf("开到文件错误 : %s", err)
		return
	}
	defer fileObj.Close() // 放这里调用
	fileObj.Read()
	// fmt.Println(fileObj)
}

// 文件操作
func main() {
	f1()
}
