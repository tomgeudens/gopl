/*
description :
  echo1 prints its command-line arguments

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/17
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	var sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
