package main
/*
	Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，
	我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等。
	logrus
	golang标准库的日志框架很简单，logrus框架的特点：
	1)完全兼容标准日志库
		六种日志级别：debug, info, warn, error, fatal, panic
	2)可扩展的Hook机制
		允许使用者通过Hook的方式将日志分发到任意地方，如本地文件系统，logstash，elasticsearch或者mq等，或者通过Hook定义日志内容和格式等
	3)可选的日志输出格式
		内置了两种日志格式JSONFormater和TextFormatter，还可以自定义日志格式
	4)Field机制
		通过Filed机制进行结构化的日志记录
	5)线程安全
		logrus不提供的功能如下：
			1)没有提供行号和文件名的支持
			2)输出到本地文件系统没有提供日志分割功能
			3)没有提供输出到ELK等日志处理中心的功能
	这些功能都可以通过自定义hook来实现
*/
import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
	// 设置日志输出的格式
	log.SetFormatter(&log.JSONFormatter{})
	// 设置日输出的地点
	fileObj,err := os.OpenFile("./xx.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		fmt.Println(" open file failed err: ",err)
	}
	//log.SetOutput(os.Stdout)
	log.SetOutput(fileObj)
	// 设置日志的级别
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("a walrus appears")
}