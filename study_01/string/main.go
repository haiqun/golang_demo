package main

//  字符串 使用双引号包裹的 ，用单引号包裹的是字符 如 '1','录' 单独的字母，汉字，符号标识一个字符
import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// z := '哈'
	h := "我是路飞"
	fmt.Println(h)
	// 多行字符串
	t := `
		我是很长的字符串
		哈哈哈哈哈
	`
	fmt.Println(t)

	// 字符串相关操作
	// 拼接  可以用 “+”号
	z1 := "我是海贼王hhhh"
	z2 := "-lufei"
	fmt.Println(z1 + z2)
	s1 := z1 + z2 + "海贼-王-是我"
	fmt.Println(s1)
	s2 := fmt.Sprintf("%s%s", z1, z2)
	fmt.Println(s2)

	// 字符串长度
	fmt.Println(len(z1)) // 一个中文 3个长度

	for i := 0; i < len(z1); i++ {
		fmt.Printf("%c%T\n", z1[i], z1[i]) // 按字节输出 byte ，中文乱码了
	}

	fmt.Println("-=========-")

	for _, c := range z1 {
		fmt.Printf("%c\n", c) // 按 rune类型输出 -  处理非ascii吗的类型的字符 代表utf8的一个字符
	}

	// 字符串分割
	ret := strings.Split(s1, "-")
	fmt.Println(ret)
	// 包含
	q1 := strings.Contains(s1, "王")
	q2 := strings.Contains(s1, "test")
	fmt.Printf("%T q1 value:%v\n", q1, q1)
	fmt.Printf("%T q2 value:%v\n", q2, q2)
	// 前缀后缀判断
	fmt.Println(strings.HasPrefix(s1, "我是")) // 前
	fmt.Println(strings.HasSuffix(s1, "我是"))

	// 判断字符串位置
	fmt.Println(strings.Index(s1, "海"))     // 第一次出现的位置
	fmt.Println(strings.LastIndex(s1, "海")) // 最后一次出现的位置

	// 判断字符串出现了几次
	s10 := "test1"
	s20 := "t1"
	// 字符串 s2 在 s1 出现的次数？
	n10 := strings.Count(s10, s20)
	fmt.Println("判断字符串出现:", n10)

	// 字符串拼接 - 数组变字符串拼接
	fmt.Println(strings.Join(ret, "/"))
	// 字符串修改

	w1 := "大白菜"             // => '大' '白' '菜'
	w2 := []rune(w1)        // 把字符串强制转换成rune切片
	w2[0] = '蓝'             // 所在这里不能用 "" 字符串，要用 ’‘ 的字符
	fmt.Println(string(w2)) // 把切片转为字符串输出

	c1 := "红"                     // string
	c2 := '红'                     // rune(int32)
	fmt.Printf("%T,%T\n", c1, c2) // 输出类型确认

	str := "hello广东"
	//含有中文的，要遍历前先转为切片
	rstr := []rune(str)
	for i := 0; i < len(rstr); i++ {
		fmt.Printf("字符=%c\n", rstr[i])
	}
	// 不转切片的遍历就是会乱码
	for i := 0; i < len(z1); i++ {
		fmt.Printf("%c%T\n", z1[i], z1[i]) // 按字节输出 byte ，中文乱码了
	}

	// 类型装换
	n1 := 10
	var f float64
	fmt.Printf("%T,value %v\n", n1, n1)
	f = float64(n1)
	fmt.Printf("%T,value %v\n", f, f)

	// 练习 - 判断字符串有几个汉字
	t2 := "dhj路飞d库"
	num := 0
	for k, i := range t2 {
		if unicode.Is(unicode.Han, i) {
			fmt.Printf("类型 %T，对应的值：%c ,rune的顺序是 %v\n", i, i, k)
			num++
		}
	}
	fmt.Println(num)

	url := "https://www.jianshu.com/p/ef946da14304?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes"
	fmt.Println("举个例子：", url)
}
