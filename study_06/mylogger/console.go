package mylogger

import (
	"fmt"
	"time"
)

type consoloLogger struct {
	level LogLevel
}

// NewconsoloLogger console输出的日志
func NewconsoloLogger(s string) consoloLogger {
	llevel, err := paramLevel(s)
	if err != nil {
		panic(err)
	}
	return consoloLogger{
		level: llevel,
	}
}

func (c consoloLogger) getOutputMsgData(level string, format string, a ...interface{}) {
	// 构造输出日志的结构;
	msg := fmt.Sprintf(format, a...)
	t := time.Now().Format("2006-01-02 15:04:05.000")
	funcName, filePath, fileLine := getFileInfo(3)
	fmt.Printf("%s [%s] [%s:%s:%d] %s \n", t, level, filePath, funcName, fileLine, msg)
}

func (c consoloLogger) levelJudgment(L LogLevel) bool {
	return L < c.level
}

func (c consoloLogger) Debug(format string, a ...interface{}) {
	if c.levelJudgment(DEBUG) {
		return
	}
	c.getOutputMsgData("DEBUG", format, a...)
}

func (c consoloLogger) Trace(format string, a ...interface{}) {
	if c.levelJudgment(TRACE) {
		return
	}
	c.getOutputMsgData("TRACE", format, a...)

}

func (c consoloLogger) Info(format string, a ...interface{}) {
	if c.levelJudgment(INFO) {
		return
	}
	c.getOutputMsgData("INFO", format, a...)
}

func (c consoloLogger) Warning(format string, a ...interface{}) {
	if c.levelJudgment(WARNING) {
		return
	}
	c.getOutputMsgData("WARNING", format, a...)

}

func (c consoloLogger) Error(format string, a ...interface{}) {
	if c.levelJudgment(ERROR) {
		return
	}
	c.getOutputMsgData("ERROR", format, a...)
}

func (c consoloLogger) Fatal(format string, a ...interface{}) {
	if c.levelJudgment(FATAL) {
		return
	}
	c.getOutputMsgData("FATAL", format, a...)
}
