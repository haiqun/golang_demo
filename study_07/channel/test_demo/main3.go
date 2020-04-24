package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch1 := make(chan os.Signal)
	signal.Notify(ch1, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	fmt.Println("22")
}