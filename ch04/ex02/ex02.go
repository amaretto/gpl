package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
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
	return sc
}()

func scanString() (s string) {
	sc.Scan()
	return sc.Text()
}

func main() {
	var option int

	flag.IntVar(&option, "len", 256, "hash length(256,384,512)")
	flag.Parse()

	s := scanString()

	switch option {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(s)))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(s)))
	}

}
