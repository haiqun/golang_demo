package main

import "log"

type SayHelloImp struct {
}

func (imp *SayHelloImp) EchoHello(name string, greeting *string) (int32, error) {
	log.Println(" cli 传过来的值 ：",name)
	*greeting = "hello " + name
	return 0, nil
}