package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

// 读取ini配置文件 - 自己写的版本

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func loadIni(myc *MysqlConfig) {
	// 打开文件
	fileObj, err := os.Open("./conf.ini")
	if err != nil {
		fmt.Printf("读取文件报错 err：%s\n", err)
		return
	}

	defer fileObj.Close()
	// 读取文件的内容 (一行行读取)
	reader := bufio.NewReader(fileObj)
	var m1 = make(map[string]string, 10)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			// fmt.Println("读取完成")
			break
		}

		if err != nil {
			fmt.Printf("bufio ReaderString 读取有误 err:%s", err)
			return
		}
		// 按照格式切割字符串 =》 key value
		ret := strings.Split(line, "=")
		if len(ret) != 2 {
			continue
		}
		m1[ret[0]] = ret[1]
	}
	// 获取结构体的信息
	t := reflect.TypeOf(*myc)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagName := field.Tag.Get("ini")
		name := field.Name
		// 判断m1中是否存在 tagName
		// ret := reflect.ValueOf(m1).MapIndex(reflect.ValueOf(tagName)).IsValid()
		_, ret := m1[tagName]
		if ret {
			// 对比结构体标签的名字 赋值
			switch name {
			case "Address":
				myc.Address = m1[tagName]
			case "Port":
				myc.Port = m1[tagName]
			case "Username":
				myc.Username = m1[tagName]
			case "Password":
				myc.Password = m1[tagName]
			}
		}
	}
	// fmt.Println(myc)
	return
}

func main() {
	var myc MysqlConfig
	loadIni(&myc)
	fmt.Println(myc)
}
