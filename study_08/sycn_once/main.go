package main

import (
	"fmt"
	"sync"
)

/*
并发安全的单例模式

sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，
而布尔值用来记录初始化是否完成。
这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次
*/
type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("23")
		instance = &singleton{}
	})
	return instance
}

func main() {
	var s1 *singleton
	for i := 0; i < 100; i++ {
		s1 = GetInstance()
	}
	fmt.Println(s1)
}
