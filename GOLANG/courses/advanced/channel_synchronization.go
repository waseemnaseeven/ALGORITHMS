package main

import (
	"fmt"
	"time"
)

/*

	- Why is channel synch important ?
		- Ensure data is properly exchanged between goroutines
		- Coordinate the execution flow to avoid race conditions and ensure predictable behavior
		- Helps manage the lifecycle of goroutines and the completion of tasks

	- Common Pitfalls
		- avoid deadlocks
		- close channels
		- avoid unnecessary blocking

*/

// func main() {
// 	done := make(chan struct{})
	
// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	<-done
// 	fmt.Println("Finished.")
// }

// // SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING GOROUTINES ARE COMPLETE
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)
// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working... \n", id)
// 			time.Sleep(time.Second)
// 			done <- id // Sending signal of completion
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-done // Wait for each goroutine to finish, being received
// 	}

// 	fmt.Println("All goroutines are complete")
// }


// SYNCHRONIZING DATA EXCHANGE
func main() {
	
	data := make(chan string)
	go func(){
		for i := range 5 {
			data <- "hello " + string('0' + i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data) // emit and receive channels
	}()

	for value := range data {
		fmt.Println("Received value: ", value, ":", time.Now())
	}
	fmt.Println("Over...")
}