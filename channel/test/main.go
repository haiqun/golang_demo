package main

import (
	"fmt"
	"time"
)

var ch8 = make(chan int, 6)

func mm1() {
	for i := 0; i < 10; i++ {
		ch8 <- 8 * i
	}
	close(ch8)
	fmt.Println("kaishi ",time.Now().Format("05.0000"))
}
func main() {

	go mm1()
	fmt.Println("kaisi2",time.Now().Format("05.0000"))
	for {
		for data := range ch8 {
			fmt.Print(data, "\t")
			fmt.Println(time.Now().Format("05.0000"))
		}
		break
	}

}