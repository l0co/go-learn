package main

import "fmt"

func main() {

	ch := make(chan string, 1)

	// with `chan<-` syntax we can create a channel which is used only to send values, but not receive
	var sender chan<- string = ch
	// with `<-chan` syntax we can create a channel which is used only to receive values, but not send
	var receiver <-chan string = ch
	// this can also be narrowed on function parameters

	sender <- "a"
	// receiver <- "a" // receiver can't be used to send values

	msg := <-receiver
	// msg := <-sender // sender can't be used to receive values

	fmt.Println(msg)
	// a

}
