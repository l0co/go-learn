package main

import (
	"fmt"
	"sort"
)

// simple sorting examples

func simpleSortExample() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs) // Strings: [a b c]

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints) // Ints:    [2 4 7]

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s) // Sorted:  true
}

// sorting by function
// array to be sorted by function should implement sort.Interface

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func sortByFunctionExample() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits)) // we cast here []string to byLength
	fmt.Println(fruits)         // [kiwi peach banana]
}

func main() {
	simpleSortExample()
	sortByFunctionExample()
}
