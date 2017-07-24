/*
description :
  server3 is a minimal echo server that echoes request information

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/24
*/
package main

import "fmt"
import "log"
import "net/http"

func echoHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "%s %s %s\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(rw, "Host = %q\n", req.Host)
	fmt.Fprintf(rw, "RemoteAddr %q\n", req.RemoteAddr)
	err := req.ParseForm()
	if err != nil {
		log.Print(err)
	}
	for k, v := range req.Form {
		fmt.Fprintf(rw, "Form[%q] = %q\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
