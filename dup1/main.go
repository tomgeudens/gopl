/*
description :
  dup1 prints the text of each line that appears more than once in the standard
  input, preceded by its count

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/19
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
