package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	d := time.Now().Add(time.Millisecond * 50) // 50 毫秒
	//  这个是 d 的时间点，关闭资源， WithTimeout 是多少时间后 关闭资源
	ctx,cancel := context.WithDeadline(context.Background(),d)

	defer cancel()

	select {
	case <-time.After(time.Second): // 1s后回调
		fmt.Println("test1")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 50毫秒后回调 ，所以就先关闭了
	}

	fmt.Println(d)
}