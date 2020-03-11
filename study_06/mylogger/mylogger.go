package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

const logFile = "xx.log"

type LogLevel uint16

type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func paramLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getFileInfo(n int) (funcName, filePath string, fileLine int) {
	pc, file, fileLine, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller  报错")
	}
	funcName = runtime.FuncForPC(pc).Name() // 方法名称
	filePath = path.Base(file)              // 调用路径
	funcName = strings.Split(funcName, ".")[1]
	return
}
