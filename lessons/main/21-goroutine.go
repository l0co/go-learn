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

	fmt.Println("awaiting")
	time.Sleep(time.Second)
	fmt.Println("done")

	// direct : 0
	// direct : 1
	// direct : 2
	// awaiting
	// going
	// goroutine : 0
	// goroutine : 1
	// goroutine : 2
	// done

	// we can now check what happens to running goroutine when the program ends

	go func() {
		for {
			time.Sleep(time.Millisecond)
			fmt.Println("infinite loop...")
		}
	}()
	time.Sleep(10 * time.Millisecond)

	// infinite loop...
	// infinite loop...
	// infinite loop...
	// infinite loop...
	// infinite loop...
	// infinite loop...
	// infinite loop...
	// infinite loop...
	// (program exits)

	// so we can leave working goroutine and it will be hard shutdown when the program exits

}
