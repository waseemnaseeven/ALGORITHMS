package main

import (
	"fmt"
	"time"
)

/*

	- Why use Buffered Channels ?
		- Asynchronous communication
		- Load Balancing
		- Flow control

	- Creating Buffered Channels
		- make(chan Type, capacity)
		- Buffer Capacity

	- Key Concepts of Channel Buffering
		- Blocking Behavior
		- Non-Blocking Operations
		- Impact on Performance

	- Best practices
		- Avoid Over-Buffering
		- Graceful Shutdown
		- Monitoring Buffer Usage

*/


// func main() {
// 	// Blocking on receive only if the buffer is empty
// 	ch := make(chan int, 2)
// 	go func(){
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2

// 	}()
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("Value: ", <-ch)

// }


func main() {
// Blocking on send only if the buffer is full
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func(){
		fmt.Println("Goroutine 2 second timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <-ch)
	}()
	// fmt.Println("Blocking starts")
	ch <- 3 // blocking because the buffer is full
	// fmt.Println("Blocking ends")
	// fmt.Println("Received: ", <-ch)
	// fmt.Println("Received: ", <-ch)
}