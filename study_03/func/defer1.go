package main

import "fmt"

// go 语言中函数return 不是原子性的，他是分两步的
// 第一步 ：返回值赋值
// defer 执行
// 第二步 ：真正的return 返回
// 函数中如果存在 defer 那么defer执行的时机是在第一步与第二步之间

func f1() int {
	x := 5 // 1. 返回值赋值
	defer func() {
		x++ // 2. x 自增 
	}()
	return x // 1 返回值复制 x = 5(单独开辟可内存空间) ; 2.defer 修改x,不影响返回值; 3.真正返回 x = 5
}

func f2() (x int) { 
	defer func() {
		x++   
	}()
	return 5 // 1.返回值赋值x=5，开辟x内存空间; 2.defer 修改 x; 3.真正返回 x 所以是 6
}

func f3() (y int) { 
	x := 5 
	defer func() {
		x++ 
	}()
	return x // 1. 返回值赋值 y = x = 5 - x,y的内存空间为不同的两个; 2. defer 修改的是 x ；3.真正返回 x = 5
}
func f4() (x int) { 
	defer func(x int) { 
		x++ 
	}(x) 
	return 5 // 1. 返回值赋值 x=5 开辟内存空间; 2.defer 修改的是内部变量x; 3.真正返回 x = 5
}

func f5() (x int)  {
	defer func (x int) int  {
		x++
		return x
	}(x)
	return 5 // 1. 返回值赋值 x =5;2.defer 修改了内部x的变量，并返回，没有接受; 3.返回 x = 5
}

// 传一个x的指针到匿名函数进去
func f6()(x int)  {
	defer func (x *int)  {
		(*x)++
	}(&x)
	return 5 // 1.返回值 x = 5 ;2.执行defer，传入内存地址，修改了x ; 3.返回 x=6
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
