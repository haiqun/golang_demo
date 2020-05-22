package main

import"fmt"

func main() {
	h, w := 2, 4

	raw := make([]int, h*w)

	for i := range raw {
		raw[i] = i
	}
	fmt.Println(raw, &raw[4]) //prints: [0 1 2 3 4 5 6 7] <ptr_addr_x>

	table := make([][]int, h)
	for i := range table {
		table[i] = raw[i*w : i*w+w]
	}

	fmt.Println(table, &table[1][0]) //prints: [[0 1 2 3] [4 5 6 7]] <ptr_addr_x>}
}



func f1(arr [3]int) {
	arr[0] = 7
	fmt.Println(arr) //prints [7 2 3]}(x)
}

func f2(arr []int) {
	arr[0] = 7
	fmt.Println(arr) //prints [7 2 3]}(x)
}


