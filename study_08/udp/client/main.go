package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入要转换的大小写的英文字符串：")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取用户输入的信息有误，err:", err)
		}
		msg := []byte(str)
		// 传送数据
		_, err = socket.Write(msg)
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			return
		}
		// 接受数据
		var data [1024]byte
		n, addrudp, err := socket.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("接收数据失败，err:", err)
			return
		}
		fmt.Printf("recv:%v count:%v\n addr:%v\n", string(data[:n]), n, addrudp)
	}

}
