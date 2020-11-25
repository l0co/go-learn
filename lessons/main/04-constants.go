package main

import (
	"fmt"
	"reflect"
)

// string constant
const s string = "constant"

func main() {
	fmt.Println(s) // contant

	// int constant
	const n = 500000000
	fmt.Println(reflect.TypeOf(n)) // int

	// float64 constant
	const d = 3e20 / n
	fmt.Println(reflect.TypeOf(d)) // float64
	fmt.Println(d)                 // 6e+11

	i64 := int64(d)
	fmt.Println(reflect.TypeOf(i64)) // int64
	fmt.Println(i64)                 // 600000000000
}
