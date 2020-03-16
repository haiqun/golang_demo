package session

// session 管理者

type session_mgr interface {
	// 初始化session用的，可能是redis可能是内存版本的
	Init(address string,option ...interface{})error
	// 创建一个session的话柄 - 用来操作 session类的
	CreateSession()(Session error)
	// 通过这个sessionId 来反查session是否存在
	GetSession(sessionId string)(Session,error)
}
