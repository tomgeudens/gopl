/*
description :
  echo2 prints its command-line arguments

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/10/02
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	sep := ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
