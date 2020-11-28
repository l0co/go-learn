package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those
// values into another goroutine.
func main() {

	messages := make(chan string)

	go func() {
		fmt.Println("initializing sending message...")
		time.Sleep(time.Second)
		fmt.Println("sending message...")
		// send to channel syntax
		messages <- "ping"
	}()

	fmt.Println("initializing receiving message...")
	// receive from channel syntax, this operation is blocking
	msg := <-messages
	fmt.Println("got message:", msg)

	// initializing receiving message...
	// initializing sending message...
	// sending message...
	// got message: ping
}
