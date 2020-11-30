package main

import "fmt"

func main() {

	// synchronous channel with buffer

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // it's possible to close channel without all messages consumed

	// range can be used to iterate messages in channel
	for elem := range queue {
		fmt.Println(elem)
	}
	// one
	// two

	// asynchronouse channel without buffer

	queue2 := make(chan string)

	go func() {
		queue2 <- "three"
		queue2 <- "four"
		close(queue2)
	}()

	for elem := range queue2 {
		fmt.Println(elem)
	}
	// three
	// four

}
