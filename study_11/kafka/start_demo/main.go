package main

/*
	zookerrper 配置文件是：
	config/zookeeper.properties
	修改配置： dataDir =》 数据文件存放路径
	修改配置： clientPort =》链接地址，集群的时候这个地址应该是多个的
	修改配置-启动-kafka的 zookeeper
	/zookeeper-server-start.sh config/zookeeper.properties


	server.properties 配置文件是：
	config/server.properties
	修改配置： log.dirs =》 文件的存放路径
	启动 kafka-server 的服务
	bin/kafka-server-start.sh config/server.properties
*/
func main() {

}
