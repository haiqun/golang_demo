package main

import (
	"context"
	"fmt"
	"time"
)

// 类型别名
// 作为withvalue传递的参数

type TraceCode string

func f1(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	// 在子goroutine中获取trace code 的值
	traceCode, ok := ctx.Value(key).(string) //.(type) 类型断言
	if !ok {
		fmt.Println("ctx value err")
		return
	}
loop:
	for {
		fmt.Println(traceCode)
		time.Sleep(time.Millisecond * 10)
		// 当我在chan能取到值的时候就退出
		select {
		// 返回一直只读的通道，输出，获取数据
		case <-ctx.Done():
			break loop
		default:
		}
	}
	fmt.Println("f1 done ")
}

func main() {
	// 设置一个50秒超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	str := fmt.Sprintf("我是WithValue开始的时间-%d", time.Now().Unix())
	// 这里记得重新赋值 ctx
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), str)
	go f1(ctx)
	time.Sleep(time.Second * 5)
	fmt.Println("over")
}
