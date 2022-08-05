package main

import (
	"bufio"
	"fmt"
)

type ByteCounter int
type WordCounter int
type RowCounter int

func main() {
	var c ByteCounter
	var w WordCounter
	var r RowCounter

	c.Write([]byte("hello"))
	fmt.Println(c)
	w.Write([]byte("hoge hoge hoge fuga foo bar hogera fugara"))
	fmt.Println(w)
	r.Write([]byte("hoge hoge hoge fuga\n foo\n bar hogera fugara"))
	fmt.Println(r)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {

	for {
		advance, _, err := bufio.ScanWords(p, true)
		if err != nil {
			panic("hoge")
		}
		*w++
		p = p[advance:]
		if len(p) == 0 {
			return len(p), nil
		}
	}
}

func (r *RowCounter) Write(p []byte) (int, error) {
	for {
		advance, _, err := bufio.ScanLines(p, true)
		if err != nil {
			panic("hoge")
		}
		*r++
		p = p[advance:]
		if len(p) == 0 {
			return len(p), nil
		}
	}
}
