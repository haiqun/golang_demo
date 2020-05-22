package main

import (
	"flag"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"os"

	"TestApp"
)

//只需初始化一次，全局的
var comm *tars.Communicator

func main() {
	comm = tars.NewCommunicator()
	obj := "TestApp.HelloGo.SayHelloObj@tcp -h 127.0.0.1 -p 9997 -t 60000"
	app := new(TestApp.SayHello)
	/*
	   // if your service has been registered at tars registry
	   comm = tars.NewCommunicator()
	   obj := "TestApp.HelloGo.SayHelloObj"
	   // tarsregistry service at 192.168.1.1:17890
	   comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 192.168.1.1 -p 17890")
	*/
	var name string
	comm.StringToProxy(obj, app)
	if len(os.Args) > 0  {
		name = os.Args[1]
	}else {
		name = *flag.String("name", "lufei", "请输入你的姓名")
	}

	//reqStr := name
	var resp string
	for i:=0;i<10;i++ {
		ret, err := app.EchoHello(name, &resp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ret: ", ret, "resp: ", resp)
	}
}