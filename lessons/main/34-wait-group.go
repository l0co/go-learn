package main

import (
	"fmt"
	"sync"
	"time"
)

// wait group does basically the same thing as we've done manually in 33-worker-pool
func worker3(id int, wg *sync.WaitGroup) {

	// defer defers execution of this call to this function end
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// a new zero-valued struct is created and assigned
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker3(i, &wg)
	}

	wg.Wait()
}
