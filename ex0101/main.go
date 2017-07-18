/*
description :
  ex0101 prints the name of the command that invoked it and its command-line
  arguments

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/18
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("invoking command       = " + os.Args[0])
	fmt.Println("command-line arguments = " + strings.Join(os.Args[1:], " "))
}
