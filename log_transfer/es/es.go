package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"strings"
	"fmt"
)

// 声明一个全局变量接受实例化的es链接

var client *elastic.Client

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
	return
}

// SendMsg 发送信息到es
func SendMsg(indexStr string,data interface{}) (err error) {
	// 构造数据
	//p1 := Person{Name: "rion", Age: 22, Married: false}
	_, err = client.Index().Index(indexStr).BodyJson(data).Do(context.Background())
	if err != nil {
		// Handle error
		//panic(err)
		return
	}
	return
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}