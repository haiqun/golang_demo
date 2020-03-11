package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "test1"
	s2 := "t1"
	// 字符串 s2 在 s1 出现的次数？
	n1 := strings.Count(s1, s2) + 1
	n2 := len(s1)
	fmt.Println(n1, n2)
}
