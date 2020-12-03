package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

// 51-spawning-process shows running the external process from within go app process
// executing processes is replacing completely currently running go app process with the other process with execve

//       execve() executes the program referred to by pathname.  This causes
//       the program that is currently being run by the calling process to be
//       replaced with a new program, with newly initialized stack, heap, and
//       (initialized and uninitialized) data segments.

func main() {

	binary, lookErr := exec.LookPath("sleep") // lookup for full path
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"sleep", "2s"}
	env := os.Environ()

	fmt.Println("still in go process")
	<-time.After(time.Second * 2)

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	fmt.Println("no longer in go process") // never called

}

// $ go run 52-executing-proces.go & while true; do ps -ef|grep "$!"; sleep 1s; done
// lukasz   11982  6617  0 12:26 pts/0    00:00:00 go run 52-executing-proces.go
// still in go process
// lukasz   11982  6617  9 12:26 pts/0    00:00:00 go run 52-executing-proces.go
// lukasz   11982  6617  4 12:26 pts/0    00:00:00 go run 52-executing-proces.go
// lukasz   11982  6617  3 12:26 pts/0    00:00:00 go run 52-executing-proces.go
// lukasz   12066 11982  0 12:26 pts/0    00:00:00 sleep 2s
// (11982 is not really replaced, but becomes a parent for 12066)
