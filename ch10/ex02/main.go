package main

import (
	_ "archive/tar"
	_ "archive/zip"
	"io"
)

func main() {
}

func readArchive(in io.Reader, out io.Writer) {
	r, err := NewReader(in)
}
