package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v ; kind:%v\n", v.Name(), v.Kind())
}

// 反射
func main() {
	str := `{"name":"test","age":18}`
	var s person
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	// fmt.Println([]byte(str))

	var a int64 = 123
	var f float32 = 1.23403
	reflectType(a)
	reflectType(f)

	// 判断值是否为空

	// 判断值是否有效、

	// 结构挺反射

}
