package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

// Generating JSON 生成json

func main()  {
	jsonObj := gabs.New()
	// or gabs.Wrap(jsonObject) to work on an existing map[string]interface{}

	jsonObj.Set(10, "outter", "inner", "value")
	jsonObj.SetP(20, "outter.inner.value2") // 这里其实就是写法不一样
	jsonObj.Set(30, "outter", "inner2", "value3")

	fmt.Println(jsonObj)
	// 优雅的打印数据
	fmt.Println(jsonObj.StringIndent("", "  "))
	fmt.Println(jsonObj.String())


	// 生成array
	jsonObj1 := gabs.New()

	jsonObj1.Array("foo", "array")
	// Or .ArrayP("foo.array")

	jsonObj1.ArrayAppend(10, "foo", "array")
	jsonObj1.ArrayAppend(20, "foo", "array")
	jsonObj1.ArrayAppend(30, "foo", "array")

	fmt.Println(jsonObj1.String()) // {"foo":{"array":[10,20,30]}}

	// Working with arrays by index:
	jsonObj2 := gabs.New()

	// Create an array with the length of 3
	jsonObj2.ArrayOfSize(3, "foo")

	jsonObj2.S("foo").SetIndex("test1", 2)
	jsonObj2.S("foo").SetIndex("test2", 1)
	fmt.Println(jsonObj2.String())

	// Create an embedded array with the length of 3
	jsonObj2.S("foo").ArrayOfSizeI(3, 2)
	jsonObj2.S("foo").ArrayOfSizeI(3, 0)

	jsonObj2.S("foo").Index(2).SetIndex(1, 0)
	jsonObj2.S("foo").Index(2).SetIndex(2, 1)
	jsonObj2.S("foo").Index(2).SetIndex(3, 2)

	fmt.Println(jsonObj2.String())

}