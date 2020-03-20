package main

import (
	"fmt"
	"math"
)

// 数字反转
func f2(x int ) int  {
	result:=0
	for x!=0{
		result=result*10+(x%10)//对原数对10取余则可以得到最后一位数
		x/=10//左移
	}
	// 判断数字溢出，直接返回0
	if result>math.MaxInt32||result<math.MinInt32{
		result = 0
	}
	return result
}



func main()  {
	n := 1233823984
	f := f2(n)
	fmt.Printf("%d,反转字符串的结果是：%d\n",n,f)
}
