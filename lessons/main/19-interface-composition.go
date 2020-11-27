package main

import "fmt"

type aint interface {
	a()
}

type bint interface {
	b()
}

// interface can be composed of different interfaces like a struct
type cint interface {
	aint
	bint
}

type intstruct struct{}

func (i intstruct) a() {
	fmt.Println("a")
}

func (i intstruct) b() {
	fmt.Println("b")
}

func test(c cint) {
	c.a()
	c.b()
}

func main() {
	s := intstruct{}
	test(s)
}
