package main

import "fmt"

func main() {

	// "while" type of for
	fmt.Println("loop1")
	i := 1
	for i <= 3 { // loops until the condition is true
		fmt.Println(i)
		i = i + 1
	}
	// 1, 2, 3

	// standard for loop
	fmt.Println("\nloop2")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 7, 8, 9

	// infinite loop
	fmt.Println("\nloop3")
	for {
		fmt.Println(1)
		break // withg break
	}
	// 1

	// continue goes to new loop iteration
	fmt.Println("\nloop4")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 { // modulo
			continue
		}
		fmt.Println(n)
	}
	// 1, 3, 5
}
