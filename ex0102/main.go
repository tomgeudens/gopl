/*
description :
  ex0102 is a variation on the echo2 program

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/10/06
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
