package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)


type test struct {
	Data string `json:"data"`
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	msg := test{Data:"Hello, Hello. How are you?"}
	marshal, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(" json Marshal err")
	}
	for i := 0; i < 20; i++ {
		conn.Write(marshal)
	}
}
