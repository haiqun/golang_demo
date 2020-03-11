package main

import "fmt"
import "encoding/json"

type Address struct {
	City    string `json:"city"`
	Provice string
	Address string `json:"address"`
}

type studens struct {
	Name    string
	Age     int
	Address // 等效 address address
}

type dog struct {
	string
}

func newStudent(name string, age int) *studens {
	return &studens{
		Name: name,
		Age:  age,
	}
}

func (s studens) gotoshools() {
	fmt.Printf("%s , 愉快的上学", s.Name)
}

func (a Address) shoolsAddr() {
	fmt.Printf("学校所在地 ：%s ", a.Address)
}

// struct 复习
func main() {
	p1 := studens{
		Name: "小明",
		Address: Address{
			Address: "广州-越秀区-秉正小学",
		},
	}
	//   fmt.Println(p1)
	p1.shoolsAddr()

	p2 := newStudent("lufei", 19)
	fmt.Println(p2.Name)
	// 匿名结构体
	d1 := dog{
		"阿黄",
	}
	fmt.Println(d1.string)

	d2 := struct { // 少了声明类型，直接调用时候声明
		name string
		age  int
	}{"阿黄", 3} // 直接赋值，这里注意用花括号，跟函数自动调用有区分

	fmt.Println(d2)

	// 序列化与反序列化

	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("报错了")
	}
	str := string(data)
	fmt.Printf("%v\n", str)

	var p3 studens
	json.Unmarshal([]byte(str), &p3)
	fmt.Println(p3)
	p3.shoolsAddr()

}
