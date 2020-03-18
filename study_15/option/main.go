package main

import "fmt"
// 选项模式  -  让传值更灵活
/*
	思路分析  ：
	1.我要给不同的类型赋值，我不能限制他们的参数，但是我可以限制他们的返回值
	2.基于一样的返回值，声明一种函数变量 ，其他赋值函数只要符合我的返回值邀请即可
	3.参数设置采用了闭包的方式，设置参数，返回一个函数
*/
type Options struct {
	str1 string
	str2 string
	str3 string
	int1 int
	int2 int
	int3 int
}

type OptF func(opt *Options)

func WithStrOption1(s string) OptF {
	return func(opt *Options) {
		opt.str1 = s
	}
}

func WithIntOption1(i int) OptF {
	return func(opt *Options) {
		opt.int1 = i
	}
}

// 实例化结构体
func InitOption(opts ...OptF )  {
	// 初始化一个空的结构体 ，接受值
	options := &Options{}
	for _,v := range opts{
		fmt.Println(v)
		// 调用函数，在函数内赋值
		v(options)
	}
	fmt.Println(options)
}


func main() {
	a:= WithStrOption1("test")
	InitOption(a,WithIntOption1(2))
}
