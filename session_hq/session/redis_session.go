package session

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis" // 这个redis的版本跟之前的不一样 ？
	"sync"
	"errors"
)

// redis版本的结构体
type RedisSession struct {
	// 通过id找到信息对象
	sessionId string
	// 链接池
	pool *redis.Pool
	// 存储key - value => 先存session再存redis
	sessionMap map[string]interface{}
	// 读写锁
	rwlook sync.RWMutex
	// 记录内存的数据是否被操作
	flag int
}

const (
	// 标识是否被处理
	SessionFlogNone = iota
	// 未处理【有数据】
	SessionFlogModify
)

// 构造函数
func NewRedisSession(id string,pool *redis.Pool) ( rs * RedisSession ) {
	return &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}),
		flag:       SessionFlogNone,
	}
}


func ( rs * RedisSession )Set(key string,value interface{})(err error) {
	// 加锁
	rs.rwlook.Lock()
	defer rs.rwlook.Unlock()
	rs.sessionMap[key] = value
	// 记录一个记录
	rs.flag = SessionFlogModify
	return
}

func ( rs * RedisSession )Get(key string) (value interface{} ,err error) {
	// 加锁 - 读
	rs.rwlook.RLocker()
	defer rs.rwlook.RUnlock()
	// 在内存取一次
	value ,ok := rs.sessionMap[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

func (rs * RedisSession)loadFromRedis()  {
	// 获取一个链接
	conn := rs.pool.Get()
	// 通过sessionid获取数据
	reply, err := conn.Do("GET", rs.sessionId)
	if err != nil {
		return
	}
	// 转字符串
	s, err := redis.String(reply, err)
	if err != nil {
		return
	}
	// 反序列化
	err = json.Unmarshal([]byte(s), &rs.sessionMap)
	if err != nil {
		return
	}
	return
}


func ( rs * RedisSession )Del(key string)(err error)  {
	// 加锁
	rs.rwlook.Lock()
	defer rs.rwlook.Unlock()
	// 删除内存版本的数据
	rs.flag = SessionFlogModify
	delete(rs.sessionMap,key)
	return
}

// redis版本在调用这个函数之后将内存数据存入redis
func ( rs * RedisSession )Save()(err error)  {
	// 锁定数据
	rs.rwlook.Lock()
	defer rs.rwlook.Unlock()
	// 数据没有变更
	if rs.flag != SessionFlogModify {
		return
	}
	// 对内存中的数据进行格式化
	data,err := json.Marshal(rs.sessionMap)
	if err != nil {
		return
	}
	// 获取redis 的链接，存入data
	conn := rs.pool.Get()
	_,err = conn.Do("SET",rs.sessionId,string(data))
	if err != nil {
		return
	}
	// 置空flag
	rs.flag = SessionFlogNone
	return
}
