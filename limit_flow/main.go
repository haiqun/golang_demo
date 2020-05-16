package main

import (
	"fmt"
	"golang_demo/limit_flow/demo1"
	"time"
)

func f1(s *demo1.Semaphore)  {
	i := 0
	for {
		if i == 12 {
			s.Release()
			fmt.Println("0 | 释放...")
		}
		i++
		time.Sleep(time.Millisecond * 500)
	}
}

func main()  {
	// 柱塞行-限流demo1
	//s := demo1.NewSemaphore(10)
	//i := 0
	//go f1(s)
	//for  {
	//	i++
	//	time.Sleep(time.Millisecond * 500)
	//	s.Acquire();
	//	fmt.Println(i," | 获取...")
	//}
	// 非阻塞-限流demo2
	//s1 := demo2.NewSemaphore(20)
	//var ws sync.WaitGroup
	//for k:=0;k<20;k++  {
	//	ws.Add(1)
	//	go func() {
	//		for q1 := 0;q1 < 5;q1++{
	//			time.Sleep(time.Millisecond * 5)
	//			s1.Acquire();
	//			fmt.Println("获取...")
	//		}
	//		ws.Done()
	//	}()
	//}
	//ws.Wait()


}
