/*
description :
  echo3 prints its command-line arguments

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/17
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
