package main

import (
	"context"
	"fmt"
	"time"
)

// example to use context with goroutine

func infinity(c chan<- bool, ctx context.Context) {

	for i := 1; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)

		select {
		// support deadline in non-blocking channel read block
		case <-ctx.Done():
			fmt.Println("is done by context")
			c <- false
		default:
			fmt.Println("working", i, "/", 100)
		}
	}

	c <- true
}

func main() {

	c := make(chan bool)
	// create context with deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	_ = cancel // unused variable workaround
	go infinity(c, ctx)

	// also can be cancelled
	// <-time.After(300 * time.Millisecond)
	// cancel()

	<-c
	fmt.Println("exitting...")

}

// working 1 / 100
// working 2 / 100
// working 3 / 100
// working 4 / 100
// working 5 / 100
// working 6 / 100
// working 7 / 100
// working 8 / 100
// working 9 / 100
// is done by context
// exitting...
