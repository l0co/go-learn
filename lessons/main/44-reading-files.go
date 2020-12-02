package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const FILENAME = "/tmp/dat"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write() {
	d1 := []byte("hello\ngo\n")
	check(ioutil.WriteFile(FILENAME, d1, 0644))
}

func main() {

	// we write the file first, but this subject example is included in lesson 45, the context of the file is:
	// hello
	// go
	write()
	defer func() {
		fmt.Println("removing")
		check(os.Remove(FILENAME))
	}()

	// read file into memory with ReadFile
	dat, err := ioutil.ReadFile(FILENAME)
	check(err)
	fmt.Print(string(dat))
	// hello
	// go

	// open file for streaming read
	f, err := os.Open(FILENAME)
	check(err)
	defer func() {
		fmt.Println("closing")
		check(f.Close())
	}()

	// read some bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1])) // 5 bytes: hello

	// seek position + read
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))
	// 2 bytes @ 6: go

	// seek position + read at least
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	// 2 bytes @ 6: go

	// reset to zero position
	_, err = f.Seek(0, 0)
	check(err)

	// read using reader
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	// 5 bytes: hello

	// (deferred)
	// closing
	// removing
}
