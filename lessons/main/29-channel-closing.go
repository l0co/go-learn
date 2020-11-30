package main

import "fmt"

// after channel is closed no more messages will be sent to this channel
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // the second return value will be false when channel is closed
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // closing the channel
	fmt.Println("sent all jobs")

	<-done
}
