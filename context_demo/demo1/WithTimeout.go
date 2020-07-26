package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var ws sync.WaitGroup

func f3(ctx context.Context)  {
Loop:
	for  {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done():
			break Loop
		default:

		}
	}
	fmt.Println("worker done!")
	ws.Done()
}

func main()  {
	//  这个是 50 毫秒后，关闭资源， WithDeadline 是某个节点 关闭资源
	ctx,cannel := context.WithTimeout(context.Background(),time.Millisecond * 50)
	ws.Add(1)
	go f3(ctx)
	time.Sleep(3)
	cannel()
	ws.Wait()
	fmt.Println("over")
}
