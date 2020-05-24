package rabbitmq

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

// ampq://用户:密码@主机:端口/vhost
const MQURL  = "amqp://fhq321:fhq123456@127.0.0.1:5672/test"

/*
给用户关联vhost 权限
rabbitmqctl set_permissions -p vhost user 后边三个.*分别代表：配置权限、写权限、读权限
rabbitmqctl set_permissions -p test fhq321 ".*" ".*" ".*"
*/

type rabbitmq struct {
	conn * amqp.Connection
	channel * amqp.Channel
	// 队列名称
	QueueName string
	// 交换器
	Exchange string
	// key
	Key string
	// MQURL
	MQURL string
}

// 创建结构体实例
func NewRabbitMQ(QueueName,Exchange,Key string) * rabbitmq  {
	mq := &rabbitmq{
		QueueName: QueueName,
		Exchange:  Exchange,
		Key:       Key,
		MQURL:     MQURL,
	}
	var err error
	mq.conn,err = amqp.Dial(mq.MQURL)
	mq.failOnErr(err,"创建链接错误！")
	mq.channel,err = mq.conn.Channel()
	mq.failOnErr(err,"获取channel错误！")
	return  mq
}

// 关闭连接
func (r *rabbitmq)Destory()  {
	r.conn.Close();
	r.channel.Close();
}

// 定义错误报告
func (r * rabbitmq) failOnErr(err error,msg string)  {
	if err != nil {
		log.Fatalf("%s:%s",msg,err)
		panic(fmt.Sprintf("%s:%s",msg,err))
	}
}

// 简单模式的mq
func NewRabbitMQSimple(Queuename string) * rabbitmq{
	return NewRabbitMQ(Queuename,"","")
}

// 简单模式的生产者
func (r * rabbitmq)PublishSimple(msg string)(err error,b bool) {
	// 申请队列 ，如果存在就跳过
	_,err = r.channel.QueueDeclare(
			r.QueueName, // 队里名称
			false, // 是否持久化，重启后是否会重载旧数据
			false, // 是否自动删除
			false, // 是否具有排他性
			false, // 是否柱塞
			nil, // 而外属性
			)
	if err != nil {
		log.Printf("QueueDeclare err :%s",err)
		return  err, false
	}
	// 推送数据
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, // 如果为true，根据exchange 类型和routkey规则，如果无法找到符合条件的队列，那么会把发送的消息返回给发送者
		false, // 如果为true，当exchange发送消息到队列后，发现队列上没有绑定消费者，则会把消息发还给消费者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		log.Printf("Publish err:%s",err)
		return err,false
	}
	return nil,true
}


// 简单模式的消费者


func (r * rabbitmq) RecieveSimple() {
	// 1 - 申请队列 ，如果存在就跳过
	log.Println("RecieveSimple starting ....")
	_,err := r.channel.QueueDeclare(
			r.QueueName, // 队里名称
			false, // 是否持久化，重启后是否会重载旧数据
			false, // 是否自动删除
			false, // 是否具有排他性 , 自己创建的是否只有自己可见 ？ false 为别人可见
			false, // 是否柱塞 ，推送完才让下一个信息推动
			nil, // 而外属性
			)
	if err != nil {
		log.Printf("QueueDeclare err : %s ",err)
	}
	// 2 - 接受信息
	consume, err := r.channel.Consume(
		r.QueueName,
		"",    // 用来区分多个消费者
		true,  // 是否自动应答，消费完ack，如果设置为false，需要自己写一个回调函数处理
		false, // 是否具有排他性
		false, // 设置为true，表示不能同一个connection中的发送信息传递给这个connection中的消费者
		false, // 消息消费是否阻塞
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	f := make(chan struct{})
	// 3 - 协程处理信息
	go func() {
		// 业务逻辑代码
		for d := range consume {
			log.Printf("received a message : %s ",d.Body)
		}
	}()
	log.Println("等待消费信息中....")
	<-f
}


