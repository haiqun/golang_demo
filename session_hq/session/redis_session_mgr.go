package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	// 地址
	addr string
	// 密码
	passwd string
	// 链接池
	pool *redis.Pool
	// 锁
	rwlook sync.RWMutex
	// map
	sessionMap map[string]Session
	
}

func myPool(addr , passwd string) *redis.Pool  {
	return &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", addr)
			if err != nil {
				return
			}
			// 判断是否有密码
			_, err = conn.Do("AUTH", passwd)
			if err != nil {
				return
			}
			return
		},
		// / TestOnBorrow()方法是一个可选项，该方法用来诊断一个连接的健康状态
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return  err
			}
			return nil
		},
		MaxIdle:      64, // 最大空闲
		MaxActive:    1000, // 最大连接数
		IdleTimeout:  240*time.Second, // 空闲连接超时时间，超过超时时间的空闲连接会被关闭
		Wait:         false,// 如果Wait被设置成true，则Get()方法将会阻塞
	}
}

// 实例化
func (rm *RedisSessionMgr)Init(address string, option ...string) error  {
	if len(option) > 0 {
		rm.passwd = option[0]
	}
	rm.addr = address
	// 创建连接池
	myPool(rm.addr, rm.passwd)
	return nil
}


func NewRedisSessionMgr() (rm * RedisSessionMgr)  {
	return &RedisSessionMgr{
		sessionMap: make(map[string]Session,32),
	}
}

// 创建一个session对象
func (rm *RedisSessionMgr)CreateSession() (s Session,err error) {
	rm.rwlook.Lock()
	defer rm.rwlook.Unlock()
	// 生成一个sessionid
	id := uuid.NewV4()
	sessionId := id.String()
	// 返回一个session对象
	s = NewRedisSession(sessionId,rm.pool)
	// 加入map记录
	rm.sessionMap[sessionId] = s
	return
}

// 获取一个session对象
func (rm *RedisSessionMgr)Get(sessionId string)(s Session,err error)  {
	rm.rwlook.Lock()
	defer rm.rwlook.Unlock()
	s,ok := rm.sessionMap[sessionId]
	if !ok {
		err =  errors.New("key not exists in session")
		return
	}
	return
}