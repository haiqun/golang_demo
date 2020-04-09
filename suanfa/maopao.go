package main

import (
	"fmt"
)

// 冒泡排序
/*
	冒泡排序是一种比较简单的排序算法，它循环走过需要排序的元素，
	依次比较相邻的两个元素，如果顺序错误就交换，直至没有元素交换，完成排序
*/

// f1 普通版本
func f1(t []int,total int)  {
	for i:=0;i<total;i++ {
		for j:=0;j<total;j++ {
			if j+1 < total{
				if t[j] > t[j+1] {
					tmp := t[j]
					t[j] = t[j+1]
					t[j+1] = tmp
				}
			}
		}
	}
	fmt.Println(t)
}

// f2 鸡尾酒排序
func f2(t []int,total int)  {
	left := 0;
	right := total
	for  {
		if left >= right {
			break
		}
		// 正向冒泡 找最大值
		for j:=left;j<right;j++ {
			// 比到最后一位的时候，就不用再跟下一位对比了
			if j+1 < total {
				// 两个连续的值对比后，前面的值大于后面的时候，需要将两个值进行替换
				if t[j] > t[j+1] {
					tmp := t[j]
					t[j] = t[j+1]
					t[j+1] = tmp
				}
			}
		}
		// 上面一轮，已经确认一个最大值，可以缩小范围
		right--
		// 反向冒泡，取最小值
		for j:=right;j>left;j-- {
			// 如果
			if j != left {
				if t[j] < t[j-1] {
					tmp := t[j]
					t[j] = t[j-1]
					t[j-1] = tmp
				}
			}
		}
		// 确认最小值 ，缩小范围
		left ++
	}
	fmt.Println(t)
}

func main() {
	t := []int{1,49,34,92,47,21,19,77,32}
	total := len(t); // 元素个数
	//f1(t,total)
	f2(t,total)
}
