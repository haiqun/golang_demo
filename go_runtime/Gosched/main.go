package  main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务 - 结果是：先执行协程，再执行主进程
		// 如果这里不使用切换，结果是：主进程跟协程并行执行
		runtime.Gosched() // 让出CPU时间片，重新等待安排任务 [等等的意思]
		fmt.Println("hello")
	}
}