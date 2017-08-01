/*
description :
  pc prints the popcount for the entered (uint64) arguments

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/31
*/
package main

import "github.com/tomgeudens/gopl/popcount"
import "github.com/tomgeudens/gopl/ex0203"
import "github.com/tomgeudens/gopl/ex0204"
import "github.com/tomgeudens/gopl/ex0205"
import "fmt"
import "os"
import "strconv"

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pc: %v\n", err)
			continue
		}

		fmt.Printf("Population (number of set bits, using popcount) for %d is %d\n", x, popcount.PopCount(x))
		fmt.Printf("Population (number of set bits, using ex0203) for %d is %d\n", x, ex0203.PopCount(x))
		fmt.Printf("Population (number of set bits, using ex0204) for %d is %d\n", x, ex0204.PopCount(x))
		fmt.Printf("Population (number of set bits, using ex0205) for %d is %d\n", x, ex0205.PopCount(x))
		fmt.Println()
	}
}
