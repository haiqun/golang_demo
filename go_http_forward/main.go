package main

import "fmt"

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

func main()  {
	findType(10)      // int
	findType("hello") // string
	var k interface{} // nil
	findType(k)
	findType(10.23) //float64
}
