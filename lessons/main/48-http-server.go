package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func shutdown(w http.ResponseWriter, req *http.Request) {
	if server != nil {
		server.Close()
	}
}

var server *http.Server

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/shutdown", shutdown)

	// simplified version to start a server (no Server instance available)
	// http.ListenAndServe(":8090", nil)

	// version with Server instance available
	server = &http.Server{Addr: ":8090", Handler: nil}
	server.ListenAndServe()

}

// $ curl http://localhost:8090
// 404 page not found
// $ curl http://localhost:8090/hello
// hello
// $ curl http://localhost:8090/headers
// User-Agent: curl/7.47.0
// Accept: */*
// $ curl http://localhost:8090/shutdown
// curl: (52) Empty reply from server
// (and the application end)
