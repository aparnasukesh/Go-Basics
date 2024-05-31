package main

import (
	"fmt"
	"sync"
)

type result struct {
	name string
	sum  int
}

func main() {

	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{5, 6, 7, 8}
	arr3 := []int{2, 3, 9, 5}

	var wg sync.WaitGroup
	results := make(chan result, 3)

	wg.Add(3)
	go sum("arr1", arr1, &wg, results)
	go sum("arr2", arr2, &wg, results)
	go sum("arr3", arr3, &wg, results)

	wg.Wait()
	close(results)

	maxSum := 0
	maxArr := ""
	for res := range results {
		fmt.Printf("Sum of %s: %d\n", res.name, res.sum)
		if res.sum > maxSum {
			maxSum = res.sum
			maxArr = res.name
		}
	}

	fmt.Printf("Array with the greatest sum is %s with a sum of %d\n", maxArr, maxSum)

}

func sum(name string, arr []int, wg *sync.WaitGroup, results chan<- result) {
	defer wg.Done()
	total := 0
	for _, value := range arr {
		total += value
	}
	results <- result{name: name, sum: total}
}
