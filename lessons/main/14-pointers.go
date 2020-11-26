package main

import "fmt"

// passing argument by value means allocating this value again on the function call stack
func zeroval(ival int) {
	ival = 0 // we can still modify this local variable copy
}

// passing argument by pointer means only allocating a pointer on the function call stack
func zeroptr(iptr *int) { // *int means a "pointer to int" type, not an int type itself
	fmt.Println("zeroptr call:", iptr, *iptr) // zeroptr call: 0xc00001a0c8 1
	*iptr = 0                                 // * gets access to the original variable declared and allocated outside this function call stack
}

func main() {
	i := 1
	fmt.Println("initial:", i) // initial: 1

	zeroval(i)
	fmt.Println("zeroval:", i) // zeroval: 1

	zeroptr(&i)                // & converts value to its pointer
	fmt.Println("zeroptr:", i) // zeroptr: 0
	{
		// once again, "pointer to int" type is *int
		var pointer *int = &i
		_ = pointer
	}

	fmt.Println("pointer:", &i) // pointer: 0xc00001a0c8
}
