package main

import (
	"github.com/micro/go-log"
	"net/http"
	"net/rpc"
)

/*
golang-rpc程序的4个条件
结构体字段首字母要大写，可以别人调用
函数名必须首字母大写
函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型
函数还必须有一个返回值error
*/

// 例题：golang实现RPC程序，实现求矩形面积和周长

type Params struct {
	Width, Height int
}

// 用来绑定方法,client调用的时候，也是用这个为主的
type Rect struct{}

// RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

// 周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

// 判断是不是四边形 
func (r *Rect)IsSquare(p Params,ret *bool) error  {
	*ret = false
	if p.Height == p.Width {
		 *ret = true
	}
	return nil
}


// 主函数
func main() {
	// 1.注册服务
	rect := new(Rect)
	// 注册一个rect的服务
	rpc.Register(rect)
	// 2.服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3.监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}