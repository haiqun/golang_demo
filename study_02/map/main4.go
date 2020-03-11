package main

import (
	"fmt"
	"strings"
)

// 练习题
func main() {
	//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	s1 := "how do you do hi hai zui wang hi hello world"
	// 字符串切割
	ret := strings.Split(s1, " ")
	// 初始化map
	map1 := make(map[string]int, 20)
	for _, v := range ret {
		// 判断map的值是否存在 ，存在+1
		_, ok := map1[v]
		if ok {
			map1[v]++
		} else {
			map1[v] = 1
		}
	}
	fmt.Println(map1)
}
