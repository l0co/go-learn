package main

import "fmt"

//noinspection GoBoolExpressions
func main() {

	// standard if
	if 7%2 == 0 { // no brackets for condition
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// 7 is odd

	// no else version
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	// 8 is divisible by 4

	// brackets are required so the following syntax doesn't work:
	// if 8%4 == 0
	//	fmt.Println("8 is divisible by 4")

	// one statement can precede condition and the declared variables is available inside brackets
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// fmt.Println(num) // no "num" variable here

	// two conditions are not possible
	// if num := 9; x := 2; num < 0 {}
}
