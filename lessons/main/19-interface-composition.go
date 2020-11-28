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

// it's also possible to compose struct of interfaces
type intstruct2 struct {
	cint
}

func main() {
	s := intstruct{}
	test(s)

	// the following code produces PANIC error, but is compilable:

	// s2 := intstruct2{}
	// s2.a()

	// implementation of `cint` interface in intstruct2 is explicite, and compiler allows it to be used here, however
	// because the method is not really implemented by this struct, it throws PANIC

}
