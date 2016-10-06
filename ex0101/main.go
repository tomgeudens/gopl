/*
description :
  ex0101 is a variation on the echo3 program

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/10/06
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
