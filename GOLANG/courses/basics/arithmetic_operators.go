package main

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	fmt.Printf("\n")

	const p float64 = 22.0 / 7.0
	fmt.Println(p)

	fmt.Println()

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println("max value of int 64: ", maxInt)

	maxInt = maxInt + 1
	fmt.Println("overflow bc max+1: ", maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
	fmt.Println("max value of uint 64: ", uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println("underflow of uint64: ", uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	/* 
		The result would lead to an underflow situation where the result is smaller than the smallest positive normalized number that float64 can handle, so result = 0
	*/
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}