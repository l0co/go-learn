package main

import (
	"fmt"
	"os"
)

// example of exit with status code

func main() {

	defer fmt.Println("exitting...") // never executed

	os.Exit(3)
}

// $ go build 54-exit.go
// $ ./54-exit; echo "$?"
// 3

// note that this works for compiled file but not for file executed with "go" (3 is not forwarded, but go returns 1)

// $ go run 54-exit.go; echo "$?"
// exit status 3
// 1
