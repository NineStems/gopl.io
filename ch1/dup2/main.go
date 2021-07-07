// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)
//book task 1.3
func main() {
	counts := make(map[string]int)
	filesDup := make(map[string]struct{})
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
			if checkDup(counts){
				filesDup[arg]= struct{}{}
			}
		}
	}
	for line := range filesDup {
			fmt.Printf("%s\n", line)
	}
}

func checkDup(counts map[string]int) bool {
	for _, n := range counts {
		if n > 1 {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
