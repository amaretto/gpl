package main

import (
	"fmt"
	"io"
)

//type ByteCounter int

type ByteCounter struct {
	writer  io.Writer
	counter int64
}

func main() {
	var bc ByteCounter
	cw, c := CountingWriter(&bc)

	cw.Write([]byte("hoge hoge fuga fuga hoge fuga hoge"))
	fmt.Println(*c)
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
	bc.counter += int64((len(p)))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var bc ByteCounter
	bc.writer = w
	return &bc, &(bc.counter)
}
