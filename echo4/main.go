/*
description :
  echo4 prints its command-line arguments

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/26
*/
package main

import "flag"
import "fmt"
import "strings"

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
