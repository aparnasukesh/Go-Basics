package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Fan in method")
	fanIn()

	fmt.Println("nomral")
	normal()
}

func normal() {
	var data []string

	start := time.Now()
	for i := 0; i < 5; i++ {
		data = append(data, quary(i))
	}
	fmt.Println(data)
	fmt.Println(time.Now().Sub(start))
}

func fanIn() {
	ch := make(chan string)
	wg := sync.WaitGroup{}

	var data []string

	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- quary(i)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		data = append(data, val)
	}

	fmt.Println(data)

	fmt.Println(time.Now().Sub(start))
}

func quary(i int) string {
	time.Sleep(1 * time.Second)
	switch i {
	case 0:
		return "first-data"
	case 1:
		return "second-data"
	case 2:
		return "thrid-data"
	case 3:
		return "fourth-data"
	case 4:
		return "fifth-data"
	default:
		return "finished"
	}
}
