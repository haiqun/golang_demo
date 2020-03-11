package main

import "fmt"

//导入语句

/*
mac 平台交叉编译
CGO_ENABLED = 0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED = 0 GOOS=windows GOARCH=amd64 go build
*/

// 函数外只能放置标识符号（变量，变量，函数，类型 ）的声明

// 声明变量 -批量-变量推荐小驼峰式 studentName

var (
	name string
	age  int
	isOk bool
)

func main() {
	// fmt.Println("hello world")
	// 非全局变量，声明后必须使用，不使用编译不过去
	var nameL string = "name+lufei"
	fmt.Println(nameL)
	// 类型推导
	var num1 = "38"
	var age1 = 20
	fmt.Println(num1)
	fmt.Println(age1)
	// 短变量声明  - 只能在函数里面使用
	test := 329
	// test := 200 # 同一个作用域（{}括号）中不能重复声明同名变量
	fmt.Println(test)
	test2 := "我是test"
	fmt.Println(test2)
	// 匿名变量 如果是多重赋值，忽略某个值使用  (变量不想使用的时候复制给"_")不占命名空间，不分配内存
	name = "鲁夫"
	age = 16
	isOk = true

	if isOk {
		fmt.Print(isOk)               // 直接打印
		fmt.Println("")                 //空行
		fmt.Printf("name:%s\n", name) // 使用%s 占位符，使用name 这个变量的值去替换占位符
		fmt.Println(age)              // 换行打印
	}
}
