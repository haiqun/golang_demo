package main

import (
	"fmt"
	"sort"
)

// 快速排序
func f1(t []int,totol int)  {
	// 取切片的最后一个元素为基数 base
	base := t[totol-1]
	// 将比base 大的放一左边，比base大的放右边
	//l := make([]int,totol)
	//r := make([]int,totol)
	//eq := make([]int,totol)

	for _,v := range t{
		if v > base {
			// 比基本数大
			//append(l,v)
		} else if v < base {
			// 比基本数小
			//append(r,v)
		} else {
			//append(eq,v)
		}
	}
	// 递归排序

	// 拼接切片
}

func main() {
	t := []int{2,23,14,89,74,28,11,33,45,89,9,60}
	n := len(t)
	var depth int
	for i := n; i > 0; i >>= 1 { // 1100 => 110 => 11 => 1
		depth++
	}
	a := depth * 2
	fmt.Println(a)
	return
	sort.Ints(t)
	fmt.Println(t)
	//totol := len(t)
	//f1(t,totol)
}
