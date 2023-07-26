package main

import "fmt"

func main() {
	fmt.Println("Array Samples In Golang")

	// Fixed sise collection of item
	// Can't change the size after the declaration

	// Example 1
	var arr1 [3]string
	arr1 = [3]string{"Coffee", "Tea", "Other"}
	fmt.Println("arr1 :", arr1)

	// Example 2
	arr2 := [3]string{"Car", "Bus", "Train"}
	fmt.Println("arr2 :", arr2)

	// Example 3
	arr3 := arr2 //this is a Copy Operation
	fmt.Println("arr3 :", arr3)

	// Example 4
	arr2[1] = "Bike"
	fmt.Println("arr2 :", arr2)
	fmt.Println("arr3 :", arr3)

	// Integer array demo
	var intArray [3]int // Declaration
	fmt.Println("intArray :", intArray)

	// Assign values
	intArray = [3]int{10, 20, 30}
	fmt.Println("intArray :", intArray)
	fmt.Println("intArray :", intArray[0])

	intArray[1] = 100
	fmt.Println("intArray :", intArray)

	// Find length of array
	length := len(intArray)
	fmt.Printf("Length of intArray : %d", length)
}
