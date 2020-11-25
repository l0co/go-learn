package main

import (
	"fmt"
	"time"
)

func main() {

	// basic switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// Write 2 as two

	// multi-options
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	// It's a weekday

	// switch without an expression allows to be used as "if" alternative
	t := time.Now()
	switch {
	case t.Hour() < 12: // full condition goes here
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	// (result depends on time)

	// type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
