/*
description :
  server1 is a minimal "echo" server

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/11/04
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
