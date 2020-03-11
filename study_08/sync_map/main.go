package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map
 
var m = make(map[string]int)
var look sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	//  不加锁会报错
	//  fatal error: concurrent map writes
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(x int) {
			key := strconv.Itoa(x)
			// 加互斥锁可以解决
			look.Lock()
			set(key, x)
			look.Unlock()
			fmt.Printf("key : %s , value : %v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
