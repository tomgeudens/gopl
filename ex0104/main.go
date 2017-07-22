/*
description :
  ex0104 is a variation on dup2

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

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// provide space for the second level map
		if len(counts[input.Text()]) == 0 {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
}

func main() {
	// provide space for the first level map
	counts := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex0104: %v\n", err)
				continue
			}

			countLines(f, counts)

			f.Close()
		}
	}

	for line, entry := range counts {
		n := 0
		for _, count := range entry {
			n += count
		}
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for filename, count := range entry {
				fmt.Printf("\t%d\t%s\n", count, filename)
			}
		}
	}
}
