package main

// 三数之和
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
// -1, 0, 1, 2, -1, -4

func threeSum(nums []int) [][]int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 将两个人的值加一起，再去找第三个数
		}
	}
}

func main() {
	n1 := []int{-1, 0, 1, 2, -1, -4}
	threeSum(n1)
}
