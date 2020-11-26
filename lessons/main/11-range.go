package main

import "fmt"

// range iterates over elements in a variety of data structures (it's like java iterator)
func main() {

	// the range returns two values, first one is the index and the second one is the element
	// and allows to iterate in `for` loop
	nums := []int{2, 3, 4}
	for i, num := range nums {
		fmt.Println("i:", i, "num:", num)
	}
	// i: 0 num: 2
	// i: 1 num: 3
	// i: 2 num: 4

	// same for map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// a -> apple
	// b -> banana

	// it's possible to only iterate through keys
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// key: a
	// key: b

	// range on string iterates over unicode character codes
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	// 0 103
	// 1 111
}
