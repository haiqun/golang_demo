package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


// 自定义类型or类型别名
type TraceCode string
var wg sync.WaitGroup

func f4(ctx context.Context)  {
	defer wg.Done()
	key := TraceCode("test")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
}

func main()  {
	ctx,cannel := context.WithTimeout(context.Background(),time.Millisecond*50)
	defer cannel()
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx =  context.WithValue(ctx,TraceCode("test"),"kkkkk")
	wg.Add(1)
	go f4(ctx)
	wg.Wait()
}