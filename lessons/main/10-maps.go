package main

import "fmt"

func main() {

	// map declaration
	var m map[string]int
	m = make(map[string]int)

	// map declaration with elements
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) // map: map[bar:2 foo:1]

	// adding / replacing elements
	m["k1"] = 7
	m["k2"] = 8
	m["k2"] = 13
	fmt.Println("map:", m) // map: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1: ", v1) // v1:  7

	// len works with maps
	fmt.Println("len:", len(m)) // len: 2

	// deleting elements
	delete(m, "k2")
	fmt.Println("map:", m) // map: map[k1:7]

	// in go function may return multiple values, here we're getting an element from map (first return value)
	// if it doesn't exists in the map, zero-valued value is returned (0 for int)
	// the second return value is a boolean telling if the element was present in the array or not
	el, prs := m["k2"]
	fmt.Println("el:", el)   // el: 0
	fmt.Println("prs:", prs) // prs: false

	// what happens on call to non-existing element?
	fmt.Println(m["non-existing"])
	// 0
	// so, zero-valued element is retuned in this scenario
	// to be more specific:
	y, exists := m["non-existing"]
	fmt.Println(y, exists)
	// 0 false

}
