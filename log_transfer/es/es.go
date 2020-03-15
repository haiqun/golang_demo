package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"strings"
	"fmt"
	"time"
)

// 声明一个全局变量接受实例化的es链接
type MsgData struct {
	Index string
	Mgs interface{}
}

var client *elastic.Client
var ch = make(chan *MsgData,100000)
// Init 实例化es的链接
func Init(address string) (err error){
	// 链接es
	if !strings.HasPrefix(address,"http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return
	}
	fmt.Println("es connect success!")
	// 监控信息渠道
	go sendMsg()
	return
}

func SendToEsChan(msg *MsgData) {
	ch <- msg
}

// SendMsg 发送信息到es
func sendMsg() {
	// 处理信息-从chan获取
	for  {
		select {
		case info := <-ch:
			// 构造数据
			_, err := client.Index().Index(info.Index).BodyJson(info.Mgs).Do(context.Background())
			if err != nil {
				fmt.Printf("es sendMsg failed : %s",err)
			}
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
