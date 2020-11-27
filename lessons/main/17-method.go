package main

import "fmt"

type intrect struct {
	width, height int
}

// this area method has a receiver type of *intrect.
func (r *intrect) area() int {
	return r.width * r.height
}

// methods can be defined for either pointer or value receiver types
func (r intrect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *intrect) resetPointer() {
	r.width = 0
	r.height = 0
}

func (r intrect) resetValue() {
	r.width = 0
	r.height = 0
}

func main() {
	r := intrect{width: 10, height: 5}

	// calling methods by struct value
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// calling methods by struct pointer
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	// this example shows that if the receiver is by value, the another variable copy is created, while by pointer it uses the same
	// variable and can modify it
	r.resetValue()
	fmt.Println(r) // {10 5}
	r.resetPointer()
	fmt.Println(r) // {0 0}
}
