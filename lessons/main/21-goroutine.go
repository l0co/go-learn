package main

import (
	"fmt"
	"time"
)

// goroutine is a lightweight thread
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// synchronous execution
	f("direct")

	// moving it to goroutine (other thread) with `go` clause
	go f("goroutine")

	// also possible for anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
