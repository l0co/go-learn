package main

import "fmt"

// takes two ints and returns int
func plus(a int, b int) int {
	return a + b
}

// you can skip types if two or more consecutive arguments have the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

// takes two ints and returns nothing
func printme(a int, b int) {
	fmt.Println("print:", a, b)
}

// function can return multiple values
func plus1(a, b, c int) (int, int, int) {
	return a + 1, b + 1, c + 1
}

// variadic functions can be called with any number of arguments
func plus2(nums ...int) {
	var x []int = nums // in fact we get them in an array
	_ = x              // a way to stop complaining about "unused variable"

	for i, num := range nums {
		fmt.Println("plus2:", i, "->", num+2)
	}
}

func main() {

	{
		res := plus(1, 2)
		fmt.Println("1+2 =", res) // 1+2 = 3
	}

	{
		res := plusPlus(1, 2, 3)
		fmt.Println("1+2+3 =", res) // 1+2+3 = 6
	}

	{
		// x := printme(1, 2) // error: function returns nothing
		printme(1, 2) // print: 1 2
	}

	// multiple values
	{
		a, b, c := plus1(1, 2, 3)
		fmt.Println("plus1:", a, b, c) // plus1: 2 3 4
		_, _, x := plus1(1, 2, 3)      // we can get only a subset of return values
		fmt.Println("plus1:", x)       // plus1: 4
	}

	// variadic function
	plus2(1, 2, 3, 4, 5)
	// plus2: 0 -> 3
	// plus2: 1 -> 4
	// plus2: 2 -> 5
	// plus2: 3 -> 6
	// plus2: 4 -> 7

}
