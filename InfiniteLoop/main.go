package main

import "fmt"

func main() {

	//  PROGRAM - INFINITE LOOP----------------------------------------------

	for {
		fmt.Println("loops demo") // without break it will execute infinite times
		break
	}

	for {
		fmt.Println("This is infinte loop") //if we want to exit from an inifinite loop we can put a conditon to exit
		break
	}
}
