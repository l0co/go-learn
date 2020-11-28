package main

import (
	"fmt"
	"math"
)

// interface declaration
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// interfaces are implicit, it's enough to implement all methods in struct to make it implementing an interface
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func interfacesTest() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

// this will also show one interesting thing about interfaces plus value vs pointer method receiver
type xint interface {
	x()
}

type xint1 struct{}
type xint2 struct{}

func (x xint1) x()    {} // by value
func (x *xint2) x()   {} // by pointer
func xreceive(x xint) {} // receiver function

func valueOrPointerInterfaceTest() {
	xint1Instance := xint1{}
	xint2Instance := xint2{}
	xreceive(xint1Instance)
	// xreceive(xint2Instance) // can't be called, I have to use pointer here:
	xreceive(&xint2Instance)
	// the explanation is funny here: cannot use xint2Instance (type xint2) as type xint in argument to xreceive:
	// xint2 does not implement xint (x method has pointer receiver)
	// so, if we have receiver method by pointer, type xint2 itself doesn't implement our interface, but POINTER to xint (*xint) DOES :)
}

func main() {
	interfacesTest()
	valueOrPointerInterfaceTest()

	// en empty interface can be everything
	var emptyInterface interface{} // by default zero-valued empty interface is nil
	fmt.Println(emptyInterface)    // <nil>
	emptyInterface = "s"
	fmt.Println(emptyInterface) // s
	emptyInterface = 1
	fmt.Println(emptyInterface) // 1

	// after assigning something to interface (can be empty) you can use type assertions for this interface
	{
		// first return value is the value itself (of given type), the second is boolean true when the variable is of this type
		val, isInt := emptyInterface.(int)
		fmt.Println(val, isInt) // 1 true
	}
	{
		val, isString := emptyInterface.(string)
		fmt.Println(val, isString) // false (val has no value in such a scenario)
		// fmt.Println(emptyInterface.(string)) // can't use this without the assignment above -> PANIC ERROR
	}

	// if you want to recover the variable from empty interface, you have to use type assertion
	type xxx struct {
		a int
	}
	emptyInterface = &xxx{1}

	// fmt.Println(emptyInterface.a) // can't do it because emptyInterface doesn't have any fields
	if xxxInstance, ok := emptyInterface.(*xxx); ok {
		fmt.Println(xxxInstance.a) // 1
	}

}
