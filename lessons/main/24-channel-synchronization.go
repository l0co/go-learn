package main

import (
	"fmt"
	"time"
)

// channels can be used to wait for the goroutine to finish and get the resulting value of goroutine execution
func worker(done chan bool) {
	fmt.Println("starting work in goroutine")
	time.Sleep(time.Second)
	fmt.Println("finishing work in goroutine")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	fmt.Println("starting to wait for goroutine to finish")
	<-done // reads from the channel but doesn't save the read value anywhere
	fmt.Println("goroutine finished")

	// starting to wait for goroutine to finish
	// starting work in goroutine
	// finishing work in goroutine
	// goroutine finished
}
