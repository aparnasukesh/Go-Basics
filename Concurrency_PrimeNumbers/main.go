package main

import (
	"fmt"
	"sync"
)

func main() {
	N := 50

	var wg sync.WaitGroup
	wg.Add(2)

	numChan := make(chan int)
	primeChan := make(chan int)

	go generateNumbers(N, &wg, numChan)
	go generatePrimeNumbers(numChan, &wg, primeChan)
	go func() {
		wg.Wait()
		close(primeChan)
	}()
	for prime := range primeChan {
		fmt.Println(prime)
	}

}
func generateNumbers(N int, wg *sync.WaitGroup, numChan chan<- int) {
	defer wg.Done()
	for i := 2; i <= N; i++ {
		numChan <- i
	}
	close(numChan)
}

func generatePrimeNumbers(numChan <-chan int, wg *sync.WaitGroup, primeChan chan<- int) {
	defer wg.Done()
	for num := range numChan {
		if isPrime(num) {
			primeChan <- num
		}
	}
}

func isPrime(num int) bool {
	flag := 0
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			flag = 1
			break
		}
	}
	if flag == 1 {
		return false
	}
	return true
}
