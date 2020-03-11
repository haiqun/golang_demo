package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller()
// 获取调用方法，文件名等信息

func f1(n int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller  报错")
	}
	name := runtime.FuncForPC(pc).Name()
	fmt.Println(name)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}
func main() {
	f1(1)
	// s := strconv.Itoa(74)
	// fmt.Println(s)
}
