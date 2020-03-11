package main

import (
	"context"
	"fmt"
	"time"
)

// WithDeadline demo
// func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
// 返回值是ctx，跟cancle函数，手动调用结束

func main() {
	// 获取5秒的时间 类型是 time.Time

	d := time.Now().Add(5 * time.Second)
	// fmt.Printf("%T", d)
	ctx, cancle := context.WithDeadline(context.Background(), d)
	// 虽然支持过期自动关闭，建议手动关闭
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间
	defer cancle()
	select {
	// 1 秒之后柱塞
	/*
		time.After 设置一个时长，这个时间间隔过了之后，马上就阻塞了
		类型这里，设置了一秒后自动阻塞
	*/
	case <-time.After(1 * time.Second):
		fmt.Println("tesT")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
