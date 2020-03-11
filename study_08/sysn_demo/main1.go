package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var (
	wg     sync.WaitGroup
	look   sync.Mutex
	relook sync.RWMutex
	x      int64
)

func read() {
	defer wg.Done()
	look.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 1)
	look.Unlock()
}

func read1() {
	defer wg.Done()
	// 读写锁
	relook.RLock()
	fmt.Println(x)
	// time.Sleep(time.Millisecond * 1)
	// 读写锁
	relook.RUnlock()
}

func write() {
	defer wg.Done()
	look.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	look.Unlock()
}

func write1() {
	defer wg.Done()
	relook.Lock()
	x = x + 1
	// time.Sleep(time.Millisecond * 5)
	relook.Unlock()
}

func main() {
	start := time.Now()
	// 互斥锁
	// for i := 0; i < 10; i++ {
	// 	go write()
	// 	wg.Add(1)
	// }
	// for i := 0; i < 1000; i++ {
	// 	go read()
	// 	wg.Add(1)
	// }
	// 读写锁
	for i := 0; i < 100; i++ {
		go write1() // 不休眠
		wg.Add(1)
	}
	// 当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待
	// 当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁
	// 写的速度太慢了，又或者读的太快了，所以需要给写时间 ？
	// time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		go read1() // 不休眠
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
