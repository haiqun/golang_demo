package main

import (
	"fmt"
	"sync"
)

type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Lock() {
	// 只有在s还有空间的时候才能发送成功
	s <- struct{}{}
}

func (s Semaphore) Unlock() {
	// 为其他信号量腾出空间
	<-s
}

var s1 sync.Mutex

var x int

func f3(i int)  {
	s1.Lock()
	x += i
	s1.Unlock()
	ws.Done()
}

var s2 = NewSemaphore(1)

func f4(i int)  {
	s2.Lock()
	x += i
	s2.Unlock()
	ws.Done()
}

var ws = sync.WaitGroup{}

func main()  {
	for j:=0;j<1000;j++ {
		ws.Add(1)
		go f3(j)
	}
	ws.Wait()
	fmt.Println(x)

	x = 0
	for j:=0;j<1000;j++ {
		ws.Add(1)
		go f4(j)
	}
	ws.Wait()
	fmt.Println(x)
}