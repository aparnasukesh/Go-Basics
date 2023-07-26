package main

import "fmt"

func Sum(a, b int) int {
	return a + b
}
func Calc(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

// Variadic Parameters
func PrintNames(name ...string) { // ... is called ellipses operator
	fmt.Println(name)

	// We can print one by one
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	fmt.Println(name[3])
}
func main() {
	fmt.Println("Function Demo")
	result := Sum(10, 20)
	fmt.Println("Result :", result)
	sum, sub := Calc(100, 200)
	fmt.Println(sum, sub)
	PrintNames("vw", "bmw", "audi", "scoda")
}
