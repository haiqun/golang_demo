package main

import (
	"fmt"
	"net"
	"strings"
)

// udp server

func main() {
	// 创建一个udp的监听
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("ListenUDP err :", err)
		return
	}
	defer listen.Close()
	// 不需要建立链接，直接收发数据
	var b [1024]byte
	for {
		// 读取信息
		n, addrUdp, err := listen.ReadFromUDP(b[:])
		if err != nil {
			fmt.Println("ReadFromUDP err :", err)
			return
		}
		str := string(b[:n])
		// 打印收到的信息
		fmt.Println(str)
		// 回复信息 - 转为大写
		str = strings.ToUpper(str)
		_, err = listen.WriteToUDP([]byte(str), addrUdp)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}

}
