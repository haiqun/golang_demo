package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 例子链接：https://www.jianshu.com/p/228c119a7d0e
// 教材 https://www.liwenzhou.com/posts/Go/14_concurrence/

// 原子操作 atomic版
/*
代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。
针对基本数据类型我们还可以使用原子操作来保证并发安全，
因为原子操作是Go语言提供的方法它在用户态就可以完成，
因此性能比加锁操作更好。
Go语言中原子操作由内置的标准库sync/atomic提供
*/

var x int64
var wg sync.WaitGroup
var look sync.Mutex

// 普通版
// func add() {
// 	defer wg.Done()
// 	look.Lock()
// 	x++
// 	look.Unlock()
// }

//atomic版
func add() {
	defer wg.Done()
	atomic.AddInt64(&x, 1)
}

func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)

	// 比较并交换
	x1 := int64(102)
	// 栗子：（如果addr和old相同,就用new代替addr）
	ret := atomic.CompareAndSwapInt64(&x1, 102, 300)
	fmt.Println(x1, ret)
}
