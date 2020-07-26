package main

import (
	"context"
	"fmt"
)

func f1(ctx context.Context) <-chan int {
	dns := make(chan int)
	n := 1
	go func() {
		for  {
			select {
			case <-ctx.Done():
				return
			case dns<-n:
				n++
			}
		}
	}()
	return  dns
}

func main()  {
	//ctx,cannel := context.WithCancel(context.Background())
	ctx,cannel := context.WithCancel(context.TODO())

	defer cannel() // defer 调用取消了子进程

	for x := range f1(ctx){
		if x == 50 {
			break // 这里跳出 ，main函数结束 调用到defer
		}
		fmt.Println(x)
	}

}