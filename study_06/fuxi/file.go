package main

import (
	"fmt"
	"os"
)

// 判断文件大小

func getFileSize() {
	// 打开一个文件
	fileObj, err := os.Open("./file.go")
	if err != nil {
		fmt.Printf("读取文件报错 ：%s", err)
		return
	}
	// 获取文件的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("获取文件详细信息报错 ：%s", err)
		return
	}
	fmt.Printf("文件的大小是 %d B1\n", fileInfo.Size())
	fmt.Printf("文件的名称是 %s \n", fileInfo.Name())
}

func main() {
	getFileSize()
}
