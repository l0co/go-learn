package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// example to receive unix signals
func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

// $ go run 53-signals.go
// awaiting signal
// ^C
// interrupt
// exiting
