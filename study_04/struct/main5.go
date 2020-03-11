package main

import "fmt"

type dog struct {
	name string
}

// 标识符 ： 变量名，函数名，类型名 ，方法名
// go 语言中，如果标识符首字母是大写的，就表示对外部可见 （暴露的，共有的），需要注释
// 小写的，外部不可见
// 如 fmt 包的 Printf，只要我引入（import ”fmt“） 就能直接调用 Printf 

// Dog 就是一个dog的结构体 
type Dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法与接受者
// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用类型名字首字母小写表示 d（形参） dog（类型）
// 函数名前面加一个（）表示是该函数的方法，不能被单独调用
func (d dog) wang() {
	fmt.Printf("%s 汪汪汪~", d.name)
}



func main() {
	a := newDog("lufei")
	a.wang()
}
