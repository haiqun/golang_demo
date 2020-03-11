package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 日志记录到文本

type fileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	maxFileSize int64
	fileObj     *os.File
	fileObjErr  *os.File
	chanMsg     chan *logMsg
}

type logMsg struct {
	lv        string
	msg       string
	funcName  string
	filePath  string
	timestamp string
	fileLine  int
}

var chanMsg chan *logMsg

// NewFileLogger 构造函数实例化日志类
func NewFileLogger(levelStr, fp, fn string, maxfs int64) *fileLogger {
	l, err := paramLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &fileLogger{
		Level:       l,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxfs,
		chanMsg:     make(chan *logMsg, 50000),
	}
	err = f1.logFileInfo()
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *fileLogger) logFileInfo() error {
	// 拼装路径
	fileLog := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed ,err %s", err)
		return err
	}
	fileObjErr, err := os.OpenFile(fileLog+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed ,err %s", err)
		return err
	}
	f.fileObj = fileObj
	f.fileObjErr = fileObjErr
	// 开启5个后台的goroutine去异步写入
	for i := 0; i < 5; i++ {
		go f.writeLogToFile()
	}
	return nil
}

// 判断文件是否需要切割 ，重新包装新的文件
func (f *fileLogger) checkSize(file *os.File) bool {
	// 获取文件的详细信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件详细信息报错 ：%s", err)
		return false
	}
	// 当前文件大小，大于等于日志文件的最大值的时候 就应该返回true
	return fileInfo.Size() >= f.maxFileSize
}

func (f *fileLogger) splitFile(file *os.File) (*os.File, error) {
	// 获取时间戳
	nowStr := time.Now().Format("20060102150405.000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件详细信息报错 ：%s", err)
		return nil, err
	}
	// 创建新的文件
	logName := path.Join(f.filePath, fileInfo.Name()) // 要切割的源文件完整的路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	// 关闭文件
	file.Close()
	// 备份 xx.log  => xxx.lo g.20060102150405.000.bak
	err = os.Rename(logName, newLogName)
	if err != nil {
		fmt.Printf("文件切割报错 ：%s", err)
		return nil, err
	}
	// 将新的文件话柄 赋值 给f.fileObj
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开新的文件话柄报错了 ：%s", err)
		return nil, err
	}
	return fileObj, nil
}

func (f *fileLogger) writeLogToFile() {
	for {
		// 判断文件是否超出方位，需要切割
		if f.checkSize(f.fileObj) {
			newfileObj, err := f.splitFile(f.fileObj)
			if err != nil {
				// panic(err)
				return
			}
			f.fileObj = newfileObj
		}
		// 获取通道的数据
		select {
		case logTmp := <-f.chanMsg:
			// 拼写日志
			logInfo := fmt.Sprintf("%s [%s] [%s:%s:%d] %s \n", logTmp.timestamp, logTmp.lv, logTmp.filePath, logTmp.funcName, logTmp.fileLine, logTmp.msg)
			// 日志写入
			fmt.Fprintf(f.fileObj, logInfo)
			// 调试代码 在控制台输出
			// fmt.Printf("%s [%s] [%s:%s:%d] %s \n", t, level, filePath, funcName, fileLine, msg)
			l, _ := paramLevel(logTmp.lv)
			if l >= ERROR { // 错误级别大于等于error 级别的再记录一次
				if f.checkSize(f.fileObjErr) {
					newfileObj, err := f.splitFile(f.fileObjErr)
					if err != nil {
						// panic(err)
						return
					}
					f.fileObjErr = newfileObj
				}
				fmt.Fprintf(f.fileObjErr, logInfo)
			}
		default:
			// 取不到日志，就休眠 500 毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (f *fileLogger) getOutputMsgData(level string, format string, a ...interface{}) {
	// 构造输出日志的结构;
	msg := fmt.Sprintf(format, a...)
	t := time.Now().Format("2006-01-02 15:04:05.000")
	funcName, filePath, fileLine := getFileInfo(3)
	logTmp := &logMsg{
		lv:        level,
		msg:       msg,
		funcName:  funcName,
		filePath:  filePath,
		timestamp: t,
		fileLine:  fileLine,
	}
	// fmt.Println(*logTmp)
	// 保证业务不被影响 ，降级处理
	select {
	case f.chanMsg <- logTmp:
	default:
		// 如果，日志阻塞，就丢掉
	}

}

func (f *fileLogger) levelJudgment(L LogLevel) bool {
	return L < f.Level
}

func (f *fileLogger) Debug(format string, a ...interface{}) {
	if f.levelJudgment(DEBUG) {
		return
	}
	f.getOutputMsgData("DEBUG", format, a...)
}

func (f *fileLogger) Trace(format string, a ...interface{}) {
	if f.levelJudgment(TRACE) {
		return
	}
	f.getOutputMsgData("TRACE", format, a...)
}

func (f *fileLogger) Info(format string, a ...interface{}) {
	if f.levelJudgment(INFO) {
		return
	}
	f.getOutputMsgData("INFO", format, a...)

}

func (f *fileLogger) Warning(format string, a ...interface{}) {
	if f.levelJudgment(WARNING) {
		return
	}
	f.getOutputMsgData("WARNING", format, a...)

}

func (f *fileLogger) Error(format string, a ...interface{}) {
	if f.levelJudgment(ERROR) {
		return
	}
	f.getOutputMsgData("ERROR", format, a...)

}

func (f *fileLogger) Fatal(format string, a ...interface{}) {
	if f.levelJudgment(FATAL) {
		return
	}
	f.getOutputMsgData("FATAL", format, a...)
}

func (f fileLogger) fileClose() {
	f.fileObj.Close()
	f.fileObjErr.Close()
}
