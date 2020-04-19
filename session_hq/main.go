package main

import (
	"fmt"
	"golang_demo/session_hq/session"
	"log"
)

func main() {

	strM := session.NewMemorySessionMgr()

	createSession, err := strM.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createSession.Set("test","hhhh"))

	fmt.Println(createSession.Get("test"))

	fmt.Println("123")
}



