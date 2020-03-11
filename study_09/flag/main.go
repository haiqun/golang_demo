package main

import (
	"flag"
	"fmt"
	"time"
)

// flag
// 不知传什么参数可以 ./flag --help
// ./flag -name xxx -age=xx  传参一定要用 '-'加字段名 参数的值可以是 空格也可以是=号赋值
func main() {

	// 创建标志位参数
	// 返回值是指针，所以需要用 * 取值
	name := flag.String("name", "lufei", "请输入你的姓名")
	age := flag.Int("age", 11, "请输入你的年龄")
	b1 := flag.Bool("hai", false, "是否海贼王")
	ctime := flag.Duration("ct", time.Second, "输入时间")
	var n1 string
	flag.StringVar(&n1, "n1", "suoluo", "n1的值")
	// 使用
	flag.Parse() // 一定要解析 ，不然无效
	fmt.Println(n1)
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*b1)
	fmt.Println(*ctime)
	fmt.Println()
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
