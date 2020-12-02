package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// line readers reads reads lines from io.Reader line by line, for example from file or stdin
// execution example: printf "hello\nworld" | go run 46-line-filters.go
// results:
// HELLO
// WORLD
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
