package main

import (
	"fmt"
)

// slice is a list
func main() {

	// slice is like array, here we create a slice with three zero-valued strings
	var s []string // slice variable is declared as array
	s = make([]string, 3)
	fmt.Println("emp:", s)           // emp: [  ]
	fmt.Println("len(emp):", len(s)) // lem(emp): 3
	if s[0] == "" {                  // zero-valued string is an empty string
		fmt.Println("empty string")
	}
	// empty string

	// declaring a slice with elements
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // dcl: [g h i]

	// setting and getting value is the same as for arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)    // set: [a b c]
	fmt.Println("get:", s[2]) // get: c

	// getting slice length is the same as for arrays
	fmt.Println("len:", len(s)) // len: 3

	// slice can have variable length
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) // apd: [a b c d e f]

	// slices can be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) // cpy: [a b c d e f]

	// our `s` slice is now: [a b c d e f]
	//                        0 1 2 3 4 5

	// it's possible to extract part of slice
	l := s[2:5]            // 2 is inclusive, 5 is exclusive !!!
	fmt.Println("sl1:", l) // sl1: [c d e]

	// an equivalent of 0:5
	l = s[:5]
	fmt.Println("sl2:", l) // sl2: [a b c d e]

	// an equivalent of 2:6
	l = s[2:]
	fmt.Println("sl3:", l) // sl3: [c d e f]

	// multi-dimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // each element is a separate slice which needs to be declared
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // [[0] [1 2] [2 3 4]]

	// slice has an underlying array with some capacity, by using the second argument of make we can define it
	{
		slice := make([]int, 3)                       // defines 3 element-s slice backed by 3 elements array
		fmt.Println("slice:", len(slice), cap(slice)) // slice: 3 3
		slice = make([]int, 0, 3)                     // defines 0 element-s slice backed by 3 elements array
		fmt.Println("slice:", len(slice), cap(slice)) // slice: 0 3
		slice = append(slice, 1, 2, 3)
		fmt.Println("slice:", len(slice), cap(slice)) // slice: 3 3
		// it's still possible to append elements to this slice regardless the capacity is insufficient
		// in such a scenario the slice will be "resliced" (a new array will be allocated)
		slice = append(slice, 4)
		fmt.Println("slice:", len(slice), cap(slice)) // slice: 4 6
	}

}
