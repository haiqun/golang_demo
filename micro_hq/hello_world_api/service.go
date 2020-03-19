package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	pb "golang_demo/micro_hq/hello_world_api/proto"
	"log"
)

// 声明绑定的结构体
type Example struct {

}
type Foo struct {

}

// 结构体绑定方法

//func (e *Example)Call(ctx context.Context, req *pb.CallRequest) (rep *pb.CallResponse, err error){
//	log.Println("收到 Example Call 请求了！")
//	if len(req.Name) < 0{
//		return rep ,errors.New("没有参数")
//	}
//	rep.Message = "hello Call :" + req.Name
//	return rep ,nil
//}
func (e *Example)Call(ctx context.Context, req *pb.CallRequest,rep *pb.CallResponse) (err error){
	log.Print("收到 Example Call 请求了！")
	if len(req.Name) < 0{
		return errors.New("没有参数")
	}
	rep.Message = "hello Call :" + req.Name
	return nil
}


func (e *Example)Call2(ctx context.Context, req *pb.CallRequest, rep *pb.CallResponse) error  {
	rep.Message = "hello Call2 :"+ req.Name
	return nil
}

func (f *Foo)Bar(ctx context.Context, req *pb.EmptyRequest, rep *pb.EmptyResponse) error  {
	log.Fatal("收到 foo bar 的请求")
	return nil
}


func main() {
	// 注册服务
	service := micro.NewService(
			micro.Name("go.micro.api.example"),
		)
	// 初始化
	service.Init()
	// 绑定
	err := pb.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		log.Fatal("RegisterExampleHandler err:",err)
		return
	}
	err = pb.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		log.Fatal("RegisterFooHandler err:",err)
		return
	}
	// 启动服务
	if err := service.Run();err != nil{
		log.Fatal("service Run err:",err)
	}
}
