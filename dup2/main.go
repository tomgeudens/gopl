/*
description :
  dup2 prints the text of each line that appears more than once in the standard
  input or a list of named files, preceded by its count

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/22
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		// process standard input
		countLines(os.Stdin, counts)
	} else {
		// process named files
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)

			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
