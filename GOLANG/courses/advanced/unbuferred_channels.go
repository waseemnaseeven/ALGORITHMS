package main

import (
	"fmt"
)

/*

	Unbeffered and buffered channels: 

*/

func main() {

    // messages := make(chan string, 2)

    // messages <- "buffered"
    // messages <- "channel"

    // fmt.Println(<-messages)
    // fmt.Println(<-messages)

	ch := make(chan int)
	go func() {
		ch <- 1
		}()
		receiver := <- ch
		fmt.Println(receiver)

}