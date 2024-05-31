package main

import "fmt"

type result struct {
	name string
	sum  int
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{7, 8, 9}

	results := make(chan result, 3)

	go arraySum("arr1", arr1, results)
	go arraySum("arr2", arr2, results)
	go arraySum("arr3", arr3, results)

	result1 := <-results
	result2 := <-results
	result3 := <-results

	allResults := []result{result1, result2, result3}

	maxSum := 0
	maxArr := ""

	for _, res := range allResults {
		if res.sum > maxSum {
			maxSum = res.sum
			maxArr = res.name
		}
	}
	fmt.Printf("Array with greatest sum is %s with sum is %d", maxArr, maxSum)
}

func arraySum(name string, arr []int, results chan<- result) {

	total := 0
	for _, val := range arr {
		total = total + val
	}

	results <- result{name: name, sum: total}
}
