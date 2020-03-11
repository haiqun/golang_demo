package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp client

func main() {
	// 1 与server端建立链接
	var ipAndPort string
	for k, v := range os.Args {
		if k == 1 {
			ipAndPort = v
		}
	}
	if len(ipAndPort) <= 0 {
		ipAndPort = "127.0.0.1:20000"
	}

	conn, err := net.Dial("tcp", ipAndPort) // "127.0.0.1:20000"
	if err != nil {
		fmt.Printf("dial %s failde ,err:%s ", ipAndPort, err)
		return
	}
	defer conn.Close()
	// 2 发送数据
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("欢迎光临，请留言：？")
	for {
		// 获取输入的内容 ，传给服务端
		fmt.Print("我：")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取输入信息有误 :", err)
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
		// 获取客户端回复的信息
		var info [512]byte
		n, err := conn.Read(info[:])
		if err != nil {
			fmt.Println("读取服务端的信息有误 :", err)
			break
		}
		fmt.Println("服务端回复：", string(info[:n]))
	}
}
