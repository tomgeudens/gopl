/*
description :
  ex0107 is a variation on the fetch program

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/24
*/
package main

import "fmt"
import "io"
import "net/http"
import "os"

func main() {
	// loop through arguments (which should be urls)
	for _, url := range os.Args[1:] {
		// get url
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex0107: %v\n", err)
			os.Exit(1)
		}

		// read data
		written, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex0107: reading %s: %v\n", url, err)
		}

		// print size data
		fmt.Printf("\nURL %s has size %d\n", url, written)
	}
}
