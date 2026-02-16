package main

import (
	"fmt"
	"time"
)

/*

	Un channel est comme un tuyau pour transmettre des données entre goroutines.
	Unbuffered channel = sender attend que receiver soit pret, synchro stricte obligatoire
	Buffered channel = 


	Channels enable safe and efficient comunication between concurrent goroutines, it helps synchronize and manage the flow of data in concurrent programs.

	Basics of Channels:
		- creating channels (make(chan Type))
		- sending and receiving data (<-)
		- channel directions:
			- send-only `ch <- value`
			- receive-only: `value := <- ch`

	Common Pitfalls
		- Avoid deadlocks
		- Avoiding Unnecessary buffering
		- channel direction
		- Graceful shutdown
		- use `defer` for unlocking

*/

func main() {
	// messages := make(chan string)

	// go func() { messages <- "ping" }()

	// msg := <-messages
	// fmt.Println(msg)

	//variable := make(chan type) '<-' operator
	greeting := make(chan string)
	greetString := "Hello"

	// greeting <- greetString // blocking because it is continuously trying to receive values, so we need a go routine function which is non-blocking function
	go func(){
		greeting <- greetString
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet : " + string(e)
		}
	}()
		
	// go func(){
	// 	receiver := <- greeting 
	// 	fmt.Println(receiver)
	// 	receiver = <- greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <- greeting 
	fmt.Println(receiver)
	receiver = <- greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <- greeting
		fmt.Println(rcvr)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("End of the program")
}