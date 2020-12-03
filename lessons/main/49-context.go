package main

import (
	"fmt"
	"net/http"
	"time"
)

// a Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines

func hello1(w http.ResponseWriter, req *http.Request) {

	// separate context is created for each request
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {

	// wait 10 second before sending a response
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")

	// at the same time observe events from context Done channel
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)

	}
}

func main() {
	http.HandleFunc("/hello", hello1)
	http.ListenAndServe(":8090", nil)
}

// $ curl localhost:8090/hello
// server: hello handler started
// ^C
// server: context canceled
// server: hello handler ended
