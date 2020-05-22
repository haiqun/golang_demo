package demo1
// 阻塞行 限流
import (
	//"log"
	"time"
)

// 采用channel作为底层数据结构，从而达到阻塞的获取和使用信号量
type Semaphore struct {
	innerChan chan struct{}
}

// 初始化信号量，本质初始化一个channel，channel的初始化大小为 信号量数值
func NewSemaphore(num uint64) *Semaphore {
	return &Semaphore{
		innerChan: make(chan struct{}, num),
	}
}
// 获取信号量，本质是 向channel放入元素，如果同时有很多协程并发获取信号量，则channel则会full阻塞，从而达到控制并发协程数的目的，也即是信号量的控制
func (s *Semaphore) Acquire() {
	for {
		select {
		case s.innerChan <- struct{}{}:
			return
		default:
			//log.Error("semaphore acquire is blocking")
			//log.Println("semaphore acquire is blocking")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
// 释放信号量 本质是 从channel中获取元素，由于有acquire的放入元素，所以此处一定能回去到元素 也就能释放成功，default只要是出于安全编程的目的
func (s *Semaphore) Release() {
	select {
	case <-s.innerChan:
		return
	default:
		return
	}
}
