package main

import (
	"context"
	"fmt"
	"time"
)

// 通过context 埋一个根节点
// 统一管理超时问题
// 涉及多个gorontine的操作
// 场景  ： http_server 入口，一个请求一个gorontine，后续的一系列的新开的goroutine操作，
// 如果请求超时，那么同一根节点的都要一起推出

/*
WithCancel 开启一个根节点 通过cancel取消
WithDeadline 设置一个时间点，到了自动关闭子goroutine
WithTimeout 设置一个超时的时长，过了时长关闭 子goroutine
*/

func f1(ctx context.Context) {
	go f2(ctx)
	loop:
	for {
		fmt.Println("TEst")
		time.Sleep(time.Millisecond * 500)
		// 当我在chan能取到值的时候就退出
		select {
		// 返回一直只读的通道，输出，获取数据
		case <-ctx.Done():
			break loop
		default:
		}
	}
}

func f2(ctx context.Context) {
	for {
		fmt.Println("TEst-kkk")
		time.Sleep(time.Millisecond * 1000)
		// 当我在chan能取到值的时候就退出
		select {
		// 返回一直只读的通道，输出，获取数据
		case <-ctx.Done():
			break
		default:
		}
	}
}

// context 解决
func main() {
	// 创建根节点的ctx，统一管理整个请求链路
	ctx, cancel := context.WithCancel(context.Background())
	go f1(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
}
