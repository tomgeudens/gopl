/*
description :
  server1 is a minimal echo server

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/24
*/
package main

import "fmt"
import "log"
import "net/http"

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "URL.Path = %q\n", req.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
