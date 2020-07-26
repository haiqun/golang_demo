package main
/**
	无需写对应的结构体，直接获取结构体的值的方式
 */
import (
	"github.com/Jeffail/gabs/v2"
	"log"
)

func main() {
	jsonParsed, err := gabs.ParseJSON([]byte(`{
		"outter":{
		"inner":{
			"value1":10,
			"value2":22,
			"value":0
		},
		"alsoInner":{
			"value1":20,
			"array1":[
				30, 40
			]
		}
	}
	}`))

	if err != nil {
		panic(err)
	}

	var value float64
	var ok bool

	value, ok = jsonParsed.Path("outter.inner.value1").Data().(float64)
	// value == 10.0, ok == true

	value, ok = jsonParsed.Search("outter", "inner", "value3").Data().(float64)
	// value == 0, ok == false
	log.Printf("%#v", value)
	log.Printf("%#v", ok)
	value, ok = jsonParsed.Search("outter", "alsoInner", "array1", "0").Data().(float64)
	// value == 40.0, ok == true
	log.Printf("%v", value)
	log.Printf("%#v", ok)

	// 获取对象 再获取值
	gObj, err := jsonParsed.JSONPointer("/outter/alsoInner/array1/1")
	if err != nil {
		panic(err)
	}
	value, ok = gObj.Data().(float64)
	// value == 40.0, ok == true
	log.Printf("value == %v , , ok == %v \n", value, ok)

	// 判断json数据结构的v字段是否存在
	exists := jsonParsed.Exists("outter", "inner", "value")
	// exists == true
	log.Printf("%#v", exists)

	exists = jsonParsed.ExistsP("does.not.exist")
	// exists == false
	log.Printf("%#v", exists)
}
