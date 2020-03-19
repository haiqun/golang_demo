package main

// 回文判断

import "fmt"

func ff(str string) bool {
	// 判断字段长度 
	if len(str) == 0 {
		return false
	}
	// 转为 rune 类型
	var r = []rune(str)
	// 遍历判断
	i,j:=len(r)-1,0
	for i> j{
		// 第一个跟最后一个相等，以此类推
		if r[i] == r[j] {
			i--
			j++
		}else{
			return false
		}
	}
	return true
}

func main() {
	str := "12321"
	t := ff(str)
	fmt.Println(t)
}
