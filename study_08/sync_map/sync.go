package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map

func main() {
	wg := sync.WaitGroup{}
	m := sync.Map{}

	//  不加锁会报错
	//  fatal error: concurrent map writes
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(x int) {
			key := strconv.Itoa(i)
			m.Store(key, i)
			v, _ := m.Load(key)
			fmt.Printf("key : %s , value : %d \n", key, v)
			wg.Done()
		}(i)
	}
	// 变量
	wg.Wait()

	//LoadOrStore
	//若key不存在，则存入key和value，返回false和输入的value
	m1 := sync.Map{}
	v, ok := m1.LoadOrStore("1", "aaa")
	fmt.Println(ok, v) //false aaa
	//若key已存在，则返回true和key对应的value，不会修改原来的value
	v, ok = m1.LoadOrStore(1, "aaab")
	fmt.Println(ok, v) //false

	m2 := sync.Map{}
	f := func(k, v interface{}) bool {
		m2.LoadOrStore(k, v)
		return true
	}
	m1.Range(f)
	v, _ = m2.Load("1")
	fmt.Printf("%T , %#v\n", v, v)
}
