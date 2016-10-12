/*
description :
  ex0104 is a variation on the dup2 program

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/10/12
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int, fileoccurences map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileoccurences[input.Text()] = append(fileoccurences[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

func main() {
	counts := make(map[string]int)
	fileoccurences := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileoccurences)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex0104: %v\n", err)
				continue
			}

			countLines(f, counts, fileoccurences)

			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for _, fo := range fileoccurences[line] {
				fmt.Printf("\toccurs in %s\n", fo)
			}
		}
	}
}
