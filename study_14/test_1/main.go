package main

import "math"

func f2(x int ) int  {
	result:=0
	for x!=0{
		result=result*10+(x%10)//对原数对10取余则可以得到最后一位数
		x/=10//左移
	}
	if result>math.MaxInt32||result<math.MinInt32{
		result = 0
	}
	return result
}



func main() {
	n := 123
	f := f2(n)
	fmt.Println(f)
}
