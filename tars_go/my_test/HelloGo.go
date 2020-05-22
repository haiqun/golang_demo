package main

import (
	"TestApp"
	"github.com/TarsCloud/TarsGo/tars"
)

func main() { //Init servant
	imp := new(SayHelloImp)                                    //New Imp
	app := new(TestApp.SayHello)                                 //New init the A JCE
	cfg := tars.GetServerConfig()                               //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".SayHelloObj") //Register Servant
	tars.Run()
}
