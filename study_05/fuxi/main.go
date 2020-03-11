package main

import "fmt"


func main() {
	p1 := struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}{"test",19}
	
	fmt.Println(p1)

	b1,err := json.Marshal(p1)

	if err !=nil {
		fmt.Println("解析报错")
	} 

	fmt.Println(b1 )
	
}