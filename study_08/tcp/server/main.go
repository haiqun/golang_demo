package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func communication(conn net.Conn) {
	defer conn.Close()
	for {
		var info [144]byte
		reader := bufio.NewReader(os.Stdin)
		// 收到客户端的信息 - 读取
		n, err := conn.Read(info[:])
		if err != nil {
			fmt.Println("Read failed，err ,", err)
			break
		}
		fmt.Println("客户端输入：", string(info[:n]))
		// 回复客户端的信息
		fmt.Print("我的回复：")
		// fmt.Println("\"Q\"为退出当前聊天")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取输入信息有误 :", err)
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "Q" {
			break
		}
		conn.Write([]byte(msg))
	}
}

var wg sync.WaitGroup

// tcp服务端
func main() {
	// 1. 本地端口启动服务
	str := "127.0.0.1:20000"
	listener, err := net.Listen("tcp", str) //  "127.0.0.1:20000"
	if err != nil {
		fmt.Println("创建本地端口 127.0.0.1:20000 有误，err ,", err)
		return
	}
	for {
		// 2. 等待别人跟我链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept failed，err ,", err)
			return
		}
		// 3. 与客户端通讯
		go communication(conn)
	}
}
