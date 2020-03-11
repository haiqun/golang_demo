package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const fileName = "./log.txt"

func writeDemo1() {
	//O_TRUNC 每次打开情况
	// os.O_APPEND 每次打开追加
	// os.O_CREATE 不存在就创建
	// os.O_WRONLY 只写
	fileObj, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		fmt.Printf("报错了 %s", err)
		return
	}
	defer fileObj.Close()
	// write
	fileObj.Write([]byte("我是写入-1111\n"))
	// writeString
	fileObj.WriteString("牛逼呀")
	// fmt.Println(fileObj)
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // 获取文件话柄
	if err != nil {
		fmt.Printf("openfile 报错了 ：%s，\n", err)
		return
	}
	defer fileObj.Close()
	write := bufio.NewWriter(fileObj) //  创建一个缓冲区的对象

	for i := 0; i < 10; i++ {
		write.WriteString("so cool~ \n") // 写入缓存区
	}
	write.Flush() // 刷到文件
}

func writeDemo3() {
	str := "testjjjj "
	err := ioutil.WriteFile("./log.txt", []byte(str), 0664)
	if err != nil {
		fmt.Printf("ioutil.WriteFile 报错了 ：%s，\n", err)
		return
	}
}

// 在文中插入 一段话

func f1(fileName string, num int, content string) {
	// 打开源文件
	os.Chmod(fileName, 0777)

	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("os.Open 打开有误 err ：%s\n", err)
		return
	}
	// 读取到指定的位置
	// var b [num]byte
	b := make([]byte, num)
	n, err := fileObj.Read(b)
	if err != nil {
		fmt.Printf("os.Read 读取有误 err ：%s\n", err)
		return
	}
	q := string(b[:n])

	q = q + content
	// 将读取到的内容写到临时文件去
	tmp := "./log.tmp"
	w1, err := os.OpenFile(tmp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("os.OpenFile 打开临时文件有误 err ：%s\n", err)
		return
	}
	_, err = w1.WriteString(q)
	// 在新文件写入要临时文件
	if err != nil {
		fmt.Printf("os.WriteString 写入临时文件有误 err ：%s\n", err)
		return
	}
	// 从源文件继续读取内容，追加到临时文件
	var x [256]byte // 每次读 256 个字符
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			fmt.Println("操作成功")
			break
		}
		if err != nil {
			fmt.Printf("os.Read 读取有误 err ：%s\n", err)
			return
		}
		q1 := string(x[:n])
		w1.WriteString(q1)
	}
	// 关闭源文件 关闭临时文件
	err = fileObj.Close()
	err = w1.Close()
	// 重新命名临时文件为源文件名
	os.Rename(tmp, fileName)
}

// 文件写入
func main() {
	// writeDemo1()
	// writeDemo2()
	// writeDemo3()
	f1("./log.txt", 100, "\"我是第100个字符之后插入的一句话\"")
}
