package main

import "fmt"

// 运算符号
func main() {
	var (
		a = 2
		b = 5
	)

	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)

	a++  // 单独的语句 ，不能放在 = 的右边进行赋值 
	b-- 
	fmt.Println(a)

	// 关系运算符 
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a < b)

	// 逻辑运算符 
	age := 7

	if age >18 && age < 60 {
		fmt.Println("努力上班")
	}else{
		fmt.Println("不用上班")
	}

	if age < 18 || age >60 {
		fmt.Println("美滋滋")
	}

	isOk := false
	if !isOk{
		fmt.Println("not ok")
	}


	// 位运算 

	/*
	& 与运算 ，两个为1 才为 1 
	| 或运算 ，两个有一个为1 就为 1
	^ 异或运算 ，两位不一样 则为 1  
	左移 5 << 1 => 101 => 1010 => 10 
	右移 
	二进制 111 => 4+2+1 => 7 
	*/

	// 赋值运算符 

	var x int
	x = 10
	x += 1
	fmt.Println(x)
	x -= 1 
	fmt.Println(x)
	x /= 2 
	fmt.Println(x)
	x %= 2 
	fmt.Println(x)
	x *= 3
	fmt.Println(x)

	x <<= 3 // x = 3 等于二进制的11左移3位变成了11000 => 16*1+8*1+4*0+2*0+1*0 = 24 
	fmt.Println(x)
	x |= 3
	fmt.Println(x)
	x ^= 10 
	fmt.Println(x)
	x &= 5
	fmt.Println(x)
	x >>= 2
	fmt.Println(x)






	



	


}
