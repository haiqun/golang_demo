package tailflog

import (
	"github.com/hpcloud/tail"
)

// 收集日志
var (
	tailObj *tail.Tail
)

// Init 实例化
func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开文件 - 文件切割的时候，切割完重新打开
		Follow:    true,                                 // 是否追随 - 文件切割的时候，切割完继续读
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 在哪里割地方开始读
		MustExist: false,                                // 文件不存在，是否报错
		Poll:      true}
	// 打开文件
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		return
	}
	return
}


//Readlog 输出读取到的内容
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
