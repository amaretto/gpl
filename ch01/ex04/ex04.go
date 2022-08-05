package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	contains := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, contains)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, contains)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("counts : %d\tline : %s\tfiles : %s\n",
				n, line, strings.Join(contains[line], ","))
		}
	}
}

func countLines(f *os.File, counts map[string]int, contains map[string][]string) {
	var line string
	input := bufio.NewScanner(f)
	for input.Scan() {
		line = input.Text()
		if !exists(contains[line], f.Name()) {
			contains[line] = append(contains[line], f.Name())
		}
		counts[line]++
	}
}

func exists(list []string, s string) bool {
	for _, e := range list {
		if s == e {
			return true
		}
	}
	return false
}
