package session

// session 类

type Session interface {
	// 将值保存起来
	Set(key string ,value interface{}) error
	// 通过key获取值
	Get(key string)(interface{},error)
	// 删除一个值
	Del(key string) error
	// ？
	Save() error
}
