package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		fmt.Println("string:", <-c)
	}()
	c <- "hello"
	time.Sleep(time.Second)
}
