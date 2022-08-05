package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	initialBufSize = 10000
	maxBufSize     = 100000000
)

var sc *bufio.Scanner = func() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, initialBufSize)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)
	return sc
}()

func main() {
	fm := make(map[string]int)
	for sc.Scan() {
		s := sc.Text()
		fm[s]++
	}
	fmt.Printf("\nword\tcount\n")
	for w, n := range fm {
		fmt.Printf("%q\t%d\n", w, n)
	}
}
