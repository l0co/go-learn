package main

import (
	"fmt"
	"time"
)

// timer publishes message via channel when is fired (once)
func main() {

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")
	// Timer 1 fired

	// we can stop timer using Stop() method
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	// (does not display anything)

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// Timer 2 stopped
}
