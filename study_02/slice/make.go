package main

import "fmt"

// 使用make 创建切片
// 切片就是一个框，框住一块连续的内存，只能保存相同的类型的元素，属于引用类型，底层是数组

func twoSum(nums []int, target int) []int {
	s1 := make([]int, 0, 2)
	// for k1, v := range nums {
	// 	for k2, v1 := range nums {
	// 		if k1 == k2 {
	// 			continue
	// 		}
	// 		n1 := v + v1
	// 		if n1 == target {
	// 			s1 = append(s1, k1, k2)
	// 			return s1
	// 		}
	// 	}
	// }
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return append(s1, i, j)
			}
		}
	}
	return s1
}

func main() {
	a := []int{2, 7, 11, 15}
	ret := twoSum(a, 9)
	fmt.Println(ret)
	// s1 := make([]int, 5, 10) // 整数类型 ，长度 5 ，容量 10
	// fmt.Println(s1, len(s1), cap(s1))
	// s2 := make([]string, 2, 10)
	// fmt.Println(s2, len(s2), cap(s2))
	// 切片之间不能比较，不能使用 == 来判断两个切片是否有包含全部相等的元素
	// 要用 len(s1) == 0判断切片是不是空的 ，不能用 nil

	// 切片的遍历
	// a1 := []int{1, 2, 3, 4, 5, 7, 43, 456, 9}
	// for k, v := range a1 {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println()
	// for i := 0; i < len(a1); i++ {
	// 	fmt.Println(i, a1[i])
	// }
	// fmt.Println()

}
