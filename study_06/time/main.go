package main

import (
	"fmt"
	"time"
)

// 时间
func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Weekday())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	// 纳秒
	fmt.Println(now.UnixNano())
	// 时间戳转换
	ret := time.Unix(now.Unix(), 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	fmt.Println(ret.Weekday())
	// 时间间隔
	fmt.Println(time.Second)
	fmt.Println(time.Hour)
	later := now.Add(time.Hour)
	fmt.Println(later)
	

	// 定时器 - 不要在直接执行，停止不了，要用go run
	// timer := time.Tick(time.Second) // 这里设置时间间隔
	// for t := range timer {
	// 	fmt.Println(t) // 每秒钟执行一次
	// }
	// 格式化时间
	//2018-09-03
	fmt.Println(now.Format("2006-01-02"))
	//2018/09/3 2:03:01
	fmt.Println(now.Format("2006-01-02 3:04:05"))
	//2018/09/3 2:03:01 AM
	fmt.Println(now.Format("2006/01/2 03:04:05 AM"))
	// 2019/09/3 2:29:01.283
	fmt.Println(now.Format("2006/01/2 15:04:05.000"))
	// 字符串转时间戳
	str, _ := time.Parse("2006-01-02", now.Format("2006-01-2"))
	fmt.Println(str)
}
