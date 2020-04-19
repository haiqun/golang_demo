package main

import "fmt"

func main() {
	ch1  := make(chan interface{})
	ch1 <- nil
	fmt.Println(123)
	close(ch1)
	for i:=10;i>0;i--  {
		a := <-ch1
		fmt.Println(a)
	}
}