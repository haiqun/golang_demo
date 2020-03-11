package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const num = 128

func fileOpen() {
	fileObj, err := os.Open("./main.go")
	// fmt.Printf("%T , %v",fileObj,fileObj)
	if err != nil {
		fmt.Println("打开文件有误：", err)
		return
	}
	defer fileObj.Close()
	var tmp [num]byte // 每次读取 128的字符
	for {
		n, err := fileObj.Read(tmp[:]) //返回的n为读取到的字符数
		if err == io.EOF {
			fmt.Println("文件读完了~ ")
			return
		}
		if err != nil {
			fmt.Println("读取文件有误", err)
			return
		}
		fmt.Printf("读取了 %d 个字节", n)
		fmt.Println(string(tmp[:n]))
		if n < num { // 这一次已经不够128字符，没有下一次了
			return
		}
	}
}

func rerdFromFileByBufio() {
	fileObj, err := os.Open("./main.go")
	// fmt.Printf("%T , %v",fileObj,fileObj)
	if err != nil {
		fmt.Println("打开文件有误：", err)
		return
	}

	defer fileObj.Close()
	// bufio 获取内容
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完成")
			return
		}
		if err != nil {
			fmt.Println("bufio ReaderString 读取有误")
			return
		}
		fmt.Print(line)
	}

}

func readFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("ioutil 读取文件有误 ：%s \n", err)
	}
	fmt.Println(string(ret))
}

// 文件操作
func main() {
	// fileOpen()
	// rerdFromFileByBufio()
	// readFileByIoutil()

	// 读取偏移量
	fileObj, _ := os.Open("./log.txt")
	fileObj.Seek(100, 0) // 指针偏移 100 位
	var b = make([]byte, 20)
	fileObj.Read(b[:])
	fmt.Println(string(b[:20]))
	// 获取当前指针
	cur_offset, _ := fileObj.Seek(0, os.SEEK_CUR)
	fmt.Printf("current offset is %d\n", cur_offset)

}
