package session

import "sync"
import "errors"

// 内存版本的session

// 内存版本的结构体
type MemorySession struct {
	// 通过id找到信息对象
	sessionId string
	// 存储key - value
	data map[string]interface{}
	// 读写锁
	rwlook sync.RWMutex
}

// 构造函数
func NewMemorySession(id string) * MemorySession  {
	return &MemorySession{
		sessionId:id,
		data: make(map[string]interface{}, 16),
	}
}


func (ms *MemorySession)Set(key string,value interface{})(err error) {
	// 加锁
	ms.rwlook.Lock()
	defer ms.rwlook.Unlock()
	ms.data[key] = value
	return
}

func (ms *MemorySession)Get(key string) (value interface{} ,err error) {
	// 加锁 - 读
	ms.rwlook.RLocker()
	defer ms.rwlook.RUnlock()
	value ,ok := ms.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

func (ms *MemorySession)Del(key string)(err error)  {
	// 加锁
	ms.rwlook.Lock()
	defer ms.rwlook.Unlock()
	delete(ms.data,key)
	return
}

// 这个是兼容redis版本的，这里什么都不做
func (ms *MemorySession)Save()(err error)  {
	return
}