package main

import (
	"fmt"
	"time"
)

/*
	goroutines is a lightweight thread manage by goruntime.
	goroutines just function that leave the main thread and run in the background and come back to join the main thread once the functions are finished/ready to return any value.
	goroutines do not stop the program flow and are non blocking

	goroutines scheduling in go: 
		- Manage by the Go runetime Scheduler
		- Uses M:N Scheduling Model: M goroutines run on N OperatingSystem threads
		- Efficient Multiplexing
	
	Common pitfalls and best practices
		- Avoiding goroutines leaks
		- Limiting goroutine creation
		- Proper error handling
		- Synchronization

*/

func main() {
	var err error

	fmt.Println("Beginning program :")
	go sayHello() // Whats happening here ? 
	fmt.Println("After sayHello function.")
	// sayHello()

	go func() {
		err = doWork()
	}()

	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)
	
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}

}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i, " | ", time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letters := range "abcde" {
		fmt.Println(string(letters), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work

	time.Sleep(1 * time.Second)

	return fmt.Errorf("an error occured in doWork.")
}

