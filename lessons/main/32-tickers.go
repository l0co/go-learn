package main

import (
	"fmt"
	"time"
)

// ticker is like timer with multiple execution with a specific duration
func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	// Tick at 2020-11-30 15:57:52.324429977 +0100 CET m=+0.500172758
	// Tick at 2020-11-30 15:57:52.824468366 +0100 CET m=+1.000211144
	// Tick at 2020-11-30 15:57:53.324463882 +0100 CET m=+1.500206659
	// Ticker stopped

}
