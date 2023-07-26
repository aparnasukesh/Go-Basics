package main

import (
	"errors"
	"fmt"
)

//Numbers
var var1 int
var var2 uint
var var3, var4 int //Multideclaration
var var5 float32   //32 bit
var var6 float64   //64 bit

// Complex Numbers a+bi
var var7 complex64
var var8 complex128

// Boolean
var var9 bool

// String
var var10 string
var var11 string = "hello"

//Constants
const RED = 0
const (
	CAR = iota
	BUS
	TRAM
)

func main() {

	// Constants
	fmt.Println(CAR)
	fmt.Println(BUS)
	fmt.Println(TRAM)

	// Short declaration
	c := 10
	fmt.Println(c)

	// Type conversion
	var3 = 10
	var5 = float32(var3)

	// Multivariable
	a, b := 1, 2
	fmt.Printf("%v+%v = %v\n", a, b, a+b)

	// Errors
	err := errors.New("hello")
	fmt.Println(err)
}
