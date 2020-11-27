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

	ipointer := &i
	fmt.Println("pointer:", ipointer) // pointer: 0xc00001a0c8

	// we can convert pointer back to variable with *, but it allocates a new variable in memory
	x := *ipointer
	fmt.Println("x:", x, &x) // x: 0 0xc00001a110

	// recap:
	// * means pointer to (when used with type) and can also be used to convert pointers back to variables
	// & means "address of"
	// we assign addresses to pointer types: var pointer *int = &i

	// an important lesson is: asssigning anything by valyue means allocating it second time and creating an independent copy
	// regardless it's primitive value, array, struct ...
	{
		arr1 := []int{1, 2, 3}
		arr2 := arr1
		fmt.Printf("%p\n", &arr1) // 0xc00009a020
		fmt.Printf("%p\n", &arr2) // 0xc00009a040 -> a new array is allocated
	}

	// however, when no assignment to variable is done, converting pointer to value doesn't mean alocating a new variable:
	{
		fmt.Println(ipointer, &(*ipointer)) // 0xc0001a2010 0xc0001a2010
		*ipointer = 2
		fmt.Println(i) // 2
	}

}
