package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 传的参数 =》与服务端的结构体一致
type Params struct {
	Width, Height int // 字段类型也要一致
}

// 主函数
func main() {
	// 1.连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	// 2.调用方法
	// 面积
	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)
	// 周长
	err3 := conn.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：", ret)
	// 判断是不是正方形
	ret1 := false
	err4 := conn.Call("Rect.IsSquare",Params{50,51},&ret1)
	if err4 != nil {
		log.Fatal(err4)
	}
	fmt.Println("是正方形吗？:" ,ret1)

}