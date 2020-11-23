package main

import (
	"fmt"
	"go-learn/examplepkg"
)

func main() {
	//fmt.Println(examplepkg.value1) // can't use this variable because starts from lowercase character and is not imported
	fmt.Println(examplepkg.Value1)
}
