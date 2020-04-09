package main

import (
	"fmt"
	"log"
	"os"
)

/*
	默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，
	比如记录该日志的文件名和行号等。log标准库中为我们提供了定制这些设置的方法。
	log标准库中的Flags函数会返回标准logger的输出配置，
	而SetFlags函数用来设置标准logger的输出配置

	log标准库提供了如下的flag选项，它们是一系列定义好的常量。
	const (
		// 控制输出日志信息的细节，不能控制输出的顺序和格式。
		// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
		Ldate         = 1 << iota     // 日期：2009/01/23 默认
		Ltime                         // 时间：01:23:23
		Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
		Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
		Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
		LUTC                          // 使用UTC时间
		LstdFlags     = Ldate | Ltime // 标准logger的初始值
	)
*/

func init()  {
	// 定制日志格式
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate | log.LUTC )
	// 修改日志前缀
	log.SetPrefix("[lufei]")
	// 设置日志输出的位置
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
}

func main() {
	v := "lufei"
	log.Printf(" %s say:  hhhhh ",v)
}
