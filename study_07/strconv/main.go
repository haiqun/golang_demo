package main

import "fmt"

import "strconv"

//strconv

func main() {
	var i int64 = 1999
	// 数字转字符串
	str := fmt.Sprintf("%d", i)
	fmt.Printf("%T,%s\n", str, str)

	// 字符串转数字
	var s1 string = "99"
	//不能这样转
	// n1 := fmt.Sprintf("%d", s1)
	//正确的转法
	n1, err := strconv.ParseInt(s1, 10, 0) // 参数（字符串，十进制，）返回值（int 64）
	if err != nil {
		fmt.Printf("转向有误,%s\n", err)
	}
	fmt.Printf("%T,%#v\n", n1, n1)
	// string 转 int
	n2, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Printf("转向有误,%s\n", err)
	}
	fmt.Printf("%T,%#v\n", n2, n2)
	// int 强转 string
	s2 := strconv.Itoa(n2)
	fmt.Printf("%T,%#v\n", s2, s2)
	// 字符串解析布尔值
	s3 := "true"
	b1, err := strconv.ParseBool(s3)
	if err != nil {
		fmt.Printf("ParseBool,%s\n", err)
	}
	fmt.Printf("%T,%#v\n", b1, b1)
	// 字符解析浮点数
	s4 := "1.3489348"
	f1, err := strconv.ParseFloat(s4, 64)
	fmt.Printf("%T,%#v\n", f1, f1)

}
