package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)
/*
	执行bin\elasticsearch ，默认在本机的9200端口启动服务。

	执行bin/kibana , 默认通达是 ：http://localhost:5601

*/

// Elasticsearch demo
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	// 链接es
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es success")
	// 构造数据
	p1 := Person{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
