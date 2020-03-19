package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "golang_demo/micro_hq/hello_world/proto"
	)

// 绑定结构体

type hello struct {
	
}


func (h *hello)Info(ctx context.Context, req *pb.InfoRequest, rep *pb.InfoResponse) error  {
	rep.Msg = "hello "+ req.Name
	return nil
}

func main()  {
	// 1 得到服务端的实例
	service := micro.NewService(
			// 服务注册 - 提空外部调用的唯一标识;
			// 需要一个编译好的micro程序
			// micro call h Hello.Info {"username":"lufei"}
			micro.Name("h"),
		)
	// 2 初始化 => context上下文，超时等
	service.Init()
	// 3 服务注册
	err := pb.RegisterHelloHandler(service.Server(), new(hello))
	if err != nil {
		fmt.Println("服务注册失败")
		return
	}
	// 4 开启服务
	service.Run()
}
