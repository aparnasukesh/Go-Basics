package main

import (
	"errors"
	"fmt"
)

func main() {
	num := -21
	err := greater(num)
	if err != nil {
		fmt.Println("error found")
	}
}

func greater(val int) error {
	if val < 0 {
		return errors.New("Value is less")
	}
	fmt.Println("val is graeter than zero")
	return nil
}
