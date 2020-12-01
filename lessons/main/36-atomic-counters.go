package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64
	var nonAtomicOps uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				nonAtomicOps++ // there's a gap between read and write ( i = i + 1 ) so we would have a race condition here
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
	fmt.Println("nonAtomicOps:", nonAtomicOps)

	// ops: 50000
	// nonAtomicOps: 22351
}
