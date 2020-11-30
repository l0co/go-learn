package main

import (
	"fmt"
	"time"
)

// ticker can be used for rate limiting
func main() {

	// reading from channel without rate limiting
	fmt.Println("no rate limiting")
	{
		c := make(chan int)

		go func() {
			for i := 1; i <= 5; i++ {
				c <- i
			}
			close(c)
		}()

		for i := range c {
			fmt.Println(time.Now(), i)
		}

		// 2020-11-30 17:48:29.897766160 +0100 CET m=+0.000045627 1
		// 2020-11-30 17:48:29.897872044 +0100 CET m=+0.000151530 2
		// 2020-11-30 17:48:29.897883108 +0100 CET m=+0.000162554 3
		// 2020-11-30 17:48:29.897890773 +0100 CET m=+0.000170208 4
		// 2020-11-30 17:48:29.897899428 +0100 CET m=+0.000178869 5
	}

	// using ticker for rate limiting
	fmt.Println("\nticker rate limiting")
	{
		c := make(chan int)

		go func() {
			for i := 1; i <= 5; i++ {
				c <- i
			}
			close(c)
		}()

		limiter := time.Tick(200 * time.Millisecond)
		for i := range c {
			<-limiter
			fmt.Println(time.Now(), i)
		}

		// 2020-11-30 17:50:28.406684962 +0100 CET m=+0.200385630 1
		// 2020-11-30 17:50:28.606675765 +0100 CET m=+0.400376442 2
		// 2020-11-30 17:50:28.806688771 +0100 CET m=+0.600389408 3
		// 2020-11-30 17:50:29.006631493 +0100 CET m=+0.800332153 4
		// 2020-11-30 17:50:29.206679341 +0100 CET m=+1.000379987 5
	}

	// using ticker for bursty rate limiting
	fmt.Println("\nburst ticker rate limiting")
	{
		c := make(chan int)

		// produces burst rate limiter, which executes each `millis` but allows to process `burst` items immediately
		var produceLimiter = func(millis time.Duration, burst int) chan time.Time {
			limiter := make(chan time.Time, burst)
			for i := 0; i < burst; i++ {
				limiter <- time.Now()
			}

			go func() {
				for range time.Tick(millis * time.Millisecond) {
					limiter <- time.Now()
				}
			}()

			return limiter
		}

		limiter := produceLimiter(200, 3)

		go func() {
			for i := 1; i <= 5; i++ {
				c <- i
			}
			close(c)
		}()

		for i := range c {
			<-limiter
			fmt.Println(time.Now(), i)
		}

		// 2020-11-30 18:13:59.16743328 +0100 CET m=+1.000471258 1
		// 2020-11-30 18:13:59.167497422 +0100 CET m=+1.000535582 2
		// 2020-11-30 18:13:59.167698236 +0100 CET m=+1.000736277 3
		// 2020-11-30 18:13:59.367654034 +0100 CET m=+1.200692048 4
		// 2020-11-30 18:13:59.567636992 +0100 CET m=+1.400674993 5
	}

}
