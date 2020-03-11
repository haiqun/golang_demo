package main

import (
	"encoding/json"
	"fmt"
)

// json 与结构体的转换
type person struct {
	Name string `json:"name"` // 这里的变量没有大写，转换之后，json数据为空
	Age  int    `json:"age"`  // 当你用json去解析的时候，他的名字叫 age
}

func main() {
	p1 := person{
		Name: "lufei",
		Age:  18,
	}
	// fmt.Println(p1)
	// 序列化
	data, err := json.Marshal(p1) // data 获取到的是byte类型
	if err != nil {
		fmt.Printf("报错了 %s", err)
	}
	fmt.Printf("%v\n", string(data)) // 所以需要转为string
	// 反序列化
	str := `{"name":"suoluo","age":19}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //支持byte的slice所以需要装换
	fmt.Printf("%#v\n", p2)
}
