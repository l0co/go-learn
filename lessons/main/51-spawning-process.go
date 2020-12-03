package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// executing process using Output
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	// > date
	// czw, 3 gru 2020, 12:15:10 CET

	// pipelining in/out
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
	// > grep hello
	// hello grep

	// executing with command line arguments
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
	// > ls -a -l -h
	// drwxr-xr-x 4 4,0K gru  3 11:03 .
	// drwxr-xr-x 6 4,0K lis 30 17:31 ..
	// -rw-r--r-- 1  128 lis 23 12:10 go.mod
	// drwxr-xr-x 2 4,0K gru  3 12:05 main
	// drwxr-xr-x 2 4,0K gru  3 11:03 test

}
