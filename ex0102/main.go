/*
description :
  ex0102 prints its command-line arguments (and the index of the argument)

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/18
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d\t%s\n", i, os.Args[i])
	}
}
