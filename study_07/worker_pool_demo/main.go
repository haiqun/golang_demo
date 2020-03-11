package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主goroutine从resultChan取出结果并打印到终端输出

type job struct {
	value int64
}

type ret struct {
	job *job
	sum int64
}

var wg sync.WaitGroup

func f1(jobChan chan<- *job) {
	defer wg.Done()
	for {
		r1 := rand.Int63() // 生成一个int64的随机数
		newJob := &job{
			value: r1,
		}
		// fmt.Println(newJob)
		jobChan <- newJob
		time.Sleep(time.Millisecond * 300) // 300 毫秒
	}
}

func f2(jobChan <-chan *job, resultChan chan<- *ret) {
	defer wg.Done()
	for {
		job := <-jobChan
		ret := &ret{
			job: job,
			sum: sun(job.value),
		}
		resultChan <- ret
	}
}

func sun(x int64) int64 {
	n := int64(0)
	for x > 0 {
		x = x / 10
		n += x % 10
	}
	return n
}

func main() {
	// channal 类型的值可以是任意类型
	jobChan := make(chan *job, 100)
	resultChan := make(chan *ret, 100)
	// 生成100个随机数
	wg.Add(1)
	go f1(jobChan)
	// // 开 24 个消费进程 这里用for去做
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go f2(jobChan, resultChan)
	}
	for x := range resultChan {
		fmt.Printf("value:%d sum:%d\n", x.job.value, x.sum)
	}
	wg.Wait() // 这个是一直消费的经常，所以不会有停止的时候 ，这句话后面的不会被执行
	fmt.Println("这里不会被执行")
}
