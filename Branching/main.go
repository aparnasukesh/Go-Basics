package main

import "fmt"

func main() {

	fmt.Println("Branching")

	// IF STATEMENTS-----------------------------------------------------------------------------

	// 1.if Statements

	i := 4
	if i < 6 {
		fmt.Println("i<6")
	}
	// 2.if else Statement

	j := 7
	if j < 6 {
		fmt.Println("j<6")
	} else {
		fmt.Println("Failed.....")
	}

	// 3.if ,else if, else Statements
	k := 6
	if k < 6 {
		fmt.Println("k<6")
	} else if k > 6 {
		fmt.Println("k>6")
	} else {
		fmt.Println("k==6")
	}

	// SWITCH STATEMENTS-----------------------------------------------------------------------------

	// 1. Using logical expression
	n := 3
	switch n {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	default:
		fmt.Println("Invalid.....")
	}

	// 2. Using conditional switch
	m := 6
	switch {
	case m < 6:
		fmt.Println("m<6")
	case m > 6:
		fmt.Println("m>6")
	default:
		fmt.Println("Default")
	}

	// 3. More than one condition
	l := 6
	switch {
	case l < 6, i < 7:
		fmt.Println("i<6")
	case i > 7:
		fmt.Println("i>6")
	default:
		fmt.Println("default....")
	}
}
