package main

import (
	"fmt"
	"strings"
)

// 练习题
/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() int {
	fs := make(map[string]int, len(users)+1)
	num := 0
	for _, v := range users {
		// 装换大小写
		v1 := strings.ToUpper(v)
		for _, v2 := range v1 {
			// v3 := string(v2) 
			// 输出每个用户的金额
			if v2 == 'E' { // 这里是字节，直接拿直接比较就可以，如果用字符串，那么需要上面的转换
				num += 1
				fs[v] += 1
				coins -= 1
			} else if v2 == 'I' {
				num += 2
				fs[v] += 2
				coins -= 2
			} else if v2 == 'O' {
				num += 3
				fs[v] += 3
				coins -= 3
			} else if v2 == 'U' {
				num += 4
				fs[v] += 4
				coins -= 4
			}
		}
	}
	// fmt.Println(fs, "总共分配出去：", num)
	// 返回剩余金额
	return coins
}
