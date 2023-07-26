package main

import (
	"fmt"
)

func main() {
	// PROGARM-1-----------------------------------------------------------------------
	// Example 1
	var arr1 []string // Slices are dinamic Array
	fmt.Println(arr1)

	arr1 = []string{"Coffee", "Tea", "Other"}
	fmt.Println("arr1 :", arr1)

	// Example 2
	arr2 := []string{"Car", "Bus", "Tram"}
	fmt.Println("arr2 :", arr2)

	// Example 3
	arr3 := arr2 // This is copy operation
	fmt.Println("arr3 :", arr3)
	arr2[1] = "Bike"
	fmt.Println("arr2 :", arr2)
	fmt.Println("arr3 :", arr3)

	// PROGRAM-2------------------------------------------------------------------------
	fmt.Println("Program 2")
	s := []string{"a", "b", "c"}
	fmt.Println("s:", s)
	s = append(s, "d", "e", "f")
	fmt.Println("s :", s)

	// PROGRAM-3------------------------------------------------------------------------
	fmt.Println("Program 3")
	S1 := []string{"hai", "hei", "hoi"}
	fmt.Println("S1:", S1)
	S1 = append(S1, "Hello", "Hey")
	fmt.Println("S1:", S1)

	fmt.Println("S1:", S1)
	S1 = S1[2:]
	fmt.Println("S1:", S1)
}
