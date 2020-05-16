package demo2
// 非阻塞行 限流
import (
	"sync"
	"sync/atomic"
)

// 信号量结构
type s struct {
	nMax uint64
	nCur uint64
}

var look sync.Mutex

// 初始化信号量，本质是设置一个最大值，超过这个值直接返回false
func NewSemaphore(num uint64) *s {
	return &s{
		nMax:num,
		nCur:0,
	}
}
// 获取信号量，本质上是加一个数
func (s *s) Acquire(){
	if s.nCur < s.nMax {
		atomic.AddUint64(&s.nCur,1)
		return
	}
	panic("满了")
}

// 释放信号量, 本质是减一个数
func (s *s) Release(){
	if s.nCur > 0 {
		// ^ 表示减一
		atomic.AddUint64(&s.nCur,^uint64(1))
		return
	}
	return
}
