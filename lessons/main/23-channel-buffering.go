package main

import "fmt"

func main() {

	/*
			sending to the channel is blocking so if you send anything to a channel in main goroutine and there's no receiving
			goroutines, this means locking the main thread and ends up with PANIC:
		 	fatal error: all goroutines are asleep - deadlock!

			messages := make(chan string)
			messages <- "buffered"
	*/

	// provided the channel is buferred
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// buffered
	// channel
}
