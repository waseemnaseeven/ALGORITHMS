package main

import "fmt"

/*
	Durant le panic: 
		- Go deroule la stack, 
		- execute les defer
		- si rien ne recupere, renvoie une stack trace
	
	Seul facon d'arreter un panic c'est recover()
*/

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		fmt.Println("After Panic")
		defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}

func main() {

	// panic(interface{})

	// Example of a valid input
	process(10)

	// Example of an invalid input
	process(-3)
}

