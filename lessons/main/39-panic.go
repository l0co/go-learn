package main

import "os"

// a panic typically means something went unexpectedly wrong
// mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully
func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

}
