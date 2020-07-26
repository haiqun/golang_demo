package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
)
/**
	合并json
 */
func main() {
	jsonParsed1, _ := gabs.ParseJSON([]byte(`{"outter":{"inner":"one"}}`))
	jsonParsed2, _ := gabs.ParseJSON([]byte(`{"outter":{"inner":{"value3":"three"}},"outter2":{"value2":"two"}}`))

	jsonParsed1.Merge(jsonParsed2)

	fmt.Println(jsonParsed1)

	// 合并过程举例1
	// {"outter":{"value1":"one"}}
	// {"outter":{"inner":{"value3":"three"}},"outter2":{"value2":"two"}}
	// {"outter":{"inner":{"value3":"three"},"value1":"one"},"outter2":{"value2":"two"}}
	// 合并过程举例2
	// {"outter":{"inner":"one"}}
	// {"outter":{"inner":{"value3":"three"}},"outter2":{"value2":"two"}}
	// {"outter":{"inner":["one",{"value3":"three"}]},"outter2":{"value2":"two"}}


	sample := []byte(`{"test":{"int":10,"float":6.66}}`)
	dec := json.NewDecoder(bytes.NewReader(sample))
	dec.UseNumber()

	val, err := gabs.ParseJSONDecoder(dec)
	if err != nil {
		fmt.Printf("Failed to parse: %v", err)
		return
	}
	//intValue := val.Path("test.int").Data() // "10"
	//intValue := val.Path("test.int").Data().(json.Number) // "10"
	intValue, err := val.Path("test.int").Data().(json.Number).Int64() // 10
	fmt.Printf("%#v \n",intValue)
}
