package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		fmt.Println("i等于5(continue)")
	// 		continue
	// 	}
	// 	if i == 8 {
	// 		fmt.Println("i等于8(break)")
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// switch n := 100; n {
	// case 1:
	// 	fmt.Println("n等于")
	// case 2:
	// 	fmt.Println("n等于2")
	// case 3, 4, 5, 6:
	// 	fmt.Println("n大于2小于7")
	// default:
	// 	fmt.Println("牛逼")
	// }
	// 跳出多层for循环
	// flag := false
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 3 {
	// 			flag = true
	// 			break
	// 		}
	// 		fmt.Printf("i的值%d，j的值%d\n", i, j)
	// 	}
	// 	if flag {
	// 		break
	// 	}
	// }
	// goto 跳出多层循环
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if j == 3 {
// 				goto xx
// 			}
// 			fmt.Printf("i的值%d，j的值%d\n", i, j)
// 		}
// 	}

// xx:
// 	fmt.Println("over!!!!")

	// case 接判断表达式
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
	

}
