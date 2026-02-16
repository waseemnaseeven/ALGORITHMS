package main

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers) // [0 0 0 0 0]

	numbers[4] = 20
	fmt.Println(numbers) // [0 0 0 0 20]

	numbers[0] = 9
	fmt.Println(numbers) // [9 0 0 0 20]

	fmt.Println()

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element:", fruits[2])

	fmt.Println()

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", copiedArray)

	fmt.Println()

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	fmt.Println()

	for index, value := range fruits {
		fmt.Printf("Index %d, Value %s\n", index, value)
	}

	fmt.Println()

	a,b := someFunction()
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println()

	// underscore is blank identifier, used to store unused values
	for _, v := range numbers {
		fmt.Printf(" Value: %d\n", v)
	}

	fmt.Println()

	fmt.Println("The length of numbers array is", len(numbers))

	// Comparing Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}

	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	Array2 := [3]int{1, 2, 3}
	var newArray *[3]int

	newArray = &Array2
	newArray[0] = 100

	fmt.Println("Original array:", Array2)
	fmt.Println("Copied array:", *newArray)
}

func someFunction() (int, int) {
	return 1, 2
}
