package main

import (
	"fmt"
	"slices"
)

func main() {
	
	// var sliceName[]elementType

	// var numbers []int
	// var number1 = []int{1,2,3}

	// number2 := []int{9,8,7}

	// slice := make([]int,5)

	a := [5]int{1,2,3,4,5}
	slice1 := a[3:5]

	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7, 8, 9, 10, 901309175091)
	fmt.Println(slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("Slicecopy:", sliceCopy)

	// var nilSlice []int

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println()

	fmt.Println("Element at index 3 of slice1", slice1[3])

	// slice1[3] = 50
	// fmt.Println("Element at index 3 of slice1", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 is equal to sliceCopy")
	}

	fmt.Println()

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}

	fmt.Println(twoD)

	// slice[low:high]
	slice2 := slice1[2:4]
	fmt.Println(slice2)

}
