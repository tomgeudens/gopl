/*
description :
  fetch prints the content found at each specified URL

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/23
*/
package main

import "fmt"
import "io/ioutil"
import "net/http"
import "os"

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", data)
	}
}
