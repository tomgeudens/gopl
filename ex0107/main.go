/*
description :
  ex0107 is a variation on the fetch program

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/10/12
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex0107: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex0107: copying %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
