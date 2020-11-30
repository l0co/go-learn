package main

import (
	"fmt"
	"time"
)

// an example of thread pool to which we randomly assign tasks
func worker2(id int, jobs <-chan int, results chan<- int) {
	// consuming messages from channel is exclusive and competitive, even using range iteration
	for j := range jobs {
		fmt.Println("worker2", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker2", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// there will be 5 tasks in total
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// work will be sprad amongst 3 workers
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	fmt.Println("Waiting to collect results")
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	fmt.Println("Exitting")

	// Waiting to collect results
	// worker2 3 started  job 1
	// worker2 1 started  job 2
	// worker2 2 started  job 3
	// worker2 2 finished job 3
	// worker2 2 started  job 4
	// worker2 3 finished job 1
	// worker2 3 started  job 5
	// worker2 1 finished job 2
	// worker2 3 finished job 5
	// worker2 2 finished job 4
	// Exitting
}
