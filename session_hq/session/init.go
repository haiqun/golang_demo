package session

import "fmt"

// 中间件让用户选择去使用哪个版本
var sessionMgr session_mgr

func Init(provider string,addr string,options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Println("不支持")
		return
	}
	err = sessionMgr.Init(addr, options...)
	return

}

