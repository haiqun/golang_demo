package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

func main() {

	// Iterating objects
	jsonParsed, err := gabs.ParseJSON([]byte(`{"object":{"first":1,"second":2,"third":3}}`))
	if err != nil {
		panic(err)
	}

	// S is shorthand for Search
	for key, child := range jsonParsed.S("object").ChildrenMap() {
		fmt.Printf("key: %v, value: %v \n", key, child.Data())
		/**
		key: first, value: 1
		key: second, value: 2
		key: third, value: 3
		*/
	}

	// Iterating arrays
	jsonParsed1, err := gabs.ParseJSON([]byte(`{"array":["first","second","third"]}`))
	if err != nil {
		panic(err)
	}

	for key, child := range jsonParsed1.S("array").Children() {
		fmt.Printf("key: %v, value: %v \n", key, child.Data())
	}

	// Searching through arrays
	jsonParsed2, err := gabs.ParseJSON([]byte(`{"array":[{"value":1},{"value":2},{"value":3}]}`))
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonParsed2.Path("array.1.value").String())

}
