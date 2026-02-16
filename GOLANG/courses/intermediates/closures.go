package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())

	fmt.Println("\nFROM GOCOURSE NOW BELOW\n")

	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())

	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))

}