package main

import (
	"fmt"
	"time"
)

func main() {
	unbufferedChannel := make(chan string, 0)

	go func() {
		time.Sleep(5 * time.Second)
		unbufferedChannel <- "hello"
	}()

	select {
	case msg := <-unbufferedChannel:
		fmt.Println(msg)
	default:
		fmt.Println("channel was too slow")
	}
}
