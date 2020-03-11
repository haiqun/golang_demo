package main

import "fmt"

// 常量
// 定义之后不能修改
// 运行期间不会改变的量
const pi = 3.14159267

// 批量声明
const (
	statusok = 200
	notfound = 404
)

// 某一行没有声明，默认跟上一行一致
const (
	n1 = 100
	n2
	n3
	n4 = 200
)

// const 出现的时候 iota 会重置为 0

// iota 常量计时器 - 只能在常量的表达式中使用 - 定义枚举值时很有用
const (
	t1 = iota // 0
	t2        //1
	t3        //2
	t4        //3
)

// 题目

const (
	k1 = iota //0
	k2        //1
	_         //2
	k3        //3
)

const (
	j1 = iota //0
	j2 = 100  //100
	j3        // 2 # 每声明一行常量，iota加1 ，所以这里是第3个变量，为2
	j4        // 3
)

// 多个变量在一行
const (
	q1, q2 = iota + 1, iota + 2 // q1 = 1,q2 = 2
	q3, q4 = iota + 1, iota + 2 // q2 = 2,q4 = 3
)

func main() {
	// fmt.Println(pi)

	// fmt.Println(n1)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// fmt.Println(n4)

	// fmt.Println(t4)

	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q3)
	fmt.Println(q4)
	
}
