package main

import "fmt"

// closure is a function type (lambda)
func intSeq() func() int {
	i := 0
	return func() int {
		i++ // it has access to outer function scope and can modify variables from this scope
		return i
	}
}

func main() {

	var nextInt func() int = intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 1 2 3

	newInts := intSeq()
	fmt.Println(newInts())
	// 1

	fmt.Println(nextInt())
	// 4

}
