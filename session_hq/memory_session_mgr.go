package session

import "sync"
import uuid "github.com/satori/go.uuid"

// 定义一个对象
type MemorySessionMgr struct {
	sessionMap map[string]*MemorySession // 存来存放 session的索引
	rwlook sync.RWMutex // 读写锁
}

// 构造对象 
func NewMemorySessionMgr() *MemorySessionMgr  {
	return &MemorySessionMgr{
		sessionMap:make(map[string]*MemorySession,1024), // 每个管理者，管理1024的session对象
	}
}

// 内存版本的，不用连接；redis版的需要
func (m *MemorySessionMgr)Init(address string, potions ...string)(err error)  {
	return
}

// 创建一个session对象
func (m *MemorySessionMgr)CreateSession() (s Session,err error) {
	m.rwlook.Lock()
	defer m.rwlook.Unlock()
	id := uuid.NewV4()
	sessionId := id.String()
	s = NewMemorySession(sessionId)
	return
}