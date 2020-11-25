package main

import "fmt"

func main() {

	// var declares a variable, then comes the variable type
	var a string = "initial"
	fmt.Println(a)

	// type can be inferred
	var a2 = "initial"
	fmt.Println(a2)

	// more variables can be declared in a single line
	var b, c int = 1, 2
	fmt.Println(b, c)

	// variables without value are zero-valued
	var d bool // false
	fmt.Println(d)
	var e int // zero
	fmt.Println(e)

	// a shorthand for `var f string = "apple"`
	f := "apple"
	fmt.Println(f)

	// variable scope
	{
		x := "x"
		fmt.Println(x)
	}
	// fmt.Println(x) // no "x" here

}
