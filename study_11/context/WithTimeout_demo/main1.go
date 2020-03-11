package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// WithTimeout
// 通常用于数据库或者网络连接的超时控制

var wg sync.WaitGroup

func f1(ctx context.Context) {
	// defer wg.Done()
LOOP:
	for {
		fmt.Println("TEst")
		time.Sleep(time.Millisecond * 20)
		// 当我在chan能取到值的时候就退出
		select {
		case <-ctx.Done(): // 50毫秒后自动调用，所以后面就不执行了
			break LOOP
		default:
		}
	}
	fmt.Println("f1 over ")
}

func main() {
	// 设置50毫秒过期
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()
	// wg.Add(1)
	go f1(ctx)
	time.Sleep(time.Second * 1)
	// wg.Wait()
	fmt.Println("over")
}
