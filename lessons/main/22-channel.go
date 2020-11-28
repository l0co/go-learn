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
		time.Sleep(time.Second) // sleep 1
		fmt.Println("sending message...")
		// send to channel syntax, this operation is blocking until someone receives the message
		messages <- "ping"
		fmt.Println("message sent")
	}()

	fmt.Println("initializing receiving message...")
	time.Sleep(time.Second * 2) // sleep 2
	// receive from channel syntax, this operation is blocking until someone receives the message
	msg := <-messages
	fmt.Println("got message:", msg)

	time.Sleep(time.Second)

	// initializing receiving message...
	// initializing sending message...
	// ... sleep 1 + sleep 2 (first second)
	// sending message...
	// ... sleep 2 (second second)
	// got message: ping
	// message sent
}
