package main

import (
	"fmt"
	session "golang_demo/session_hq"
)

func main() {
	sMgr := session.NewMemorySessionMgr()
	fmt.Println(sMgr)
}