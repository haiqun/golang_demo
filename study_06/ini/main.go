package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// 读取ini配置文件

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 1 参数校验
	// data 参数必须是指针类型 , 因为需要在函数中对它进行赋值
	t := reflect.TypeOf(data)
	// fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr { // Ptr 就是指针类型
		err := fmt.Errorf("参数必须是指针类型")
		return err
	}
	// data 指针的结构必须是结构体类型的，因为配置中的赋值需要键值对的类型
	// Elem 获取指针对应的值，值再获取他的类型
	if t.Elem().Kind() != reflect.Struct {
		err := fmt.Errorf("指针的结构必须是结构体类型的")
		return err
	}
	// 读取文件的字节类型数据 => 读取整个文件
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	// 一行行的读取数据类型
	lineSlice := strings.Split(string(b), "/n")
	// fmt.Println(lineSlice)
	// 如果是注释就跳过
	for idx, line := range lineSlice {

	}
	// 如果是[开头的就跳过

	// 如果不是[开头的，那就键值对 = 分割

	return nil

}

func main() {
	var myc MysqlConfig
	loadIni("./conf.ini", &myc)
	fmt.Println(myc)
}
