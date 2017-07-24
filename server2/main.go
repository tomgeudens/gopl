/*
description :
  server2 is a minimal echo and counter server

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/24
*/
package main

import "fmt"
import "log"
import "net/http"
import "sync"

var mu sync.Mutex
var count int

func echoHandler(rw http.ResponseWriter, req *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(rw, "URL.Path = %q\n", req.URL.Path)
}

func countHandler(rw http.ResponseWriter, req *http.Request) {
	mu.Lock()
	fmt.Fprintf(rw, "Count %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", echoHandler)
	http.HandleFunc("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
