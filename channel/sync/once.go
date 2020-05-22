package main

import (
	"fmt"
	"sync"
	"time"
)

type Once chan struct{}

func NewOnce() Once {
	o := make(Once, 1)
	// 只允许一个goroutine接收，其他goroutine会被阻塞住
	o <- struct{}{}
	return o
}

func (o Once) Do(f func()) {
	_, ok := <-o
	if !ok {
		// Channel已经被关闭
		// 证明f已经被执行过了，直接return.
		return
	}
	// 调用f, 因为channel中只有一个值
	// 所以只有一个goroutine会到达这里
	f()

	// 关闭通道，这将释放所有在等待的
	// 以及未来会调用Do方法的goroutine
	close(o)
}

func f1()  {
	fmt.Println("f1111")
}

func f2()  {
	fmt.Println("f222")
}

func main()  {
	// 自己实现的 sync.once
	o := NewOnce()
	for i:=0;i<10;i++ {
		go o.Do(f1)
	}
	time.Sleep(time.Second*1)
	// 自带包
	o1 := sync.Once{}
	for i:=0;i<10;i++ {
		go o1.Do(f2)
	}
	time.Sleep(time.Second*1)
	fmt.Println("end")
}
