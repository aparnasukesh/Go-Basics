package main

import "fmt"

func main() {

	// COLLECTIONS OF ARRAYS-----------------------------------------------------------------------------------------

	numbers := [4]int{100, 111, 222, 333}
	for i, v := range numbers {
		fmt.Println(i, v) // here i representing index and v representing value

	}

	// if we don't want index then we have to declare '_'underscore there as default

	for _, v := range numbers {
		fmt.Println(v)
	}

	for i, _ := range numbers {
		fmt.Println(i)
	}

	// COLLECTIONS OF MAP----------------------------------------------------------------------------------------------

	letters := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}
	for k, v := range letters {
		fmt.Println(k, v)
	}
}
