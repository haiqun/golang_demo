package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"fmt"
	"time"
)

// Init 实例化单个链接
func Init() (* mongo.Client){
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
// ...
func CancelConnect(client * mongo.Client)  {
	// 断开连接
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}



// pool 连接池模式
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error)  {
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 通过传进来的uri连接相关的配置
	o := options.Client().ApplyURI(uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	o.SetMaxPoolSize(num)
	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 返回 client
	return client.Database(name), nil
}