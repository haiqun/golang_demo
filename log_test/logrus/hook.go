package main

// hook的原理是，在logrus写入日志时拦截，修改logrus.Entry
type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}