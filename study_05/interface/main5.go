package main

import "fmt"
import "encoding/json"

// 空接口作为map的值
// 接口断言
func main() {

	var info = make(map[int]map[string]interface{}, 20)
	info[1] = make(map[string]interface{}, 4)
	info[1]["Name"] = "lufei"
	info[1]["Age"] = 19
	info[1]["Gender"] = "男"
	info[1]["Marred"] = false
	info[1]["Hobby"] = []string{
		"游泳", "篮球", "足球",
	}
	// info["B"] = make(map[string]interface{}, 4)
	// info["B"]["Name"] = "songluo"
	// info["B"]["Age"] = 19
	// info["B"]["Gender"] = "男"
	// info["B"]["Marred"] = false
	// info["B"]["Hobby"] = []string{
	// 	"羽毛球", "跳水", "剑",
	// }

	fmt.Println(info)
	// for _,v := range info {
	// 	fmt.Printf("%T 对应 %#v\n",v,v)
	// }
	// 格式化 json
	data, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("格式化字符串有误：%s\n", err)
	}
	// fmt.Printf("%s", data)
	// 解析 json
	var json_decode_info map[int]map[string]interface{}
	ret := json.Unmarshal([]byte(data), &json_decode_info)
	if ret != nil {
		fmt.Printf("解析json格式化有误 ：%s \n", ret)
	}
	fmt.Println(json_decode_info)

	// 接口断言
	x := info[1]["Age"]
	// fmt.Printf("%v\n", x)
	switch v := x.(type) {
	case string:
		fmt.Printf("%#v is %T", v, v)
	case int:
		fmt.Printf("%#v is %T", v, v)
	}
}
