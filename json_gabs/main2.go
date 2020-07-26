package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

func main() {
	// json 格式化
	jsonParsedObj, _ := gabs.ParseJSON([]byte(`{
	"outter":{
		"values":{
			"first":10,
			"second":11
		}
	},
	"outter2":"hello world"
}`))

	jsonOutput := jsonParsedObj.String()

	fmt.Println(jsonOutput)

	jsonOutput1 := jsonParsedObj.Search("outter").String() // {"outter":{"values":{"first":10,"second":11}},"outter2":"hello world"}

	fmt.Println(jsonOutput1) // {"values":{"first":10,"second":11}}



}
