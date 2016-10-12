/*
description :
  ex0103 measures the difference in duration of the different echo programs

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/10/12
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)

	var s string
	var sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	fmt.Println("echo1, elapsed", time.Since(start))

	time.Sleep(1000)
	start = time.Now()
	fmt.Println(start)

	s = ""
	sep = ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
	fmt.Println("echo2, elapsed", time.Since(start))

	time.Sleep(1000)
	start = time.Now()
	fmt.Println(start)

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("echo3, elapsed", time.Since(start))
}
