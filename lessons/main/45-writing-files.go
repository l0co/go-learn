package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check1(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// writing file using WriteFile
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check1(err)
	defer func() {
		fmt.Println("removing /tmp/dat1")
		check1(os.Remove("/tmp/dat1"))
	}()
	// the content of /tmp/dat1 is:
	// hello
	// go

	// creating new file
	f, err := os.Create("/tmp/dat2") // file is being created and then opened for read&write (O_RDWR, 0666)
	defer func() {
		fmt.Println("removing /tmp/dat2")
		check1(os.Remove("/tmp/dat2"))
	}()

	check1(err)
	defer func() {
		fmt.Println("closing /tmp/dat2")
		check1(f.Close())
	}()

	// write some bytes
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check1(err)
	fmt.Printf("wrote %d bytes\n", n2)
	// wrote 5 bytes

	// write string
	n3, err := f.WriteString("writes\n")
	check1(err)
	fmt.Printf("wrote %d bytes\n", n3)
	// wrote 7 bytes

	check1(f.Sync()) // sync memory changes into the file

	// write using writer
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check1(err)
	fmt.Printf("wrote %d bytes\n", n4)
	// wrote 9 bytes

	check1(w.Flush())
	// the content of /tmp/dat2 is:
	// some
	// writes
	// buffered

	// (defer)
	// closing /tmp/dat2
	// removing /tmp/dat2
	// removing /tmp/dat1
}
