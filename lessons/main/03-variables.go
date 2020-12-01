package main

import (
	"fmt"
	"reflect"
)

func main() {

	// var declares a variable, then comes the variable type
	var a string = "initial"
	fmt.Println(a) // initial

	// type can be inferred
	var a2 = "initial"
	fmt.Println(a2) // initial

	// more variables can be declared in a single line
	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2

	// assigned too
	b, c = 3, 4
	fmt.Println(b, c) // 3 4

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

	// variable type case
	{
		var x int = 1
		fmt.Println(x, reflect.TypeOf(x)) // 1 int
		y := int64(x)
		fmt.Println(y, reflect.TypeOf(y)) // 1 int64
	}

}
