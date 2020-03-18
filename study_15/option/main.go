package main

import "fmt"
// 选项模式  -  让传值更灵活
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
