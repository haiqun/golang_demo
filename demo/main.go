package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {

	quit := make(chan os.Signal)
	t := time.NewTicker(time.Second * 3) // 每3秒钟执行一次
	go func() {
		for  {
			select {
			case <-t.C:
				fmt.Println(time.Now().Format("15:04:05"))
			}
		}
	}()

	<-quit
}