package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"os"
)

func main() {
	var (
		t = flag.String("type", "jpeg", "type of decode file")
	)
	flag.Parse()
	if err := transform(*t, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", *t, err)
		os.Exit(1)
	}
}

func transform(t string, in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch t {
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	default:
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	}
}
